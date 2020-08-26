package main

import (
	"log"

	"github.com/JnecUA/GolangRestAPIExemple/internal/app/apiserver"
)

func main() {
	config := apiserver.DefaultConfig()
	App := apiserver.Init(config)
	if err := App.Start(); err != nil {
		log.Fatal(err)
	}

}
