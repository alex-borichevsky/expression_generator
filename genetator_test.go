package expression_generator

import (
	"github.com/borichevskiy/expression_calculator"
	"math/rand"
	"testing"
	"time"
)

func TestGenerateExpression(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 3; i++ {
		var length = rand.Intn(10) // rand length of expression
		_, err := expression_calculator.Evaluate(GenerateExpression(length))
		if err != nil {
			t.Error("error in expression generation")
		}
	}

}
