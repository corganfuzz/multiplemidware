package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type city struct {
	Name string
	Area uint64
}

//MiddleWare to check if content type as JSON

func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Currently midware is checking content type ")

		// Filtering requests by MIME type

		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupported Media Type. Only JSON accepted"))
			return
		}
		handler.ServeHTTP(w, r)
	})
}

// Middlerware to add server timestamp for response cookie

func setServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)

		//Setting cookie to each and every response

		cookie := http.Cookie{
			Name:  "Server-Time(UTC)",
			Value: strconv.FormatInt(time.Now().Unix(), 10),
		}
		http.SetCookie(w, &cookie)
		log.Println("Currently in the set server time middleware")
	})
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
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/city", filterContentType(setServerTimeCookie(mainLogicHandler)))
	http.ListenAndServe(":8000", nil)
}
