package main

import (
"log"
"sumi/app/apiserver"
)

func main() {
	config := apiserver.LoadConfig()
	err := apiserver.Start(config)
	if err != nil {
		log.Fatal(err)
	}
}
