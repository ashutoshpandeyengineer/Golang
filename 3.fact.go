package main
import "fmt"
func main(){
	fmt.Print("Enter the number to find the factorial")
	var v int
	fmt.Scan(&v)
	if v>1 {
		factorial :=1
		i:=1
		for i<=v{
			factorial=factorial*i
			i=i+1
		}
		fmt.Println(factorial)
	} else {
		fmt.Println("The Factorial is ",1)
	}
}