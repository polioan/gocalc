package eval

import (
	"errors"
	"strconv"
	"strings"

	"github.com/polioan/gocalc/internal/roman"
)

type notationKind uint8

const (
	arabicNotation notationKind = iota
	romanNotation
)

type evaluationResult struct {
	value int
	kind  notationKind
}

func (e evaluationResult) String() string {
	if e.kind == arabicNotation {
		return strconv.Itoa(e.value)
	}
	if e.kind == romanNotation {
		result, err := roman.FromArabic(e.value)
		if err == nil {
			return result
		}
	}
	return "" // unreachable
}

func tokenize(expression string) []string {
	tokens := strings.Split(expression, " ")
	var filteredTokens []string
	for _, token := range tokens {
		trimmed := strings.TrimSpace(token)
		if trimmed != "" {
			filteredTokens = append(filteredTokens, trimmed)
		}
	}
	return filteredTokens
}

func parseOperand(a string) (int, notationKind, error) {
	var (
		num int
		err error
	)

	num, err = strconv.Atoi(a)
	if err == nil {
		return num, arabicNotation, nil
	}

	num, err = roman.ToArabic(a)
	if err == nil {
		return num, romanNotation, nil
	}

	return 0, arabicNotation, errors.New("invalid operand")
}

func parseOperands(a, b string) (int, int, notationKind, error) {
	aAsNumber, aKind, aErr := parseOperand(a)
	if aErr != nil {
		return 0, 0, arabicNotation, aErr
	}
	bAsNumber, bKind, bErr := parseOperand(b)
	if bErr != nil {
		return 0, 0, arabicNotation, bErr
	}
	if aKind != bKind {
		return 0, 0, arabicNotation, errors.New("don't mix Roman and Arabic notation")
	}
	return aAsNumber, bAsNumber, aKind, nil
}

func Evaluate(expression string) (evaluationResult, error) {
	result := evaluationResult{}

	if strings.Contains(expression, "\n") {
		return result, errors.New("invalid multi-line expression")
	}

	tokens := tokenize(expression)

	switch tokenCount := len(tokens); {
	case tokenCount < 3:
		return result, errors.New("invalid expression, too few arguments")
	case tokenCount > 3:
		return result, errors.New("invalid expression, too many arguments")
	}

	a := tokens[0]
	b := tokens[2]

	aAsNumber, bAsNumber, kind, err := parseOperands(a, b)

	if err != nil {
		return result, err
	}

	result.kind = kind

	if aAsNumber <= 0 {
		return result, errors.New("first number is too small")
	}
	if bAsNumber <= 0 {
		return result, errors.New("second number is too small")
	}
	if aAsNumber > 10 {
		return result, errors.New("first number is too big")
	}
	if bAsNumber > 10 {
		return result, errors.New("second number is too big")
	}

	switch op := tokens[1]; op {
	case "+":
		result.value = aAsNumber + bAsNumber
	case "-":
		temp := aAsNumber - bAsNumber
		if result.kind == romanNotation && temp < 0 {
			return result, errors.New("invalid expression, Roman numbers cannot be negative")
		}
		result.value = temp
	case "*":
		result.value = aAsNumber * bAsNumber
	case "/":
		result.value = aAsNumber / bAsNumber
	default:
		return result, errors.New("invalid operator")
	}

	return result, nil
}
