package day1

import (
	"fmt"
	"sync"
)

type BlockQueue struct {
	capacity int
	queue    []interface{}
	lock     sync.Mutex
	notFull  sync.Cond
	notEmpty sync.Cond
}

func NewBlockQueue(capacity int) *BlockQueue {
	lock := sync.Mutex{}
	bq := &BlockQueue{
		capacity: capacity,
		queue:    make([]interface{}, 0, capacity),
		lock:     lock,
	}
	bq.notEmpty = sync.Cond{L: &bq.lock}
	bq.notFull = sync.Cond{L: &bq.lock}
	return bq
}
func (bq *BlockQueue) take() interface{} {
	bq.lock.Lock()
	defer bq.lock.Unlock()
	for len(bq.queue) == 0 {
		bq.notEmpty.Wait()
	}
	item := bq.queue[0]
	bq.queue = bq.queue[1:]
	bq.notFull.Signal()
	return item
}

func (bq *BlockQueue) put(item interface{}) {
	bq.lock.Lock()
	defer bq.lock.Unlock()
	for len(bq.queue) == bq.capacity {
		bq.notFull.Wait()
	}
	bq.queue = append(bq.queue, item)
	bq.notEmpty.Signal()
}
func main() {
	bq := NewBlockQueue(5)
	var wg sync.WaitGroup
	// 生产者
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 3; i++ {
			bq.put(i)
			fmt.Printf("Enqueued: %d\n", i)
		}
	}()

	// 消费者
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			item := bq.take()
			fmt.Printf("Dequeued: %v\n", item)
		}
	}()

	// 主线程等待，防止程序退出
	wg.Wait()
}
