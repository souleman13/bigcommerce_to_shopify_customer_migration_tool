# Big Commerce to Shopify (via Matrixify) Migration Tool - Multi-Address Customers

Takes a Big Commerce customer export .csv file and converts it to Matrixify (Shopify app) import format.

This tool is meant to be a starting point and does not intake call customer data, it is focused on breaking apart multi address cells. However, with some simple modification it can also extract additional data if needed.

## How to
* go to your big commerce portal, navigate to orders and click 'export'
* add big commerce customers export csv to project. Rename the file 'bigcommerce-customers-export.csv'. This file should live in the same directory as main.go
* run the tool using 'go run .'
* output file will appear named 'matrixify-customers-import.csv'

## Gotcha's
* ensure data doesn't have a '|' inside a data row. This is used as a seperator between address records, including one inside an address record will cause an error or produce incorrect data.