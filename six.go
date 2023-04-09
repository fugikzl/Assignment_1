package main
import "fmt"

func power(a int, p int) int {
    res:= 1;
    for i:=0; i < p; i++{
        res = res * a;
    }
    return res
}

func main(){
    var a,n int
    fmt.Scan(&a,&n)
    fmt.Println(power(a,n))
}