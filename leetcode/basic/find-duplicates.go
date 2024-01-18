package basic

import "fmt"

func FindDuplicate(arr []int) bool {
	n := len(arr)

	for i := 0; i < n; i++ {
		idx := arr[i] % n
		arr[idx] += n
	}

	fmt.Println(arr)
	for i := 0; i < n; i++ {
		if (arr[i] / n) >= 2 {
			return true
		}
	}

	return false
}
