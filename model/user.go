package model

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"go-gin-backend-admin/pkg/util"
	"sync"
)

type UserInfo struct {
	Id         uint64 `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	CreatedOn  string `json:"createdOn"`
	ModifiedOn string `json:"modifiedOn"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}

// Token represents a JSON web token.
type Token struct {
	Token string `json:"token"`
}

// User represents a registered user.
type User struct {
	BaseModel
	Password string `gorm:"column:password;not null" valid:"Required; MinSize(5) MaxSize(128)"`
	Username string `gorm:"column:username;not null" valid:"Required; MinSize(1) MaxSize(32)"`
	Email    string `gorm:"column:email" valid:"Required; Email; MaxSize(100)"`
	Phone    string `gorm:"column:phone" valid:"Required; Phone"`
}

func (c User) TableName() string {
	return c.BaseModel.TableName("user")
}

// Create creates a new user account.
func (u *User) Create() error {
	return db.Create(&u).Error
}

// DeleteUser deletes the user by the user identifier.
func DeleteUser(id uint64) error {
	user := User{}
	user.BaseModel.ID = id
	return db.Delete(&user).Error
}

// Update updates an user account information.
func (u *User) Update() error {
	return db.Save(u).Error
}

// GetUser gets an user by the user identifier.
func GetUser(username string) (*User, error) {
	u := &User{}
	d := db.Where("username = ?", username).First(&u)
	return u, d.Error
}

// GetUserById gets an user by the user identifier.
func GetUserById(id uint64) (*User, error) {
	u := &User{}
	d := db.Where("id = ?", id).First(&u)
	return u, d.Error
}

// ListUser List all users
func ListUser(username string, offset, limit int) ([]*User, uint64, error) {
	if limit == 0 {
		limit = 10
	}

	users := make([]*User, 0)
	var count uint64

	where := fmt.Sprintf("username like '%%%s%%'", username)
	if err := db.Model(&User{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := db.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
func (u *User) Compare(pwd string) (err error) {
	err = util.ComparePwd(u.Password, pwd)
	return err
}

// Encrypt the user password.
func (u *User) Encrypt() (err error) {
	u.Password, err = util.EncryptPwd(u.Password)
	return err
}

// Validate the fields.
func (u *User) Validate() (ok bool, err error) {
	valid := validation.Validation{}
	ok, err = valid.Valid(u)
	return ok, err
}
