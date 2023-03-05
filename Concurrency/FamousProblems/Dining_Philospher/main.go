/*
this code is to demonstrait dining philosopher concurrency problem..
*/

// Package name definition
package main

import (
	"fmt"
	"sync"
	"time"
)

// Defining Philosopher struct

type Philosopher struct {
	SlNo                int    // Slno
	Name                string // Name of the Philosopher
	LeftFork, RightFork uint8  // left and right forks
}

// Constants to experiment with variaing number of Philosophers
const (
	CountPhilosopher int           = 5               // Philosopher count
	eatCount         int           = 3               // How many times philopher is eating
	eatTime          time.Duration = 1 * time.Second // time philosper takes to eat
	thinkTime        time.Duration = 1 * time.Second // time philospher thinking after eating
)

var (
	Philosophers    []Philosopher            = []Philosopher{} // List of philosphers
	PhilosoperNames [CountPhilosopher]string = [CountPhilosopher]string{"A",
		"B",
		"C",
		"D",
		"E"} // Names of philosophers
)

func init() {
	fmt.Println("Executing init function")
	//Initialising philospher list
	for count := 0; count < CountPhilosopher; count++ {
		var leftFork, rightFork uint8
		if count == CountPhilosopher-1 {
			leftFork = uint8(count)
			rightFork = 0

		} else {
			leftFork = uint8(count)
			rightFork = uint8(count + 1)
		}
		var philosopher Philosopher = Philosopher{SlNo: count + 1,
			Name:      PhilosoperNames[count],
			LeftFork:  leftFork,
			RightFork: rightFork}
		Philosophers = append(Philosophers, philosopher)
	}

}

func createLockMap(list []Philosopher) map[uint8]*sync.Mutex {
	// Creating lock map to simulate waiting for others
	fmt.Println("Creating lock map...")

	lockMap := make(map[uint8]*sync.Mutex)

	for idx, _ := range list {
		lockMap[uint8(idx)] = &sync.Mutex{}
	}
	return lockMap
}

func Dine(philosopher Philosopher, lockMap map[uint8]*sync.Mutex, waitGroup *sync.WaitGroup, seated *sync.WaitGroup) {

	defer waitGroup.Done()
	fmt.Println("Philosopher:", philosopher.Name, "\t is seating on the table...")
	seated.Done()
	seated.Wait()

	for cnt := 0; cnt < eatCount; cnt++ {
		fmt.Println(philosopher.Name, " is waiting for left fork")
		lockMap[philosopher.LeftFork].Lock()
		lockMap[philosopher.RightFork].Lock()
		fmt.Println("\t", philosopher.Name, " has started eating")
		time.Sleep(eatTime)
		fmt.Println("\t", philosopher.Name, " has completed eating ", cnt+1, " times")
		fmt.Println("\t", philosopher.Name, " is thinking ")
		time.Sleep(thinkTime)
		lockMap[philosopher.LeftFork].Unlock()
		lockMap[philosopher.RightFork].Unlock()

	}

	fmt.Println("Phiolospher:", philosopher.Name, " is done with eating")
	fmt.Println("Philospher:", philosopher.Name, " is leaving table..")
}

func main() {

	fmt.Println("Dining Philosophers!!!")
	fmt.Println("Philosopher sitting layour")
	wg := sync.WaitGroup{}
	seated := sync.WaitGroup{}

	wg.Add(len(Philosophers))

	seated.Add(len(Philosophers))

	for _, val := range Philosophers {
		fmt.Println("Sl No:", val.SlNo, "\tName:", val.Name, "\tLeftFork:", val.LeftFork, "\tRightFork:", val.RightFork)
	}

	lockMap := createLockMap(Philosophers)

	for _, val := range Philosophers {
		go Dine(val, lockMap, &wg, &seated)
	}
	wg.Wait()

}
