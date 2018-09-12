
function HandleRequest(req, res, state) {
    res.status(500).json({ error: 'I am encounter a error' });
}

module.exports = HandleRequest
