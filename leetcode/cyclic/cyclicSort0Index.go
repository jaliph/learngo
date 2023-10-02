package cyclic

func Cyclic0Index(arr []int) []int {
	i := 0
	for i < len(arr) {
		j := arr[i]
		if arr[i] < len(arr) && arr[i] != arr[j] {
			arr[i], arr[j] = arr[j], arr[i]
		} else {
			i++
		}
	}
	return arr
}
