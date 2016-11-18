package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nekketsuuu/lambda2pi/lib/convert"
)

// The version number will be set as the git version tag during a build
var version string = "[version unknown]"

var (
	showVersion    bool
	inputFile      string
	outputFile     string
	evalModeString string
)

func setFlags() {
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
