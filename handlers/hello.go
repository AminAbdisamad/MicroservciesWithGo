package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	logger *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (hello *Hello) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	hello.logger.Println("API Running")
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "Something went wrong", http.StatusBadRequest)
		hello.logger.Println(err)
		return
	}
	fmt.Fprintf(writer, "Hello %s", data)

}
