/*
The goal of this activity is to explore the use of threads by creating a program for sorting integers that uses four goroutines to create four sub-arrays and then merge the arrays into a single array

Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.

Students will receive five points if the program sorts the integers and five additional points if there are four goroutines that each prints out a set of array elements that it is storing.

16 15 14 13 12 11 10 9 8 7 6 5 4 3 2 1
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func sortSynchronous(wg *sync.WaitGroup, slice []int) {
	bubbleSort(slice)
	fmt.Println("Sorted by goroutine: ", slice)
	wg.Done()
}

// you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing.
func bubbleSort(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				swap(slice, j)
			}
		}
	}
}

// You should write a Swap() function which performs this operation.
// Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a position in the slice.
func swap(slice []int, j int) {
	slice[j], slice[j+1] = slice[j+1], slice[j]
}

func main() {

	//The program should prompt the user to input a series of integers.
	fmt.Printf("Enter a series of integers separated by space: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Split the input by commas and convert to integers
	strArray := strings.Split(input, " ")

	//Create array of ints to store the recently read values
	arrayOfInts := make([]int, len(strArray))

	for i, str := range strArray {
		num, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			fmt.Println("Invalid input. Please enter a series of integers separated by spaces.")
			return
		}
		arrayOfInts[i] = num
	}

	n := len(arrayOfInts)

	//Now divide the array in 4 parts
	partSize := n / 4

	partitions := make([][]int, 4)

	for i := 0; i < 4; i++ {
		start := i * partSize
		end := start + partSize
		if i == 3 {
			// Last partition takes any remainder
			end = n
		}
		partitions[i] = arrayOfInts[start:end]
	}

	var wg sync.WaitGroup

	wg.Add(4)

	go sortSynchronous(&wg, partitions[0])

	go sortSynchronous(&wg, partitions[1])

	go sortSynchronous(&wg, partitions[2])

	go sortSynchronous(&wg, partitions[3])

	wg.Wait()

	allPartitions := append(partitions[0], partitions[1]...)

	allPartitions = append(allPartitions, partitions[2]...)

	allPartitions = append(allPartitions, partitions[3]...)

	bubbleSort(allPartitions)

	fmt.Println("\nAll partitions sorted in the main using bubble:")
	fmt.Println(allPartitions)
}
