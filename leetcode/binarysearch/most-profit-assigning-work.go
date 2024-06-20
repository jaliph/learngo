package binarysearch

import "sort"

type Task struct {
	d int
	p int
}

func MaxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	tasks := []Task{}

	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := range difficulty {
		tasks = append(tasks, Task{d: difficulty[i], p: profit[i]})
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].d < tasks[j].d
	})

	maxProfit := 0
	for i := range tasks {
		maxProfit = Max(tasks[i].p, maxProfit)
		tasks[i].p = maxProfit
	}

	res := 0
	for _, w := range worker {
		wProfit := 0

		l := 0
		r := len(tasks) - 1

		for l <= r {
			mid := l + ((r - l) / 2)
			if tasks[mid].d <= w {
				wProfit = Max(wProfit, tasks[mid].p)
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		res += wProfit
	}

	return res

}
