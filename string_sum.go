package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
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

func StringSum(input string) (output string, err error) {

	input = strings.TrimSpace(input)
	if len(input) == 0 {
		//Use when the input is empty, and input is considered empty if the string contains only whitespace
		return "", fmt.Errorf("inpute errore: %w", errorEmptyInput)
	}

	// Колличество слогаемых
	const amountNumber = 2

	// Массив слагаемых в виде string
	var numbers [amountNumber]string

	// Порядковый номер слагаемого
	addendum := 0

	// Перебираем руны и  заполняем numbers
	for _, itemRune := range input {
		itemString := string(itemRune)
		if itemString == "+" || itemString == "-" {
			if len(numbers[addendum]) == 0 {
				numbers[addendum] += itemString
			} else {
				addendum++
				if addendum == amountNumber {
					// Use when the expression has number of operands not equal to two
					return "", fmt.Errorf("new number errore: %w", errorNotTwoOperands)
				}
				numbers[addendum] += itemString
			}
		} else if itemString == " " {
			continue
		} else {
			_, err := strconv.Atoi(itemString)
			if err != nil {
				// For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace) the function should return an empty string and an appropriate error from strconv package wrapped into your own error with fmt.Errorf function
				return "", fmt.Errorf("contains characters, that are not numbers, +, - or whitespace: %w", err)
			} else {
				numbers[addendum] += itemString
			}
		}
	}

	if addendum < amountNumber-1 {
		// Use when the expression has number of operands not equal to two
		return "", fmt.Errorf("minimum numbers errore: %w", errorNotTwoOperands)
	}

	// сумма всех слогаемых
	var sum int
	for _, number := range numbers {
		intNumber, err := strconv.Atoi(number)
		if err != nil {
			return "", fmt.Errorf("error converting string to number: %w", err)
		}
		sum += intNumber
	}

	output = strconv.Itoa(sum)
	return output, nil
}
