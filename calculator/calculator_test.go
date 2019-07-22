package calculator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenguyenit/testing/calculator"
)

func TestSum(t *testing.T) {
	if calculator.Sum(1, 3) != 4 {
		t.Error("Expected 1 + 3 to equal 4")
	}

	assert.Equal(t, 3, calculator.Sum(2))

}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculator.Sum(2)
	}
}

func TestSumWithDataTable(t *testing.T) {
	//Table testing
	var testData = []struct {
		input    int
		expected int
	}{
		{4, 4},
		{5, 4},
	}

	for _, test := range testData {
		actualResult := calculator.Sum(test.input)
		assert.Equal(t, actualResult, test.expected)
	}
}
