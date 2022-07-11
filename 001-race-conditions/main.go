package main

import "fmt"

/**
A race condition occurs when two or more operations must execute in
the correct order, but the program has not been written so that this
order is guaranteed to be maintained.
*/

func main() {
	var counter int

	go func() {
		counter++
	}()

	if counter == 0 {
		fmt.Println("Counter is: ", counter)
	}
}

/**
Depending on the order in which the three parts of this code execute (`counter++`, `if` and `fmt.Println`), the printed value can be 0, 1 or even nothing.
*/
