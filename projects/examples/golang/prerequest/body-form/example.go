package api
import "callapi"
import "net/http"
import "net/url"
import "strconv"
import "bytes"
import "io"
import "io/ioutil"

type CallBack struct {}

func (c  CallBack)PreRequest(req *http.Request, ctx *callapi.RestBirdCtx) {

	data := url.Values{}

	data.Set("aa", "abcd")
	data.Set("bb", "cdefg")

	var body io.Reader
	body = bytes.NewBufferString(data.Encode())

	rc, ok := body.(io.ReadCloser)
	if !ok && body != nil {
		rc = ioutil.NopCloser(body)
	}

	req.Body = rc

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
}
