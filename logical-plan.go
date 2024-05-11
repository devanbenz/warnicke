package main

import (
	"fmt"
	"strings"
)

type LogicalPlan interface {
	getSchema() Schema
	getChildren() []LogicalPlan
}

type Scan struct {
	path       string
	datasource DataSource
	projection []string
}

type Projection struct {
	input LogicalPlan
	expr  []LogicalExpr
}

type Filter struct {
	input LogicalPlan
	expr  LogicalExpr
}

type Aggregate struct {
}

func format(logicalPlan LogicalPlan, indent int) (string, error) {
	var printout strings.Builder
	for i := 0; i < indent; i++ {
		_, err := fmt.Fprint(&printout, "\t")
		if err != nil {
			return "", err
		}
	}
	_, err := fmt.Fprintf(&printout, "%s\n", logicalPlan)
	if err != nil {
		return "", err
	}

	for _, v := range logicalPlan.getChildren() {
		fp, err := format(v, indent+1)
		_, err = fmt.Fprintf(&printout, "%s", fp)
		if err != nil {
			return "", err
		}

	}

	return fmt.Sprintf("%s", &printout), nil
}

func deriveSchema(datasource DataSource, projection []string) Schema {
	schema := datasource.schema()
	if len(projection) == 0 {
		return schema
	}

	return schema.schemaSelect(projection)
}

func (s *Scan) getSchema() Schema {
	return deriveSchema(s.datasource, s.projection)
}

func (s *Scan) getChildren() []LogicalPlan {
	return []LogicalPlan{}
}

func (s *Scan) String() string {
	if len(s.projection) == 0 {
		return fmt.Sprintf("Scan: %s; projection=None", s.path)
	}

	return fmt.Sprintf("Scan: %s; projection=%s", s.path, s.projection)
}

func (p *Projection) getSchema() Schema {
	var schemaFields []Field

	for _, expr := range p.expr {
		schemaFields = append(schemaFields, expr.toField())
	}

	return Schema{
		fields: schemaFields,
	}
}

func (p *Projection) getChildren() []LogicalPlan {
	return []LogicalPlan{
		p.input,
	}
}

func (f *Filter) getSchema() Schema {
	// selections do not change the schema at all
	return f.input.getSchema()
}

func (f *Filter) getChildren() []LogicalPlan {
	return []LogicalPlan{
		f.input,
	}
}

func (f *Filter) String() string {
	return fmt.Sprintf("Filter: %s", f.expr)
}
