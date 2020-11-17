package main
import (
	"fmt"
	"math"
)
func main(){
	var m,n int
	fmt.Scanf("%d",&n)
	tmp:= 2
	for i:= 3; i <= n; i+= 2 {
		for j:= 3;; j+= 2 {
			if j > int(math.Sqrt(float64(i))){
				if tmp+2 == i {m++}
				tmp = i
				break
			}
			if i%j == 0 {
				break
			}
		}
	}
	fmt.Println(m)
}
