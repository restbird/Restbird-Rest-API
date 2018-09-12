package api
import (
	 "callapi"
	 "net/http"
)



type CallBack struct {}

func (c  CallBack)PreRequest(req *http.Request, ctx *callapi.RestBirdCtx) {

	//will add one more value 'abcd' to header 'aa'
	req.Header.Add("aa", "abcd")

	//will overwrite or set header 'cc' 's value to 'cdef'
	req.Header.Set("cc", "cdef")

	//delete header 'dd'
	req.Header.Del("dd")
}
