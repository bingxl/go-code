package main

import (
	"fmt"
	"os"

	"github.com/bingxl/gotest/leetcode"
)

func main() {

	index := leetcode.TwoSum([]int{2, 7, 11, 5}, 9)

	fmt.Println(index)

	fmt.Println(os.Args)

}
