package main

import (
	"fmt"
)

const x = 99

func init() {
	fmt.Println("This is my func do%d", x)
}

func main() {

	switch {
	case x < 100:
		fmt.Println("Lower")
	case x > 100:
		fmt.Println("Hiogher")
	}
	for x := 0; x < 100; x++ {
		fmt.Printf("%d", x)
	}

	z := []int{10, 20, 30}
	m := map[string]int{
		"Hello": 41,
		"gELLO": 56,
	}
	for i, v := range z {
		fmt.Println(i, v)
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
