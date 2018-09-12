
function HandleRequest(req, res, state) {
    console.log("hello, I am advanced Request handler");
    
    user = req.params.user;
    
    res.send('hello, ' + user);
}

module.exports = HandleRequest
