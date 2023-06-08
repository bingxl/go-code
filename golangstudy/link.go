package main

import (
	"container/list"
	"fmt"
)

/**
 *golang 双向列表的操作
 */

func linktest() {
	// 初始化列表，var l list.List 也能初始化
	// var l = list.New(), 此方式获得的 l 是指针类型
	var l list.List

	// 从表尾添加元素
	e4 := l.PushBack(4)
	// l.PushBackList(es *List), 将es的节点添加到列尾
	// 从表头添加元素
	e1 := l.PushFront(1)
	// l.PushFrontList(es *List)

	// 在元素之后插入元素
	e2 := l.InsertAfter(2, e1)
	// 在元素之前插入元素
	e3 := l.InsertBefore(3, e4)

	// 获取元素的前驱或后继节点
	e := e2.Prev()
	e = e3.Next()

	// 获取列表首尾节点
	l.Back()
	l.Front()

	// 获取列表节点数
	l.Len()
	l.MoveBefore(e, e2) // 将e移动到e2之前
	printlink(l)
	l.MoveAfter(e, e3)
	l.MoveToFront(e)
	l.MoveToBack(e)

	printlink(l)

	l.Remove(e)
	printlink(l)

	// 清空或初始化列表
	l.Init()
	printlink(l)
}

func printlink(l list.List) {

	if l.Len() == 0 {
		fmt.Println("列表暂无节点")
		return
	}
	index := 0
	// 遍历列表
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d=>%v\t", index, e.Value)
	}
	fmt.Printf("\n")
}
