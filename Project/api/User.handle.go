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
		idclient.Int32 = *req.Idcliente
		idclient.Valid = true
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

type updateUserRequest struct {
	User      string `json:"usuario"`
	Correo    string `json:"correo"`
	Idusuario int32  `json:"idusuario"`
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
		Idusuario: req.Idusuario,
	}

	result, err := server.dbtx.UpdateUser(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rows, err := result.RowsAffected()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if rows == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado o sin cambios"})
		return
	}

	ctx.JSON(http.StatusOK, args)
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

	// Obtener el usuario para saber si tiene cliente o administrador asociado
	user, err := server.dbtx.GetUserById(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Si tiene cliente, eliminarlo
	if user.Idcliente.Valid {
		_, err := server.dbtx.DeleteClient(ctx, user.Idcliente.Int32)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar cliente"})
			return
		}
	}

	// Si tiene administrador, eliminarlo
	if user.Idadministrador.Valid {
		_, err := server.dbtx.DeleteAdmin(ctx, user.Idadministrador.Int32)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar administrador"})
			return
		}
	}

	// Eliminar el usuario
	result, err := server.dbtx.DeleteUser(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rows, _ := result.RowsAffected()
	ctx.JSON(http.StatusOK, gin.H{"rows_affected": rows})
}
