package service

import "remember/entity"

type SystemMessageService interface {
	GetAllSystemMessages(user *entity.User) interface{}

	AddSystemMessage(message *entity.SystemMessage) error
}
