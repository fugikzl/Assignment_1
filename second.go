package main
import "fmt"

func findAverage(arr []int, l int) float32 {
	summ := 0
	for i:= 0; i < l; i++{
		summ = summ + arr[i]
	}
	return float32(summ)/float32(l)
}

func main(){
	var l int
	fmt.Scan(&l)
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		fmt.Scan(&arr[i])
	}
	average := findAverage(arr, l)
	fmt.Println(average)
}