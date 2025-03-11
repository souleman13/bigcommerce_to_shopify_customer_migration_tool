# Big Commerce to Shopify (via Matrixify) Migration Tool - Multi-Address Customers

Takes a Big Commerce customer export .csv file and converts it to Matrixify (Shopify app) import format.

This tool is meant to be a starting point and does not intake call customer data, it is focused on breaking apart multi address cells. However, with some simple modification it can also extract additional data if needed.

## How to
* add big commerce customers export csv to project. Rename the file 'bigcommerce-customers-export.csv'. This file should live in the same directory as main.go
* run the tool using 'go run .'
* output file should appear named 'matrixify-customers-import.csv'