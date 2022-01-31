//Load HTTP module
const http = require("http");
const fs = require('fs');
const { parse } = require('querystring');
const db = require('./database.js')

const urls = {'index': get_js_file, 'main.css': get_file, 'users': get_js_file, 'login': get_js_file}
const redirects = {'': 'index', 'index.html': 'index', 'users.html': 'users'}

// const hostname = '127.0.0.1';
const hostname = '192.168.1.72';
const port = 3000;


function read_file(filename) {
	return fs.readFileSync(filename, 'utf8')
}

function handle_post_request(url, data, response) {
	db.insert_user(data.name)

	response.statusCode = 200;
	response.setHeader('Content-Type', 'text/html');
	response.end(read_file('www/' + url));
}

let loaded_files = {}
function load_file(filename) {
	if (!(filename in loaded_files)) {
		loaded_files[filename] = read_file('www/' + filename)
	}
	return loaded_files[filename]
}

function get_file(req, res) {
	var page = req.url.split('?')[0];
	res.statusCode = 200;
	if (page.endsWith('.html')) {
		res.setHeader('Content-Type', 'text/html');
	} else if (page.endsWith('.css')) {
		res.setHeader('Content-Type', 'text/css');
	} else {
		res.setHeader('Content-Type', 'text/plain');
	}
	// TODO: Add cashing
	res.end(load_file(req.url.split('?')[0]))
}

let loaded_modules = {}
function load_module(name) {
	if (!(name in load_module)) {
		loaded_modules[name] = require(`./www/${name}.js`)
	}
	return loaded_modules[name]
}

function get_js_file(req, res) {
	load_module(req.url.split('?')[0]).handle(req, res)
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

			let message = '<!DOCTYPE HTML><html><head><link rel="stylesheet" href="main.css"><title>Users</title></head>';

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
	try {
		// parse url
		var page = req.url.split('?')[0];
		if (page[0] == '/') { // remove leading '/'
			req.url = req.url.substring(1);
			page = page.substring(1);
		}
		// redirect url
		if (page in redirects) {
			page = redirects[page]
			if (req.url.split('?').length == 2)
				req.url = page + req.url.split('?')[1]
			else
				req.url = page
		// if url isn't valid
		} else if (!(page in urls)) {
			console.log("new message " + req.method + "@" + page)
			res.statusCode = 404;
			res.setHeader('Content-Type', 'text/html')
			res.end('Not Found :(')
			return
		}

		// Simple Read
		if (req.method === "GET") {
			urls[page](req, res)
		// post data
		} else if (req.method === "POST") {
			let body = '';

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
	} catch (error) {
		console.log(error)
	}
});

//listen for request on port 3000, and as a callback function have the port listened on logged
server.listen(port, hostname, () => {
	console.log(`Server running at http://${hostname}:${port}/index.html`);
});