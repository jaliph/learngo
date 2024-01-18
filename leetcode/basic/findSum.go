package basic

func GetSum(t []int) int {
	sum := 0
	for _, val := range t {
		sum += val
	}
	return sum
}
