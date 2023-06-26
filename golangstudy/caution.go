package main

import "fmt"

func print(n int) int {
	fmt.Println(n)
	return n
}

func caution() {

	// switch fallthrough
	// case 中多条件时从左向右执行，当有满足的case时 后续的条件不会执行， 如下列 print(1), print(2) 会执行，
	// print(2) 满足条件后 print(3） 不再执行

	// fallthrough 会直接跳到下一个case 的语句块执行，不会执行 case 的条件， 如下列的 print(4) 不会执行

	switch print(2) {
	case print(1), print(2), print(3):
		fmt.Println("case 1 执行")
		fallthrough

	case print(4):
		fmt.Println("case 2 执行")
	}

}
