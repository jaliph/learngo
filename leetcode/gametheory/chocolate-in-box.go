// https://www.hackerrank.com/challenges/chocolate-in-box/problem?isFullScreen=true&h_r=next-challenge&h_v=zen
package gametheory

func ChocolateInBox(arr []int32) int32 {
	var nim_sum int32 = 0

	for _, v := range arr {
		nim_sum ^= v
	}

	var res int32 = 0
	for _, v := range arr {
		if v^nim_sum < v {
			res++
		}
	}

	return res
}
