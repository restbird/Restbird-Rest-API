# encoding: utf-8
import json
import urllib

def HandleRequest(req, state) :
    allqueries = urllib.parse.parse_qs(urllib.parse.urlparse(req.path).query)

    for k, v in allqueries :
        print (k + "=" + v + "\n")

    jsonstr = json.dumps(allqueries)
    req.send_response(200)
    req.send_header('Content-type', 'Application/json')
    req.end_headers()
    req.wfile.write(jsonstr.encode())
    req.wfile.flush()
