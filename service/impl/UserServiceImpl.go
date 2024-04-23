package impl

import (
	"errors"
	"log"
	"reflect"
	"remember/common"
	"remember/entity"
	"remember/mapper"
	"remember/utils"
)

type UserServiceImpl struct {
	mapper mapper.UserMapper
}

func (u *UserServiceImpl) Delete(controller, username string) error {
	userC := u.mapper.GetUserByUsername(controller)
	user := u.mapper.GetUserByUsername(username)
	if user.ID == 0 {
		return errors.New("用户：" + username + "已经删除过了")
	}
	userCRoles := userC.Roles
	userRoles := user.Roles

	if controller == username {
		err := u.mapper.Delete(user)
		if err != nil {
			return err
		}
		utils.DeleteInfoByUsername(username)
		log.Println("用户：" + controller + "成功注销")
		return errors.New("注销成功")
	} else {
		if userCRoles == common.Admins.Name() && userRoles != common.Admins.Name() {
			err := u.mapper.Delete(user)
			if err != nil {
				return err
			}
		} else if userCRoles == common.Admin.Name() && userRoles == common.User.Name() {
			err := u.mapper.Delete(user)
			if err != nil {
				return err
			}
		} else {
			log.Println("用户：" + controller + "没有权限删除用户：" + username)
			return errors.New("用户：" + controller + "没有权限删除用户：" + username)
		}
		log.Println("用户：" + controller + "成功删除用户：" + username)
		return nil
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
		user.Roles = common.User.Name() // 通过注册的用户只能为 user
		err := u.mapper.Insert(user)
		log.Println("新用户：" + user.Username + "注册成功")
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

func (u *UserServiceImpl) ChangePassword(user *entity.User, token, OldPassword, NewPassword string) error {
	// 验证旧密码是否正确
	if user.Password != utils.PasswordEncrypt(OldPassword) {
		return errors.New("密码错误")
	}
	// 新旧密码不能相同
	if OldPassword == NewPassword {
		return errors.New("新旧密码不能相同")
	}
	// 修改信息
	user.Password = utils.PasswordEncrypt(NewPassword)
	err := u.mapper.Update(user)
	if err != nil {
		return err
	}
	// 退出登录
	err = u.Logout(user.Username, token)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserServiceImpl) GetUserInfo(user *entity.User) *entity.User {
	return user.Desensitization()
}

func (u *UserServiceImpl) ChangeUserInfo(user *entity.User, mp map[string]any) error {
	v := reflect.ValueOf(user)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for key, value := range mp {
		if reflect.ValueOf(value).Type() == v.FieldByName(key).Type() {
			v.FieldByName(key).Set(reflect.ValueOf(value))
		}
	}
	err := u.mapper.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserServiceImpl) ChangeUserInfoById(id int, mp map[string]any) error {
	user := u.mapper.GetUserById(id)
	if user.ID == 0 {
		return errors.New("用户不存在")
	}
	return u.ChangeUserInfo(user, mp)
}

func (u *UserServiceImpl) AddUser(user *entity.User) error {
	user.Password = utils.PasswordEncrypt(user.Password)
	err := u.mapper.Insert(user)
	if err != nil {
		return err
	}
	return nil
}
