
function HandleRequest(req, res, state) {
    var myres = {
        hello: 'hello world',
        counter: 100,
        array : ["aa", "bb", "cc"],
    }

    res.json(myres);
}

module.exports = HandleRequest
