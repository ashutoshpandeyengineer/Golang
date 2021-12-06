package main
import "fmt"
func main(){
//	const row int
//	fmt.Scan(&row)
//	const column int
//	fmt.Scan(&column)
	var array1 [3][3]int
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
		fmt.Println("Enter the ",i,j,"value")
		var v int
		fmt.Scan(&v)
		array1[i][j]=v
		}
	}
	fmt.Println(array1)
}