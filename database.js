const sqlite3 = require('sqlite3').verbose();
const fs = require('fs');
const bcrypt = require('bcrypt');

let db = new sqlite3.Database('database.db');

// init database
initDatabase()


function insert_user(name, password, permissions, callback) {
	bcrypt.hash(password, 4).then((hash) => {
		db.run("INSERT INTO USERS (NAME, PASSWORD, PERMISSION_LEVEL) VALUES (?, ?, ?)", [name, hash, permissions], callback)
	})
}

function login_user(name, password, successes_callback, failed_callback) {
	db.get("SELECT ID, NAME, PASSWORD, PERMISSION_LEVEL FROM USERS WHERE NAME = ?", [name], (errors, rows) => {
		if (errors) {
			failed_callback(errors)
		} else if (rows === undefined) {
			failed_callback(new Error("Incorrect username"))
		} else if (!bcrypt.compareSync(password, rows.PASSWORD)) {
			failed_callback(new Error("Incorrect password"))
		} else {
			successes_callback(rows.ID, name, rows.PERMISSION_LEVEL)
		}
	})
}

function get_users(callback){
	db.all("SELECT ID, NAME FROM USERS", [], callback)
}

function get_user_name(id, callback) {
	db.all("SELECT ID, NAME FROM USERS WHERE ID = ?", [id], callback)
}


function initDatabase() {
	db.run(
		`CREATE TABLE IF NOT EXISTS USERS (
			ID INTEGER  PRIMARY KEY AUTOINCREMENT NOT NULL,
			NAME TEXT UNIQUE NOT NULL,
			PASSWORD TEXT,
			PERMISSION_LEVEL TEXT
		)`
	)
}


module.exports = {
	insert_user: insert_user,
	get_users: get_users,
	get_user_name: get_user_name,
	login_user: login_user
}