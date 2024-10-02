package main
import "fmt"

func main() {
	name := "Carlos"
	address := &name
	fmt.Println(name, address)

	*address = "oi mundo"

	fmt.Println(name, address)
}