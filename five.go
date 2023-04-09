package main
import "fmt"

func fibbonachi(n int) int{

	if(n == 0){
		return 0
	}else if(n == 1){
		return 1
	}else{
		return fibbonachi(n-1)
	}
}

func main(){
	var n int

	fmt.Scan(&n)
	fmt.Println(fibbonachi(n))
}