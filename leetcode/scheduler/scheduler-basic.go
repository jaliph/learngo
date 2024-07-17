package scheduler

import "time"

type Job interface {
	Execute()
}

type SchedulerBasic interface {
	Scheduler(job Job, time time.Duration)
}

type SumJob struct {
	a   int64
	b   int64
	sum int64
}

func (p *SumJob) Execute() {
	p.sum = p.a + p.b
}

func NewSumJob(a, b int64) *SumJob {
	return &SumJob{
		a: a,
		b: b,
	}
}

type delayedSchedulerBasic struct {
}

func NewScheduler() *delayedSchedulerBasic {
	return &delayedSchedulerBasic{}
}

func (d *delayedSchedulerBasic) Schedule(job Job, duration time.Duration) {
	go func() {
		time.Sleep(duration)
		job.Execute()
	}()
}
