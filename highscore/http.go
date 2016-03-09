package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func decodeJson(r io.Reader, dst interface{}) interface{} {
	json.NewDecoder(r).Decode(&dst)
	return dst
}

func JsonDecoder(paramsTemplate interface{}) func(r io.Reader) interface{} {
	return func(r io.Reader) interface{} {
		return decodeJson(r, paramsTemplate)
	}
}

func Responder(response *HandlerResponse, w http.ResponseWriter) {
	switch response.ResponseType {
	case DATA:
		out, _ := json.Marshal(response.Data)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, string(out))
	case TEXT:
		io.WriteString(w, response.Data.(string))
	}
}

func HttpHandler(decoder func(body io.Reader) interface{}, handler func(params interface{}) HandlerResponse) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := decoder(r.Body)
		var response = handler(params)

		Responder(&response, w)
	}
}
