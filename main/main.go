package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)

	go func() {
		ch1 <- 10
		close(ch1)
	}()

	go func() {
		ch2 <- 20
		close(ch2)
	}()

	for ch1 != nil || ch2 != nil {
		select {
		case num, ok := <-ch1:
			if ok {
				fmt.Println("Received from ch1:", num)
			} else {
				ch1 = nil // 当ch1关闭时，将ch1设为nil，避免再尝试从中接收
			}
		case num, ok := <-ch2:
			if ok {
				fmt.Println("Received from ch2:", num)
			} else {
				ch2 = nil // 当ch2关闭时，将ch2设为nil，避免再尝试从中接收
			}

		}
	}

}
