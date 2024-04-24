package main

import (
	"fmt"
	"strconv"
	"sync"
)

type SafeCounter struct {
	mu     sync.Mutex
	NumMap map[string]int
}

func (s *SafeCounter) Add(num int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.NumMap[strconv.Itoa(num)] = num
}

func (s *SafeCounter) Read(num int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	fmt.Printf("`%v` : %v\n", num, s.NumMap[strconv.Itoa(num)])
}

func (s *SafeCounter) Remove(num int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.NumMap, strconv.Itoa(num))
}

func main() {
	s := SafeCounter{NumMap: map[string]int{}}
	var wg sync.WaitGroup
	var n int
	fmt.Printf("`map`ga nechta son yozmoqchisiz: ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.Add(i)
		}(i)
	}
	wg.Wait()

	fmt.Printf("\n\n`map`ga %v ta son yozildi.\n", n)
	fmt.Printf("`map`ga yozilgan sonlar:\n\n")

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.Read(i)
		}(i)
	}
	wg.Wait()

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.Remove(i)
		}(i)
	}
	wg.Wait()

	fmt.Printf("\n'map'dagi barcha sonlar o'chirildi. Bo'sh map: ")

	fmt.Println(s.NumMap)
}
