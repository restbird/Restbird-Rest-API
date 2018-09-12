package api
import "net/http"
import "fmt"

// three different methos to access query parameters
func HandleRequest(w http.ResponseWriter, r *http.Request, pam map[string]string ) error {

    //r.FormValue to access the url query parameters
    fmt.Fprint(w, "fileds=" + r.FormValue("fields") + "\n")

    //r.URL.Query().Get() only get the 1st item of the query param
    fmt.Fprint(w, "fields=" + r.URL.Query().Get("fields") + "\n")

    //r.URL.Query()[] get array of all values of the query param
    paramarray := r.URL.Query()["fields"]

    for _, value := range paramarray {
        fmt.Fprint(w, "fields="+value + "\n")
    }

    return nil
}
