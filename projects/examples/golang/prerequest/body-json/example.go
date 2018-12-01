package api
import "callapi"
import "net/http"
import "strconv"
import "bytes"
import	"encoding/json"
import "io"
import "io/ioutil"

type CallBack struct {}

type JsonBody struct {
	Name string  `json:"name"`
	Password string  `json:"password"`
}

func (c  CallBack)PreRequest(req *http.Request, ctx *callapi.RestBirdCtx) {

	reqdata := JsonBody{Name:"admin", Password:"admin"}

	reqjsonbytes, err := json.Marshal(reqdata)
	if err != nil {
	    return
	}


	var body io.Reader
	body = bytes.NewReader(reqjsonbytes)

	rc, ok := body.(io.ReadCloser)
	if !ok && body != nil {
		rc = ioutil.NopCloser(body)
	}


	req.Body = rc

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(reqjsonbytes)))
}
