package lab2
import (
	. "gopkg.in/check.v1"
	"testing"
	"errors"
	)
	
	func Test(t *testing.T) { TestingT(t) }
	
	type prefixToInfixSuite struct{}
	
	var _ = Suite(&prefixToInfixSuite{})
	
	func (s *prefixToInfixSuite) TestEmptyInput(c *C) {
	res, err := prefixToInfix("")
	c.Assert(res, Equals, "")
	c.Assert(err, DeepEquals, errors.New("invalid expression"))
	}
	
	func (s *prefixToInfixSuite) TestSingleNumber(c *C) {
	res, err := prefixToInfix("5")
	c.Assert(res, Equals, "5")
	c.Assert(err, IsNil)
	}
	
	func (s *prefixToInfixSuite) TestInvalidInput(c *C) {
	res, err := prefixToInfix("+ 2")
	c.Assert(res, Equals, "")
	c.Assert(err, DeepEquals, errors.New("invalid expression"))
	}
	
	func (s *prefixToInfixSuite) TestSimplePrefixExpression(c *C) {
	res, err := prefixToInfix("+ 2 3")
	c.Assert(res, Equals, "2 + 3")
	c.Assert(err, IsNil)
	}
	
	func (s *prefixToInfixSuite) TestNestedPrefixExpression(c *C) {
	res, err := prefixToInfix("+ * 2 3 4")
	c.Assert(res, Equals, "2 * 3 + 4")
	c.Assert(err, IsNil)
	}
	
	func (s *prefixToInfixSuite) TestComplexPrefixExpression(c *C) {
	res, err := prefixToInfix("+ * / 12 3 3 5")
	c.Assert(res, Equals, "12 / 3 * 3 + 5")
	c.Assert(err, IsNil)
	}
	
	func (s *prefixToInfixSuite) TestPowerOperator(c *C) {
	res, err := prefixToInfix("^ 2 3")
	c.Assert(res, Equals, "2 ^ 3")
	c.Assert(err, IsNil)
	}
	
	func (s *prefixToInfixSuite) TestParentheses(c *C) {
	res, err := prefixToInfix("+ * 2 3 4")
	c.Assert(res, Equals, "2 * 3 + 4")
	c.Assert(err, IsNil)
	
	res, err = prefixToInfix("+ * 2 3 - 4 5")
	c.Assert(res, Equals, "2 * 3 + (4 - 5)")
	c.Assert(err, IsNil)
	
	res, err = prefixToInfix("+ * 2 / 6 2 - 4 5")
	c.Assert(res, Equals, "2 * (6 / 2) + (4 - 5)")
	c.Assert(err, IsNil)

	res, err = prefixToInfix("+ * ^ 9 2 5 - 22 / 10 2")
	c.Assert(res, Equals, "9 ^ 2 * 5 + (22 - (10 / 2))")
	c.Assert(err, IsNil)

	res, err = prefixToInfix("- / * 5 8 2 + 10 2")
	c.Assert(res, Equals, "(5 * 8) / 2 - (10 + 2)")
	c.Assert(err, IsNil)

	res, err = prefixToInfix("+ - / * 6 2 12 2 * 3 2")
	c.Assert(res, Equals, "6 * 2 / 12 - 2 + (3 * 2)")
	c.Assert(err, IsNil)

	}
