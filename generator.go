package generator

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

	var any = func(str string) string {
		return string(str[rand.Intn(len(str))])
	}

	var space = func() string {
		if rand.Intn(2) == 0 {
			return ""
		}

		return " "
	}

	var operand = func() {
		fmt.Fprint(&builder, space(), any("0123456789"), space())
	}

	var operator = func() {
		fmt.Fprint(&builder, space(), any("+-"), space())
	}

	operand()

	for i := 0; uint(i) < length; i++ {
		operator()
		operand()
	}

	return builder.String()
}
