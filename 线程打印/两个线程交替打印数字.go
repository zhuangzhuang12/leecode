package 线程打印

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	co := sync.NewCond(&mu)
	count := 1
	go func(threadNum int) {
		for count <= 10 {
			mu.Lock()

			co.Wait()
			if count%2 == threadNum {
				fmt.Printf("thread:%d  print:%d", threadNum, count)
				count++
			}
			co.Signal()
			mu.Unlock()
		}
	}(1)

	go func(threadNum int) {
		for count <= 10 {
			mu.Lock()

			co.Wait()
			if count%2 == threadNum {
				fmt.Printf("thread:%d  print:%d", threadNum, count)
				count++
			}
			co.Signal()
			mu.Unlock()
		}
	}(0)
}
