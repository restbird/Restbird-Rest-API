package api
import "callapi"
import "net/http"
import "io/ioutil"
import "strings"
import "fmt"

func (c  CallBack) ResponseValidate(resp *http.Response, ctx *callapi.RestBirdCtx) bool {
	var data string
	var body []byte
	var err error

	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		fmt.Println("read body failed.")
		return true
	}

	data = string(body)
	if ret := strings.Contains(data, "hello world"); ret {
		return true
	}

	fmt.Println(data)

	return false
}
