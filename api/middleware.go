package api

import (
	"errors"
	"github.com/ariandi/kilat-be-go1/dto"
	util "github.com/ariandi/kilat-be-go1/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

// AuthMiddleware creates a gin middleware for authorization
func AuthMiddleware(config util.Config, roleName []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientIp := ctx.Request.Header.Get("x-forwarded-for")

		if clientIp == "" {
			clientIp = ctx.ClientIP()
		}

		logrus.Info("client ip is ", clientIp)
		whiteListIP := strings.Split(config.WhiteListIP, ",")
		checkIp := false

		for _, whiteList := range whiteListIP {
			if whiteList == clientIp {
				checkIp = true
			}
		}

		if !checkIp {
			err := errors.New("ip address not allowed")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse(err))
			return
		}

		if len(roleName) > 0 {
			logrus.Println("[AuthMiddleware]test : ", roleName[0])
		}

		ctx.Next()
	}
}
