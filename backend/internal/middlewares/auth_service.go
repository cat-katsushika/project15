package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SessionMiddleware(secret string) gin.HandlerFunc {
	store := cookie.NewSrore([]byte(secret))
	return sessions.Sessions("mysession", store)
}