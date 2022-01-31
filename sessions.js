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
		} catch {
			sessions[session] = undefined
		}
	})

	return session
}


function get_seesion(session) {
	return sessions[session]
}

module.exports = {
	new_session: new_session,
	get_seesion: get_seesion
}