package basic

func IsHappy(n int) bool {
	slow := n
	fast := n

	for {
		slow = process(slow)
		fast = process(process(fast))

		if slow == 1 {
			return true
		}

		if fast == slow {
			return false
		}
	}
}

func process(n int) int {
	p := 0

	for n > 0 {
		p += (n % 10) * (n % 10)
		n = n / 10
	}

	return p
}
