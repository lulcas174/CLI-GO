package main

import (
	"cli/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Application started...")

	aplication := app.Generate()
	erro := aplication.Run(os.Args)

	if erro != nil {
		log.Fatal(erro)
	}

}
