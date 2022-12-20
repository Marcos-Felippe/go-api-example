package repository

import (
	"database/sql"

	"github.com/projetosgo/exemploapi/entity"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (r *UserRepository) Save(user *entity.User) error {
	stmt, err := r.Db.Prepare("INSERT INTO users (id, name, email) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.ID, user.Name, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("SELECT count(*) FROM users").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *UserRepository) GetOne(id string) (*entity.User, error) {
	var user entity.User

	row := r.Db.QueryRow("SELECT * FROM users WHERE id = ?", id)
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAll() ([]entity.User, error) {
	var users []entity.User

	rows, err := r.Db.Query("SELECT * FROM users")
	if err != nil {
		return users, err
	}

	for rows.Next() {

		var user entity.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) Update(user *entity.User) (*entity.User, error) {
	stmt, err := r.Db.Prepare("UPDATE users SET name = ?, email = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(user.Name, user.Email, user.ID)
	if err != nil {
		return nil, err
	}

	user, err = r.GetOne(user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Delete(id string) error {
	stmt, err := r.Db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
