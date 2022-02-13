package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Data struct {
	Amount      int
	Description string
	Date        string
	ID          string
}

type Reconsil struct {
	Date       string
	Debit      int
	Credit     int
	DescDebit  string
	DescCredit string
}

func Write(inputData []Reconsil) {
	file, err := os.Create("reconsiliation.csv")
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Successfully create file .csv")
	w := csv.NewWriter(file)
	defer w.Flush()

	//Using WriteAll
	var data [][]string
	data = append(data, []string{
		"Date",
		"Description",
		"Debit",
		"Credit",
	})
	for _, record := range inputData {
		row1 := []string{record.Date, record.DescDebit, strconv.Itoa(record.Debit), " "}
		row2 := []string{"", record.DescCredit, " ", strconv.Itoa(record.Credit)}
		data = append(data, row1, row2)
	}
	w.WriteAll(data)
}

func Read(file string) ([]Data, error) {
	index := map[string]int{}
	csvFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully Opened CSV file", file)
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(csvLines[0]); i++ {
		index[string(csvLines[0][i])] = i
	}

	Source := []Data{}
	for i, line := range csvLines {
		if i == 0 {
			continue
		}
		amountInt, _ := strconv.Atoi(line[index["Amount"]])
		dataFromcsv := Data{
			Date:        line[index["Date"]],
			ID:          line[index["ID"]],
			Amount:      amountInt,
			Description: line[index["Description"]],
		}
		Source = append(Source, dataFromcsv)
	}
	return Source, nil
}

func main() {
	var result []Reconsil
	dataSource, err := Read("source.csv")
	if err != nil {
		fmt.Println(err.Error())
	}
	dataProxy, err := Read("proxy.csv")
	if err != nil {
		fmt.Println(err.Error())
	}
	hashmap := map[string]Data{}
	for _, colSource := range dataSource {
		hashmap[colSource.ID] = colSource
	}
	for _, colProxy := range dataProxy {
		_, isExist := hashmap[colProxy.ID]
		if isExist {
			amountSource := hashmap[colProxy.ID]
			if colProxy != hashmap[colProxy.ID] {
				diff_Amout := colProxy.Amount - amountSource.Amount
				result = append(result, Reconsil{
					Date:       colProxy.Date,
					DescDebit:  "Pettycash on Source",
					Debit:      diff_Amout,
					DescCredit: colProxy.Description,
					Credit:     diff_Amout,
				})
			}
		} else {
			result = append(result, Reconsil{
				Date:       colProxy.Date,
				DescDebit:  colProxy.Description,
				Debit:      colProxy.Amount,
				DescCredit: "Pettycash on Source",
				Credit:     colProxy.Amount,
			})
		}
	}
	Write(result)
}
