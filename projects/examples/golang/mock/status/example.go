package api
import "net/http"

func HandleRequest(w http.ResponseWriter, r *http.Request, pam map[string]string ) error {
	// w.WriteHeader(http.StatusInternalServerError)
	w.WriteHeader(500)

	w.Write([]byte("Something bad happened!"))

	return nil
}
