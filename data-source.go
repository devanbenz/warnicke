package main

type DataSource interface {
	schema() Schema
	projection(projection []string) chan RecordBatch
}

type ParquetDataSource struct {
	path string
}

func (p ParquetDataSource) readParquet() {
}
