package main
import "fmt"

func min(arr []int, n int) int{
	if(n==1){
		return arr[0]
	}
	m := min(arr, n-1)

	if(m < arr[n-1]){
		return m
	}else{
		return arr[n-1]
	}
}

func main() {
	var l int
	fmt.Scan(&l)
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		fmt.Scan(&arr[i])
	}
	min := min(arr, l)
	fmt.Println(min)
}