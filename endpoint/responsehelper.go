package endpoint

import (
	"encoding/json"
	"net/http"
	"strings"

	router "sbcash_api/BARoutingHelper"
)

func GETResponse(w http.ResponseWriter, bodyData interface{}) {
	router.ResponseByteAsJSON(w, http.StatusOK, bodyData)
}

func NOTFOUNDResponse(w http.ResponseWriter, bodyData interface{}) {
	var response map[string]string = make(map[string]string)

	response["message"] = "Item não encontrado"
	if bodyData != nil {
		var buf strings.Builder
		e := json.NewEncoder(&buf)
		e.Encode(bodyData)
		data := buf.String()
		response["data"] = data
	}

	router.ResponseByteAsJSON(w, http.StatusNotFound, response)
}

func ERRORResponse(w http.ResponseWriter, bodyData interface{}) {
	var response map[string]string = make(map[string]string)

	response["message"] = "Erro interno"
	if bodyData != nil {
		var buf strings.Builder
		e := json.NewEncoder(&buf)
		e.Encode(bodyData)
		data := buf.String()
		response["data"] = data
	}

	router.ResponseByteAsJSON(w, http.StatusInternalServerError, response)
}

func POSTResponse(w http.ResponseWriter, bodyData interface{}) {
	var response map[string]string = make(map[string]string)

	response["message"] = "Criado com sucesso"
	if bodyData != nil {
		var buf strings.Builder
		e := json.NewEncoder(&buf)
		e.Encode(bodyData)
		data := buf.String()
		response["data"] = data
	}

	router.ResponseByteAsJSON(w, http.StatusCreated, response)
}

func PUTResponse(w http.ResponseWriter, bodyData interface{}) {
	var response map[string]string = make(map[string]string)

	response["message"] = "Alterado com sucesso"
	if bodyData != nil {
		var buf strings.Builder
		e := json.NewEncoder(&buf)
		e.Encode(bodyData)
		data := buf.String()
		response["data"] = data
	}

	router.ResponseByteAsJSON(w, http.StatusAccepted, response)
}

func DELETEResponse(w http.ResponseWriter, bodyData interface{}) {
	var response map[string]string = make(map[string]string)

	response["message"] = "Removido com sucesso"
	if bodyData != nil {
		var buf strings.Builder
		e := json.NewEncoder(&buf)
		e.Encode(bodyData)
		data := buf.String()
		response["data"] = data
	}

	router.ResponseByteAsJSON(w, http.StatusMovedPermanently, response)
}
