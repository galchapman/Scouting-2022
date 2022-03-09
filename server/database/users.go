package database

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
)

func (db *Database) NewUser(name string, password string, screenName string, role string) error {
	roleValue := parseRole(role)
	if roleValue == -1 {
		return errors.New("Invalid role: " + role)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.db.Exec("INSERT INTO USERS (NAME, PASSWORD, SCREEN_NAME, ROLE) VALUES ($1, $2, $3, $4)", name, hashedPassword, screenName, roleValue)
	if err != nil {
		return err
	}

	return err
}

func (db *Database) GetUser(ID int) (User, error) {
	var user User
	user = User{ID: ID}

	res := db.db.QueryRow("SELECT NAME, PASSWORD, SCREEN_NAME, ROLE FROM USERS WHERE ID = $1", ID)

	err := res.Scan(&user.Name, &user.hashedPassword, &user.ScreenName, &user.Role)
	if err != nil {
		return user, err
	}
	return user, err
}

func (db *Database) GetUserByName(name string) (User, error) {
	var user User
	user = User{Name: name}

	res := db.db.QueryRow("SELECT ID, PASSWORD, SCREEN_NAME, ROLE FROM USERS WHERE NAME = $1", name)

	err := res.Scan(&user.ID, &user.hashedPassword, &user.ScreenName, &user.Role)
	if err != nil {
		return user, err
	}
	return user, err
}

func (user *User) TryLoggingIn(password string) error {
	return bcrypt.CompareHashAndPassword(user.hashedPassword, []byte(password))
}

func (db *Database) GetScouters() ([]User, error) {
	var users []User

	rows, err := db.db.Query("SELECT ID, NAME, PASSWORD, SCREEN_NAME, ROLE FROM USERS WHERE ROLE != $1 AND ROLE != $2", AdminRole, ManagerRole)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.hashedPassword, &user.ScreenName, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (db *Database) GetNextGame(scouter User, currentGame int) (Game, error) {
	var gameID int

	row := db.db.QueryRow("SELECT ID FROM GAMES WHERE ID > $1 AND (RED_SCOUTER1 = $2 OR RED_SCOUTER2 = $2 OR BLUE_SCOUTER1 = $2 OR BLUE_SCOUTER2 = $2) ORDER BY ID LIMIT 1", currentGame, scouter.ID)

	err := row.Scan(&gameID)
	if err != nil {
		return Game{}, err
	}

	return db.GetGame(gameID)
}
