package api

import (
	"callapi"
	"fmt"

	api13 "restbird/rest/project1/case1/api13"
	api2 "restbird/rest/project2/case3/api2"

	//go script api
	api8 "restbird/rest/project3/case5/api8"

	api9 "restbird/rest/project4/case6/api9"
)

type CallBack struct{}

func (c CallBack) GoScripts(ctx *callapi.RestBirdCtx) bool {
	var err error
	var ret bool
	var msg string
	//call rest project 'project1/case1' 's api13 with default ctx(env)
	if err, ret, msg = callapi.DoHttpRequest("project1/case1", "api13", api13.CallBack{}, ctx); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ret, msg)

	//call rest project 'project2/case3' 's api2 with env "Box"
	if err, ret, msg = callapi.DoHttpRequestWithEnv("project2/case3", "api2", api2.CallBack{}, "Box"); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ret, msg)

	//call rest project 'project1/case1' 's api13 with my current ctx(env)
	callapi.DoGoScript("project3/case5", "api8", api8.CallBack{}, ctx)

	//call rest project 'project2/case3' 's api2 with env "Box"
	callapi.DoGoScriptWithEnv("project4/case6", "api9", api9.CallBack{}, "Office")

	return true
}
