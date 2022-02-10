const util = require('../util')
const sessions = require('../sessions')
const db = require('../database')
const { parse } = require('querystring');
const fs = require('fs');


function mint(user_id) {
	return sessions.new_session(user_id);
}


function handle(req, res) {
	if (req.method === "GET") {
		var prams = util.get_prams(req.url)

		if ('token' in prams) {
			res.statusCode = 303;
			res.setHeader('Set-Cookie', `session=${prams.token}`)
			res.setHeader('Location', '/')
			res.end()
		} else if ('mint' in prams) {
			res.statusCode = 200;
			res.setHeader('Content-Type', 'text/plain')
			res.end(mint(prams.mint))
		} else {
			res.statusCode = 200;
			res.setHeader('Content-Type', 'text/html')
			res.end(fs.readFileSync('www/login.html', 'utf-8'))
		}
	} else if (req.method === "POST") {
		let body = '';
		req.on('data', function (data) {
			body += data.toString();

			// likly dDos
			if (body.length > 1e6)
				request.connection.destroy();
		});
		req.on('end', () => {
			var body_parsed = parse(body)
			if ('username' in body_parsed && 'password' in body_parsed) {
				db.login_user(body_parsed['username'], body_parsed['password'], (id, name, permissions_level) => {
					res.statusCode = 303;
					res.setHeader('Set-Cookie', `session=${sessions.login(id, name, permissions_level)}`)
					res.setHeader('Location', '/')
					res.end()
				}, (error) => {
					res.statusCode = 402;
					res.setHeader('Content-Type', 'text/plain');
					res.end("Internal server error: " + error.toString());
				})
			} else {
				res.statusCode = 400;
				res.setHeader('Content-Type', 'text/plain');
				res.end("Your bad");
			}
		})
	}
}


module.exports = {
	handle: handle
}