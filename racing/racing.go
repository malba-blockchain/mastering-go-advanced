/*
MY PROGRAM
The race condition occurs becase in my program I'm sending transactions which have a delay of 2 seconds to arrive.

In this case I sent 10 batches of transactions to 3 different countries and it was common to find some transactions of batch "i" getting settled when there was a batch of transations of i+1 being initiated.


WHAT THE RACE CONDITION IS
Here multiple goroutines are accessing and modifying the shared variables of the countries canada, switzerland, and USA without proper synchronization. This leads to unpredictable and inconsistent results.


HOW IT CAN OCCUR
Multiple goroutines accessing shared variables: The transaction function is called concurrently for different countries, and each call accesses and modifies the corresponding shared variable.

No synchronization: There's no mechanism to ensure that only one goroutine can modify a shared variable at a time. This allows multiple goroutines to potentially modify the same variable simultaneously.

Data races: When multiple goroutines try to modify the same variable at the same time, it can lead to data races, where the final value of the variable depends on the interleaving of the goroutines' execution.
*/

package main

import (
	"fmt"
	"time"
)

func main() {

	var canada = 0
	var switzerland = 0
	var USA = 0

	for i := 0; i < 10; i++ {

		fmt.Println("\nCurrent batch: ", i)
		go transaction("Canada ", canada+i)
		go transaction("Switzerland ", switzerland+i)
		transaction("USA ", USA+i)
	}
}

func transaction(name string, countryCounter int) {
	fmt.Print("\n>Transaction initiated: ", name, countryCounter)
	time.Sleep(2 * time.Second)
	fmt.Print("\n@Transaction settled: ", name, countryCounter)
}
