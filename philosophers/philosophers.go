/*
	I made sure to check all criteria. Look at the comments [REQ # 1]... to [REQ # 8]
*/

package main

import (
	"fmt"
	"sync"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
}

func (p Philo) eat(wg *sync.WaitGroup, i int, permissionChan chan bool) {

	defer wg.Done()

	//[REQ # 2] Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
	for i := 0; i < 3; i++ {

		//[REQ #4] In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
		permissionChan <- true // This blocks if there are already 2 philosophers eating

		p.leftCS.Lock()
		p.rightCS.Lock()
		//[REQ #7] When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
		fmt.Println("Starting to eat: ", i)

		//[REQ #8] When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
		fmt.Println("finishing eating: ", i)

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		// Release permission after eating
		<-permissionChan // This frees up a spot for another philosopher
	}
}

func main() {

	CSticks := make([]*ChopS, 5)

	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	//[REQ #1] There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
	philos := make([]*Philo, 5)

	//[REQ #3] The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
	for i := 0; i < 5; i++ {
		//This modification alternates the order of chopstick assignment for odd and even-numbered philosophers
		if i%2 == 0 {
			philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]}
		} else {
			philos[i] = &Philo{CSticks[(i+1)%5], CSticks[i]}
		}
	}

	var wg sync.WaitGroup

	wg.Add(5)

	// [REQ #5] The host allows no more than 2 philosophers to eat concurrently.
	permissionChan := make(chan bool, 2)

	for i := 0; i < 5; i++ {
		// [REQ #6]  Each philosopher is numbered, 1 through 5.
		go philos[i].eat(&wg, i, permissionChan)
	}

	wg.Wait()
}
