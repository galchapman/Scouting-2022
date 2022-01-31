const sessions = require("../sessions.js");
const fs = require('fs');
const util = require('../util.js');

function handle(req, res) {
	var index = fs.readFileSync('www/index.html', 'utf8');
	var body = '';
	var cookies = util.get_cookies(req.headers.cookie)

	if (!req.headers.cookie) {
		body = "You got no cookie :(";
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