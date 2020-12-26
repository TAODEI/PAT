package main

import "fmt"

func s(x int) (int, [4]int) {
	var a [4]int
	a[0] = x / 1000
	a[1] = x / 100 % 10
	a[2] = x / 10 % 10
	a[3] = x % 10
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 4; j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	y := a[0]*1000 + a[1]*100 + a[2]*10 + a[3]
	z := a[3]*1000 + a[2]*100 + a[1]*10 + a[0]
	return y - z, a
}
func main() {
	var x int
	fmt.Scan(&x)
	a1 := x / 1000
	b := x / 100 % 10
	c := x / 10 % 10
	d := x % 10
	if a1 == b && b == c && c == d {
		fmt.Printf("%d - %d = 0000\n", x, x)
		return
	}
	//var a [4]int
	for h, a := s(x); ; h, a = s(h) {

		fmt.Printf("%d%d%d%d - %d%d%d%d = %d\n", a[0], a[1], a[2], a[3], a[3], a[2], a[1], a[0], h)
		if h == 6174 {
			return
		}
	}
}
