package main

import (
	"github.com/apache/arrow/go/v9/arrow"
	"slices"
)

type Schema struct {
	fields []Field
}

type Field struct {
	name  string
	value arrow.DataType
}

func (f *Field) toArrow() arrow.Field {
	return arrow.Field{
		Name: f.name,
		Type: f.value,
	}
}

func (s *Schema) toArrow() {
	for _, k := range s.fields {
		k.toArrow()
	}
}

func (s *Schema) projection(indices []int) Schema {
	var fields []Field
	for _, v := range indices {
		field := s.fields[v]
		fields = append(fields, field)
	}

	return Schema{
		fields: fields,
	}
}

func (s *Schema) schemaSelect(names []string) Schema {
	var selectedSchema []Field
	for k, v := range names {
		field := slices.ContainsFunc(s.fields, func(f Field) bool {
			return f.name == v
		})

		if field {
			selectedSchema = append(selectedSchema, s.fields[k])
		}
	}

	return Schema{
		selectedSchema,
	}
}
