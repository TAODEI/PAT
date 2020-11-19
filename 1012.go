package main
import "fmt"
func main(){
	n := 0
	fmt.Scanf("%d",&n)
	var x,a1,a2,a3,a5,n2,n4 int
    var a4 float64
	for i:= 0; i<n; i++{
		fmt.Scanf("%d",&x)
		switch x%5 {
		case 0:if x%2==0 { a1+=x }
		case 1:
			n2++
			if n2%2==1 {
				a2+=x
			}else{
				a2-=x
			}
		case 2: a3++
		case 3:
			n4++
			a4+= float64(x)
		case 4: if x > a5 { a5 = x }
		}
	}
	if a1!=0 {
		fmt.Printf("%d ",a1)
	}else{
		fmt.Printf("N ")
	}
	if a2!=0 {
		fmt.Printf("%d ",a2)
	}else{
		fmt.Printf("N ")
	}
	if a3!=0 {
		fmt.Printf("%d ",a3)
	}else{
	    fmt.Printf("N ")
	}
	if a4!=0 {
		fmt.Printf("%.1f ",a4/float64(n4))
	}else{
        fmt.Printf("N ")
	}
	if a5!=0 {
		fmt.Printf("%d\n",a5)
	}else{
        fmt.Printf("N\n")
	}
}


