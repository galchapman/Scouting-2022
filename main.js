//Load HTTP module
const http = require("http");
const fs = require('fs');
const sessions = require('./sessions.js')
const util = require('./util.js');
const adduser = require("./www/adduser.js");

const urls = {'index': get_js_file, 'main.css': get_file, 'users': get_js_file, 'login': get_js_file, 'adduser': get_js_file, 'favicon.ico': get_file_bin, 'robots': get_js_file, 'sql': get_js_file, 'scout-form.html': get_file,'field-form.html': get_file,}
const redirects = {'': 'index', 'index.html': 'index', 'users.html': 'users', 'adduser.html': 'adduser', 'sql.html': 'sql'}

const permissions = {
	'users': 'Manager',
	'adduser': 'Admin',
	'sql': 'Admin'
}

// const hostname = '127.0.0.1';
const hostname = '10.77.76.10';
const port = 3000;


function read_file_binary(filename) {
	return fs.readFileSync(filename)
}


function read_file(filename) {
	return fs.readFileSync(filename, 'utf8')
}

let loaded_files = {}
function load_file(filename) {
	// if (!(filename in loaded_files)) {
		loaded_files[filename] = read_file('www/' + filename)
	// }
	return loaded_files[filename]
}

function load_file_bin(filename) {
	// if (!(filename in loaded_files)) {
		loaded_files[filename] = read_file_binary('www/' + filename)
	// }
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

function get_file_bin(req, res) {
	var page = req.url.split('?')[0];
	res.statusCode = 200;
	if (page.endsWith('.ico')) {
		res.setHeader('Content-Type', 'image/x-icon')
	} else {
		res.setHeader('Content-Type', 'text/plain');
	}
	// TODO: Add cashing
	res.end(load_file_bin(req.url.split('?')[0]))
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


function check_permission(permission, req) {
	if (process.argv.length > 2 && 'test' == process.argv[2])
		return true;

	var cookies = util.get_cookies(req.headers.cookie)
	var session = sessions.get_seesion(cookies["session"])
	if (permission == undefined) {
		return true;
	} else if (session == undefined) {
		return false;
	} else if (session.permission_level == "Admin") {
		return true;
	} else {
		return session.permission_level == permission;
	}
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

		// check permissions
		if (check_permission(permissions[page], req)) {
			urls[page](req, res)
		} else {
			res.statusCode = 403;
			res.setHeader('Content-Type', 'text/plain')
			res.end("You don't have access to this page")
		}
	} catch (error) {
		console.log(error)
	}
});

//listen for request on port 3000, and as a callback function have the port listened on logged
server.listen(port, hostname, () => {
	console.log(`Server running at http://${hostname}:${port}/index.html`);
});
