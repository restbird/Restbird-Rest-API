package api
import "callapi"
import "net/http"
import "io/ioutil"
import	"encoding/json"
import "fmt"

type MyDATA struct {
	Aa    string `json:"aa, omitempty"`
	Bb    int `json:"bb, omitempty"`
	Cc    bool `json:"cc, omitempty"`
}

func (c  CallBack) ResponseValidate(resp *http.Response, ctx *callapi.RestBirdCtx) bool {
	var body []byte
	var err error
	var data MyDATA

	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		fmt.Println("read body failed.")
		return true
	}

	if err = json.Unmarshal(body, &data); err != nil {
		fmt.Println("conver body to json failed:", err.Error())
		return true
	}

	if data.Aa == "hello world" {
		return true
	}

	return false
}
