package string_sum

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
	// Use when the expression has improper characters
	errorImproperChar = errors.New("expecting only digits and arithmetic operators, but received some extra")
	inputFailed       = "input failed: %w"
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

/*string_sum_test.go:36: letters in second operand:
wrong type of error is wrapped into the returned error: got *errors.errorString, want *strconv.NumError*/

func StringSum(input string) (output string, err error) {
	// space extracting
	//input = strings.ReplaceAll(input, " ", "")

	// empty input check
	if input == "" {
		return "", fmt.Errorf(inputFailed, errorEmptyInput)
	}
	// operands number check
	re, _ := regexp.Compile(`\d+`)
	operandSlice := re.FindAllString(input, -1)
	if len(operandSlice) != 2 {
		return "", fmt.Errorf(inputFailed, errorNotTwoOperands)
	}

	var numStack []int
	sign := '+'
	num := 0
	total := 0
	// bitwise string iteration
	for i := 0; i < len(input); i++ {
		chr := input[i]
		// improper sign check
		if !(chr >= '0' && chr <= '9') && !strings.Contains("+- ", string(chr)) {
			return "", fmt.Errorf(inputFailed, errorImproperChar)
		}
		// operand forming
		if chr >= '0' && chr <= '9' {
			num = num*10 + int(chr-'0')
		}
		// numStack filling
		if i+1 == len(input) || strings.Contains("+-", string(chr)) {
			switch sign {
			case '+':
				numStack = append(numStack, num)
			case '-':
				numStack = append(numStack, -num)
			}
			sign = int32(chr)
			num = 0
		}
	}
	// numStack sum
	for _, v := range numStack {
		total += v
	}
	// convert int to string
	output = strconv.Itoa(total)

	return output, nil
}
