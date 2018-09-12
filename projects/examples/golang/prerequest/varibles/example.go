package api

import "callapi"
import "net/http"
import "fmt"
import "time"

type CallBack struct {}

func (c  CallBack)PreRequest(req *http.Request, ctx *callapi.RestBirdCtx) {

	hellovars := ctx.GetVars("hello")
	fmt.Println("hello var is: " + hellovars)

	ctx.SetVars("hello", "I am not world")

	var mytime  time.Time
	if err := callapi.GetGlobalVars("currTime", &mytime); err != nil {
		return err
	}

	mytime =  time.Now()
	if err := callapi.SetGlobalVars("currTime", mytime); err != nil {
		return err
	}

	if err, mytoken := callapi.GetGlobalString("mytoken"); err != nil {
		return err
	 } else {
		callapi.SetGlobalString("mytoken", "hello world token")
	}

}
