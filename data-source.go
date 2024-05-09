package main

type DataSource interface {
	schema() Schema
	projection(projection []string) chan RecordBatch
}
