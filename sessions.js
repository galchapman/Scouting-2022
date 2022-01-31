const db = require('./database')

var sessions = {}

function new_session(id) {
	var session = Math.random().toString(36).substring(2) + Math.random().toString(36).substring(2)

	sessions[session] = new Object

	db.get_user_name(id, (err, rows) => {
		try {
			sessions[session] = new Object
			sessions[session].name = rows[0].NAME
			sessions[session].id = rows[0].ID
			sessions[session].permission_level = undefined
		} catch {
			sessions[session] = undefined
		}
	})

	return session
}


function login(id, name, permission_level) {
	var session = Math.random().toString(36).substring(2) + Math.random().toString(36).substring(2)
	
	sessions[session] = new Object
	sessions[session].id = id
	sessions[session].name = name
	sessions[session].permission_level = permission_level

	return session
}


function get_seesion(session) {
	return sessions[session]
}

module.exports = {
	new_session: new_session,
	get_seesion: get_seesion,
	login: login
}