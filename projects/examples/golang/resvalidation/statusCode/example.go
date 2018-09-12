package api

import "callapi"
import "net/http"

func (c  CallBack) ResponseValidate(resp *http.Response, ctx *callapi.RestBirdCtx) bool {
	if resp.StatusCode == 200 {
		return true
	}

	return false
}
