package main

import "strings"

type AddressLineItem struct {
	Email string
	Item string
}


func main() {
	//load initial csv file into memory
	data, err := readCSVFile("bigcommerce-customers-export.csv")
	if err!= nil {
		panic(err)
    }
	//parse initial csv into usuable format
	reader, err := parseCSV(data)
    if err!= nil {
        panic(err)
    }
	
	//remove header line
	_, err = reader.Read()
	if err != nil {
		panic(err)
	}

	//read records into array
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	//create csv writer to new file
	writer, file, err := createCSVWriter("matrixify-customers-import.csv")
	if err != nil {
        panic(err)
    }
	//defer file close to end of execution
	defer file.Close()

	//for addresses
	for _, record := range records {
		addresses := strings.Split(record[1], "|")
		if len(addresses) > 1 {
			for _, address := range addresses {
				newAddressLine := AddressLineItem{
					Email: record[0],
					Item: address,
				}
				writeCSVRecord(writer, []string{newAddressLine.Email,newAddressLine.Item})
			}
		} else {
			newAddressLine := AddressLineItem{
				Email: record[0],
				Item: record[1],
			}
			writeCSVRecord(writer, []string{newAddressLine.Email,newAddressLine.Item})
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
        panic(err)
    }
}