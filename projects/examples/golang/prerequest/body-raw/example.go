package api
import "callapi"
import "net/http"
import "strconv"
import "bytes"
import "io"
import "io/ioutil"

type CallBack struct {}

func (c  CallBack)PreRequest(req *http.Request, ctx *callapi.RestBirdCtx) {

	data := []byte("this is raw body content")

	var body io.Reader
	body = bytes.NewReader(data)

	rc, ok := body.(io.ReadCloser)
	if !ok && body != nil {
		rc = ioutil.NopCloser(body)
	}

	req.Body = rc

	req.Header.Add("Content-Length", strconv.Itoa(len(data)))
}
