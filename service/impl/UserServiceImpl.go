package impl

import (
	"errors"
	"remember/entity"
	"remember/mapper"
	"remember/utils"
)

type UserServiceImpl struct {
	mapper mapper.UserMapper
}

func (u *UserServiceImpl) Login(user *entity.User) error {
	userDb := u.mapper.GetUserByUsername(user.Username)
	if userDb.ID == 0 {
		return errors.New("用户不存在")
	}
	pass := utils.PasswordEncrypt(user.Password)
	if pass != userDb.Password {
		return errors.New("密码错误")
	}
	return nil
}

func (u *UserServiceImpl) Registration(user *entity.User) (err error) {
	p := u.mapper.Select("getUserByUsername", user.Username)
	p1 := p.(*entity.User)
	if p1.ID == 0 {
		// 是新用户 存入数据库
		// 将密码加密 SHA256
		user.Password = utils.PasswordEncrypt(user.Password)
		err := u.mapper.Insert(user)
		if err != nil {
			return err
		}
	} else {
		return errors.New("用户名已经存在")
	}
	return nil
}

func (u *UserServiceImpl) GetAllUsers() interface{} {
	users := u.mapper.Select("getAllUsers")
	return users
}
