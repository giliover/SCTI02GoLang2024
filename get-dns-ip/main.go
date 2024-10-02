package main
import (
	"go.mod/app"
	"log"
	"os"
)


func main(){
	r := app.Start()
	err := r.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	} 
}