package api

import "callapi"
import "net/http"
import "strings"

func (c  CallBack) ResponseValidate(resp *http.Response, ctx *callapi.RestBirdCtx) bool {
	if ret := strings.Contains(resp.Status, "200 OK"); ret {
		return true
	}

	return false
}
