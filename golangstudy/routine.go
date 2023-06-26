package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 使用通道同步 goroutine
func channelSync(w *sync.WaitGroup) {
	defer w.Done()

	running := func(c chan int) {
		var times = 0
		for {
			times++
			fmt.Println("running function tick:", times)
			// time.Sleep(time.Second)
			c <- times
			if times == 20 {
				c <- 0
				break
			}
		}
	}

	// 使用通道进行通信和同步
	c := make(chan int, 10)
	go running(c)

	for v := range c {
		fmt.Println("main 接收到 c 通道数据：", v)
		if v == 0 {
			break
		}
	}
}

// 使用waitGroup 来等待协程结束
func waitGroupSync(w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("waitGroupSync start")
	var wg = sync.WaitGroup{}

	task := func(n int) {
		defer wg.Done()
		fmt.Println("run task: ", n)
		time.Sleep(2 * time.Second)
		fmt.Println("end task: ", n)
	}

	wg.Add(2)
	go task(1)
	go task(2)
	wg.Wait()

	fmt.Println("waitGroupSync end")
}

// 互斥锁操作共享数据
func mutexSync(w *sync.WaitGroup) {
	defer w.Done()
	var noLockNumber = 100
	var lockNumber = 100

	var mutex sync.Mutex
	var wg sync.WaitGroup

	taskWithNoLock := func(n int) {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			v := noLockNumber
			// 协程调度
			runtime.Gosched()
			noLockNumber = v + i

		}

	}
	taskWithLock := func(n int) {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			mutex.Lock()
			v := lockNumber
			// 协程调度
			runtime.Gosched()
			lockNumber = v + i
			mutex.Unlock()
		}

	}

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go taskWithNoLock(i)
		go taskWithLock(i)
	}

	wg.Wait()
	// task 函数中不加锁时得到的 noLockNumber 不确定（多次运行数值不同）
	// noLockNumber 未加锁，得到的结果是不确定的， lockNumber 加了锁，结果是确定的
	fmt.Printf("finally noLockNumber: %d, localNumber: %d ", noLockNumber, lockNumber)
}

func routine() {
	// 多 cpu 执行
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := &sync.WaitGroup{}

	// wg.Add(1)
	// go channelSync(wg)

	// wg.Add(1)
	// go waitGroupSync(wg)

	wg.Add(1)
	go mutexSync(wg)

	// 阻塞等待协程结束
	wg.Wait()
}
