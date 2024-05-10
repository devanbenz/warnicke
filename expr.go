package main

import "slices"

type LogicalExpr interface {
	toField(input LogicalPlan) Field
}

type ColumnExpr struct {
	name string
}

func (e *ColumnExpr) toField(input LogicalPlan) Field {
	s := input.getSchema()
	ss := s.schemaSelect([]string{})

	return Field{
		name: e.name,
		value:
	}
}
