package api
import "net/http"

func HandleRequest(w http.ResponseWriter, r *http.Request, pam map[string]string ) error {

	http.Redirect(w, r, "http://www.google.com/", 307)

	return nil
}
