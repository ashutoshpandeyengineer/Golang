package main 
import "fmt"
func main(){
	statepopulation:=map[string]int{
		"texas":123456 ,
		"Dallas":789456 ,
		"Seattle":123789 ,
	}
	fmt.Println(statepopulation)
	statepopulation["Abu dhabi"]=45582
	delete(statepopulation,"texas")
	fmt.Println(statepopulation)
}