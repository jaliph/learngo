package concurrency

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func Boring(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			ch <- fmt.Sprintf("Hello, %v", msg)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return ch
}

func FanIn(ch1, ch2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			ch <- <-ch1
		}
	}()
	go func() {
		for {
			ch <- <-ch2
		}
	}()
	return ch
}

func FanInSelect(ch1, ch2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case msg := <-ch1:
				ch <- msg
			case msg := <-ch2:
				ch <- msg
			case <-time.After(time.Duration(1 * time.Nanosecond)):
				ch <- "too slow"
			default:
				ch <- "noone is ready!!"
			}
		}
	}()
	return ch
}

func MainDriver() {
	ch := FanInSelect(Boring("ann"), Boring("lisa"))

	for i := 0; i < 20; i++ {
		fmt.Println("you say, ", <-ch)
	}
	fmt.Println("I'm done now")
}
