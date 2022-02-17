package main

import (
	"sync"
)

const N = 10

var wg = &sync.WaitGroup{}

func main() {
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(i int) {
			// wg.Add(1) //将wg.Add(1)语句放置在main()函数体中才行，否则可能会出现main函数已经执行完毕，而匿名函数还没有执行
			println(i)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
}

/**
3
9
6
0
4
7
8
5
2
1

*/
