package main

import (
	"fmt"
)
func main (){
	fmt.Println("Enter the frist operand")
	var a int
	fmt.Scan(&a)
	fmt.Println("Enter the second operand")
    var b int
	fmt.Scan(&b)
	fmt.Println("Enter the operator")
	var c string
	fmt.Scan(&c)
	switch c {
	case "+" :
		fmt.Printf("The sum is %v ",a+b)
	case "-" :
		fmt.Printf("The diffrence is %v ",a-b)
	case "*" :
		fmt.Printf("The product is %v ",a*b)
	case "/" :
		fmt.Printf("%v",a/b)
	default:
		fmt.Println("Invalid operand")
	
	}
		
	

}