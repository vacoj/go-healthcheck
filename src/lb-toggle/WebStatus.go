package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func statusWeb(rw http.ResponseWriter, req *http.Request) {
	blob, err := json.Marshal(&STATUS)
	if err != nil {
		fmt.Println(err, err.Error())
	}
	io.WriteString(rw, string(blob))
}

func statusSimpleWeb(rw http.ResponseWriter, req *http.Request) {
	if STATUS.HealthStatus.OK && STATUS.SmokeStatus.OK && STATUS.State {
		rw.WriteHeader(http.StatusOK)
	} else {
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func statusSimple2Web(rw http.ResponseWriter, req *http.Request) {
	if STATUS.HealthStatus.OK && STATUS.SmokeStatus.OK && STATUS.State {
		rw.WriteHeader(http.StatusOK)
	} else {
		var err = errors.New("intentionally erroring for sake of not returning anything to the caller")
		panic(err)
	}
}
