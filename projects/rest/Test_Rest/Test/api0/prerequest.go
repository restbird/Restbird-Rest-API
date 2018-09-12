package api
import "callapi"
import "net/http"
//import "net/url"

type CallBack struct {}

func (c  CallBack)PreRequest(req *http.Request, ctx *callapi.RestBirdCtx) {
//    newurl :="http://www.google.com/"
//	req.URL, _ = url.Parse(newurl)
//	req.Host = req.URL.Host
	
//	req.Header.Add("aa", "bb")

}
