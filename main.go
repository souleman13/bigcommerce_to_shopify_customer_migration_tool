package main

import "strings"

type LineItem struct {
	Email string
	Item string
}

func main() {
	//load initial csv file into memory
	data, err := readCSVFile("OrderProductExports.csv")
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
	writer, file, err := createCSVWriter("Order-Product-Brokendown-v2.csv")
	if err != nil {
        panic(err)
    }
	//defer file close to end of execution
	defer file.Close()

	//for order products
	for _, record := range records {
		products := strings.Split(record[1], "|")
		if len(products) > 1 {
			for _, product := range products {
				newAddressLine := LineItem{
					Email: record[0],
					Item: product,
				}
				writeCSVRecord(writer, []string{newAddressLine.Email,newAddressLine.Item})
			}
		} else {
			newAddressLine := LineItem{
				Email: record[0],
				Item: record[1],
			}
			writeCSVRecord(writer, []string{newAddressLine.Email,newAddressLine.Item})
		}
	}

	//for addresses
	// for _, record := range records {
	// 	addresses := strings.Split(record[1], "|")
	// 	if len(addresses) > 1 {
	// 		for _, address := range addresses {
	// 			newAddressLine := LineItem{
	// 				Email: record[0],
	// 				Item: address,
	// 			}
	// 			writeCSVRecord(writer, []string{newAddressLine.Email,newAddressLine.Item})
	// 		}
	// 	} else {
	// 		newAddressLine := LineItem{
	// 			Email: record[0],
	// 			Item: record[1],
	// 		}
	// 		writeCSVRecord(writer, []string{newAddressLine.Email,newAddressLine.Item})
	// 	}
	// }

	writer.Flush()
	if err := writer.Error(); err != nil {
        panic(err)
    }
}