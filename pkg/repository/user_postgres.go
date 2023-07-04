package repository

import (
	"fmt"

	simplerestapi "github.com/2hard4me/simple-rest-api"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) Create(user simplerestapi.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", userTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserPostgres) GetAll() ([]simplerestapi.User, error) {
	var users []simplerestapi.User
	query := fmt.Sprintf("SELECT * FROM %s", userTable)
	err := r.db.Select(&users, query)

	return users, err
}

func (r *UserPostgres) GetById(id int) (simplerestapi.User, error) {
	var user simplerestapi.User
	query := fmt.Sprintf("SELECT id, name, username, password_hash FROM %s WHERE id = $1", userTable)
	
	err := r.db.Get(&user, query, id)

	return user, err
}

func (r *UserPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", userTable)
	_, err := r.db.Exec(query, id)
	
	return err
}

func (r *UserPostgres) Update(id int, input simplerestapi.UpdateUserInput) error {
	query := fmt.Sprintf("UPDATE %s SET name=$1, username=$2, password_hash=$3 WHERE id=$4", userTable)
	_, err := r.db.Exec(query, input.Name, input.Username, input.Password, id)
	
	return err
}

