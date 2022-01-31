function get_prams(url) {
	if (url.split('?').length == 1) {
		return {}
	}
	var parms = url.split('?')[1].split('&')

	var dict = {};
	parms.forEach(element => {
		dict[element.split('=')[0]] = element.split('=')[1]
	});

	return dict
}

function get_cookies(cookies_string) {
	var cookies = {};
	if (cookies_string != undefined) {
		cookies_string.split(';').forEach(cookie => {
			if (cookie[0] == ' ') {
				cookie = cookie.substring(1)
			}

			var parts = cookie.split('=');
			cookies[parts[0]] = parts[1];
		});
	}
	return cookies
}


module.exports = {
	get_prams: get_prams,
	get_cookies: get_cookies
}