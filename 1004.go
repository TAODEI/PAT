package main

import (
	"fmt"
)

type student struct {
	xh string
	cj int
	name string
}

func main(){
	var n int
	fmt.Scanf("%d",&n)
	a := make([]student,n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &a[i].name)
		fmt.Scanf("%s", &a[i].xh)
		fmt.Scanf("%d", &a[i].cj)
	}
	for i := 0; i < n-1; i++ {
		min := i
		for j := i+1; j < n; j++ {
			if a[j].cj < a[min].cj {
				min = j
			}
		}
		a[i],a[min] = a[min],a[i]
	}
	//事实上并不需要排序 只需要最大和最小
	fmt.Printf("%s %s\n",a[n-1].name,a[n-1].xh)
	fmt.Printf("%s %s\n",a[0].name,a[0].xh)
}


