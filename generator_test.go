package generator

import (
	"strings"
	"testing"
	"testing/quick"

	"github.com/borichevskiy/expression_calculator"
)

func TestGenerate(t *testing.T) {
	fn := func(length uint8) bool {
		t.Log("length:", length)
		expr := Generate(uint(length))
		t.Log("expr:", expr)
		var operands, operators uint

		for _, r := range expr {
			if strings.ContainsRune("0123456789", r) {
				operands++

				continue
			}

			if strings.ContainsRune("+-", r) {
				operators++

				continue
			}
		}

		t.Log("ops:", operators, ";digits:", operands)

		return (operators == uint(length)) && (operands == uint(length)+1)
	}

	if err := quick.Check(fn, &quick.Config{MaxCount: 10}); err != nil {
		t.Errorf("count operators and operands: %s", err)
	}

	fn2 := func(n uint8) bool {
		_, err := expression_calculator.Evaluate(Generate(uint(n)))

		return err == nil
	}

	if err := quick.Check(fn2, &quick.Config{MaxCount: 10}); err != nil {
		t.Errorf("check by calculator: %s", err)
	}
}
