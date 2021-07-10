package main

import (
	"math/rand"
	"strings"
	"time"
)

// length- quantity of operations in expression
func GenerateExpression(length int) string {
	rand.Seed(time.Now().UnixNano())
	builder := strings.Builder{}
	var (
		operators = "+-"
		operands  = "0123456789"
	)

	generate := func(str string) string {
		return string(str[rand.Intn(len(str))])
	}
	IsSpace := func() string {
		if rand.Intn(2) == 0 {
			return ""
		}
		return " "
	}
	edger := func(str string) {
		builder.WriteString(IsSpace())
		builder.WriteString(str)
		builder.WriteString(IsSpace())
	}
	operand := func() {
		edger(generate(operands))
	}
	operator := func() {
		edger(generate(operators))
	}
	operand()

	for i := 0; i < length; i++ {
		operator()
		operand()
	}
	return builder.String()
}
