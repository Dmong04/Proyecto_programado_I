package api

import (
	"errors"
	"net/http"
	"project/security"
	"strings"

	"github.com/gin-gonic/gin"
)

func authMiddleware(tokenBuilder security.Builder) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("authorization")
		if len(authHeader) == 0 {
			err := errors.New("no existe un token de autorización")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			err := errors.New("formato de token inválido")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		if strings.ToLower(fields[0]) != "bearer" {
			err := errors.New("tipo de autorización no soportado")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		accessToken := fields[1]
		payload, err := tokenBuilder.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		ctx.Set("authorized", payload)
		ctx.Next()
	}
}

func roleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payloadInterface, exists := ctx.Get("authorized")
		if !exists {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}
		payload := payloadInterface.(*security.Payload)
		if payload.Role != requiredRole {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "No se cuenta con permisos para este recurso"})
			return
		}
		ctx.Next()
	}
}
