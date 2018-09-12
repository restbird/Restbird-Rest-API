package api

import "callapi"
import "net/http"
import "strings"

func (c  CallBack) ResponseValidate(resp *http.Response, ctx *callapi.RestBirdCtx) bool {
	var  valueofaa string

	// check if header 'aa' is exist or not
	if vv, exist := resp.Header['aa']; !exist {
		return false
	}

	//check if heaer 'aa' 's value contains 'hello world'
	for k, vv := range resp.Header {
		for _, v := range vv {
			if k == "aa" {
				if ret := strings.Contains(v, "hello world"); ret {
					return true
				}
			}
		}
	}


	return false
}
