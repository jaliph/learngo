package maths

func MajorityElement2(nums []int) []int {
	res := []int{}

	ct1 := 0
	c1 := 0

	ct2 := 0
	c2 := 0

	for _, v := range nums {
		// c1's count should not be contributed by c2
		if ct1 == 0 && v != c2 {
			ct1 = 1
			c1 = v
		} else if ct2 == 0 && v != c1 {
			// c2's count should not be contributed by c1
			ct2 = 1
			c2 = v
		} else if v == c1 {
			ct1++
		} else if v == c2 {
			ct2++
		} else {
			ct1--
			ct2--
		}
	}

	threshold := len(nums) / 3
	ct1 = 0
	ct2 = 0

	for _, v := range nums {
		if v == c1 {
			ct1++
		} else if v == c2 {
			ct2++
		}
	}

	if ct1 > threshold {
		res = append(res, c1)
	}

	if ct2 > threshold {
		res = append(res, c2)
	}

	return res
}

func MajorityElement2Basic(nums []int) []int {
	res := []int{}

	ct1 := 0
	c1 := 0

	ct2 := 0
	c2 := 0

	for _, v := range nums {
		if v == c1 {
			ct1++
		} else if v == c2 {
			ct2++
		} else if ct1 == 0 {
			ct1 = 1
			c1 = v
		} else if ct2 == 0 {
			ct2 = 1
			c2 = v
		} else {
			ct1--
			ct2--
		}
	}

	threshold := len(nums) / 3
	ct1 = 0
	ct2 = 0

	for _, v := range nums {
		if v == c1 {
			ct1++
		} else if v == c2 {
			ct2++
		}
	}

	if ct1 > threshold {
		res = append(res, c1)
	}

	if ct2 > threshold {
		res = append(res, c2)
	}

	return res
}
