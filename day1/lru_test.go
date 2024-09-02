package day1

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_constructor(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want *LRUCache
	}{
		{
			name: "test1",
			args: args{
				capacity: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructor(tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				lruCache := constructor(4)
				lruCache.Put(1, 1)
				lruCache.Put(2, 2)
				lruCache.Put(3, 3)
				lruCache.Put(4, 4)
				lruCache.Put(5, 5)

				lruCache.Get(1)
				for lruCache.head.next != nil {
					fmt.Printf("%d ", lruCache.head.key)
					lruCache.head = lruCache.head.next
				}
			}
		})
	}
}
