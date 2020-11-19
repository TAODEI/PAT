package main
import (
	"fmt"
	"os"
	"bufio"
	"strings"
)
func main(){
	var s string
	inputReader := bufio.NewReader(os.Stdin)
	s ,_= inputReader.ReadString('\n')
	a := strings.Fields(s)
	for i := len(a); i>0; i-- {
		if i!=1 {
		fmt.Printf("%s ",a[i-1])
		}else{
		fmt.Printf("%s\n",a[0])
		}
	}
}
