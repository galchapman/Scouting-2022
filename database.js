const sqlite3 = require('sqlite3').verbose();
const fs = require('fs');

let db = new sqlite3.Database('database.db');

// init database
initDatabase()


function insert_user(name, password) {
	db.run("INSERT INTO USERS (name) VALUES (?)", [name])
}

function get_users(callback){
	db.all("SELECT ID, NAME FROM USERS", [], callback)
}


function initDatabase() {
	db.run(
		`CREATE TABLE IF NOT EXISTS USERS (
			ID INTEGER  PRIMARY KEY AUTOINCREMENT NOT NULL,
			NAME TEXT UNIQUE NOT NULL,
			PASSWORD TEXT NOT NULL
		)`
	)
}


module.exports = {
	insert_user: insert_user,
	get_users: get_users
}