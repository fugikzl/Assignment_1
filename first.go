package main
import "fmt"

func findMin(arr []int, l int) int {
	min := arr[0]
	for i := 0; i < l; i++{
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func main() {
	var l int
	fmt.Scan(&l)
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		fmt.Scan(&arr[i])
	}
	min := findMin(arr, l)
	fmt.Println(min)
}