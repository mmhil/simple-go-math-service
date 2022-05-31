package domath

import (
	"fmt"
	"strconv"
)

func ReturnStringsOfAllMathFunctions(num1 int, num2 int) [6]string {
	num1str := strconv.Itoa(num1)
	num2str := strconv.Itoa(num2)
	var resultArray = [6]string{
		fmt.Sprintf("%q + %q = %q \n", num1str, num2str, strconv.Itoa(AddIntegers(num1, num2))),
		fmt.Sprintf("%q - %q = %q \n", num1str, num2str, strconv.Itoa(SubtractIntegers(num1, num2))),
		fmt.Sprintf("%q * %q = %q \n", num1str, num2str, strconv.Itoa(MultiplyInt(num1, num2))),
		fmt.Sprintf("%q / %q = %q \n", num1str, num2str, strconv.Itoa(DivideInt(num1, num2))),
		fmt.Sprintf("Remainder of %q / %q = %q \n", num1str, num2str, strconv.Itoa(GetRemainderInt(num1, num2))),
		fmt.Sprintf("(%q + %q) / %q = %q \n", num1str, num2str, num2str, strconv.Itoa(addThenDivideInt(num1, num2))),
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
