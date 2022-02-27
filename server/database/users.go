package database

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
)

func (db *Database) NewUser(name string, password string, screenName string, role string) (User, error) {
	roleValue := parseRole(role)
	if roleValue == -1 {
		return User{}, errors.New("Invalid role: " + role)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	_, err = db.db.Exec("INSERT INTO USERS (NAME, PASSWORD, SCREEN_NAME, ROLE) VALUES ($1, $2, $3, $4)", name, hashedPassword, screenName, roleValue)
	if err != nil {
		return User{}, err
	}

	user := User{Name: name, ScreenName: name, hashedPassword: hashedPassword}

	res := db.db.QueryRow("SELECT ID FROM USERS WHERE NAME = $1", name)
	err = res.Scan(user.ID)
	if err != nil {
		return User{}, err
	}

	return user, err
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
