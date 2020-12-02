package main

import (
    "Week02/service"
	"log"
)

func main() {
	svr := service.New()
	num, err := svr.Update()
	if err != nil {
		log.Println("HTTP 500")
	}
	log.Println(num)
}

