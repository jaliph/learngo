package cyclic

func cyclicSort(arr []int) {
	i := 0
	for i < len(arr) {
		j := arr[i] - 1
		if j < len(arr) && arr[j] != arr[i] {
			temp := arr[j]
			arr[j] = arr[i]
			arr[i] = temp
		} else {
			i++
		}
	}
}
