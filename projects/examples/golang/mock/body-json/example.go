package api
import "net/http"
import "encoding/json"

type MyDATA struct {
	Aa    string `json:"aa, omitempty"`
	Bb    int `json:"bb, omitempty"`
	Cc    bool `json:"cc, omitempty"`
}

func HandleRequest(w http.ResponseWriter, r *http.Request, pam map[string]string ) error {
	var err error
	var body []byte

	mydata := MyDATA {
		Aa: "hello world",
		Bb: 100,
		Cc: true,
	}

	if body, err = json.Marshal(mydata); err != nil {
		return err
	}

	// content-type header
	w.Header().Add("Content-Type", "application/json")

	w.Write(body)
	return nil
}
