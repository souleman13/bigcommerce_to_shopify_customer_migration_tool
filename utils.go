package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func readCSVFile(filename string) ([]byte, error) {
    f, err := os.Open(filename)
    if err!= nil {
        return nil, err
    }
    defer f.Close()
    data, err := io.ReadAll(f)
    if err!= nil {
        return nil, err
    }
    return data, nil
}

func parseCSV(data []byte) (*csv.Reader, error) {
    reader := csv.NewReader(bytes.NewReader(data))
    return reader, nil
}

func createCSVWriter(filename string) (*csv.Writer, *os.File, error) {
    f, err := os.Create(filename)
    if err != nil {
        return nil, nil, err
    }
    writer := csv.NewWriter(f)
    return writer, f, nil
}

func writeCSVRecord(writer *csv.Writer, record []string) {
    err := writer.Write(record)
    if err != nil {
        fmt.Println("Error writing record to CSV:", err)
    }
}