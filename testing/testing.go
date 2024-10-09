package main

import (
	"fmt"
	"sync"
)

const (
	maxPhilosophers = 2
	maxEats         = 3
)

type ChopStick struct {
	sync.Mutex
}

type Philosopher struct {
	id       int
	left     *ChopStick
	right    *ChopStick
	noOfEats int
}

func (p *Philosopher) eat(ch chan *Philosopher, wg *sync.WaitGroup) {
	for p.noOfEats > 0 {
		ch <- p
		p.left.Lock()
		p.right.Lock()

		fmt.Printf("starting to eat %d\n", p.id)
		fmt.Printf("finishing eating %d\n", p.id)
		p.noOfEats--

		p.left.Unlock()
		p.right.Unlock()
		wg.Done()
	}
}

type Host struct {
	ch chan *Philosopher
}

func (h *Host) givePermission() {

	for {
		<-h.ch
	}
}

func main() {
	chopSticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		chopSticks[i] = new(ChopStick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			id:       i + 1,
			left:     chopSticks[i],
			right:    chopSticks[(i+1)%5],
			noOfEats: maxEats,
		}
	}

	host := Host{
		ch: make(chan *Philosopher, maxPhilosophers),
	}

	var wg sync.WaitGroup
	wg.Add(maxEats * len(philosophers))
	go host.givePermission()

	for _, p := range philosophers {
		go p.eat(host.ch, &wg)
	}

	wg.Wait()
	close(host.ch)
}
