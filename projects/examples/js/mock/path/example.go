// GET /hello/:group/:name/
function HandleRequest(req, res, state) {
    console.log(req.params.group);
    console.log(req.params.name);

    res.send("group is: " + req.params.group + ", name is:" + req.params.name);
}

module.exports = HandleRequest
