package main
import "fmt"
import "os"

func main() {
	var inp string
	fmt.Scan(&inp)
	l := len(inp)
	for i:=0; i < l; i++{
		j:=int(inp[i])
		if( j < 48 || j > 57 ){
			fmt.Println("NO")
			os.Exit(0)
		}
	}
	fmt.Println("YES")
	os.Exit(0)

}