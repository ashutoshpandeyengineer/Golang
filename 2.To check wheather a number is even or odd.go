package main

import "fmt"
func main(){
    fmt.Println("Enter The value ")
    var x int
    fmt.Scanln(&x)
    if x%2 == 0 {
        fmt.Print(x, " is even")
    } else {
        fmt.Print(x, " is odd")
    }
}
    