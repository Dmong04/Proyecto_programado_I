package api

import (
	"database/sql"
	"net/http"
	"project/dto"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	User            string `json:"user" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	Idadministrador *int32 `json:"idadministrador"`
	Idcliente       *int32 `json:"idcliente"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	idadmin := sql.NullInt32{}
	if req.Idadministrador != nil {
		idadmin.Int32 = *req.Idadministrador
		idadmin.Valid = true
	}
	idclient := sql.NullInt32{}
	if req.Idcliente != nil {
		idadmin.Int32 = *req.Idcliente
		idadmin.Valid = true
	}
	args := dto.CreateUserParams{
		Usuario:         req.User,
		Correo:          req.Email,
		Contraseña:      req.Password,
		Idadministrador: idadmin,
		Idcliente:       idclient,
	}
	result, err := server.dbtx.CreateUser(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var lastId, _ = result.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{"generated_id": lastId})
}

type updateUserRequest struct {
	User   string `json:"user"`
	Correo string `json:"email"`
	ID     int32  `json:"id"`
}

func (server *Server) getAllUsers(ctx *gin.Context) {
	users, err := server.dbtx.GetAllUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (server *Server) getUserById(ctx *gin.Context) {
	var req struct {
		ID int32 `uri:"id" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := server.dbtx.GetUserById(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (server *Server) updateUser(ctx *gin.Context) {
	var req updateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := dto.UpdateUserParams{
		Usuario:   req.User,
		Correo:    req.Correo,
		Idusuario: req.ID,
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

func (server *Server) getUserByUserName(ctx *gin.Context) {
	var req struct {
		User string `uri:"user" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := server.dbtx.GetUserByUserName(ctx, req.User)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}
func (server *Server) deleteUser(ctx *gin.Context) {
	var req struct {
		ID int32 `uri:"id" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	result, err := server.dbtx.DeleteUser(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rows, _ := result.RowsAffected()
	ctx.JSON(http.StatusOK, gin.H{"rows_affected": rows})
}
