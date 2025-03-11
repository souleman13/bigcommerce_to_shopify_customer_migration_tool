package main

import "strings"

type Customer struct {
	Email string
	FirstName string
	LastName string
	Phone string
	Note string
	Tags string
}
type Address struct {
	AddressFirstName string
	AddressLastName string
	AddressCompany string
	AddressPhone string
	AddressLine1 string
	AddressLine2 string
	AddressCity string
	AddressProvinceCode string
	AddressCountry string
	AddressZip string
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
		customer := Customer {
			Email : record[4],
			FirstName : record[1],
			LastName : record[2],
			Phone : record[5],
			Note : "Joined Big Commerce: "+record[9]+" - "+record[6],
			Tags : "BigCommerce,",
		}
		addresses := strings.Split(record[10], "|")
		if len(addresses) > 1 {
			for _, address := range addresses {
				//Address First Name: Daniel, Address Last Name: Ramirez, Address Company: , 
				// Address Line 1: 221 E. Indianola Ave, Address Line 2: , City/Suburb: Phoenix, 
				// State Abbreviation: AZ, Zip/Postcode: 85012, Country: United States, Address Phone: 
				newAddress := Address{}
				addressParts := strings.Split(address, ",")
				for _, addressValue := range addressParts {
					values := strings.Split(addressValue, ":")
					//bc col order
					newAddress.AddressFirstName = values[0]
					newAddress.AddressLastName = values[1]
					newAddress.AddressCompany = values[2]
					newAddress.AddressLine1 = values[3]
					newAddress.AddressLine2 = values[4]
					newAddress.AddressCity = values[5]
					newAddress.AddressProvinceCode = values[6]
					newAddress.AddressZip = values[7]
					newAddress.AddressCountry = values[8]
					newAddress.AddressPhone = values[9]
				}
				
				writeCSVRecord(writer, []string{
					//basic cols - email, command, first, last, phone, note, tags, tags command
					customer.Email, "MERGE", customer.FirstName, customer.LastName, customer.Phone,customer.Note,customer.Tags, "MERGE",
					//address cols - command, first, last, company, phone, line1, line2, city, provinceCode, country, zip
					"MERGE", newAddress.AddressFirstName, newAddress.AddressLastName, newAddress.AddressCompany, newAddress.AddressPhone,
					newAddress.AddressLine1, newAddress.AddressLine2, newAddress.AddressCity, newAddress.AddressProvinceCode, 
					newAddress.AddressCountry, newAddress.AddressZip,
				})
			}
		} else {
			addressVals := Address{
				AddressFirstName : "",
				AddressLastName : "",
				AddressCompany : "",
				AddressPhone : "",
				AddressLine1 : "",
				AddressLine2 : "",
				AddressCity : "",
				AddressProvinceCode : "",
				AddressCountry : "",
				AddressZip : "",
			}
			writeCSVRecord(writer, []string{
				//basic cols - email, command, first, last, phone, note, tags, tags command
				customer.Email, "MERGE", customer.FirstName, customer.LastName, customer.Phone,customer.Note,customer.Tags, "MERGE",
				//address cols - command, first, last, company, phone, line1, line2, city, provinceCode, country, zip
				"MERGE", addressVals.AddressFirstName, addressVals.AddressLastName, addressVals.AddressCompany, addressVals.AddressPhone,
				addressVals.AddressLine1, addressVals.AddressLine2, addressVals.AddressCity, addressVals.AddressProvinceCode, 
				addressVals.AddressCountry, addressVals.AddressZip,
			})
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
        panic(err)
    }
}