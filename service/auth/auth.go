package auth

import "go-gin-backend-admin/model"

type Auth struct {
	Username string
	Password string
}

// Check Api Auth
func (a *Auth) Check() (bool, error) {
	return model.CheckAuth(a.Username, a.Password)
}
