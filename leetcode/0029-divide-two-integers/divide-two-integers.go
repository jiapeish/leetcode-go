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

	// begin non-recursively
	low, high := 0, abs(dividend)
	dividend, divisor = abs(dividend), abs(divisor)
	for low <= high {
		quotient := low + (high-low)/2
		if quotient*divisor <= dividend && dividend <= (quotient+1)*divisor {
			if dividend == (quotient+1)*divisor {
				output = quotient + 1
				break
			}
			output = quotient
			break
		}

		if quotient*divisor > dividend {
			high = quotient - 1
		}

		if (quotient+1)*divisor < dividend {
			low = quotient + 1
		}
	}
	// end non-recursively

	// begin recursively
	//output = binarySearchQuotient(0, abs(dividend), abs(dividend), abs(divisor))

	if output > math.MaxInt32 {
		return sign * math.MaxInt32
	}
	if output < math.MinInt32 {
		return sign * math.MinInt32
	}

	return sign * output
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
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