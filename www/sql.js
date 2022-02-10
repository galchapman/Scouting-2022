const fs = require('fs');
const db = require('../database')

function handle(req, res) {
	if (req.method == "GET") {
		res.statusCode = 200;
			res.setHeader('Content-Type', 'text/html')
			res.end(fs.readFileSync('www/sql.html', 'utf-8'))
	} else if (req.method == "POST") {
		let body = '';
		req.on('data', function (data) {
			body += data.toString();

			// likly dDos
			if (body.length > 1e6)
				request.connection.destroy();
		});
		req.on('end', () => {
			db.runSQL(body, (error, resualt) => {
				if (error) {
					res.statusCode = 200;
					res.setHeader('Content-Type', 'text/html');
					res.end(error.toString());
				} else {
					var text;
					if (resualt.length == 0) {
						text = "Empty";
					} else {
						resualt[0]
						let keys = Object.keys(resualt[0]);
						text = "<table><tr>"
						keys.forEach(key => {
							text += "<th>" + key + "</th>"
						});
						text += "</tr>"

						resualt.forEach(row => {
							text += "<tr>"
							keys.forEach(key => {
								text += "<td>" + row[key] + "</td>"
							})
							text += "</tr>"
						})

						text += "</table>"
					}
					res.statusCode = 200;
					res.setHeader('Content-Type', 'text/html');
					res.end(text);
				}
			})
		});
	}
}

module.exports = {
	handle: handle
}