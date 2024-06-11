package basic

func GenerateCombination(str string) []string {
	res := []string{}
	arr := make([]string, len(str))
	for i, r := range str {
		arr[i] = string(r)
	}

	N := 1 << len(str)

	for i := 0; i < N; i++ {
		s := ""
		for j := 0; j < len(str); j++ {
			if (1<<j)&i != 0 {
				s += arr[j]
			}
		}
		res = append(res, s)
	}

	return res
}
