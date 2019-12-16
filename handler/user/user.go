package user

import (
	"go-gin-backend-admin/model"
)

type LoginRequest struct {
	Username string `json:"username" valid:"Required; MinSize(1) MaxSize(32)" example:"admin"`
	Password string `json:"password" valid:"Required; MinSize(5) MaxSize(128)" example:"123456"`
}

type CreateRequest struct {
	LoginRequest
	Email string `json:"email" valid:"Required; Email; MaxSize(100)" example:"82345611@gmail.com"`
	Phone string `json:"phone" valid:"Required; Phone" example:"15845451515"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}

type SwaggerListResponse struct {
	TotalCount uint64           `json:"totalCount"`
	UserList   []model.UserInfo `json:"userList"`
}
