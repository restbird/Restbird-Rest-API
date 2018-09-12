# encoding: utf-8
import json

# GET /hello/:group/:name
def HandleRequest(req, state) :
    echoparams = {}

    print (req.params)
    print (req.params.get('group', None))
    print (req.params.get('name', None))

    for k, v in req.params.items():
        echoparams[k] = v


    req.send_response(200)

    req.send_header('Content-type', 'Application/json')
    req.end_headers()

    jsonstr = json.dumps(echoparams)
    req.wfile.write(jsonstr.encode())
    req.wfile.flush()
