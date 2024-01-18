package basic

// https://leetcode.com/problems/find-the-original-array-of-prefix-xor

func findArray(pref []int) []int {
	prev := pref[0]
	for i := 1; i < len(pref); i++ {
		pref[i] = pref[i] ^ prev
		prev = prev ^ pref[i]
	}
	return pref
}
