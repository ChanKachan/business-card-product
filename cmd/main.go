package main

import (
	"log"

	"github.com/ChanKachan/business-card-product/internal/app"
)

func main() {
	if err := app.Start(); err != nil {
		log.Println(err)
		return
	}
}
