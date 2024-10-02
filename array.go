package main
import (
	"fmt"
)

func main() {
	arrayStatic := [4]int32{1,2,3}	
	fmt.Println(arrayStatic)
	var arrayDinamic []int32
	fmt.Println(arrayDinamic)

	var arrayDinamic2 []int32
	arrayDinamic2 = append(arrayDinamic2, 1, 2, 3, 4)
	fmt.Println(arrayDinamic2)

	for _, v:=range arrayDinamic2 {
		fmt.Println(v)
	}
}