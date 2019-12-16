package model

import (
	"fmt"
	"sync"
	"time"

	"go-gin-backend-admin/pkg/util"

	"github.com/astaxie/beego/validation"
)

type UserInfo struct {
	Id         uint64    `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	CreatedOn  time.Time `json:"createdOn"`
	ModifiedOn time.Time `json:"modifiedOn"`
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
type UserModel struct {
	BaseModel
	Password string `json:"password" gorm:"column:password;not null" valid:"Required; MinSize(5) MaxSize(128)"`
	Username string `json:"username" gorm:"column:username;not null" valid:"Required; MinSize(1) MaxSize(32)"`
	Email    string `json:"email" gorm:"column:email" valid:"Required; Email; MaxSize(100)"`
	Phone    string `json:"phone" gorm:"column:phone" valid:"Required; Phone"`
}

func (c *UserModel) TableName() string {
	return "users"
}

// Create creates a new user account.
func (u *UserModel) Create() error {
	return db.Create(&u).Error
}

// DeleteUser deletes the user by the user identifier.
func DeleteUser(id uint64) error {
	user := UserModel{}
	user.BaseModel.ID = id
	return db.Delete(&user).Error
}

// Update updates an user account information.
func (u *UserModel) Update() error {
	return db.Save(u).Error
}

// GetUser gets an user by the user identifier.
func GetUser(username string) (*UserModel, error) {
	u := &UserModel{}
	d := db.Where("username = ?", username).First(&u)
	return u, d.Error
}

// GetUserById gets an user by the user identifier.
func GetUserById(id uint64) (*UserModel, error) {
	u := &UserModel{}
	d := db.Where("id = ?", id).First(&u)
	return u, d.Error
}

// ListUser List all users
func ListUser(username string, offset, limit int) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = 10
	}

	users := make([]*UserModel, 0)
	var count uint64

	where := fmt.Sprintf("username like '%%%s%%'", username)
	if err := db.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := db.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
func (u *UserModel) Compare(pwd string) (err error) {
	err = util.ComparePwd(u.Password, pwd)
	return err
}

// Encrypt the user password.
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = util.EncryptPwd(u.Password)
	return err
}

// Validate the fields.
func (u *UserModel) Validate() (ok bool, err error) {
	valid := validation.Validation{}
	ok, err = valid.Valid(u)
	return ok, err
}
