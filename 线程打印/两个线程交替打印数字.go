package 线程打印

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	co := sync.NewCond(&mu)
	count := 1

	group := sync.WaitGroup{}
	group.Add(2)
	go func(threadNum int) {
		defer group.Done()
		for count <= 10 {
			mu.Lock()

			if count%2 == threadNum {
				fmt.Printf("thread:%d  print:%d\n", threadNum, count)
				count++
				co.Signal()
			} else {
				co.Wait()
			}

			mu.Unlock()
		}
	}(1)

	go func(threadNum int) {
		defer group.Done()
		for count <= 10 {
			mu.Lock()

			if count%2 == threadNum {
				fmt.Printf("thread:%d  print:%d\n", threadNum, count)
				count++
				co.Signal()
			} else {
				co.Wait()
			}

			mu.Unlock()
		}
	}(0)
	co.Signal()
	group.Wait()
}
