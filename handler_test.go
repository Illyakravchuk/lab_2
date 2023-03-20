package lab2

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"gopkg.in/check.v1"
)

type ComputeHandlerSuite struct{}

var _ = check.Suite(&ComputeHandlerSuite{})

func (s *ComputeHandlerSuite) TestCompute(c *check.C) {
	testCases := []struct {
		Input   string
		Output  string
		Err     error
	}{
		{
			Input:  "+ 2 3",
			Output: "2 + 3",
			Err:    nil,
		},
		{
			Input:  "* - 6 3 2",
			Output: "(6 - 3) * 2",
			Err:    nil,
		},
		{
			Input:  "invalid input",
			Output: "",
			Err:    errors.New("invalid expression"),
		},
	}

	for _, tc := range testCases {
		input := strings.NewReader(tc.Input)
		output := &bytes.Buffer{}
		ch := &ComputeHandler{
			Input:  input,
			Output: output,
		}

		err := ch.Compute()

		c.Check(output.String(), check.Equals, tc.Output)
		c.Check(err, check.DeepEquals, tc.Err)
	}
}

func Test_comp(t *testing.T) { check.TestingT(t) }