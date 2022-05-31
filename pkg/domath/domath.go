package domath

func DoAllIntMathFunctions(num1 int, num2 int) [6]int {
	var resultArray = [6]int{
		AddIntegers(num1, num2),
		SubtractIntegers(num1, num2),
		MultiplyInt(num1, num2),
		DivideInt(num1, num2),
		GetRemainderInt(num1, num2),
		addThenDivideInt(num1, num2),
	}
	return resultArray
}

func AddIntegers(num1 int, num2 int) int {
	return num1 + num2
}

func SubtractIntegers(num1 int, num2 int) int {
	return num1 - num2
}

func MultiplyInt(num1 int, num2 int) int {
	return num1 * num2
}

func DivideInt(num1 int, num2 int) int {
	return num1 / num2
}

func GetRemainderInt(num1 int, num2 int) int {
	return num1 % num2
}

func addThenDivideInt(num1 int, num2 int) int {
	return (num1 + num2) / num2
}
