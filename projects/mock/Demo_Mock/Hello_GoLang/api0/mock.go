package api

import (
	"State"
	"fmt"
	"net/http"
)

type MyCounter struct {
	Count int `json:"count"`
}

func HandleRequest(w http.ResponseWriter, r *http.Request, pam map[string]string) error {

	// set or add headers
	w.Header().Set("aa", "123")
	w.Header().Add("bb", "456")

	//get/set Stat
	var counter = &MyCounter{}

	if err := State.Get("counter", &counter); err != nil {
		counter = &MyCounter{Count: 1}
	} else {
		counter.Count++
	}

	State.Save("counter", *counter)

	fmt.Fprintf(w, "Hello, "+pam["user"]+"\ncurrent counter:%d\n", counter.Count)

	return nil
}
