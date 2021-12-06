package main

import (
	"fmt"
)
func main(){
	fmt.Println("Enter the no of elements")
	var s int
	fmt.Scan(&s)
	a:=[10]int{}
	for i:=0;i<s;i++{
	fmt.Println("Enter the elements ")
	var e int
	fmt.Scan(&e)
	a[i]=e
	}
    fmt.Println(a)
}