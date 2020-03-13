package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to Meant4")
}

type JsonResponse struct {
	Data interface{} `json:"data"`
	Status interface{} `json:"status"`
}

type CalculateParams struct {
	AInput   int  `json:"a"`
	BInput int `json:"b"`
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/calculate", CalculateFactorial)
	log.Fatal(http.ListenAndServe(":8989", router))
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
	if err != nil || msg.AInput<1 || msg.BInput<1{
	WriteHandler(w,"Error","Incorrect Input",http.StatusBadRequest)
	return
	}
	
    c1 := make(chan int)
	c2 := make(chan int)
    go func() {
		c1 <- factorial(msg.AInput)
    }()
    go func() {
        c2 <- factorial(msg.BInput)
    }()
	var finalc1,finalc2 int
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
			finalc1=msg1
        case msg2 := <-c2:
			finalc2=msg2
        } 
    }
 WriteHandler(w,"Success",strconv.Itoa(finalc1*finalc2),http.StatusCreated)
}

func factorial(n int)(result int) {
	if (n > 0) {
		result = n * factorial(n-1)
		return result
	}
	return 1
}
         
func WriteHandler(w http.ResponseWriter,status string,data string, StatusCode int){
	response := &JsonResponse{Status:status,Data: data}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(StatusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	} 
}
