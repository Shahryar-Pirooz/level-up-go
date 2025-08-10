package main

import (
	"flag"
	"log"
)

// isBalanced returns whether the given expression
// has balanced brackets.
var brackets = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

func isBalanced(expr string) bool {

	for k, v := range brackets {
		var (
			keyExist = false
			valExist = false
		)
		for _, x := range expr {
			if x == k {
				keyExist = true
			}
			if x == v {
				valExist = true
			}
		}
		if keyExist && !valExist {
			return false
		}
	}
	return true
}

// printResult prints whether the expression is balanced.
func printResult(expr string, balanced bool) {
	if balanced {
		log.Printf("%s is balanced.\n", expr)
		return
	}
	log.Printf("%s is not balanced.\n", expr)
}

func main() {
	expr := flag.String("expr", "", "The expression to validate brackets on.")
	flag.Parse()
	printResult(*expr, isBalanced(*expr))
}
