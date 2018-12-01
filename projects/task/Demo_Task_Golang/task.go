package main

import (
	"callapi"
	"fmt"

	api0 "restbird/rest/Demo_Rest/Rest_Golang/api0"
	api1 "restbird/rest/Demo_Rest/Rest_Golang/api1"
)

func main() {
	// for {
	// select {
	// case <-time.Tick(time.Millisecond * 5000):
	// 	fmt.Println("hello, curren time is :", time.Now())

	var err error
	var ret bool
	var msg string
	//call rest project 'Demo_Rest/Rest_Golang' 's api0 with env "Demo_Rest_Env"
	if err, ret, msg = callapi.DoHttpRequestWithEnv("Demo_Rest/Rest_Golang", "api0", api0.CallBack{}, "Demo_Rest_Env"); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ret, msg)

	ctx := &callapi.RestBirdCtx{}
	*ctx = callapi.DefaultRestBirdCtx

	//call rest project  'Demo_Rest/Rest_Golang' 's api11 with default ctx(env)
	if err, ret, msg = 	callapi.DoGoScript( "Demo_Rest/Rest_Golang", "api1", api1.CallBack{}, ctx); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ret, msg)

	// 	}
	// }

}
