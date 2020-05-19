package main

import "plugin"

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	p, err := plugin.Open("./test/test.so")
	panicErr(err)
	v, err := p.Lookup("V")
	panicErr(err)
	f, err := p.Lookup("F")
	panicErr(err)
	*v.(*int) = 10
	f.(func())()
}
