package main

// Package examples contains examples demonstrating the usage of the tablewriter library.
package examples_test

import (
	"context"
	"fmt"
	"log"

	"github.com/njchilds90/go-tablewriter"
)

// ExampleNewTableWriter demonstrates how to create a new TableWriter instance.
func ExampleNewTableWriter() {
	// Create a new table writer with default settings.
	tw := tablewriter.NewTableWriter()
	// Output: 
}

// ExampleTableWriterWrite demonstrates how to write data to a TableWriter instance.
func ExampleTableWriterWrite() {
	tw := tablewriter.NewTableWriter()
	data := [][]string{
		{"Name", "Age"},
		{"John", "25"},
		{"Alice", "30"},
	}
	// Write the data to the table writer.
	err := tw.Write(context.Background(), data)
	if err != nil {
		log.Fatal(err)
	}
	// Output: 
}
