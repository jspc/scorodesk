package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type IncomingRequest struct {
	// project title; taken from the email subject
	Title string

	// email address of the requestor
	Requestor string

	// taken from the body of the original ticket
	Description string
}

func Incoming(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	sr := IncomingRequest{}

	err := decoder.Decode(&sr)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)

		ErrorLogger.Print(err)
		fmt.Fprint(w, err)

		return
	}

	defer r.Body.Close()

}
