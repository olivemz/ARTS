package main

import "fmt"

type Roman struct {
	rom string
	val int
}

var romArray = []Roman{
	Roman{
		rom: "M",
		val: 1000,
	},
	Roman{
		rom: "CM",
		val: 900,
	},
	Roman{
		rom: "D",
		val: 500,
	},
	Roman{
		rom: "CD",
		val: 400,
	},
	Roman{
		rom: "C",
		val: 100,
	},
	Roman{
		rom: "XC",
		val: 90,
	},
	Roman{
		rom: "L",
		val: 50,
	},
	Roman{
		rom: "XL",
		val: 40,
	},
	Roman{
		rom: "X",
		val: 10,
	},
	Roman{
		rom: "IX",
		val: 9,
	},
	Roman{
		rom: "V",
		val: 5,
	},
	Roman{
		rom: "IV",
		val: 4,
	},
	Roman{
		rom: "I",
		val: 1,
	},
}

func intToRoman(num int) string {
	result := ""

	return result
}

func main() {
	x := intToRoman(3940)

	fmt.Print(x)
}
