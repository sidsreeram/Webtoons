package main

import (
	"log"
    "github.com/webtoons/pkg/config"
    "github.com/webtoons/pkg/di"


)


func main() {
    config, configErr := config.LoadConfig()
    if configErr != nil {
        log.Fatal("Error in configuration", configErr)
    }
   

    server, diErr := di.InitializeAPI(config)
    if diErr != nil {
        log.Fatal("Cannot initialize API", diErr)
    }

    if server != nil {
        server.Start()
    } else {
        log.Fatal("Server is nil")
    }
}