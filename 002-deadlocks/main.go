package main

import (
	"fmt"
	"sync"
	"time"
)

type data struct {
	mu    sync.Mutex
	value int
}

func main() {
	var wg sync.WaitGroup
	var data1, data2 data
	data1.value = 10
	data2.value = 20
	wg.Add(2)
	go printSum(&data1, &data2, &wg)
	go printSum(&data2, &data1, &wg) // notice the order of parameters!
	wg.Wait()
}

func printSum(d1, d2 *data, wg *sync.WaitGroup) {
	defer wg.Done()
	d1.mu.Lock()
	defer d1.mu.Unlock()

	time.Sleep(2 * time.Second)
	d2.mu.Lock()
	defer d2.mu.Unlock()

	fmt.Println(d1.value, d2.value)
}

/*
Here, a deadlock occurs because the first goroutrine (go printSum())
locks d1 and sleeps, during which time the second goroutine (go printSum())
locks d2 and waits for d1 to release, resulting in a circular dependency.
*/

/*
Turns out, deadlocks are well studied in computer science and occur when one
or more of the so-called Coffman Conditions exist:
- Mutual Exclusion: A concurrent process holds exclusive rights to a resource at any one time.
- Wait For Condition: A concurrent process must simultaneously hold a resource and be waiting for an additional resource.
- No Preemption: A resource held by a concurrent process can only be released by that process (this is seen in our example).
- Circular Wait: A concurrent process (P1) must be waiting on a chain of other concurrent processes (P2), which are in turn waiting on it (P1) (also present in our example).
*/
