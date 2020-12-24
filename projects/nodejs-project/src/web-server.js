const http = require("http")
const host = 'localhost'
const port = 8002

const requestHandler = (req, res) => {
    res.end("i am working!")
}

const server = http.createServer(requestHandler)
server.listen(port, (e) => console.log(e ? e : ""))