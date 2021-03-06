package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nekketsuuu/lambda2pi/lib/parser"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
L:
	for {
		fmt.Printf(">> ")
		if !scanner.Scan() {
			break L
		}
		l, err := parser.ParseExpr(scanner.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		} else {
			fmt.Printf("%v\n", l.String())
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Bad input: %v\n", err)
		os.Exit(1)
	}
}
