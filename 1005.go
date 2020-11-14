package main
import "fmt"
func main(){
	var n,i int
	var a,b []int
	fmt.Scanf("%d",&n)
	var x [100][100]int

	for i = 0; i < n; i++ {
		tmp := 0
		fmt.Scanf("%d",&tmp)
		a = append(a,tmp)
	}

	for i = 0; i < n; i++ {
		h := a[i]
		for j:= 0; h != 1; j++ {
			if h%2 == 0 {
				h/= 2
			}else{
				h = (3*h+1)/2
			}
			x[i][j] = h
		}
	}
	hh := 0
	for i = 0; i < n; i++ {
		for j:= 0; j < n; j++ {
			if hh == 1 {break}
			for k:= 0; x[j][k] != 1; k++ {
				if a[i] == x[j][k] {
					hh = 1
					break
				}
			}
		}
		if hh == 1 {
             hh = 0
		     continue
		}
		b=append(b,a[i])
	}

	for i = 0; i < len(b)-1; i++ {
		min := i
		for j := i+1; j < len(b); j++ {
			if b[j] < b[min] {
				min = j
			}
		}
		b[i],b[min] = b[min],b[i]
	}

	for i = len(b)-1; i>= 0; i-- {
		if i!=0 {
		fmt.Printf("%d ",b[i])
		}else{fmt.Printf("%d",b[i])		
		}
	}
	fmt.Println()
}
