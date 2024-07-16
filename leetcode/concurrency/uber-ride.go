package concurrency

import (
	"fmt"
	"sync"
	"time"
)

type UberRide struct {
	lock            sync.Mutex
	cond            *sync.Cond
	democratCount   int
	republicanCount int
	seatedCount     int
}

func NewUberRide() *UberRide {
	ride := &UberRide{}
	ride.cond = sync.NewCond(&ride.lock)
	return ride
}

func (ride *UberRide) seatDemocrat() {
	ride.lock.Lock()
	defer ride.lock.Unlock()

	ride.democratCount++
	ride.seatedCount++

	for !ride.isRideReady() {
		ride.cond.Wait()
	}

	ride.cond.Broadcast()
	ride.seated()
}

func (ride *UberRide) seatRepublican() {
	ride.lock.Lock()
	defer ride.lock.Unlock()

	ride.republicanCount++
	ride.seatedCount++

	for !ride.isRideReady() {
		ride.cond.Wait()
	}

	ride.cond.Broadcast()
	ride.seated()
}

func (ride *UberRide) isRideReady() bool {
	if ride.seatedCount < 4 {
		return false
	}

	if ride.democratCount == 4 || ride.republicanCount == 4 {
		return true
	}

	if ride.democratCount == 2 && ride.republicanCount == 2 {
		return true
	}

	return false
}

func (ride *UberRide) seated() {
	fmt.Println("A passenger has been seated.")

	if ride.seatedCount == 4 {
		ride.drive()
	}
}

func (ride *UberRide) drive() {
	fmt.Println("All passengers are seated. Driving now!")
	ride.reset()
}

func (ride *UberRide) reset() {
	ride.democratCount = 0
	ride.republicanCount = 0
	ride.seatedCount = 0
}

func UberDriver() {
	ride := NewUberRide()

	// Simulate Democrats and Republicans arriving
	for i := 0; i < 6; i++ {
		go ride.seatDemocrat()
	}

	for i := 0; i < 6; i++ {
		go ride.seatRepublican()
	}

	// Wait for goroutines to complete
	time.Sleep(5 * time.Second)
}
