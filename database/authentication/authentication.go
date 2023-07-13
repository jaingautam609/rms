package authentication

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"rms/models"
)

//todo: use hash password
func Login(db *sqlx.DB, email string, password string) (int, error) {
	var user models.Users
	SQL := `SELECT id, password FROM users WHERE user_email = $1`
	err := db.Get(&user, SQL, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return user.Id, errors.New("user not found")
		}
		return user.Id, err
	}
	if user.Password != password {
		return user.Id, errors.New("incorrect password")
	}
	return user.Id, nil
}

func Create(db *sqlx.DB, createdBy int, addUser models.AddUser) (int, error) {
	SQL := `insert into users(user_name,user_email,password,created_by) values ($1,$2,$3,$4) returning id`
	var userId int
	err := db.QueryRowx(SQL, addUser.Name, addUser.Email, addUser.Password, createdBy).Scan(&userId)
	if err != nil {
		return userId, err
	}
	return userId, nil
}
func AddRole(db *sqlx.DB, userId int, role string) error {
	SQL := `insert into user_role(user_id,user_type) values ($1,$2)`
	_, err := db.Exec(SQL, userId, role)
	if err != nil {
		return err
	}
	return nil
}
func AddAddress(db *sqlx.DB, userId int, addUser models.AddUser) error {
	SQL := `insert into address(user_id, latitude, longitude, location) values($1,$2,$3,$4)`
	_, err := db.Exec(SQL, userId, addUser.Latitude, addUser.Longitude, addUser.Location)
	if err != nil {
		return err
	}
	return nil
}
func ValidateAdmin(db *sqlx.DB, userId int) (bool, error) {
	var userType string
	SQL := `select user_type from user_role where user_id=$1 and user_type='admin'`

	//err := db.Get(&userType, SQL, userId)
	err := db.QueryRowx(SQL, userId).Scan(&userType)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
