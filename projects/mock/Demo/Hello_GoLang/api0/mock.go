package api
import "net/http"
import "fmt"

func HandleRequest(w http.ResponseWriter, r *http.Request, pam map[string]string ) error {
    fmt.Fprint(w, "Hello, " + pam["user"])
	return nil
}
