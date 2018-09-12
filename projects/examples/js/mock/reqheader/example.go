
function HandleRequest(req, res, state) {
    var ctype = req.get('Content-Type');
    var host = req.get('Host');


    res.send('hello, I am advanced Request handler, ' + 'content-type is :' + ctype + ', host is :' + host);
}

module.exports = HandleRequest
