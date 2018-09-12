package api
import "net/http"
import "fmt"

type MyCounter struct {
	Count     int      `json:"count"`
}

func HandleRequest(w http.ResponseWriter, r *http.Request, pam map[string]string ) error {

	var counter = &MyCounter{}

	if err := State.Get("counter", &counter); err != nil {
		counter = &MyCounter{Count:1}
	} else {
		counter.Count++
	}

	State.Save("counter", *counter)

	fmt.Fprintf(w, "current counter:%d\n", counter.Count)

	return nil
}
