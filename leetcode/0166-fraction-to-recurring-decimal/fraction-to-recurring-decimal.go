package _166_fraction_to_recurring_decimal

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	if numerator * denominator < 0 {
		return "-" + fractionToDecimal(abs(numerator), abs(denominator))
	}

	numerator, denominator = abs(numerator), abs(denominator)






}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}