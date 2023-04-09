package main
import "fmt"

func factorial(n int) float64{
	if(n == 1){
		return 1 
	}else{
		return float64(n) * factorial(n-1)
	}
}

func main(){
	var n int
	fmt.Scan(&n)
	fmt.Println(factorial(n))
}