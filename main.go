package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"io"
	"io/ioutil"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
type CalculateParams struct {
	aParam   int `json:"a"`
	bParam  int `json:"b"`
}

type JsonResponse struct {
	Data interface{} `json:"data"`
	Status interface{} `json:"status"`
}

func CalculateFactorial(rw http.ResponseWriter, r *http.Request, hp httprouter.Params) {
	var calculateParams CalculateParams

    fmt.Fprintln(rw, "Calculating Factorial")
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        log.Fatal(err)
    }
    if err := r.Body.Close(); err != nil {
        log.Fatal(err)
    }
    if err := json.Unmarshal(body,calculateParams); err != nil {
		response := &JsonResponse{Data: &calculateParams}
		rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
		rw.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(rw).Encode(response); err != nil {
			panic(err)
		}
    }
 
}


func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/calculate", CalculateFactorial)

	log.Fatal(http.ListenAndServe(":8080", router))
}