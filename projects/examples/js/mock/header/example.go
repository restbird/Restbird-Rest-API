
function HandleRequest(req, res, state) {
    res.append('hello', 'world');
    res.append('Link', ['<http://localhost/>', '<http://localhost:3000/>']);
    res.append('Set-Cookie', 'foo=bar; Path=/; HttpOnly');
    res.append('Warning', '199 Miscellaneous warning');

    res.send('hello, I am advanced Request handler');
}

module.exports = HandleRequest
