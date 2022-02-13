Coding Technical Assessment: Data Reconcilation

The purpose of this technical assessment is to gauge your coding skill and your thinking on a data reconciliation sample problem.

You may write-up the assessment in any mainstream coding language, but as Xendit primarily uses Typescript and Go, submissions in these languages are preferred.

Evaluation criteria

Coding style and code readability. Choose good variable names, with the right balance between brevity and verboseness. Write short functions with meaningful names.
See clean code by Robert C. Martin (https://www.youtube.com/watch?v=7EmboKQH8lM&t=3336).

Unit tests. Ensure you code tests for the required behaviours. See, for example, the unit test package for Go: https://pkg.go.dev/testing

Function. Your code must work correctly.

Instructions
Reconcile source.csv and proxy.csv for the month of July 2021.

Think of source.csv as a bank statement and proxy.csv as your own record of the transactions.

*A. Produce an output report, in csv format, of mismatched, i.e. non-reconciled entries. The report format should follow proxy.csv's format but have an additional Remarks column appended. The Remarks column should highlight the reason for the discrepancy. Think about how to make the Remarks column user-friendly.

B. Produce a summary report, in text format, listing the following data:

date range for the report.
number of source records processed.
numbers and types of discrepancies.
C. Package your code and output reports (A and B above) as a zip file.

Development environment and constraints
You may use any mainstream language (eg. nodejs, python, java etc.), but submissions in TypeScript (https://www.typescriptlang.org/) and Go (https://golang.org/) are preferred.

You are free to consult any resource and use any library, but the work must be your own.

You should be able to complete this assignment within a half-day. However you should not take more than 3 days to complete the assignment. You are free to complete this assignment at your own pace and timing.
