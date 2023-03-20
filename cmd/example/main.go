package main

import (
	"flag"
	"fmt"
	"io"
	"os"
    "strings"
	"github.com/Illyakravchuk/lab_2"
)

func main() {
	var (
		inputExpression = flag.String("e", "", "Expression to compute")
		inputFromFile   = flag.String("f", "", "Take input from file")
		outputToFile    = flag.String("o", "", "Save output to file")
	)

	flag.Parse()

	var input io.Reader
	var output io.Writer = os.Stdout

	if *inputExpression != "" {
		input = strings.NewReader(*inputExpression)
	}

	if *inputFromFile != "" {
		f, err := os.Open(*inputFromFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening file:", err)
			os.Exit(1)
		}
		defer f.Close()
		input = f
	}

	if *outputToFile != "" {
		file, err := os.Create(*outputToFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error creating file:", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
	}

	if input == nil {
		fmt.Fprintln(os.Stderr, "No input provided")
		os.Exit(1)
	}

	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error computing expression:", err)
		os.Exit(1)
	}
}