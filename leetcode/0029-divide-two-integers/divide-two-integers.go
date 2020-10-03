package _029_divide_two_integers

import "math"

func divide(dividend int, divisor int) int {
	sign, output := -1, 0
	if divisor == 1 {
		return dividend
	}
	if dividend > 0 && divisor > 0 || dividend < 0 && divisor < 0 {
		sign = 1
	}

	output = binarySearchQuotient(0, abs(dividend), abs(dividend), abs(divisor))

	if output > math.MaxInt32 {
		return sign * math.MaxInt32
	}
	if output < math.MinInt32 {
		return sign * math.MinInt32
	}

	return sign * output
}

func binarySearchQuotient(low, high, dividend, divisor int) int {
	quotient := low + (high-low)/2
	if quotient*divisor <= dividend && dividend <= (quotient+1)*divisor {
		if dividend == (quotient+1)*divisor {
			return quotient + 1
		}
		return quotient
	}

	if quotient*divisor > dividend {
		return binarySearchQuotient(low, quotient-1, dividend, divisor)
	}

	if (quotient+1)*divisor < dividend {
		return binarySearchQuotient(quotient+1, high, dividend, divisor)
	}
	return 0
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
