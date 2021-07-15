package expression_generator

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// length- quantity of operations in expression
func Generate(length uint) string {
	const (
		operands  = "0123456789"
		operators = "+-"
	)
	rand.Seed(time.Now().UnixNano())
	var builder = strings.Builder{}

	var any = func(str string) string {
		return string(str[rand.Intn(len(str))])
	}
	var space = func() string {
		if rand.Intn(2) == 0 {
			return ""
		}
		return " "
	}
	var edger = func(str string) { fmt.Fprint(&builder, space(), str, space()) }

	var operand = func() {
		edger(any(operands))
	}
	var operator = func() {
		edger(any(operators))
	}
	operand()

	for i := 0; uint(i) < length; i++ {
		operator()
		operand()
	}
	return builder.String()
}
