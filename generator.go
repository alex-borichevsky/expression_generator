package expression_generator

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// length- quantity of operations in expression
func Generate(length uint) string {
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
	var edger = func(str string) { fmt.Fprint(&builder, isSpace(), str, isSpace()) }

	var operand = func() {
		var operands = "0123456789"
		edger(generate(operands))
	}
	var operator = func() {
		var operators = "+-"
		edger(generate(operators))
	}
	operand()

	for i := 0; uint(i) < length; i++ {
		operator()
		operand()
	}
	return builder.String()
}
