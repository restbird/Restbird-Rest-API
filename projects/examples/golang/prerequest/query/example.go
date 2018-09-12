package api
import (
	 "callapi"
	 "net/http"
	 "net/url"
)

type CallBack struct {}

func (c  CallBack)PreRequest(req *http.Request, ctx *callapi.RestBirdCtx) {

	m, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		return
	}

	//will add one more value 'abcd' to parameter 'aa'
	m.Add("aa", "abcd")

	//will overwrite parameter 'bb' 's value to 'bcdef'
	m.Set("bb", "bcdef")

	//delete the parameter 'cc'
	m.Del("cc")

	req.URL.RawQuery = m.Encode()
}
