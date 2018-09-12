package api
import (
	 "callapi"
	 "net/http"
	 "net/url"
)


type CallBack struct {}

//An example to modify part of URL
// if you need to modify the path of url or use a new url, please follow url replace example
func (c  CallBack)PreRequest(req *http.Request, ctx *callapi.RestBirdCtx) {

	// change method to any other stand HTTP method
	req.URL.Method = "PUT"

	// change scheme to http or https
	req.URL.Scheme = "https"

	// change the host name, make sure to chang req.Host too
	req.URL.Host =  "google.com"
	req.Host = "google.com"
}
