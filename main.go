package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

const PORT = "2022"

func main() {
	http.HandleFunc(
		"/", func(writer http.ResponseWriter, request *http.Request) {
			log.Println("API Running")
			data, err := ioutil.ReadAll(request.Body)
			if err != nil {
				log.Println(err)
				return
			}
			log.Printf("Data Received: %s\n\n", data)

		},
	)
	err := http.ListenAndServe(":2022", nil)
	if err != nil {
		log.Println(err)
		return
	}

}
