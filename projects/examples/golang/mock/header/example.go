package api
import "net/http"
import "fmt"

func HandleRequest(w http.ResponseWriter, r *http.Request, pam map[string]string ) error {

	// set or add headers
	w.Header().Set("aa", "abcd")
	w.Header().Add("bb", "cdef")

	// use fmt.Fprint to write response body
	fmt.Fprint(w, "Welcome!\n")

	return nil
}
