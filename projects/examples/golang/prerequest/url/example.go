package api
import (
	 "callapi"
	 "net/http"
	 "net/url"
)

type CallBack struct {}

func (c  CallBack)PreRequest(req *http.Request, ctx *callapi.RestBirdCtx) {

	newurl :="http://www.google.com/"
	if u, err := url.Parse(newurl) ; err != nil {
		return
	} else {
		req.URL = u
		req.Host = u.Host
	}
}
