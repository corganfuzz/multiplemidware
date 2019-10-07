package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type city struct {
	Name string
	Area uint64
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	// Check if POST method exists

	if r.Method == "POST" {
		var tempCity city
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&tempCity)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		//Resource Creatin logic goes here

		fmt.Printf("Got %s city with area of %d sq miles!\n", tempCity.Name, tempCity.Area)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("201 - Created"))
	} else {

		//Method not Allowed

		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method not Allowed"))
	}
}

func main() {
	fmt.Println("Server is running in port 8000...")
	http.HandleFunc("/city", mainLogic)
	http.ListenAndServe(":8000", nil)
}
