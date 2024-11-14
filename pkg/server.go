package pkg

import (
	"fmt"
	"net/http"
)

type ServerConfig struct {
	Port string
}

const defaultPort = "8080"

var mux = http.NewServeMux()

// RunServer Run the server on a port passed in the env
func RunServer(conf ServerConfig) {
	InitRoutes()
	port := conf.Port
	if port == "" {
		port = defaultPort
	}
	displayWebServerRunningMessage(port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		fmt.Println(err)
	}
}

// AddRoute add an http handler
func AddRoute(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	mux.HandleFunc(pattern, handler)
}

// displayWebServerRunningMessage display a message indicates the port and running state of the server
func displayWebServerRunningMessage(port string) {
	fmt.Println("***************************************************")
	fmt.Println("******** WEB SERVER RUNNING ON PORT : " + port + " ********")
	fmt.Println("***************************************************")
}
