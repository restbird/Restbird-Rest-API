package api

import "callapi"
import "net/http"
import "fmt"
import "io/ioutil"
import "encoding/json"
import "time"

type MyDATA struct {
	Info  string   `json:"info,omitempty"`
	Code  int      `json:"code,omitempty"`
	Array []MyCase `json:"array,omitempty"`
}

type MyCase struct {
	Name     string `json:"name,omitempty"`
	Casepath string `json:"casepath,omitempty"`
	Categoty string `json:"category,omitempty"`
	Type     int    `json:"type,omitempty"`
	Lang     int    `json:"lang,omitempty"`
}

func (c CallBack) ResponseValidate(resp *http.Response, ctx *callapi.RestBirdCtx) bool {

	//check status code
	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		return false
	}
	//Get body json data
	var body []byte
	var err error
	var data MyDATA

	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		fmt.Println("read body failed.")
		return false
	}

	if err = json.Unmarshal(body, &data); err != nil {
		fmt.Println("conver body to json failed:", err.Error())
		return false
	}

	fmt.Println(data)
	//get/set env
	hellovars := ctx.GetVars("hello")
	fmt.Println("hello var is: " + hellovars)

	ctx.SetVars("hello", "from rest api response validation")
	//get/set global env
	var mytime time.Time
	if err := callapi.GetGlobalVars("currTime", &mytime); err == nil {
		fmt.Println("currTime is: ", mytime)
	}

	mytime = time.Now()
	if err := callapi.SetGlobalVars("currTime", mytime); err != nil {
		return false
	}
	return true
}
