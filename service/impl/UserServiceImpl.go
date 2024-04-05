package impl

import (
	"errors"
	"remember/common"
	"remember/entity"
	"remember/mapper"
	"remember/utils"
	"time"
)

type UserServiceImpl struct {
	mapper mapper.UserMapper
}

func (u *UserServiceImpl) Delete(controller, username string) error {
	userC := u.mapper.GetUserByUsername(controller)
	user := u.mapper.GetUserByUsername(username)
	userCRoles := userC.Roles
	userRoles := user.Roles

	now := time.Now()
	user.DeletedAt = &now
	user.UpdatedAt = now

	if controller == username {
		err := u.mapper.Update(user)
		if err != nil {
			return err
		}
		utils.DeleteInfoByUsername(username)
		return errors.New("注销成功")
	} else {
		if userCRoles == common.Admins.Name() && userRoles != common.Admins.Name() {
			err := u.mapper.Update(user)
			if err != nil {
				return err
			}
		} else if userCRoles == common.Admin.Name() && userRoles == common.User.Name() {
			err := u.mapper.Update(user)
			if err != nil {
				return err
			}
		} else {
			return errors.New("用户：" + controller + "没有权限删除用户：" + username)
		}
		return errors.New("用户：" + controller + "成功删除用户：" + username)
	}
}

func (u *UserServiceImpl) Logout(username, token string) error {
	err := utils.DeleteInfo(username, token)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserServiceImpl) GetAllUser() []entity.User {
	return u.mapper.GetAllUser()
}

func (u *UserServiceImpl) Login(user *entity.User) (string, error) {
	userDb := u.mapper.GetUserByUsername(user.Username)
	if userDb.ID == 0 {
		return "", errors.New("用户不存在")
	}
	pass := utils.PasswordEncrypt(user.Password)
	if pass != userDb.Password {
		return "", errors.New("密码错误")
	}
	token, err := utils.CreatToken(userDb)
	if err != nil {
		return "", err
	}
	err = utils.AddInfo(userDb.Username, token)
	if err != nil {
		return "", err
	}
	return token, nil
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
