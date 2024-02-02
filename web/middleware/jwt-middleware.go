package middleware

import (
	"gin-sites/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JWT token 验证
func JWTMiddleware(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	claims, err := utils.ParseToken(utils.GetToken(token))

	if err != nil || claims.IsFresh == 1 {
		utils.Res(c, http.StatusUnauthorized, "鉴权失败: "+err.Error(), nil)
		c.Abort()
		return
	}
	setUser(c, claims)

	c.Next()
}

func setUser(c *gin.Context, claims *utils.CustomClaims) {
	c.Set("username", claims.Username)
	c.Set("userID", claims.UserID)
}
