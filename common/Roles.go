package common

type Role int

const (
	User Role = iota
	Admin
	Admins
)

func (r Role) Name() string {
	if r < User || r > Admins {
		return "user"
	}
	return [...]string{"user", "admin", "admins"}[r]
}

func (r Role) Ge(role Role) bool {
	if role.Name() == Admins.Name() {
		return true
	}
	if r.Name() == role.Name() {
		return true
	}
	return role.Name() == Admin.Name() && r.Name() == User.Name()
}

func ValueOf(name string) Role {
	switch name {
	case "user":
		return User
	case "admin":
		return Admin
	case "admins":
		return Admins
	}
	return User
}
