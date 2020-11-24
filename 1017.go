package main

import (
	"fmt"
)
var b int
var a []int
func xx(x int,r int) (int,q int){
	q=(r*10+x)/b
	a=append(a,q)
	return r*10+x-q*b,q
}
func main() {
	var s string
	var x []int
	fmt.Scanf("%s %d", &s, &b)
	for i:=0;i<len(s);i++{
		x=append(x,int(s[i])-48)
	}
	q:=0
	r:=q
	for i:=0;i<len(x);i++ {
		r,q=xx(x[i],r)
	}
	for i:=0;i<len(a);i++ {
		if a[0]==0&&i==0 {
			continue
		}else{
			fmt.Printf("%d",a[i])
		}
	}
	fmt.Printf(" %d\n",r)

}
