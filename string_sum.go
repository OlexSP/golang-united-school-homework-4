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

func StringSum(input string) (output string, err error) {

	// empty input check
	if input == "" {
		return "", fmt.Errorf(inputFailed, errorEmptyInput)
	}
	// operands number check (just need it for the test)
	re, _ := regexp.Compile(`\d+`)
	operandSlice := re.FindAllString(input, -1)
	if len(operandSlice) != 2 {
		return "", fmt.Errorf(inputFailed, errorNotTwoOperands)
	}

	var numStack []int
	sign := '+'
	num := 0
	total := 0
	//  string iteration
	for i := 0; i < len(input); i++ {
		chr := input[i]
		// improper sign check. the test wants an error from strconv.Itoa() :D
		if !(chr >= '0' && chr <= '9') && !strings.Contains("+- ", string(chr)) {
			_, err = strconv.Atoi(input)
			return "", err //fmt.Errorf("input failed: %w", errorImproperChar) було/стало ))
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
