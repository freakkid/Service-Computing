package main

import (
	"net/http"
	"os"

	"github.com/freakkid/Service-Computing/hw3/server"
	flag "github.com/spf13/pflag"
)

const (
	PORT string = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	pPort := flag.StringP("port", "p", PORT, "PORT for listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	http.HandleFunc("/user/register", server.RegisterHandler)
	http.HandleFunc("/todo/add", server.AddItemHandler)
	http.HandleFunc("/todo/delete", server.DeleteItemHandler)
	http.HandleFunc("/todo/show", server.ShowListHandler)
	http.HandleFunc("/", server.OtherHandler)

	server.ListenAndServe(":"+port, nil)

}
