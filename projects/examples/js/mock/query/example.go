
function HandleRequest(req, res, state) {
    // GET /search?q=tobi+ferret
    console.log(req.query.q);
    // => "tobi ferret"

    // GET /shoes?order=desc&shoe[color]=blue&shoe[type]=converse
    console.log(req.query.order);
    // => "desc"

    console.log(req.query.shoe.color);
    // => "blue"

    console.log(req.query.shoe.type);
    // => "converse"

    res.send('hello, I am advanced Request handler');
}

module.exports = HandleRequest
