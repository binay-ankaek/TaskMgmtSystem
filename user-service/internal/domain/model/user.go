package model

type User struct {
	ID       string `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Password string `db:"password" json:"password"`
	Email    string `db:"email" json:"email"`
	Address  string `db:"address" json:"address"`
}
