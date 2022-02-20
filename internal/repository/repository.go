package repository

import "database/sql"

type Repository interface {
	SignIn(name, pass string) (string, error)
	SignUp(name, pass string) error
}

func New(conn *sql.DB) Repository {
	return &repository{
		conn: conn,
	}
}

type repository struct {
	conn *sql.DB
}

func (r *repository) SignIn(name, pass string) (string, error) {
	var uuid string
	err := r.conn.QueryRow("select sign_in($1,$2);", name, pass).Scan(&uuid)
	return uuid, err
}

func (r *repository) SignUp(name, pass string) error {
	_, err := r.conn.Exec("call sign_up($1,$2);", name, pass)
	return err
}
