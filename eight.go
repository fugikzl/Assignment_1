package main
import "fmt"

func checkForSymbol(inp string) string{
	l := len(inp)
	for i:=0; i < l; i++{
		j:=int(inp[i])
		if( j < 48 || j > 57 ){
			return "NO"
		}
	}
	return "YES"
}

func main() {
	var inp string
	fmt.Scan(&inp)
	fmt.Println(checkForSymbol(inp))
}