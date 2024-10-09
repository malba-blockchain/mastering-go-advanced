package main

import (
	"fmt"
	"os/exec"
)

func main() {

	err := rock() // Assuming executeRock() returns an error if rock exits with status 2
	if err != nil {
		if exitStatus, ok := err.(*exec.ExitError); ok && exitStatus.ExitCode() == 2 {
			// Handle exit status 2 here, e.g., log a message
			for i := 0; 1 == 1; i++ {
				go fmt.Println("New go routine", i)
				i = i * i
			}
			fmt.Println("Main go routine")
		} else {
			// Handle other errors
			for i := 0; 1 == 1; i++ {
				go fmt.Println("New go routine", i)
				i = i * i
			}
			fmt.Println("Main go routine")
		}
	}

}

func rock() error {
	for i := 0; 1 == 1; i++ {
		go fmt.Println("New go routine", i)
		i = i * i
	}
	fmt.Println("Main go routine")

	return nil
}
