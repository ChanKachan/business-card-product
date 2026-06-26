package main

import (
	"log"
	"visit/internal/app"
)

func main() {
	if err := app.Start(); err != nil {
		log.Println(err)
		return
	}
}
