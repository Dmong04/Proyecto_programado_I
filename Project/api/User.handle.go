package api

import (
	"database/sql"
	"net/http"
	"project/dto"
	"time"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	User     string `json:"user" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	req.Role = "user_role"
	args := dto.CreateUserParams{
		Usuario:    req.User,
		Correo:     req.Email,
		Contraseña: req.Password,
		Role:       req.Role,
	}
	result, err := server.dbtx.CreateUser(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var lastId, _ = result.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{"generated_id": lastId})
}

type loginRequest struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type loginResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}
type userResponse struct {
	ID   int32  `json:"id"`
	User string `json:"user"`
	Role string `json:"role"`
}

func (server *Server) login(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := server.dbtx.GetUserByUserName(ctx, req.User)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if user.Password != req.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autorizado"})
		return
	}
	//IMPORTANTE QUE LA DURACIÓN SE ESTABLEZCA EN EL .ENV
	accessToken, err := server.tokenBuilder.CreateToken(user.Email, time.Minute*5)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	resp := loginResponse{
		AccessToken: accessToken,
		User: userResponse{
			ID:   user.ID,
			User: user.User,
			Role: user.Role,
		},
	}
	ctx.JSON(http.StatusOK, resp)
}

type updateUserRequest struct {
	User   string `json:"user"`
	Correo string `json:"email"`
	ID     int32  `json:"id"`
}

func (server *Server) updateUser(ctx *gin.Context) {
	var req updateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := dto.UpdateUserParams{
		Correo:  req.Correo,
		Usuario: req.User,
	}
	result, err := server.dbtx.UpdateUser(ctx, args)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var rows, _ = result.RowsAffected()
	ctx.JSON(http.StatusOK, gin.H{"rows_affected": rows})
}

type updatePasswordRequest struct {
	Password string `json:"password"`
	ID       int32  `json:"id"`
}

func (server *Server) updatePassword(ctx *gin.Context) {
	var req updatePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := dto.UpdateUserPasswordParams{
		Idusuario:  req.ID,
		Contraseña: req.Password,
	}
	result, err := server.dbtx.UpdateUserPassword(ctx, args)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var rows, _ = result.RowsAffected()
	ctx.JSON(http.StatusOK, gin.H{"rows_affected": rows})
}

type updateRoleRequest struct {
	Role string `json:"role"`
	ID   int32  `json:"id"`
}

func (server *Server) updateRole(ctx *gin.Context) {
	var req updateRoleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := dto.UpdateUserRoleParams{
		Idusuario: req.ID,
		Role:      req.Role,
	}
	result, err := server.dbtx.UpdateUserRole(ctx, args)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var rows, _ = result.RowsAffected()
	ctx.JSON(http.StatusOK, gin.H{"rows_affected": rows})
}
