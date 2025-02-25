package main

import "strings"

type AddressLine struct {
	Email string
	Address string
}

func main() {
	//load initial csv file into memory
	data, err := readCSVFile("ExportedAddresses.csv")
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
	writer, file, err := createCSVWriter("Matrixify-Multi-Address_Import.csv")
	if err != nil {
        panic(err)
    }
	//defer file close to end of execution
	defer file.Close()

	for _, record := range records {
		addresses := strings.Split(record[1], "|")
		if len(addresses) > 1 {
			for _, address := range addresses {
				newAddressLine := AddressLine{
					Email: record[0],
					Address: address,
				}
				writeCSVRecord(writer, []string{newAddressLine.Email,newAddressLine.Address})
			}
		} else {
			newAddressLine := AddressLine{
				Email: record[0],
				Address: record[1],
			}
			writeCSVRecord(writer, []string{newAddressLine.Email,newAddressLine.Address})
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
        panic(err)
    }
}