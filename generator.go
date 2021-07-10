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

	generate := func(str string) string {
		return string(str[rand.Intn(len(str))])
	}
	IsSpace := func() string {
		if rand.Intn(2) == 0 {
			return ""
		}
		return " "
	}
	var edger = func(str string) {
		builder.WriteString(IsSpace())
		builder.WriteString(str)
		builder.WriteString(IsSpace())
	}
	var operand = func() {
		var operands = "0123456789"
		edger(generate(operands))
	}
	var operator = func() {
		var operators = "+-"
		edger(generate(operators))
	}
	operand()

	for i := 0; i < length; i++ {
		operator()
		operand()
	}
	return builder.String()
}
