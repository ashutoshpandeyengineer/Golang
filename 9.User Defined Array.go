package main
import "fmt"
func main () {
	fmt.Printf("Enter the size of the array")
	var  s int
	fmt.Scan(&s)
	a:=[5]int{}
	for i:=0 ; i<s ; i++{
		fmt.Println("Enter the ",i,"value")
		var v int
		fmt.Scan(&v)
		a[i]=v
	}
	fmt.Println("The array is ", a)
}