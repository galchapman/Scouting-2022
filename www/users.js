const db = require('../database')
const fs = require('fs');


function handle(req, res) {
	var index = fs.readFileSync('www/users.html', 'utf8');

	db.get_users((error, rows) => {
		if (error) {
			res.statusCode = 402;
			res.setHeader('Content-Type', 'text/plain');
			res.end("Internal server error: " + error);
		} else {
			var users = '';

			rows.forEach(row => {
				users += `<tr><td>${row.ID}</td><td onclick="SelectUser(${row.ID})">${row.NAME}</td></tr>`
			});

			res.statusCode = 200;
			res.setHeader('Content-Type', 'text/html');
			res.end(index.replace('${USERS}', users));
		}
	})
}

module.exports = {
	handle: handle
}