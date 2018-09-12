package api
import "callapi"
import "net/http"
import "io/ioutil"
import	"encoding/json"
import "fmt"

type MyDATA struct {
	Code    int `json:"code,omitempty"`
	Info    string `json:"info,omitempty"`
	His     MyHistory `json:"his,omitempty"`
}
type MyHistory struct {
    Id    string `json:"id,omitempty"`
}
func (c  CallBack) ResponseValidate(resp *http.Response, ctx *callapi.RestBirdCtx) bool {
    var body []byte
    var err error
    var data MyDATA = MyDATA{}
    
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
	    fmt.Println("read body failed.")
		return true
	} 
	if err = json.Unmarshal(body, &data); err != nil {
	    fmt.Println("conver body to json failed")
		return true
	}	
	
	ctx.SetVars("rest_history_id", data.His.Id)
	
	fmt.Println("rest_history_id: " + ctx.GetVars("rest_history_id"))
	
	return true
}