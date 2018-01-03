package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"encoding/json"
	"fmt"
)

/**
http://localhost:8080/hello
http://localhost:8080/json
 */
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello go"))
	})

	router.HandleFunc("/json", handleJSON)

	err := http.ListenAndServe(":8080", router)
	handlerError(err)
}

func handlerError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func handleJSON(writer http.ResponseWriter, request *http.Request) {
	row := make(map[string]interface{})
	row["id"] = 1
	row["name"] = "aaa"
	row["age"] = 11
	row["address"] = "上海"
	row["created_at"] = time.Now()

	json.NewEncoder(writer).Encode(row)
}
