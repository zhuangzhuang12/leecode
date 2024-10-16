package day2

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	mu := &sync.Mutex{}
	cond := sync.NewCond(mu)

	current := 'a' // 当前要打印的数字

	printNum := func(id rune) {
		defer wg.Done()
		for i := 0; i < 2; i++ { // 每个协程打印两轮
			mu.Lock()
			for current != id {
				cond.Wait() // 等待条件满足
			}
			fmt.Print(id) // 打印当前协程的编号
			if current == 'a' {
				current = 'b'
			} else {
				current = 'a'
			}
			cond.Broadcast() // 通知其他协程
			mu.Unlock()
		}
	}

	wg.Add(2)

	// 启动四个协程
	go printNum('a')
	go printNum('b')

	// 等待所有协程完成
	wg.Wait()
}
