package users

import (
	"errors"
	"fmt"
	"golangNorthwindRestApi/datasources/mysql/users_db"
	"golangNorthwindRestApi/utils/mysql_utils"
	"golangNorthwindRestApi/utils/rest_errors"
	"strings"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser                = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindByStatus           = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?"
)

func (user *User) Save() rest_errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return rest_errors.NewInternalServerError("error when tying to save user", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		return rest_errors.NewInternalServerError("error when tying to save user", errors.New("database error"))
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return rest_errors.NewInternalServerError("error when tying to save user", errors.New("database error"))
	}
	user.Id = userId
	return nil
}

func (user *User) Get() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return rest_errors.NewInternalServerError("error when tying to Get user by Id", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		return rest_errors.NewInternalServerError("error when tying to get user", errors.New("database error"))
	}
	return nil
}

func (user *User) Update() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return rest_errors.NewInternalServerError("error when tying to update user", errors.New("database error"))
	}
	defer stmt.Close()

	_, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if updateErr != nil {
		return rest_errors.NewInternalServerError("error when tying to save user", errors.New("database error"))
	}
	return nil
}

func (user *User) Delete() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to delete user", errors.New("database error"))
	}
	defer stmt.Close()

	_, deleteErr := stmt.Exec(user.Id)
	if deleteErr != nil {
		return rest_errors.NewInternalServerError("error when trying to delete user", errors.New("database error"))
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, rest_errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to search by status", errors.New("database error"))
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to search by status", errors.New("database error"))
	}

	defer rows.Close()
	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, rest_errors.NewInternalServerError("error when trying to search by status", errors.New("database error"))
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}

func (user *User) FindByEmailAndPassword() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		return rest_errors.NewInternalServerError("error when tying to find user", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, StatusActive)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRows) {
			return rest_errors.NewNotFoundError("invalid user credentials")
		}
		return rest_errors.NewInternalServerError("error when tying to find user", errors.New("database error"))
	}
	return nil
}
