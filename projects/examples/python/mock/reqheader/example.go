# encoding: utf-8
import json

def HandleRequest(req, state) :
    echoheaders = {}

    print (req.headers.get('Host', None))
    print (req.headers.get('Content-Type', None))

    for k, v in req.headers.items() :
        echoheaders[k] = v

    req.send_response(200)

    req.send_header('Content-type', 'Application/json')
    req.end_headers()

    jsonstr = json.dumps(echoheaders)
    req.wfile.write(jsonstr.encode())
    req.wfile.flush()
