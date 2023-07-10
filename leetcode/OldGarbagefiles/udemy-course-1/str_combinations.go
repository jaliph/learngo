package main

import "fmt"

func main() {
	fmt.Println(combinations([]rune{'a', 'b', 'c', 'd'}))
	fmt.Println(combinations2([]rune{'a', 'b', 'c', 'd'}))
	fmt.Println(combinations3([]rune{'a', 'b', 'c', 'd'}))
}

func combinations(arr []rune) [][]string {
	var subset [][]string
	subset = append(subset, []string{})

	for _, v := range arr {
		length := len(subset)
		for j := 0; j < length; j++ {
			clone := make([]string, len(subset[j]))
			copy(clone, subset[j])
			clone = append(clone, string(v))
			subset = append(subset, clone)
		}
	}
	return subset
}

func combinations2(arr []rune) []string {
	var combinations_recur func(s string, o string, result *[]string)
	combinations_recur = func(s string, o string, result *[]string) {
		// fmt.Println(s, o, result)
		if len(s) == 0 {
			*result = append(*result, o)
			return
		}
		newString := s[1:]

		combinations_recur(newString, o+string(s[0]), result)
		combinations_recur(newString, o, result)
	}
	res := []string{}
	combinations_recur("abcd", "", &res)

	return res
}

func combinations3(arr []rune) [][]string {
	var results [][]string
	var combinations_recur func(arr []rune, i int, result *[]string)
	combinations_recur = func(arr []rune, i int, result *[]string) {

		if i == len(arr) {
			results = append(results, *result)
			return
		}
		result2 := make([]string, len(*result))
		copy(result2, *result)
		result2 = append(result2, string(arr[i]))
		combinations_recur(arr, i+1, result)
		combinations_recur(arr, i+1, &result2)
	}
	res := []string{}
	combinations_recur(arr, 0, &res)

	return results
}
