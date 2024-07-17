package scheduler_test

import (
	"fmt"
	"scheduler"
	"testing"
	"time"
)

type MulJob struct {
	x int
	y int
	z int
}

func (m *MulJob) Execute() {
	fmt.Println("I'm triggered")
	m.z = m.x * m.y
	fmt.Println("I'm Done", m)
}

func BenchmarkNewDelayedScheduler(b *testing.B) {
	delayedScheduler := scheduler.NewDelayedScheduler()

	m1 := MulJob{x: 1, y: 2}
	m2 := MulJob{x: 1, y: 2}
	m3 := MulJob{x: 1, y: 2}

	for i := 0; i < b.N; i++ {
		delayedScheduler.Schedule(&m1, time.Second*1)
		delayedScheduler.Schedule(&m2, time.Second*2)
		delayedScheduler.Schedule(&m3, time.Second*5)
	}

	// go func() {
	// 	time.Sleep(time.Second * 10)
	// 	delayedScheduler.Stop()
	// }()
	// <-delayedScheduler.Done
}
