package eval

import (
	"testing"
)

func TestEvaluatePositive(t *testing.T) {
	tests := [...]struct {
		expression, want string
	}{
		{"1 + 2", "3"},
		{"10  + 10", "20"},
		{" 10  +   10 ", "20"},
		{"1 - 10", "-9"},
		{"10 * 10", "100"},
		{"10 / 3", "3"},
		{"VI / III", "II"},
		{"VI * VI", "XXXVI"},
		{"VI - I", "V"},
		{"III + II", "V"},
		{"  I + I ", "II"},
	}

	for _, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			result, _ := Evaluate(tt.expression)
			if result.String() != tt.want {
				t.Errorf("got %s, want %s", result, tt.want)
			}
		})
	}
}

func TestEvaluateNegative(t *testing.T) {
	tests := [...]struct {
		expression, want string
	}{
		{"2 + 2\n", "invalid multi-line expression"},
		{"2 ^ 2", "invalid operator"},
		{"V hhhhhhh V", "invalid operator"},
		{"11 + 1", "first number is too big"},
		{"10 + 11", "second number is too big"},
		{"11 + 11", "first number is too big"},
		{"XX + I", "first number is too big"},
		{"II + XX", "second number is too big"},
		{"0 + 3", "first number is too small"},
		{"4 + -5", "second number is too small"},
		{"I - II", "invalid expression, Roman numbers cannot be negative"},
		{"I - V", "invalid expression, Roman numbers cannot be negative"},
		{"2 +", "invalid expression, too few arguments"},
		{"2", "invalid expression, too few arguments"},
		{"2 + 2 + 2", "invalid expression, too many arguments"},
		{"II + 2 + V", "invalid expression, too many arguments"},
		{"II + 2", "don't mix Roman and Arabic notation"},
		{"4 + V", "don't mix Roman and Arabic notation"},
		{"3.0 + I", "invalid operand"},
		{"3.0 + 2", "invalid operand"},
		{"10 + 2.1", "invalid operand"},
		{"h + 2", "invalid operand"},
	}
	for _, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			_, err := Evaluate(tt.expression)
			if err.Error() != tt.want {
				t.Errorf("got %s, want %s", err, tt.want)
			}
		})
	}
}
