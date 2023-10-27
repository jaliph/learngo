package channels

import "fmt"

func GenFibs() {
	// channel1 := updatePosition("Akash ::")
	// channel2 := updatePosition("Prashant ::")
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-channel1)
	// 	fmt.Println(<-channel2)
	// }
	ch := fibonacci()
	for i := 0; i < 15; i++ {
		fmt.Println(i, " has ", <-ch)
	}
}

// func updatePosition(s string) chan string {
// 	ch := make(chan string)

// 	go func() {
// 		for i := 0; ; i++ {
// 			time.Sleep(time.Second * 1) // some work
// 			ch <- s + " - " + strconv.Itoa(i)
// 		}
// 	}()

// 	return ch
// }

func fibonacci() chan int {
	ch := make(chan int)

	go func() {
		for i, j := 0, 1; ; {
			ch <- i
			i, j = j, i+j
		}
	}()

	return ch
}
