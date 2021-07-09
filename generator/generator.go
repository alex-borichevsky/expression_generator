package generator
import (
	"math/rand"
	"strings"
	"time"
)

func GenerateExpression(length int) string{
	rand.Seed(time.Now().UnixNano())
	builder := strings.Builder{}
	operators := "+-"
	operands := "0123456789"

	generate := func(str string) string{
		return string(str[rand.Intn(len(str))])
	}
	space := func() string{
		length := 2
		if rand.Intn(length) == 0 {return  ""}
		return " "
	}
	handler := func(str string){
		builder.WriteString(space())
		builder.WriteString(str)
		builder.WriteString(space())
	}
	operand := func(){
		handler(generate(operands))
	}
	operator := func(){
		handler(generate(operators))
	}
	operand()

	for i:= 0; i < length; i++{
		operator()
		operand()
	}
	return builder.String()
}