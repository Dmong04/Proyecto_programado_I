package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type loginUserRequest struct {
	ID       int32  `json:"idUsuario"`
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	LoggedUser  userResponse `json:"logged_user"`
	Role        string       `json:"role"`
}

type createUserResponse struct {
	Email      string    `json:"email"`
	User       string    `json:"user"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type userResponse struct {
	Email      string    `json:"email"`
	User       string    `json:"user"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

func (server *Server) login(ctx *gin.Context) {
	var request loginUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	admin, err := server.dbtx.GetUserByUserName(ctx, request.User)
	if err == nil {
		if admin.Password != request.Password {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Contraseña incorrecta"})
			return
		}
		if admin.Role == "Admin" {
			accessToken, err := server.tokenBuilder.CreateToken(admin.User, "Admin", time.Hour)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
			response := loginUserResponse{
				AccessToken: accessToken,
				LoggedUser: userResponse{
					User:  admin.User,
					Email: admin.Email,
				},
				Role: "Admin",
			}
			ctx.JSON(http.StatusOK, response)
			return
		}
	}
	if err == sql.ErrNoRows {
		client, err := server.dbtx.GetUserByUserName(ctx, request.User)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no encontrado"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		if client.Password != request.Password {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Contraseña incorrecta"})
			return
		}
		if client.Role == "Client" {
			accessToken, err := server.tokenBuilder.CreateToken(client.User, "Client", time.Hour)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
			response := loginUserResponse{
				AccessToken: accessToken,
				LoggedUser: userResponse{
					User:  client.User,
					Email: client.Email,
				},
				Role: "Client",
			}
			ctx.JSON(http.StatusOK, response)
			return
		}
	}
	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
}
