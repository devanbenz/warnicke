package main

import (
	"errors"
	"github.com/apache/arrow/go/v9/arrow"
)

type LiteralValueVector struct {
	arrowType arrow.Type
	value     interface{}
	size      int
}

type RecordBatch struct {
	schema arrow.Schema
	fields []ColumnVector
}

type ColumnVector interface {
	getType() arrow.Type
	getValue(i int) interface{}
	getSize() int
}

func (d *LiteralValueVector) getType() arrow.Type {
	return d.arrowType
}

func (d *LiteralValueVector) getValue(i int) interface{} {
	if i < 0 || i > d.size {
		panic("out of bounds error")
	}

	return d.value
}

func (d *LiteralValueVector) getSize() int {
	return d.size
}

func (b *RecordBatch) getRowCount() (int, error) {
	if len(b.fields) == 0 {
		return 0, errors.New("no fields in column vector")
	}

	return b.fields[0].getSize(), nil
}

func (b *RecordBatch) getColCount() int {
	return len(b.fields)
}

func (b *RecordBatch) getField(i int) ColumnVector {
	return b.fields[i]
}
