package service

import "remember/entity"

type UserService interface {
	GetAllUsers() interface{}

	GetAllUser() []entity.User

	Registration(user *entity.User) error

	Login(user *entity.User) (string, error)

	Logout(username, token string) error

	Delete(controller, username string) error

	ChangePassword(user *entity.User, token, OldPassword, NewPassword string) error

	GetUserInfo(user *entity.User) *entity.User

	ChangeUserInfo(user *entity.User, mp map[string]any) error

	AddUser(user *entity.User) error
}
