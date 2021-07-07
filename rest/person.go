package rest

import (
	"encoding/json"
	"fmt"
	"golangserver/entity"
	"io/ioutil"
	"net/http"
)

func PostPerson(rw http.ResponseWriter, r *http.Request) {
	reqBody, err := r.GetBody()
	if hasError(rw, err, "Internal issue") {
		return
	}

	bodyBytes, err := ioutil.ReadAll(reqBody)
	if hasError(rw, err, "Internal issue") {
		return
	}

	var person = entity.Person{}
	err = json.Unmarshal(bodyBytes, &person)
	if hasError(rw, err, "Internal issue") {
		return
	}

	fmt.Println(person)
	rw.Write(bodyBytes)
}

func hasError(rw http.ResponseWriter, err error, message string) bool {
	if err != nil {
		rw.Write([]byte(message))
		return true
	}
	return false
}
