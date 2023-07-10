package main

import "fmt"

func main() {
	for i := 1; i <= 12; i++ {
		fmt.Println(Season(i))
	}
}

func Season(month int) string {
	var season string
	switch month {
	case 1, 2, 12:
		season = "Winter"
	case 3, 4, 5:
		season = "Spring"
	case 6, 7, 8:
		season = "Summer"
	case 9, 10, 11:
		season = "Autumn"
	default:
		season = "Season unknown"
	}
	return season
}
