
function HandleRequest(req, res, state) {
    var statecount = state.get("counter");

    if (statecount) {
        statecount.count  = statecount.count + 1;
        state.save("counter", statecount);
    } else {
        statecount = {
            count: 1
        };
        state.save("counter", statecount);
    }
    
      var myres = {
        hello: 'hello world',
        counter: statecount.count,
        array : ["aa", "bb", "cc"],
    }

    res.json(myres);

}

module.exports = HandleRequest
