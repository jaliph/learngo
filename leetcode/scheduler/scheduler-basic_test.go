package scheduler_test

import (
	"scheduler"
	"testing"
	"time"
)

func BenchmarkScheduler(b *testing.B) {
	delayedScheduler := scheduler.NewScheduler()

	sumJob1 := scheduler.NewSumJob(2, 3)
	sumJob2 := scheduler.NewSumJob(3, 4)
	sumJob3 := scheduler.NewSumJob(4, 5)

	for i := 0; i < b.N; i++ {
		delayedScheduler.Schedule(sumJob1, 5*time.Second)
		delayedScheduler.Schedule(sumJob2, 1*time.Second)
		delayedScheduler.Schedule(sumJob3, 3*time.Second)
	}

}
