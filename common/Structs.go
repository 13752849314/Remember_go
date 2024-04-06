package common

import "time"

type ChangeUserP struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ChangeUserI struct {
	Phone    string     `json:"phone"`
	Email    string     `json:"email"`
	Birthday *time.Time `json:"birthday"`
}
