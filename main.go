package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	states := GenerateStates()
	scanner := bufio.NewScanner(os.Stdin)
	var expressions []string

	for scanner.Scan() {
		expressions = append(expressions, scanner.Text())
	}

	fmt.Println("REGEXP result:")

	for _, expr := range expressions {
		analyzer := NewAnalyzer(expr, states)
		analyzer.Run()

		fmt.Printf("%s -> %s \n", expr, analyzer.matchedExp)
	}
}
