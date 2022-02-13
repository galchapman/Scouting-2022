package database

func (db *Database) init() error {
	_, err := db.db.Exec(`
	CREATE TABLE IF NOT EXISTS USERS (
	    ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	    NAME TEXT NOT NULL UNIQUE,
	    PASSWORD TEXT NOT NULL,
	    SCREEN_NAME TEXT NOT NULL,
	    ROLE INTEGER NOT NULL
	);

	CREATE TABLE IF NOT EXISTS GROUPS (
		TEAM INTEGER NOT NULL PRIMARY KEY,
	    NAME TEXT NOT NULL UNIQUE
	);
`)
	if err != nil {
		return err
	}

	return nil
}
