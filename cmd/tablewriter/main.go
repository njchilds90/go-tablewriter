package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"njchilds90/go-tablewriter"
)

func main() {
	// Initialize the flag set
	var (
		rows int
		columns int
	)
	flag.IntVar(&rows, "rows", 5, "number of rows")
	flag.IntVar(&columns, "columns", 3, "number of columns")
	flag.Parse()

	// Create a new table writer
	tw := tablewriter.New(os.Stdout)

	// Set the header
	headers := make([]string, columns)
	for i := range headers {
		headers[i] = fmt.Sprintf("Column %d", i+1)
	}
	tw.SetHeader(headers)
	tw.SetAlignment(tablewriter.CENTER)

	// Generate some data
	data := make([][]string, rows)
	for i := range data {
		data[i] = make([]string, columns)
		for j := range data[i] {
			data[i][j] = fmt.Sprintf("Row %d, Column %d", i+1, j+1)
		}
	}
	tw.Render(data)

	if err := tw.Render(); err != nil {
		log.Printf("Failed to render table: %v", err)
		os.Exit(1)
	}
}
