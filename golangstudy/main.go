package main

import "fmt"

func main() {
	// run("structTest", structTest)
	// interfaceTest()

	// run("caution", caution)

	run("routine", routine)

}

func run(name string, f func()) {
	fmt.Printf("================ %s start ========================\n", name)
	f()
	fmt.Printf("================ %s end ========================\n\n", name)
}
