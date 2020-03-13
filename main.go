package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to Meant4")
}

type JsonResponse struct {
	Data interface{} `json:"data"`
	Status interface{} `json:"status"`
}
type CalculateParams struct {
	AInput   int64  `json:"a"`
	BInput int64 `json:"b"`
}

func CalculateFactorial(w http.ResponseWriter, r *http.Request, hp httprouter.Params) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var msg CalculateParams
	err = json.Unmarshal(b, &msg)
	if err != nil {
	response := &JsonResponse{Status:"Error",Data: "Incorrect Input"}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	} 
	return
	}
	response := &JsonResponse{Status:"Success",Data: "Success"}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	} 
}


func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/calculate", CalculateFactorial)

	log.Fatal(http.ListenAndServe(":8080", router))
}