package api
import "net/http"
import "fmt"

func HandleRequest(w http.ResponseWriter, r *http.Request, pam map[string]string ) error {

	fmt.Fprint(w, "Here is all req headers:\n")

	//get the req Headers
	for k, v := range r.Header {
		for _, vv := range v {
			fmt.Fprint(w, k+"=" + vv)
			fmt.Fprint(w, "\n")
		}
	}
	return nil
}
