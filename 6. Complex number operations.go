//User Defined declaration for complex numbers 
package main 
import "fmt"
func main(){
	fmt.Println("***Welcome for the operations on complex numbers***")
	var re float64
	var i float64
	fmt.Println("Enter the real part")
	fmt.Scan(&re)
	fmt.Println("Enter the imaginary part")
	fmt.Scan(&i)
	var n complex128=complex(re,i)
	fmt.Printf("The complex no entered is %v",n)
    

}

//Addition of complex numbers
