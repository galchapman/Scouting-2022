const fs = require('fs');
const files = fs.readdirSync('robots/');
const util = require('../util')


function handle(req, res) {
	if (req.method == "GET") {
		var prams = util.get_prams(req.url)

		if ('team' in prams && files.includes(prams.team + '.jpeg')) {
			res.statusCode = 200;
			res.setHeader('Content-Type', 'image/jpeg')
			res.end(fs.readFileSync('robots/' + prams.team + '.jpeg'));
		} else {
			// res.statusCode = 404;
			// res.end('NotFound');
			res.statusCode = 200;
			res.setHeader('Content-Type', 'image/jpeg')
			res.end(fs.readFileSync('robots/18833.jpeg'));
		}
	} else {
		res.statusCode = 400;
		res.end('ERROR')
	}
}


module.exports = {
	handle: handle
}