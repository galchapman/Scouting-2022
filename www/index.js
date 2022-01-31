const sessions = require("../sessions.js");
const fs = require('fs');
const util = require('../util.js');

function handle(req, res) {
	var index = fs.readFileSync('www/index.html', 'utf8');
	var body = '';
	var cookies = util.get_cookies(req.headers.cookie)

	if (!req.headers.cookie) {
		body = "You got no cookie :(";
	} else if (sessions.get_seesion(cookies['session']) === undefined) {
		body = `Your session in no longer valid. please update it`
		res.setHeader('Set-Cookie', 'session=deleted; path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT')
	} else {
		body = `You loged in as ${sessions.get_seesion(cookies['session']).name}`
	}

	res.statusCode = 200;
	res.setHeader('Content-Type', 'text/html');
	res.end(index.replace('${body}', body));
}


module.exports = {
	handle: handle
}