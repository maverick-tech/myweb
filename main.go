package main

import (
	"fmt"
	"net/http"
	"os"
)

func webHandlerFunc(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "<h1>Hello, World</h1>")
}

func GetPort() (port string){
	return os.Getenv("PORT")
}

func main() {
	port:=GetPort()

	http.HandleFunc("/", webHandlerFunc)
	http.ListenAndServe(":" + port, nil)
}

