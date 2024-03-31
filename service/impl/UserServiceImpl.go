package impl

import (
	"remember/mapper"
)

type UserServiceImpl struct {
	mapper mapper.UserMapper
}

func (u *UserServiceImpl) GetAllUsers() interface{} {
	users := u.mapper.Select("getAllUsers")
	return users
}
