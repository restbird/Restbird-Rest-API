package main
import "callapi"
import "fmt"
import "time"

import api4 "restbird/Restbird_API/RestProject/api4"
import api5 "restbird/Restbird_API/RestProject/api5"

func main() {
		for {
			select {
			case <-time.Tick(time.Millisecond * 60000 * 10): //timer runs every 10 minutes
				fmt.Println("hello, curren time is :", time.Now())
				
				callapi.DoHttpRequestWithEnv("Restbird_API/RestProject", "api4", api4.CallBack{},  "Restbird")
				
				callapi.DoHttpRequestWithEnv("Restbird_API/RestProject", "api5", api5.CallBack{},  "Restbird")
				
	//			ctx := callapi.DefaultRestBirdCtx
	//			callapi.DoHttpRequest("CASB/test", "api0", api0.CallBack{}, &ctx)
			}
		}

}

