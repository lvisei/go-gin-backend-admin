package mock

import (
	"go-gin-backend-admin/handler"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v4"
	"github.com/gin-gonic/gin"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func SysLogCount(c *gin.Context) {
	handlerG := handler.Gin{C: c}

	handlerG.Response(nil, map[string]interface{}{
		"online":    random(160, 1000),
		"newVisits": random(160, 1000),
		"totalUser": random(500, 1000),
		"messages":  random(10, 100),
	})
}

func ResponseOk(c *gin.Context) {
	handlerG := handler.Gin{C: c}

	handlerG.Response(nil, true)
}

func Search(c *gin.Context) {
	handlerG := handler.Gin{C: c}

	gofakeit.Seed(time.Now().UnixNano())

	type user struct {
		Id         string   `json:"id"`
		Username   string   `json:"username"`
		Name       string   `json:"name"`
		Department string   `json:"department"`
		Starttime  string   `json:"starttime"`
		State      int      `json:"state"`
		Sex        int      `json:"sex"`
		Age        int      `json:"age"`
		Email      string   `json:"email"`
		Areacode   []string `json:"areacode"`
		Areaname   string   `json:"areaname"`
	}

	var userList []user

	for i := 0; i < 10; i++ {
		user := user{
			gofakeit.UUID(),
			gofakeit.FirstName(),
			gofakeit.Username(),
			gofakeit.State(),
			gofakeit.Date().Format("2006-01-02 15:04:05"),
			gofakeit.Number(0, 1),
			gofakeit.Number(1, 2),
			gofakeit.Number(16, 80),
			gofakeit.Email(),
			[]string{"hangzhou"},
			"杭州",
		}
		userList = append(userList, user)
	}

	handlerG.Response(nil, map[string]interface{}{
		"userList": userList,
		"count":    40,
	})
}
