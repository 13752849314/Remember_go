package service

import "remember/entity"

type UserService interface {
	GetAllUsers() interface{}
	Registration(user *entity.User) error
	Login(user *entity.User) error
}
