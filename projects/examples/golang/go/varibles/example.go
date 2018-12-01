package api

import "callapi"
import "fmt"
import "time"

type CallBack struct {}

func (c  CallBack) GoScripts(ctx *callapi.RestBirdCtx) bool {

	hellovars := ctx.GetVars("hello")
	fmt.Println("hello var is: " + hellovars)

	ctx.SetVars("hello", "I am not world")

	var mytime  time.Time
	if err := callapi.GetGlobalVars("currTime", &mytime); err != nil {
		return false
	}

	mytime =  time.Now()
	if err := callapi.SetGlobalVars("currTime", mytime); err != nil {
		return false
	}

	if err, mytoken := callapi.GetGlobalString("mytoken"); err != nil {
		return false
	 } else {
		fmt.Println(mytoken)
		callapi.SetGlobalString("mytoken", "hello world token")
	}

	return true
}
