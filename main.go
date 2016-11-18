package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nekketsuuu/lambda2pi/lib/convert"
)

func usage() {
	fmt.Fprintf(os.Stderr, "[Description of %s]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, `This program converts a lambda term into a pi term.
Precise syntax are described under the description of options.

[Options]
`)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, `
[Syntax]
The syntax of the lambda calculus:

    M ::= x      // variables
        | \x. x  // lambda abstraction
        | M M    // application
        | (M)    // parenthesis

Note that a period is required in each lambda abstraction. The period means the longest match rule of the abstraction. Therefore, (\x. y \z. w) means (\x. (y (\z. w))), not ((\x. y) (\z. w)).

The syntax of the pi calculus:

    M ::= O           // the empty process
                      // (Note that the symbol is not zero,
                      //  but the 15th latin alphabet, O.)
        | x           // names
        | x?x.M       // input
        | x!x.M       // output
        | M|M         // parallel execution
        | *M          // replication (*M is congruent to M|*M)
        | new x in M  // name restriction
        | (M)         // parenthesis

Note that there is no summation.
`)
}

var (
	showVersion    bool
	inputFile      string
	outputFile     string
	evalModeString string
)

func setFlags() {
	flag.Usage = usage
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.StringVar(&outputFile, "o", "out.pi", "an output filename")
	flag.StringVar(&evalModeString, "mode", "CallByValue", "an evaluation strategy for the lambda calculus (value: CallByValue, CallByName)")
	flag.Parse()
	inputFile = flag.Arg(0)
}

func main() {
	setFlags()
	if err := Run(); err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] lambda2pi: %v\n", err)
	}
}

func Run() error {
	if showVersion {
		fmt.Println("lambda2pi", version)
	} else {
		mode, err := convert.ToMode(evalModeString)
		if err != nil {
			return err
		}

		if inputFile != "" {
			return fileMode(mode)
		} else {
			return replMode(mode)
		}
	}
	return nil
}

func fileMode(mode convert.EvalMode) error {
	pi, err := convert.ConvertFromFile(inputFile, mode)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(outputFile, []byte(pi.String()), 0664)
	if err != nil {
		return err
	}
	return nil
}

func replMode(mode convert.EvalMode) error {
	fmt.Println("lambda2pi", version, "REPL mode")
	scanner := bufio.NewScanner(os.Stdin)
REPL:
	for {
		fmt.Printf(">> ")
		if !scanner.Scan() {
			break REPL
		}

		p, err := convert.ConvertFromString(scanner.Text(), mode)
		if err != nil {
			fmt.Fprintln(os.Stderr, "[ERROR]", err)
		} else {
			fmt.Println(p.String())
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
