package service

import "remember/entity"

type UserService interface {
	GetAllUsers() []entity.User
}
