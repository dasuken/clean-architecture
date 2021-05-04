package usecase

import "github.com/noguchidaisuke/clean-architecture/domain"

type UserRepository interface {
	Store(domain.User) (int, error)
	FindById(int) (domain.User, error)
}
