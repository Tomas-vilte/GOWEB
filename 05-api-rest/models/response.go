package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string
	respWrite   http.ResponseWriter
}

func CreateDefaultResponse(rw http.ResponseWriter) Response {
	return Response{
		Status:      http.StatusOK,
		respWrite:   rw,
		contentType: "application/json",
	}
}

func (resp *Response) Send() {
	resp.respWrite.Header().Set("Content-Type", resp.contentType)
	resp.respWrite.WriteHeader(resp.Status)
	// Convertir los usuarios a formato JSON.
	output, err := json.Marshal(&resp)
	if err != nil {
		fmt.Println(err)
	}

	// Devolver los usuarios en formato JSON.
	fmt.Fprintln(resp.respWrite, string(output))
}

func SendData(rw http.ResponseWriter, data interface{}, message string) {
	response := CreateDefaultResponse(rw)
	response.Data = data
	response.Message = message
	response.Send()

}

func (resp *Response) NoFound() {
	resp.Status = http.StatusNotFound
	resp.Message = "Resource No Found"

}

func SendNoFound(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.NoFound()
	response.Send()
}

func (resp *Response) UnproccesableEntity() {
	resp.Status = http.StatusUnprocessableEntity
	resp.Message = "UnproccesableEntity No Found"
}

func SendUnproccesableEntity(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.UnproccesableEntity()
	response.Send()
}

func SendServerError(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusInternalServerError)
	rw.Header().Set("Content-Type", "application/json")
	response := Response{Status: http.StatusInternalServerError, Message: "Internal server error"}
	json.NewEncoder(rw).Encode(response)
}
