package main 
import "fmt"
func main(){
	fmt.Println("***Welcome for the operations on complex numbers***")
	fmt.Println("The the Data for 1 operator")
	var re float64
	var i float64
	fmt.Println("Enter the real part")
	fmt.Scan(&re)
	fmt.Println("Enter the imaginary part")
	fmt.Scan(&i)
	var n complex128=complex(re,i)
	fmt.Println("The the Data for 2 operator")
	var re2 float64
	var i2 float64
	fmt.Println("Enter the real part")
	fmt.Scan(&re2)
	fmt.Println("Enter the imaginary part")
	fmt.Scan(&i2)
	var m complex128=complex(re2,i2)
	var c string
	fmt.Println("Enter the operator")
	fmt.Scan(&c)
    switch c {
	case "+" :
		fmt.Printf("%v",m+n)
	case "-" :
		fmt.Printf("%v",m-n)
	case "*" :
		fmt.Printf("%v",m*n)
	default:
		fmt.Println("Invalid syntax")

	}


}