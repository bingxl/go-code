package main

import "fmt"

type sayer interface {
	say(string)
}

type dog struct{}

func (d *dog) say(c string) {
	fmt.Println("dog say: ", c)
}

func interfaceTest() {
	// dog 使用指针实现了sayer接口， 所以只有 *dog 能赋值给接口
	var a sayer = &dog{}
	a.say("你好啊")
}
