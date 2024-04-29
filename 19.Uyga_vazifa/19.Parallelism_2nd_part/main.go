package main

import (
	"fmt"
	"sync"
)

func f1(out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- 1
}
	
func f2(out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- 2
}

func f3(out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- 3
}

func f4(out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- 4
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(4)

	go f1(ch, &wg)
	go f2(ch, &wg)
	go f3(ch, &wg)
	go f4(ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	var result []int
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			select {
			case v, ok := <-ch:
				if !ok {
					return
				}
				result = append(result, v)
			}
		}
	}()

	<-done
	fmt.Println("Natijalar:", result)
}
