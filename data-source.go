package main

import "github.com/apache/arrow/go/v9/arrow"

type DataSource interface {
	schema() arrow.Schema
	projection(projection []string) chan RecordBatch
}
