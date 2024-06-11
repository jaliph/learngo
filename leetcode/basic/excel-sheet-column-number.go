// https://leetcode.com/problems/excel-sheet-column-number/

package basic

func TitleToNumber(columnTitle string) int {
	res := 0

	for _, v := range columnTitle {
		num := v - 64
		res = (res * 26) + int(num)
	}
	return res
}
