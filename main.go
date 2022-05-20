package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"time"
)

type Response struct {
	 Code int `json:"code"`
	Message string `json:"message"`
	 Timestamp int64 `json:"timestamp"`
}


func readEvents(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	statusCode := rand.Intn((299-200)+200)
	resp := &Response{statusCode,"Some message",time.Now().Unix()}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	b,err := json.Marshal(resp)
	m := mux.Vars(r) // m map[string]string
	fmt.Println(m)
	query := r.URL.Query()
	fmt.Println(query)
	if err != nil{
		fmt.Errorf(err.Error())
	}
	w.Write(b)
}


func main(){
	router := mux.NewRouter()

	router.HandleFunc("/device/{deviceId}/config", readEvents).Methods("GET")

	http.ListenAndServe(":5300", router)
}
