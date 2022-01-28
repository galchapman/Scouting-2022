//Load HTTP module
const http = require("http");
const fs = require('fs');
const { parse } = require('querystring');
const db = require('./database.js')

const urls = {'index.html': get_file, 'main.css': get_file, 'users.html': get_users}
const redirects = {'': 'index.html'}

const hostname = '127.0.0.1';
const port = 3000;


function read_file(filename) {
	return fs.readFileSync(filename, 'utf8')
}

function handle_post_request(url, data, responce) {
	db.insert_user(data.name)

	responce.statusCode = 200;
	responce.setHeader('Content-Type', 'text/html');
	responce.end(read_file('www/' + url));
}

function get_file(req, res) {
	res.statusCode = 200;
	if (req.url.endsWith('.html')) {
		res.setHeader('Content-Type', 'text/html');
	} else if (req.url.endsWith('.css')) {
		res.setHeader('Content-Type', 'text/css');
	} else {
		res.setHeader('Content-Type', 'text/plain');
	}
	res.end(read_file('www/' + req.url))
}

function get_users(req, res) {
	db.get_users((error, rows) => {
		if (error) {
			res.statusCode = 402;
			res.setHeader('Content-Type', 'text/plain');
			res.end("Internal server error: " + error);
		} else {
			res.statusCode = 200;
			res.setHeader('Content-Type', 'text/html');
			console.log(rows);

			let message = '<!DOCTYPE HTML><html><head><link rel="stylesheet" href="main.css"></head>';

			message += "<body><table class=\"fl-table\"><tr><th>ID</th><th>Name</th></tr>"

			rows.forEach(row => {
				message += "<tr><td>" + row.ID + "</td><td>" + row.NAME + "</td></tr>"
			});
			message += "</table></body></html>";


			res.end(message)
		}
	})
}

//Create HTTP server and listen on port 3000 for requests
const server = http.createServer((req, res) => {
	// parse url
	if (req.url[0] == '/') // remove leading '/'
		req.url = req.url.substring(1);
	// redirect url
	if (req.url in redirects) {
		req.url = redirects[req.url]
	// if url isn't valid
	} else if (!(req.url in urls)) {
		console.log("new message " + req.method + "@" + req.url)
		res.statusCode = 404;
		res.setHeader('Content-Type', 'text/html')
		res.end('Not Found :(')
		return
	}

	// Simple Read
	if (req.method == "GET") {
		urls[req.url](req, res)
	// post data
	} else if (req.method == "POST") {
		var body = '';

		req.on('data', function (data) {
			body += data.toString();

			// likly dDos
			if (body.length > 1e6)
				request.connection.destroy();
		});
		req.on('end', () => {
			handle_post_request(req.url, parse(body), res)
		})
	}
});

//listen for request on port 3000, and as a callback function have the port listened on logged
server.listen(port, hostname, () => {
	console.log(`Server running at http://${hostname}:${port}/index.html`);
});