package main

import (
	"fmt"
	"strings"
)

type LogicalPlan interface {
	getSchema() Schema
	getChildren() []LogicalPlan
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
