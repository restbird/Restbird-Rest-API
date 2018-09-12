
function HandleRequest(req, res, state) {
    console.log("hello, I am advanced Request handler");

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

    res.send('hello, I am advanced Request handler:'+ statecount.count);
}

module.exports = HandleRequest
