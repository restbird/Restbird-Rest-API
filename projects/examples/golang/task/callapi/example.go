package main
import "callapi"

// rest api
import api13 "restbird/project1/case1/api13"
import api2 "restbird/project2/case3/api2"

//go script api
import api8 "restbird/project3/case5/api8"
import api9 "restbird/project4/case6/api9"

func main() {
	ctx :=  &callapi.RestBirdCtx{}
	*ctx = callapi.DefaultRestBirdCtx

	//call rest project 'project1/case1' 's api13 with default ctx(env) 
	callapi.DoHttpRequest("project1/case1", "api13", api13.CallBack{}, ctx)

	//call rest project 'project2/case3' 's api2 with env "Box"
	callapi.DoHttpRequestWithEnv("project2/case3", "api2", api2.CallBack{},  "Box")

	//call rest project 'project1/case1' 's api13 with default ctx(env) 
	callapi.DoGoScript("project3/case5", "api8", api8.CallBack{}, ctx)

	//call rest project 'project2/case3' 's api2 with env "Box"
	callapi.DoGoScriptWithEnv("project4/case6", "api9", api9.CallBack{},  "Office")
}
