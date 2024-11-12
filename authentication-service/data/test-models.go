package data

import (
	"database/sql"
	"time"
)

type PostgresTestRepository struct {
	Conn *sql.DB
}

func NewPostgresTestRepository(db *sql.DB) *PostgresTestRepository {
	return &PostgresTestRepository{
		Conn: db,
	}
}

// GetAll akan mengembalikan semua data semua user yang terdaftar
func (u *PostgresTestRepository) GetAll() ([]*User, error) {
	users := []*User{}

	return users, nil
}

// GetByEmail akan mengembalikan data satu user dengan email yang diminta
func (u *PostgresTestRepository) GetByEmail(email string) (*User, error) {
	user := User{
		ID:        1,
		FirstName: "First",
		LastName:  "Last",
		Email:     "me@here.com",
		Password:  "",
		Active:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &user, nil
}

func (u *PostgresTestRepository) GetOne(id int) (*User, error) {
	user := User{
		ID:        1,
		FirstName: "First",
		LastName:  "Last",
		Email:     "me@here.com",
		Password:  "",
		Active:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &user, nil
}

// Update akan melakukan perubahan data user berdasarkan data yang ada pada instance user sekarang
func (u *PostgresTestRepository) Update(user User) error {
	return nil
}

func (u *PostgresTestRepository) Insert(user User) (int, error) {
	return 2, nil
}

// DeleteByID akan menghapus data user pada database berdasarkan id yang diberikan
func (u *PostgresTestRepository) DeleteByID(id int) error {
	return nil
}

// ResetPassword akan mengubah password dari user berdasarkan User.ID
func (u *PostgresTestRepository) ResetPassword(password string, user User) error {
	return nil
}

// PasswordMatches akan mengembalikan hasil perbandingan
// antara password yang diberikan oleh user dengan password
// yang telah tersimpan pada user terkait
func (u *PostgresTestRepository) PasswordMatches(plainText string, user User) (bool, error) {
	return true, nil
}
