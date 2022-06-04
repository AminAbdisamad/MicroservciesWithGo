package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const PORT = ":2022"

func home(writer http.ResponseWriter, request *http.Request) {
	log.Println("API Running")
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "Something went wrong", http.StatusBadRequest)
		log.Println(err)
		return
	}
	fmt.Fprintf(writer, "Hello %s", data)

}
func main() {
	http.HandleFunc("/", home)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Println(err)
		return
	}

}
