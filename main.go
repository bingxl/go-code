package main

import (
	"fmt"
	"os"

	"github.com/bingxl/gotest/cli"
	"github.com/bingxl/gotest/leetcode"
)

func main() {

	index := leetcode.TwoSum([]int{2, 7, 11, 5}, 9)

	fmt.Println(index)

	fmt.Println(os.Args)
	fmt.Println(cli.ColorStr("fgpurple,fggreen", "你好阿"))
	fmt.Println(cli.ColorStr("hightlight", "你好阿"))
	// fmt.Printf("%c[0;41;36m%s%c[0m\n", 0x1B, "testPrintColor", 0x1B)

	var commandParams = []cli.FlagT{
		{Name: "test", DefVal: "test", Usage: "this is a test to command line"},
	}
	re := cli.Parse(commandParams)
	fmt.Println(re)

}
