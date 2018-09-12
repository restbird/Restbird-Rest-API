package api
import "callapi"
import "net/http"

func (c  CallBack) ResponseValidate(resp *http.Response, ctx *callapi.RestBirdCtx) bool {
	return true
}
