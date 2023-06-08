package main

import "fmt"

type Addr struct {
	Province string
	City     string
	ZipCode  int
}

func Structtest() {
	addr := Addr{
		Province: "四川",
		City:     "成都",
		ZipCode:  12323,
	}

	addr2 := &Addr{}
	addr3 := new(Addr)

	fmt.Printf("type(addr): %T, type(addr2): %T, type(addr3): %T", addr, addr2, addr3)
}
