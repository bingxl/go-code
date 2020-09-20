package leetcode

import "strings"

func p12v1(num int) string {
	type roman struct {
		symbol string
		value  int
	}
	romans := []roman{
		{"M", 1000},
		{"D", 500},
		{"C", 100},
		{"L", 50},
		{"X", 10},
		{"V", 5},
		{"I", 1},
	}

	getRoman := func(ten, five, base roman, v int) string {
		reman := v / base.value
		var result string = ""
		switch reman {
		case 1, 2, 3:
			{
				result = strings.Repeat(base.symbol, reman)
			}
		case 4:
			{
				result = base.symbol + five.symbol
			}
		case 5, 6, 7, 8:
			{
				result = five.symbol + strings.Repeat(base.symbol, reman-5)
			}
		case 9:
			{
				result = base.symbol + ten.symbol
			}
		}
		return result
	}
	reman := num
	result := strings.Repeat(romans[0].symbol, num/romans[0].value)
	reman = num % romans[0].value
	result += getRoman(romans[0], romans[1], romans[2], reman)
	reman = reman % romans[2].value
	result += getRoman(romans[2], romans[3], romans[4], reman)
	reman = reman % romans[4].value
	result += getRoman(romans[4], romans[5], romans[6], reman)
	return result
}
