package main
import (
	"fmt"
)

func main() {
	arrayStatic := [4]int32{1,2,3}	
	fmt.Println(arrayStatic)
	var arrayDinamic []int32
	fmt.Println(arrayDinamic)
}