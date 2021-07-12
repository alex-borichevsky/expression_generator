package expression_generator

import (
	"github.com/borichevskiy/expression_calculator"
	"testing"
	"testing/quick"
)

func TestGenerate(t *testing.T) {
	config := &quick.Config{}
	config.MaxCount = 10

	equ := func(length uint) bool {
		length = 5 //random number
		operators := map[rune]int{'1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}
		operands := map[rune]string{'+': "+", '-': "-"}
		expr := Generate(length)
		operandsNumber := 0
		operatorsNumber := 0
		for _, val := range expr {
			if _, ok := operators[val]; ok {
				operatorsNumber++
			}
			if _, ok := operands[val]; ok {
				operandsNumber++
			}
			continue
		}
		return operandsNumber+1 == operatorsNumber
	}

	if err := quick.Check(equ, config); err != nil {
		t.Errorf("generate error : %s", err)
	}
	var n uint = 20 // random number
	if _, err := expression_calculator.Evaluate(Generate(n)); err != nil {
		t.Errorf("failed Eval(); %s", err)
	}
}
