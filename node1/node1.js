const http = require('http');
const os = require('os');
http.createServer((req, res)=>{
    res.writeHead(200, {'Content-Type': 'text/plain'})
    res.end(" hello iam node 1, hostname:  "+ os.hostname);
}).listen(8085)
console.log('created server listen 8085');
