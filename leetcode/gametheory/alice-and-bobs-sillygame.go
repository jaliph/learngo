// https://www.hackerrank.com/challenges/alice-and-bobs-silly-game/problem?isFullScreen=true

package gametheory

import "fmt"

var primes []bool
var pArray []int

const SIZE = 10 * 10 * 10 * 10 * 10

func SillyGame(n int32) string {
	// Write your code here
	findPrimes()
	findPreFix()
	ct := countPrimes(int(n))
	fmt.Println(ct)

	if ct == 0 || ct&1 == 0 {
		return "Bob"
	}
	return "Alice"
}

func countPrimes(r int) int {
	return pArray[r]
}

func findPreFix() {
	pArray = make([]int, SIZE+1)
	for i := 1; i <= SIZE; i++ {
		if !primes[i] {
			pArray[i] = pArray[i-1] + 1
		} else {
			pArray[i] = pArray[i-1]
		}
	}
}

func findPrimes() {
	primes = make([]bool, SIZE+1)
	primes[0] = true
	primes[1] = true
	for i := 2; i <= SIZE; i++ {
		for j := i * i; j <= SIZE; j += i {
			if !primes[j] {
				primes[j] = true
			}
		}
	}
}
