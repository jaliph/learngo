package stacks

// https://leetcode.com/problems/minimum-deletions-to-make-string-balanced
func MinimumDeletions(s string) int {

	Min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	bCount := 0
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'b' {
			bCount++
		} else {
			// two choices
			// delete all the bs, or delete as, as counted in res
			res = Min(res+1, bCount)
		}
	}
	return res
}
