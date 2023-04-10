package main
import "fmt"

//  Task 1
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

// Task 2
func findAverage(arr []int, l int) float32 {
	summ := 0
	for i:= 0; i < l; i++{
		summ = summ + arr[i]
	}
	return float32(summ)/float32(l)
}

// Task 4
func factorial(n int) float64{
	if(n == 1){
		return 1 
	}else{
		return float64(n) * factorial(n-1)
	}
}

// Task 5
func fibbonachi(n int) int{

	if(n == 0){
		return 0
	}else if(n == 1){
		return 1
	}else{
		return fibbonachi(n-1) + fibbonachi(n-2)
	}
}

// Task 6
func power(a int, p int) int {
    res:= 1;
    for i:=0; i < p; i++{
        res = res * a;
    }
    return res
}

// Task 8
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

func main(){
	fmt.Println("Hello world!")
}