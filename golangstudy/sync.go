package main

import (
	"fmt"
	"sync"
)

func syncTest() {
	wg := sync.WaitGroup{}
	// wg.Go 功能与wg.Add wg.Done相同，但是wg.Go 中的函数不能painc
	wg.Go(func() {
		fmt.Println("hello")
	})
	wg.Wait()
}
