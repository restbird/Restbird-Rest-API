package api
import "net/http"
import "fmt"

// the path will be :  /xxxx/:user/:book/
func HandleRequest(w http.ResponseWriter, r *http.Request, pam map[string]string ) error {

    fmt.Fprint(w, pam["user"])
    fmt.Fprint(w, pam["book"])

    return nil
}
