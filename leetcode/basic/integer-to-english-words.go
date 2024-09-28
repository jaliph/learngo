package basic

import (
	"strings"
)

func NumberToWords(num int) string {

	map10 := map[int]string{
		0:  "Zero",
		1:  "One",
		2:  "Two",
		3:  "Three",
		4:  "Four",
		5:  "Five",
		6:  "Six",
		7:  "Seven",
		8:  "Eight",
		9:  "Nine",
		10: "Ten",
		11: "Eleven",
		12: "Twelve",
		13: "Thirteen",
		14: "Fourteen",
		15: "Fifteen",
		16: "Sixteen",
		17: "Seventeen",
		18: "Eighteen",
		19: "Nineteen",
	}

	map100 := map[int]string{
		2: "Twenty",
		3: "Thirty",
		4: "Forty",
		5: "Fifty",
		6: "Sixty",
		7: "Seventy",
		8: "Eighty",
		9: "Ninety",
	}

	// var getHundred func(n int) string

	getHundred := func(n int) string {
		var sb strings.Builder

		if n >= 100 {
			num := n / 100
			sb.WriteString(" " + map10[num] + " Hundred")
			n = n % 100
		}

		if n >= 20 {
			num := n / 10
			sb.WriteString(" " + map100[num])
			n = n % 10
		}

		if n > 0 {
			sb.WriteString(" " + map10[n])
		}

		// fmt.Println(sb.String())
		return strings.Trim(sb.String(), " ")
	}

	var final strings.Builder

	if num >= 1e9 {
		n := num / 1e9
		final.WriteString(" " + getHundred(n) + " Billion")
		num = num % 1e9
	}

	if num >= 1e6 {
		n := num / 1e6
		final.WriteString(" " + getHundred(n) + " Million")
		num = num % 1e6
	}

	if num >= 1e3 {
		n := num / 1e3
		final.WriteString(" " + getHundred(n) + " Thousand")
		num = num % 1e3
	}

	final.WriteString(" " + getHundred(num))

	return strings.Trim(final.String(), " ")
}
