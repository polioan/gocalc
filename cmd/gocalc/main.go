package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/polioan/gocalc/internal/eval"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter expression: ")
	scanner.Scan()
	expression := scanner.Text()
	result, err := eval.Evaluate(expression)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println("Result: ", result)
	}
}
