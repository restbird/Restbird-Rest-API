# encoding: utf-8
import json

def HandleRequest(req, state) :
    data = {
        'hello': 'world',
        'count': 10,
        'array': ["aaa", "bbb", "ccc"]
    }

    jsonstr = json.dumps(data)

    req.send_response(200)

    req.send_header('Content-type', 'Application/json')
    req.end_headers()

    req.wfile.write(jsonstr.encode())
    req.wfile.flush()
