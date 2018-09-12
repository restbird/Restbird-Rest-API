# encoding: utf-8

def HandleRequest(req, state) :
    req.send_response(200)

    req.send_header('Content-type', 'text/html')
    req.send_header('aaaa', 'bbbb')
    req.send_header('cccc', 'dddd')
    req.end_headers()

    req.wfile.write("\nwelcome to advanced Mock Handler\n".encode())
    req.wfile.flush()
