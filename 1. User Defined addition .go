package main
import "fmt"
func main(){
	fmt.Println("Enter The frist number")
	var x int
	fmt.Scanln(&x)
	fmt.Println("Enter the second number")
	var y int
	fmt.Scanln(&y)
	fmt.Printf("The value of sum is %v", x+y)
}