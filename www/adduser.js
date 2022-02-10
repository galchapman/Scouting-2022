const fs = require('fs');
const db = require('../database')
const { parse } = require('querystring')


function handle(req, res) {
	if (req.method === "GET") {
		res.statusCode = 200;
		res.setHeader('Content-Type', 'text/html')
		res.end(fs.readFileSync('www/adduser.html', 'utf-8'))
	} else {
		let body = '';
		req.on('data', function (data) {
			body += data.toString();

			// likly dDos
			if (body.length > 1e6)
				request.connection.destroy();
		});
		req.on('end', () => {
			var body_parsed = parse(body)
			if ('username' in body_parsed && 'password' in body_parsed && 'permission-level' in body_parsed) {
				db.insert_user(body_parsed['username'], body_parsed['password'], body_parsed['permission-level'], 
				(_, error) => {
					if (error == null) {
						res.statusCode = 303;
						res.setHeader('Location', 'adduser.html')
						res.end()
					} else {
						res.statusCode = 402;
						res.setHeader('Content-Type', 'text/plain');
						res.end("Internal server error: " + error.toString());
					}
				})
			}
		})
	}
}


module.exports = {
	handle: handle
}