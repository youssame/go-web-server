package main

import (
	"go-webserver/pkg"
	"os"
)

func main() {
	pkg.RunServer(pkg.ServerConfig{
		Port: os.Getenv("PORT"),
	})
}
