package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

type Customer struct {
	Email string
	FirstName string
	LastName string
	Phone string
	Note string
	Tags string
}

//values listed in order of big commerce export
type Address struct {
	AddressFirstName string
	AddressLastName string
	AddressCompany string
	AddressLine1 string
	AddressLine2 string
	AddressCity string
	AddressProvinceCode string
	AddressZip string
	AddressCountry string
	AddressPhone string
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
	//will create a new file if does not exist
	writer, file, err := createCSVWriter("matrixify-customers-import.csv")
	if err != nil {
        panic(err)
    }
	defer file.Close()

	//create header line for output csv
	writeCSVRecord(writer,[]string{
		"email","command","first name","last name","phone","note","tags","tags command",
		"address command","address first name","address last name","address company", "address line 1","address line 2",
		"address city","address province code","address ","address country","address zip",
	})

	//for addresses
	for _, record := range records {
		//check record has address content
		if record[10] != "" {
			//create customer values for record
			customer := Customer {
			Email : record[4],
			FirstName : record[1],
			LastName : record[2],
			Phone : record[5],
			Note : "Joined Big Commerce: "+record[9]+" - "+record[6],
			Tags : "BigCommerce,",
			}
			//check for multiple addresses in cell
			addresses := strings.Split(record[10], "|")
			if len(addresses) > 1 {
				//loop over multiple addresses, create new line for each
				for _, address := range addresses {
					newAddress := breakdownAddress(address)
					writeNewRecord(writer, customer, newAddress)
				}
			} else {
				//create single address line
				address := record[10]
				newAddress := breakdownAddress(address)
				writeNewRecord(writer, customer, newAddress)
			} 
		} else {
			fmt.Println("NO ADDRESS")
		}
		
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
        panic(err)
    }
}

func writeNewRecord(writer *csv.Writer, customer Customer, address Address){
	writeCSVRecord(writer, []string{
		//basic cols - email, command, first, last, phone, note, tags, tags command
		customer.Email, "MERGE", customer.FirstName, customer.LastName, customer.Phone,customer.Note,customer.Tags, "MERGE",
		//address cols - command, first, last, company, phone, line1, line2, city, provinceCode, country, zip
		"MERGE", address.AddressFirstName, address.AddressLastName, address.AddressCompany, address.AddressPhone,
		address.AddressLine1, address.AddressLine2, address.AddressCity, address.AddressProvinceCode, 
		address.AddressCountry, address.AddressZip,
	})
}

func breakdownAddress(address string) Address {
	//Address First Name: Daniel, Address Last Name: Ramirez, Address Company: , 
	// Address Line 1: 221 E. Indianola Ave, Address Line 2: , City/Suburb: Phoenix, 
	// State Abbreviation: AZ, Zip/Postcode: 85012, Country: United States, Address Phone: 
	addressParts := strings.Split(address, ",")
	addressValues := []string{}
	for _, addressKeyValue := range addressParts {
		value := strings.Split(addressKeyValue, ": ")
		if len(value) > 1 {
			addressValues = append(addressValues,value[1])
		} else {
			addressValues = append(addressValues,"")
		}
	}
	return Address{
		AddressFirstName : addressValues[0],
		AddressLastName : addressValues[1],
		AddressCompany : addressValues[2],
		AddressLine1 : addressValues[3],
		AddressLine2 : addressValues[4],
		AddressCity : addressValues[5],
		AddressProvinceCode : addressValues[6],
		AddressZip : addressValues[7],
		AddressCountry : addressValues[8],
		AddressPhone : addressValues[9],
	}
}