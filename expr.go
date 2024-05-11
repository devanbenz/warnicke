package main

import (
	"fmt"
	"github.com/apache/arrow/go/v9/arrow"
)

type LogicalExpr interface {
	toField() Field
}

type ColumnExpr struct {
	name string
	lp   LogicalPlan
}

type AggregateExpr struct {
	name string
	expr LogicalExpr
}

type LiteralStrExpr struct {
	value string
}

type LiteralIntExpr struct {
	value int
}

type BinaryBoolExpr struct {
	name string
	op   string
	l    LogicalExpr
	r    LogicalExpr
}

type BinaryMathExpr struct {
	name string
	op   string
	l    LogicalExpr
	r    LogicalExpr
}

func (e *ColumnExpr) toField() Field {
	s := e.lp.getSchema()
	ss := s.schemaSelect([]string{e.name})

	if len(ss.fields) == 0 {
		return Field{}
	}
	return ss.fields[0]
}

func (e *ColumnExpr) String() string {
	return fmt.Sprintf("#%s", e.name)
}

func (e *LiteralStrExpr) toField() Field {
	return Field{
		name:  e.value,
		value: &arrow.StringType{},
	}
}

func (e *LiteralStrExpr) String() string {
	return fmt.Sprintf("%s", e.value)
}

func (e *LiteralIntExpr) toField() Field {
	return Field{
		name:  fmt.Sprintf("%d", e.value),
		value: &arrow.Int64Type{},
	}
}

func (e *LiteralIntExpr) String() string {
	return fmt.Sprintf("%d", e.value)
}

// /Binary Boolean Expressions ////////////////
// ////////////////////////////////////////////
func (b *BinaryBoolExpr) toField() Field {
	return Field{
		name:  b.name,
		value: &arrow.BooleanType{},
	}
}

func (b *BinaryBoolExpr) String() string {
	return fmt.Sprintf("%s %s %s", b.l.toField().name, b.op, b.l.toField().name)
}

func gt(l LogicalExpr, r LogicalExpr) BinaryBoolExpr {
	return BinaryBoolExpr{
		name: "gt",
		op:   ">",
		l:    l,
		r:    r,
	}
}

func gte(l LogicalExpr, r LogicalExpr) BinaryBoolExpr {
	return BinaryBoolExpr{
		name: "gte",
		op:   ">=",
		l:    l,
		r:    r,
	}
}

func lt(l LogicalExpr, r LogicalExpr) BinaryBoolExpr {
	return BinaryBoolExpr{
		name: "lt",
		op:   "<",
		l:    l,
		r:    r,
	}
}

func lte(l LogicalExpr, r LogicalExpr) BinaryBoolExpr {
	return BinaryBoolExpr{
		name: "lte",
		op:   "<=",
		l:    l,
		r:    r,
	}
}

func eq(l LogicalExpr, r LogicalExpr) BinaryBoolExpr {
	return BinaryBoolExpr{
		name: "eq",
		op:   "=",
		l:    l,
		r:    r,
	}
}

func neq(l LogicalExpr, r LogicalExpr) BinaryBoolExpr {
	return BinaryBoolExpr{
		name: "neq",
		op:   "!=",
		l:    l,
		r:    r,
	}
}

func and(l LogicalExpr, r LogicalExpr) BinaryBoolExpr {
	return BinaryBoolExpr{
		name: "and",
		op:   "AND",
		l:    l,
		r:    r,
	}
}

func or(l LogicalExpr, r LogicalExpr) BinaryBoolExpr {
	return BinaryBoolExpr{
		name: "or",
		op:   "OR",
		l:    l,
		r:    r,
	}
}

// Binary Math Expression ////////
// ///////////////////////////////
func (b *BinaryMathExpr) toField() Field {
	return Field{
		name:  b.name,
		value: b.l.toField().value,
	}
}

func (b *BinaryMathExpr) String() string {
	return fmt.Sprintf("%s %s %s", b.l.toField().name, b.op, b.l.toField().name)
}

func add(l LogicalExpr, r LogicalExpr) BinaryMathExpr {
	return BinaryMathExpr{
		name: "add",
		op:   "+",
		l:    l,
		r:    r,
	}
}

func subtract(l LogicalExpr, r LogicalExpr) BinaryMathExpr {
	return BinaryMathExpr{
		name: "sub",
		op:   "-",
		l:    l,
		r:    r,
	}
}

func multiply(l LogicalExpr, r LogicalExpr) BinaryMathExpr {
	return BinaryMathExpr{
		name: "mult",
		op:   "*",
		l:    l,
		r:    r,
	}
}

func divide(l LogicalExpr, r LogicalExpr) BinaryMathExpr {
	return BinaryMathExpr{
		name: "div",
		op:   "/",
		l:    l,
		r:    r,
	}
}

func modulus(l LogicalExpr, r LogicalExpr) BinaryMathExpr {
	return BinaryMathExpr{
		name: "mod",
		op:   "%",
		l:    l,
		r:    r,
	}
}

// Aggregate Expressions ////////
// ///////////////////////////////
func (a *AggregateExpr) toField() Field {
	return Field{
		name:  a.name,
		value: a.expr.toField().value,
	}
}

func (a *AggregateExpr) String() string {
	return fmt.Sprintf("%s(%s)", a.name, a.expr)
}

func minAgg(input LogicalExpr) AggregateExpr {
	return AggregateExpr{
		name: "MIN",
		expr: input,
	}
}

func maxAgg(input LogicalExpr) AggregateExpr {
	return AggregateExpr{
		name: "MAX",
		expr: input,
	}
}

func sumAgg(input LogicalExpr) AggregateExpr {
	return AggregateExpr{
		name: "SUM",
		expr: input,
	}
}

func avgAgg(input LogicalExpr) AggregateExpr {
	return AggregateExpr{
		name: "AVG",
		expr: input,
	}
}

func count(input LogicalExpr) AggregateExpr {
	return AggregateExpr{
		name: "COUNT",
		expr: input,
	}
}
