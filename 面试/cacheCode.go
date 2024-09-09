package 面试

import (
	"fmt"
	"sync"
)

//业务中有日志要写数据库，一个个写有性能损耗，写个缓存组件使产生的数据每100条或每1s能批量入库

type Cache struct {
	capacity int
	data     []interface{}
	sngle    chan bool
	lock     sync.Mutex
	wait     []interface{}
}

func (c *Cache) add(staement interface{}) {
	lock := c.lock.TryLock()
	if !lock {
		c.wait = append(c.wait, staement)
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Print("r:", r)
		}
		c.lock.Unlock()
	}()
	c.data = append(c.data, staement)
	if len(c.data) >= c.capacity {
		c.sngle <- true
	}
}

func (c *Cache) init(capacity int) {
	c.capacity = capacity
}
func main() {
	c := &Cache{capacity: 100}
	for {
		select {
		case <-c.sngle:
			for _, datum := range c.data {
				//执行datum数据库语句
				fmt.Println(datum)
			}
			c.lock.Lock()
			c.data = []interface{}{}
			c.lock.Unlock()
		}
	}

}
