package _166_fraction_to_recurring_decimal

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	if numerator * denominator < 0 {
		return "-" + fractionToDecimal(abs(numerator), abs(denominator))
	}

	numerator, denominator = abs(numerator), abs(denominator)

	if numerator > denominator {
		decimals := fractionToDecimal(numerator % denominator, denominator)
		return strconv.Itoa(numerator / denominator) + decimals[1:]
	}

	digits := make([]byte, 2, 1024)
	digits[0] = '0'
	digits[1] = '.'

	index := 2
	record := make(map[int]int, 1024)

	for {
		if i, ok := record[numerator]; ok {
			return fmt.Sprintf("%s(%s)", string(digits[:i]), string(digits[i:]))
		}

		record[numerator] = index
		numerator *= 10
		index++

		digits = append(digits, byte(numerator/denominator) + '0')
		numerator %= denominator

		if numerator == 0 {
			return string(digits)
		}
	}

}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}