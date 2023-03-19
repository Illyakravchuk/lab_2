package lab2

import (
	"bufio"
	"io"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	scanner := bufio.NewScanner(ch.Input)
	if scanner.Scan() {
		expression := scanner.Text()
		trimmed := strings.TrimSpace(expression)
		res, err := prefixToInfix(trimmed)
		if err != nil {
			return err
		}
		_, err = ch.Output.Write([]byte(res))
		if err != nil {
			return err
		}
	}
	return scanner.Err()
}