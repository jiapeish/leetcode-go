package _202_happy_number

func isHappy(n int) bool {
	slow, fast := n, compute(n)

	for slow != fast {
		slow = compute(slow)
		fast = compute(compute(fast))

	}

	if slow == 1 {
		return true
	}

	return false

}

func compute(n int) int {
	res := 0

	for n != 0 {
		res += (n % 10) * (n % 10)
		n /= 10
	}

	return res

}
