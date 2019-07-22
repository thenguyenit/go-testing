package calculator

//Sum to represent a sum function
func Sum(numbers ...int) (result int) {

	for _, number := range numbers {
		result += number
	}

	return result
}
