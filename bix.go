package bix

import (
	"fmt"
	"log"
	"net/http"
)

//BASIC DOCS SERVE FUNC
func Startbix() {
	http.HandleFunc("/", HelloServer)
	fmt.Println("BixGo Server Start...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func BasicStart() {
	fmt.Println("Server Starting")

	http.HandleFunc("/", HttpFileHandler)

	http.ListenAndServe(":8080", nil)
}

func HttpFileHandler(response http.ResponseWriter, request *http.Request) {
	//fmt.Fprintf(w, "Hi from e %s!", r.URL.Path[1:])
	http.ServeFile(response, request, "index.html")
}
