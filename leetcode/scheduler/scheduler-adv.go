package scheduler

import (
	"container/heap"
	"fmt"
	"sync"
	"time"
)

type DelayedJob struct {
	Job      Job
	Priority int64
	Index    int
}

type JobQueue []*DelayedJob

func (jq JobQueue) Len() int {
	return len(jq)
}

func (jq JobQueue) Less(i, j int) bool {
	return jq[i].Priority < jq[j].Priority
}

func (jq JobQueue) Swap(i, j int) {
	jq[i], jq[j] = jq[j], jq[i]
	jq[i].Index = i
	jq[j].Index = j
}

func (jq *JobQueue) Push(x any) {
	delayedJob := x.(*DelayedJob)
	delayedJob.Index = len(*jq) + 1
	*jq = append(*jq, x.(*DelayedJob))
}

func (jq *JobQueue) Pop() any {
	old := *jq
	x := (old)[len(old)-1]
	(old)[len(old)-1] = nil // avoid memory leak
	*jq = old[:len(old)-1]
	x.Index = -1
	return x
}

type DelayedScheduler struct {
	Ticker   *time.Ticker
	Mu       sync.Mutex
	JobQueue JobQueue
	Done     chan bool
}

func NewDelayedScheduler() *DelayedScheduler {
	delayedScheduler := &DelayedScheduler{
		Ticker:   time.NewTicker(1 * time.Second),
		Mu:       sync.Mutex{},
		JobQueue: make(JobQueue, 0),
		Done:     make(chan bool),
	}
	heap.Init(&delayedScheduler.JobQueue)
	go delayedScheduler.Start()
	return delayedScheduler
}

func (ds *DelayedScheduler) Start() {
	for {
		select {
		case <-ds.Done:
			ds.Ticker.Stop()
			return
		case t := <-ds.Ticker.C:
			fmt.Println("Ticked at", t)
			ds.Mu.Lock()
			if len(ds.JobQueue) > 0 {
				job := heap.Pop(&ds.JobQueue).(*DelayedJob)

				now := time.Now().Unix()
				fmt.Println(job, job.Priority, now)
				if job.Priority-now > 0 {
					heap.Push(&ds.JobQueue, job)
				} else {
					go job.Job.Execute()
				}
			}
			ds.Mu.Unlock()
		}
	}
}

func (ds *DelayedScheduler) Stop() {
	ds.Done <- true
}

func (ds *DelayedScheduler) Schedule(job Job, duration time.Duration) {
	ds.Mu.Lock()
	now := time.Now().Unix()
	delayedJob := &DelayedJob{
		Job:      job,
		Priority: now + int64(duration.Seconds()),
	}
	heap.Push(&ds.JobQueue, delayedJob)
	ds.Mu.Unlock()
}
