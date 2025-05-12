package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	AccessToken string       `json:"access_token"`
	LoggedUser  userResponse `json:"logged_user"`
	Role        string       `json:"role"`
}

type userResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	User  string `json:"user"`
}

func (server *Server) login(ctx *gin.Context) {
	var request loginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	admin, err := server.dbtx.GetAdminByUser(ctx, request.User)
	if err == nil {
		if admin.Contraseña != request.Password {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Contraseña incorrecta"})
			return
		}
		accessToken, err := server.tokenBuilder.CreateToken(admin.Usuario, "Admin", time.Hour)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		response := loginResponse{
			AccessToken: accessToken,
			LoggedUser: userResponse{
				Name:  admin.Nombre,
				Email: admin.Correo,
				User:  admin.Usuario,
			},
			Role: "Admin",
		}
		ctx.JSON(http.StatusOK, response)
		return
	}
	if err == sql.ErrNoRows {
		client, err := server.dbtx.GetClientByUser(ctx, request.User)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no encontrado"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		if client.Contraseña != request.Password {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Contraseña incorrecta"})
			return
		}
		accessToken, err := server.tokenBuilder.CreateToken(client.Usuario, "Client", time.Hour)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		response := loginResponse{
			AccessToken: accessToken,
			LoggedUser: userResponse{
				Name:  client.Nombre,
				Email: client.Correo,
				User:  client.Usuario,
			},
			Role: "Client",
		}
		ctx.JSON(http.StatusOK, response)
		return
	}
	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
}
