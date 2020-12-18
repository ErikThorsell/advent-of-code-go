package util

import (
	"fmt"
	"log"

	"github.com/apaxa-go/eval"
)

// ShuntingYard -> https://en.wikipedia.org/wiki/Shunting-yard_algorithm
func ShuntingYard(expression []string, precedence map[string]int) int {

	outputQueue := []string{}
	operatorStack := []string{}

	for _, t := range expression {
		if IsInt(t) {
			outputQueue = append(outputQueue, t)
		} else if t == "(" {
			operatorStack = append(operatorStack, t)
		} else if t == ")" {
			topOfStack := peek(operatorStack)
			for {
				if topOfStack == "(" || topOfStack == "" {
					break
				}
				operatorStack, outputQueue = applyOperator(operatorStack, outputQueue)
				topOfStack = peek(operatorStack)
			}
			operatorStack = operatorStack[:len(operatorStack)-1]
		} else { // t is an operator (Wikipedia also lists "function", we ain't gonna bother with that.)
			topOfStack := peek(operatorStack)
			for {
				if topOfStack == "" || topOfStack == "(" || topOfStack == ")" || precedence[t] > precedence[topOfStack] {
					break
				}
				operatorStack, outputQueue = applyOperator(operatorStack, outputQueue)
				topOfStack = peek(operatorStack)
			}
			operatorStack = append(operatorStack, t)
		}
	}
	for {
		if len(operatorStack) == 0 {
			break
		}
		operatorStack, outputQueue = applyOperator(operatorStack, outputQueue)
	}
	return ToInt(outputQueue[0])
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
