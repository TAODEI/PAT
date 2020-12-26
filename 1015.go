package main

import "fmt"

var ps [4][]p

type p struct {
	a int
	b int
	c int
	d int
}

func ssort(arr []p) []p {
	var gap int
	for gap = len(arr) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(arr); i++ {
			for j, current := i, arr[i]; ; j -= gap {
				if !(j-gap >= 0 && current.d <= arr[j-gap].d) {
					arr[j] = current
					break
				}
				arr[j] = arr[j-gap]
			}
			// for j, current := i, arr[i]; ; j -= gap {
			// 	if !(j-gap >= 0 && current.d == arr[j-gap].d) {
			// 		if current.b<arr[j-gap].b{
			// 			arr[j] = current
			// 			break
			// 	}
			// 	arr[j] = arr[j-gap]
			// }
		}
	}
	return arr
}

func sSort(arr []p) []p {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i].d < arr[j].d {
				arr[i], arr[j] = arr[j], arr[i]
				continue
			}
			if arr[i].d > arr[j].d {
				continue
			} else {
				if arr[i].b < arr[j].b {
					arr[i], arr[j] = arr[j], arr[i]
				}
				if arr[i].b == arr[j].b {
					if arr[i].a > arr[j].a {
						arr[i], arr[j] = arr[j], arr[i]
					}
				}
			}
		}
	}
	return arr
}
func main() {
	var n, l, m, sum int
	fmt.Scan(&n, &l, &m)
	for i := 0; i < n; i++ {
		var x p
		fmt.Scan(&x.a, &x.b, &x.c)
		if x.b < l || x.c < l {
			continue
		}
		sum++
		x.d = x.b + x.c
		if x.b >= m && x.c >= m {
			ps[0] = append(ps[0], x)
			continue
		}
		if x.b >= m && x.c < m {
			ps[1] = append(ps[1], x)
			continue
		}
		if x.c < m && x.b < m && x.b >= x.c {
			ps[2] = append(ps[2], x)
			continue
		}
		ps[3] = append(ps[3], x)
	}
	ps[0] = sSort(ps[0])
	ps[1] = sSort(ps[1])
	ps[2] = sSort(ps[2])
	ps[3] = sSort(ps[3])
	fmt.Println(sum)
	for i := 0; i < 4; i++ {
		for j := 0; j < len(ps[i]); j++ {
			fmt.Printf("%d %d %d\n", ps[i][j].a, ps[i][j].b, ps[i][j].c)
		}
	}
}
