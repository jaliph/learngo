package dp

import (
	"fmt"
	"sort"
)

func jobScheduling(startTime []int, endTime []int, profit []int) int {

	jobs := func() [][]int {
		res := [][]int{}
		for i := range startTime {
			res = append(res, []int{startTime[i], endTime[i], profit[i]})
		}
		return res
	}()

	jobLen := len(jobs)

	Max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][0] < jobs[j][0]
	})

	var solve func(int) int
	dp := map[int]int{}
	solve = func(idx int) int {
		if idx == jobLen {
			return 0
		}

		if v, ok := dp[idx]; ok {
			return v
		}

		nextIdx := idx + 1
		for nextIdx < jobLen && jobs[nextIdx][0] < jobs[idx][1] {
			nextIdx++
		}

		res := Max(solve(idx+1), jobs[idx][2]+solve(nextIdx))
		dp[idx] = res
		return res
	}

	return solve(0)
}

func Driver() {
	startTime := []int{1, 2, 3, 4, 6}
	endTime := []int{3, 5, 10, 6, 9}
	profit := []int{20, 20, 100, 70, 60}

	fmt.Println(jobScheduling(startTime, endTime, profit))
}
