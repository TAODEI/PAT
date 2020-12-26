package main

import "fmt"

func main() {
	var a, b, d, x, y, i int64
	fmt.Scan(&a, &b, &d)
	c := a + b
	for i = 1; c != 0; i *= 10 {
		x = c - d*(c/d)
		c = c / d
		y += i * x
	}
	fmt.Println(y)
}
