package expression_generator

import (
	"math/rand"
	"strings"
	"time"
)

// length- quantity of operations in expression
func GenerateExpression(length int) string {
	rand.Seed(time.Now().UnixNano())
	var builder = strings.Builder{}

	var generate = func(str string) string {
		return string(str[rand.Intn(len(str))])
	}
	var isSpace = func() string {
		if rand.Intn(2) == 0 {
			return ""
		}
		return " "
	}
	var edger = func(str string) {
		builder.WriteString(isSpace())
		builder.WriteString(str)
		builder.WriteString(isSpace())
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
