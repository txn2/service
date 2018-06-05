package ginack

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"github.com/txn2/service"

	"time"
)

func Ack(c *gin.Context) service.Ack {
	t := time.Now()
	u, _ := uuid.NewV4()

	// get uuid from header
	ru := c.Request.Header.Get("uuid")

	ack := service.Ack{
		Uuid:        u.String(),
		RequestUuid: ru,
		ServerCode:  200,
		Success:     true,
		Version:     5,
		DateTime:    t.Format(time.RFC3339),
		Location:    c.Request.URL.String(),
	}

	return ack
}
