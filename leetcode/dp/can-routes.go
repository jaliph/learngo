// https://leetcode.com/problems/count-all-possible-routes/
package dp

func countRoutes(locations []int, start int, finish int, fuel int) int {

	const MOD = 1e9 + 7
	abs := func(a, b int) int {
		if a-b < 0 {
			return b - a
		}
		return a - b
	}

	dp := map[[2]int]int{}

	var solve func(int, int) int
	solve = func(idx, fuel int) int {
		if fuel < 0 {
			return 0
		}

		if v, ok := dp[[2]int{idx, fuel}]; ok {
			return v
		}

		ans := 0
		if idx == finish {
			ans++
		}

		for j := 0; j < len(locations); j++ {
			if idx == j {
				continue
			}

			ans = ((ans % MOD) + solve(j, (fuel-abs(locations[idx], locations[j]))%MOD)) % MOD
		}

		dp[[2]int{idx, fuel}] = ans
		return ans
	}

	return solve(start, fuel)
}
