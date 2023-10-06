// https://www.hackerrank.com/challenges/tower-breakers-revisited-1/problem
package gametheory

var primeDivisorCnt []int

const SIZE_PRIMES = 1e6 + 1

func PopulatePrimeDivisorCount() {
	primeDivisorCnt = make([]int, SIZE_PRIMES+1)

	for i := 2; i <= SIZE_PRIMES; i++ {
		if primeDivisorCnt[i] == 0 {
			primeDivisorCnt[i] = 1
			for j := i * 2; j <= SIZE_PRIMES; j = j + i {
				jj := j
				for jj%i == 0 {
					jj /= i
					primeDivisorCnt[j]++
				}
			}
		}
	}
}

func TowerBreakersRevist(arr []int32) int32 {
	nim_sum := 0
	for _, v := range arr {
		nim_sum ^= primeDivisorCnt[v]
	}
	if nim_sum > 0 {
		return 1
	} else {
		return 2
	}
}
