package common

type ChangeUserP struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ChangeUserI struct {
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Birthday MyDate `json:"birthday"`
}
