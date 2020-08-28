package main

import (
	"database/sql"
	"fmt"
)

func main() {
	fmt.Println(sql.ErrConnDone)
	// fmt.Println(sql.ErrNoRows)
	// fmt.Println(sql.ErrTxDone)

	fmt.Println(sql.Drivers())

}
