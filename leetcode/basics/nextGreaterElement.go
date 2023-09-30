package basic

// https://leetcode.com/problems/next-greater-element-i/
func NextGreaterElement(nums1 []int, nums2 []int) []int {
	n1Map := make(map[int]int)

	for i, v := range nums1 {
		n1Map[v] = i
	}

	res := make([]int, len(nums1))
	q := []int{}
	for idx, v := range nums2 {

		// create monotonic decreasing stack
		for len(q) > 0 && nums2[q[len(q)-1]] < v {
			i := q[len(q)-1]
			q = q[:len(q)-1]

			// check if nums2[i] which is popped is present in nums1
			if _, ok := n1Map[nums2[i]]; ok {
				// fmt.Println(nums2[i], "-> ", v)
				res[n1Map[nums2[i]]] = v
			}
		}
		// push to monotonic decreasing stack
		q = append(q, idx)
		// fmt.Println(q, res)
	}

	for len(q) > 0 {
		i := q[len(q)-1]
		q = q[:len(q)-1]
		if _, ok := n1Map[nums2[i]]; ok {
			res[n1Map[nums2[i]]] = -1
		}
	}
	return res
}

func NextGreaterElementBrute(nums1 []int, nums2 []int) []int {

	n1Map := make(map[int]int)

	for i, v := range nums2 {
		n1Map[v] = i
	}

	res := make([]int, len(nums1))

	for i, v := range nums1 {
		idx := n1Map[v]
		for idx < len(nums2) {
			if nums2[idx] > v {
				break
			}
			idx++
		}

		if idx < len(nums2) {
			res[i] = nums2[idx]
		} else {
			res[i] = -1
		}
	}

	return res

}
