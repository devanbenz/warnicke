package main

import (
	"fmt"
	"github.com/parquet-go/parquet-go"
	"log"
)

func main() {
	type RowType struct {
		FirstName string
		LastName  string
		Big       bool
		Little    bool
	}

	rows, err := parquet.ReadFile[RowType]("cats2.parquet")
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range rows {
		fmt.Printf("%v\n", c)
	}
}
