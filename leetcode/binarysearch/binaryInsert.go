package binarysearch

import "fmt"

// decreasing order
func BinaryInsert(target int, arr *[]int) {
	l, r := 0, len(*arr)-1

	var mid int
	for l <= r {
		mid = l + (r-l)/2
		if target == (*arr)[mid] {
			l = mid
			break
		} else if target > (*arr)[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	fmt.Println(*arr, l, target)
	if l == len(*arr) {
		(*arr) = append((*arr), target)
	} else {
		(*arr) = append((*arr)[:l+1], (*arr)[l:]...)
		(*arr)[l] = target
	}
}

func Driver() {
	sorted := &[]int{}
	arr := []int{3, 4, 6, 8}

	for _, v := range arr {
		BinaryInsert(v, sorted)
	}
	fmt.Println(sorted)
}
