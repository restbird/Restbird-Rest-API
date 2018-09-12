# encoding: utf-8

def HandleRequest(req, state) :
    mycounter = state.get("counter")
    if mycounter == None:
        mycounter = {
            'counter' :1,
        }
    else:
        mycounter['counter'] = mycounter['counter'] + 1

    state.save("counter", mycounter)

    req.send_response(200)

    req.send_header('Content-type', 'text/html')
    req.end_headers()

    req.wfile.write("\nwelcome to advanced Mock Handler\n".encode())
    req.wfile.write(str(mycounter['counter']).encode())
    req.wfile.flush()
