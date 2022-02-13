package database

func (db *Database) init() error {
	_, err := db.db.Exec(`
	CREATE TABLE IF NOT EXISTS USERS (
	    ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	    NAME TEXT NOT NULL,
	    PASSWORD TEXT NOT NULL,
	    SCREEN_NAME TEXT NOT NULL,
	    ROLE TEXT NOT NULL
	)
`)
	if err != nil {
		return err
	}

	return nil
}
