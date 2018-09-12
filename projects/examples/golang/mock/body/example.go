package api
import "net/http"

func HandleRequest(w http.ResponseWriter, r *http.Request, pam map[string]string ) error {

	// use w.Wrtie to write response body 
	w.Write([]byte("hello, this  is a test"))

	return nil
}
