package channels

import "fmt"

func DoParallel_1(n int) int {

	ch := make(chan int)

	go func(c chan<- int, n int) {
		sum := 0
		for i := 0; i <= n; i++ {
			fmt.Println("im running 1")
			sum += i
		}
		ch <- sum
	}(ch, n)

	for i := 0; i < n; i++ {
		fmt.Println("im running 2")
	}

	res := <-ch
	fmt.Println("Sum is ", res)
	return res
}
