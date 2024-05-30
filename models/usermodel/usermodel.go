package usermodel

import (
	"goweb/config"
	"goweb/entities"
)

func GetAll() []entities.User {
	rows, err := config.DB.Query("SELECT * FROM users")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var users []entities.User

	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	return users
}

func Store(user entities.User) bool {
	_, err := config.DB.Exec("INSERT INTO users (name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		panic(err)
	}

	return true
}

func GetById(id int) entities.User {
	row := config.DB.QueryRow("SELECT * FROM users WHERE id = ?", id)

	user := entities.User{}
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		panic(err)
	}

	return user
}

func Update(user entities.User) bool {
	_, err := config.DB.Exec("UPDATE users SET name = ?, email = ?,  updated_at = ? WHERE id = ?", user.Name, user.Email, user.UpdatedAt, user.Id)

	if err != nil {
		panic(err)
	}

	return true
}

func Delete(id int) bool {
	_, err := config.DB.Exec("DELETE FROM users WHERE id = ?", id)

	if err != nil {
		panic(err)
	}

	return true
}

func GetByEmail(email string) entities.User {
	row := config.DB.QueryRow("SELECT * FROM users WHERE email = ?", email)

	user := entities.User{}
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		panic(err)
	}

	return user
}
