const util = require('../util')
const sessions = require('../sessions')


function mint(user_id) {
	return sessions.new_session(user_id);
}


function handle(req, res) {

	var prams = util.get_prams(req.url)

	if ('token' in prams) {
		res.statusCode = 303;
		res.setHeader('Set-Cookie', `session=${prams.token}`)
		res.setHeader('Location', '/')
		res.end('')
	} else if ('mint' in prams) {
		res.statusCode = 200;
		res.setHeader('Content-Type', 'text/plain')
		res.end(mint(prams.mint))
	} else {
		res.statusCode = 200;
		res.setHeader('Content-Type', 'text/html')
		res.end('AAA')
	}
}


module.exports = {
	handle: handle
}