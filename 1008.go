package main
import "fmt"
func main(){
	var m,n,i,tmp int
	var a []int
	fmt.Scanf("%d%d",&n,&m)
	for i=0;i<n;i++ {
		fmt.Scanf("%d",&tmp)
		a=append(a,tmp)
	}
	for i=0;i<m;i++ {
	a = append(a[:i], append([]int{a[n-m+i*2]}, a[i:]...)...)
	}
	for i=0;i<len(a)-m;i++ {
		if i!=len(a)-m-1 {
			fmt.Printf("%d ",a[i])
		}else{
			fmt.Printf("%d\n",a[i])
		}
	}
}

