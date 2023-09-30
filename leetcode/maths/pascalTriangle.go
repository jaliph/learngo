package maths

func Generate(numRows int) [][]int {
	pascals := [][]int{}

	prev := []int{1}
	pascals = append(pascals, prev)
	for i := 0; i < numRows; i++ {
		row := []int{prev[0]}

		for j := 1; j < len(prev); j++ {
			row = append(row, (prev[j-1] + prev[j]))
		}

		row = append(row, prev[0])
		prev = row

		pascals = append(pascals, row)
	}

	return pascals
}
