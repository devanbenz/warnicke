package main

import (
	"errors"
	"fmt"
	"strings"
)

type Plan struct {
	schema   Schema
	children []LogicalPlan
}

type LogicalPlan interface {
	schema() Schema
	children() []LogicalPlan
}

func (p *Plan) String() string {
	var str strings.Builder
	for _, v := range p.schema.fields {
		fmt.Fprintf(&str, "%s", v.name)
	}
	return fmt.Sprintf("%s", str)
}

func (p *Plan) print(indent int) (string, error) {
	var printout strings.Builder
	i := 0
	for i <= indent {
		_, err := fmt.Fprint(&printout, "\t")
		if err != nil {
			return "", errors.New("error with string builder writing")
		}
	}
	_, err := fmt.Fprintf(&printout, "%s", p)
	if err != nil {
		return "", err
	}

	return "", nil
}
