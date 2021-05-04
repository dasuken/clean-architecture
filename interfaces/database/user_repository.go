package database

import (
	"github.com/noguchidaisuke/clean-architecture/domain"
)

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(u domain.User) (id int, err error) {
	result, err := repo.Execute("INSERT INTO users(name) VALUES ('Jamal')")
	if err != nil {
		return
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}

func (repo *UserRepository) FindById(userId int) (user domain.User, err error) {
	row, err := repo.Query("SELECT * FROM users WHERE id = ?", userId)
	defer row.Close()
	if err != nil {
		return
	}

	var id int
	var name string
	var createdAt string
	row.Next()
	if err = row.Scan(&id, &name, &createdAt); err != nil {
		return
	}

	user.Id = id
	user.Name = name
	user.CreatedAt = createdAt

	return
}