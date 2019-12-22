package mock

import (
	"go-gin-backend-admin/handler"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func SysLogCount(c *gin.Context) {
	handlerG := handler.Gin{C: c}

	handlerG.Response(nil, map[string]interface{}{
		"todayCount": random(160, 1000),
		"monthCount": random(160, 1000),
		"yearCount":  random(160, 1000),
		"online":     random(160, 1000),
		"dayrate":    random(160, 1000),
		"monthrate":  random(160, 1000),
	})
}
