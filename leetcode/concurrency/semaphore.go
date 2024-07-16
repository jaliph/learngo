package concurrency

// Semaphore type that encapsulates the semaphore logic.
type Semaphore struct {
	ch chan struct{}
}

// NewSemaphore creates a new Semaphore with a given maximum count.
func NewSemaphore(maxCount int) *Semaphore {
	return &Semaphore{
		ch: make(chan struct{}, maxCount),
	}
}

// Acquire waits to obtain a permit from the semaphore.
func (s *Semaphore) Acquire() {
	s.ch <- struct{}{}
}

// Release returns a permit to the semaphore.
func (s *Semaphore) Release() {
	<-s.ch
}
