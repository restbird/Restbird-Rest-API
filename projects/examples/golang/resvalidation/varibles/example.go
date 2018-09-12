package api

import "callapi"
import "net/http"
import "fmt"
import "time"

func (c  CallBack) ResponseValidate(resp *http.Response, ctx *callapi.RestBirdCtx) bool {

	hellovars := ctx.GetVars("hello")
	fmt.Println("hello var is: " + hellovars)

	ctx.SetVars("hello", "I am not world")

	var mytime  time.Time
	if err := callapi.GetGlobalVars("currTime", &mytime); err != nil {
		return  true
	}

	mytime =  time.Now()
	if err := callapi.SetGlobalVars("currTime", mytime); err != nil {
		return true
	}

	if err, mytoken := callapi.GetGlobalString("mytoken"); err != nil {
		return true
	 } else {
		callapi.SetGlobalString("mytoken", "hello world token")
	}

	return true
}
