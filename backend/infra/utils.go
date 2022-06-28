package infra

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type ResponsePayload struct {
	Message string `json:"message"`
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("Unable to read the request body: %s", err.Error())
		SendResponse(w, "Please send a valid json.", http.StatusBadRequest)
		return err
	}

	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Printf("Unable to parse JSON: %s", err.Error())
		SendResponse(w, "Please send a valid json.", http.StatusBadRequest)
		return err
	}

	return nil
}

func SendResponse(w http.ResponseWriter, msg string, status int) error {
	res, ok := json.Marshal(ResponsePayload{Message: msg})

	if ok != nil {
		return ok
	}

	w.WriteHeader(status)
	w.Write(res)

	return nil
}
