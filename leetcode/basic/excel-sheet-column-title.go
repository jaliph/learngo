// https://leetcode.com/problems/excel-sheet-column-title/

package basic

import "fmt"

func ConvertToTitle(columnNumber int) string {

	str := ""
	for columnNumber > 0 {
		res := (columnNumber - 1) % 26
		str = fmt.Sprintf("%c", res+65) + str
		columnNumber = (columnNumber - 1) / 26
	}
	return str
}
