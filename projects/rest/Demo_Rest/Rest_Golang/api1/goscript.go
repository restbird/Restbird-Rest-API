package api

import (
	"callapi"
	"fmt"
	api0 "restbird/rest/Demo_Rest/Rest_Golang/api0"
)

type CallBack struct{}

func (c CallBack) GoScripts(ctx *callapi.RestBirdCtx) bool {

	var err error
	var ret bool
	var msg string
	//call rest project 'Demo_Rest/Rest_Golang' 's api0 with default ctx(env)
	if err, ret, msg = callapi.DoHttpRequest("Demo_Rest/Rest_Golang", "api0", api0.CallBack{}, ctx); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ret, msg)

	//call rest project 'Demo_Rest/Rest_Golang' 's api0 with env "Demo_Rest_Env"
	if err, ret, msg = callapi.DoHttpRequestWithEnv("Demo_Rest/Rest_Golang", "api0", api0.CallBack{}, "Demo_Rest_Env"); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ret, msg)

	return true
}
