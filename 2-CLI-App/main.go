package main

import (
	"cli-app/app"
	"log"
	"os"
)

func main() {

	aplication := app.Generate()
	if error := aplication.Run(os.Args); error != nil {
		log.Fatal(error)
	}

}
