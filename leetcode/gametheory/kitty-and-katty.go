package gametheory

// https://www.hackerrank.com/challenges/kitty-and-katty/

// whoever has the last chance, decides who wins

func KittyKatty(n int) string {
	if n == 1 {
		return "Kitty"
	} else {
		if n&1 == 1 {
			return "Katty"
		} else {
			return "Kitty"
		}
	}
}
