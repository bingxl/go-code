package main

import (
	"fmt"
	"time"
)

func runing() {
	var times int
	for {
		times++
		fmt.Println("tick", times)
		time.Sleep(time.Second)
	}
}

func routinetest() {
	go runing()

	var input string
	fmt.Scanln(&input)
}
