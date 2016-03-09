package main

import (
	"allthethings/highscore/score"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	DATA = iota
	TEXT = iota
)

type HandlerResponse struct {
	ResponseType int
	Data         interface{}
}

type MyRequestParams struct {
}

func main() {
	router := NewGorillaRouter(mux.NewRouter())

	hello := func(params interface{}) HandlerResponse {
		return HandlerResponse{DATA, []string{"test"}}
	}

	index := func(params interface{}) HandlerResponse {
		return HandlerResponse{TEXT, "<html><body><h1>Scores</h1></body></html>"}
	}

	scoreRequestHandler := func(params interface{}) HandlerResponse {
		var request = params.(score.ScoreRequest)
		fmt.Println("request:", request)
		values, _ := score.GetAll()
		return HandlerResponse{DATA, values}
	}

	router.GET("/", HttpHandler(JsonDecoder(MyRequestParams{}), index))
	router.GET("/hello.json", HttpHandler(JsonDecoder(MyRequestParams{}), hello))
	router.GET("/scores.json", HttpHandler(JsonDecoder(score.ScoreRequest{}), scoreRequestHandler))

	fmt.Println("serving on 8001")
	log.Fatal(http.ListenAndServe(":8001", router.Router()))
}
