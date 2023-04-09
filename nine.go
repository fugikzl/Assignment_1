package main
import "fmt"

func factorial(n int) float64{
	if(n == 1){
		return 1 
	}else{
		return float64(n) * factorial(n-1)
	}
}

func binom(n int, k int) float64{
	return factorial(n)/((factorial(n-k)*factorial(k)))
}

func main(){
	var n,k int
	fmt.Scan(&n, &k)
	fmt.Println(binom(n,k))
}