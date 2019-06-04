package routes

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Permission() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		v := session.Get("count")
		if v == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"errorCode": "SE001", "message": "Unauthenticated user"})
			c.Abort()
			return
		}
		c.Next()
	}
}
