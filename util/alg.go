package util

import (
	"fmt"
	"log"

	"github.com/apaxa-go/eval"
)

// ShuntingYard -> https://en.wikipedia.org/wiki/Shunting-yard_algorithm
func ShuntingYard(expression []string, precedence map[string]int) int {

	values := []string{}
	operators := []string{}

	for _, t := range expression {
		if IsInt(t) {
			values = append(values, t)
		} else if t == "(" {
			operators = append(operators, t)
		} else if t == ")" {
			topOfStack := peek(operators)
			for {
				if topOfStack == "(" || topOfStack == "" {
					break
				}
				operators, values = applyOperator(operators, values)
				topOfStack = peek(operators)
			}
			operators = operators[:len(operators)-1]
		} else { // t is an operator (Wikipedia also lists "function", we ain't gonna bother with that.)
			topOfStack := peek(operators)
			for {
				if topOfStack == "" || topOfStack == "(" || topOfStack == ")" || precedence[t] > precedence[topOfStack] {
					break
				}
				operators, values = applyOperator(operators, values)
				topOfStack = peek(operators)
			}
			operators = append(operators, t)
		}
	}
	for {
		if len(operators) == 0 {
			break
		}
		operators, values = applyOperator(operators, values)
	}
	return ToInt(values[0])
}

func peek(stack []string) string {
	if len(stack) > 0 {
		return stack[len(stack)-1]
	}
	return ""
}

func applyOperator(operatorStack, outputQueue []string) ([]string, []string) {
	op, operatorStack := operatorStack[len(operatorStack)-1], operatorStack[:len(operatorStack)-1]
	left, outputQueue := outputQueue[len(outputQueue)-1], outputQueue[:len(outputQueue)-1]
	right, outputQueue := outputQueue[len(outputQueue)-1], outputQueue[:len(outputQueue)-1]

	src := fmt.Sprintf("%v%v%v", left, op, right)

	expr, err := eval.ParseString(src, "")
	if err != nil {
		log.Fatal("Error:", err)
	}

	r, err := expr.EvalToInterface(nil)
	if err != nil {
		log.Fatal("Error:", err)
	}

	outputQueue = append(outputQueue, fmt.Sprintf("%v", r))
	return operatorStack, outputQueue
}
