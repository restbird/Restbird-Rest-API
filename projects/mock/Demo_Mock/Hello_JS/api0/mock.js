
function HandleRequest(req, res, state) {
    console.log("hello, I am advanced Request handler");
    res.send('hello, I am advanced Request handler');
}

module.exports = HandleRequest
