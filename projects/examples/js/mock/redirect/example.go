
function HandleRequest(req, res, state) {
    res.redirect('http://www.google.com');
    //res.redirect(301, 'http://www.yahoo.com');
}

module.exports = HandleRequest
