package main

import (
	"fmt"
	"net/http"
	"os"
)

func webHandlerFunc(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "<h1>Hello, World</h1> <h3>This is new feature</h3> <h3> This is feature 2</h3>")
}

func GetPort() (port string){
	return os.Getenv("PORT")
}

func main() {
	port:=GetPort()

	http.HandleFunc("/", webHandlerFunc)
	http.ListenAndServe(":" + port, nil)
}

