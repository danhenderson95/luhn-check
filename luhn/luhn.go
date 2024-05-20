package luhn

import "strconv"

func LuhnCheck(input string) bool {
	digits := splitStringIntoDigits(input)

	sum := 0
	isSecond := false

	for i := len(digits) - 1; i >= 0; i-- {
		digit := digits[i]

		if isSecond {
			digit *= 2

			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		isSecond = !isSecond
	}

	return sum%10 == 0
}

func splitStringIntoDigits(str string) []int {
	digits := make([]int, len(str))

	for i, char := range str {
		digit, _ := strconv.Atoi(string(char))
		digits[i] = digit
	}

	return digits
}
