package main

import(
"fmt"
"sync"
)

type SafeMap struct {
	sync.RWMutex
	internalMap map[int]int
}

func (m *SafeMap) read(key int) (int, bool){
	m.RLock()
	value, exists := m.internalMap[key]
	m.RUnlock()
	return value, exists
}

func (m *SafeMap) write(key, value int) {
	m.Lock()
	m.internalMap[key] = value
	m.Unlock()
}

var(
	memo = SafeMap{internalMap: make(map[int]int)}
)

func fibo(toFib int) int {
	if toFib <= 2 {
		return toFib
	}
	_, exists := memo.read(toFib)
	if !exists {
		memo.write(toFib, fibo(toFib-1)+fibo(toFib-2))
	}
	value, _ := memo.read(toFib)
	return value
}

func workerPool(jobs<-chan int, returnChan chan<-int){
	for job := range jobs{
		returnChan <- fibo(job)
	}
}

func main() {
	fibonacciThresh := 90 // value to compute the fibonacci sequence up to
	// AUDIENCE WORK:
	// create a buffered channel to use
	jobChannel := make(chan int, fibonacciThresh)
	returnChannel := make(chan int, fibonacciThresh)

	// spawn 3 worker pools and give them the correct job channel, return channel and memo map
	go workerPool(jobChannel, returnChannel)
	go workerPool(jobChannel, returnChannel)
	go workerPool(jobChannel, returnChannel)

	for i := 1; i <= fibonacciThresh; i++ {
		jobChannel <- i
	}
	close(jobChannel)
	fmt.Println("=================================================")
	fmt.Printf( "|      Fibonacci Sequence from 0 - %v           |\n", fibonacciThresh)
	fmt.Println("=================================================")

	for i := 0; i < fibonacciThresh; i++ {
		fmt.Printf("%v\n", <-returnChannel)
	}
}
