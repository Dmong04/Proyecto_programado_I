package api

import (
	"database/sql"
	"net/http"
	"project/dto"

	"github.com/gin-gonic/gin"
)

func (server *Server) GetAllAdmins(ctx *gin.Context) {
	admins, err := server.dbtx.GetAllAdmins(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, admins)
}

type createAdminRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) CreateAdmin(ctx *gin.Context) {
	var req createAdminRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := dto.CreateAdminParams{
		Nombre:     req.Name,
		Correo:     req.Email,
		Usuario:    req.User,
		Contraseña: req.Password,
	}
	admin, err := server.dbtx.CreateAdmin(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, admin)
}

type getAdminByIDRequest struct {
	ID int32 `json:"id" binding:"required,min=1"`
}

func (server *Server) GetCategoryByID(ctx *gin.Context) {
	var req getAdminByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	admin, err := server.dbtx.GetAdminById(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, admin)
}

// Obtener al usuario admin por nombre
type getAdminByNameRequest struct {
	Nombre string `json:"name" binding:"required,min=1"`
}

func (server *Server) GetAdminByName(ctx *gin.Context) {
	var req getAdminByNameRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	admin, err := server.dbtx.GetAdminByName(ctx, req.Nombre)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, admin)
}

type updateAdminRequest struct {
	ID int32 `json:"id" binding:"required"`
}

type updateAdminBody struct {
	Name  string `json:"name" binding:"required,min=1"`
	Email string `json:"email" binding:"required,email"`
	User  string `json:"user" binding:"required,alphanum"`
}

func (server *Server) UpdateAdmin(ctx *gin.Context) {
	var req updateAdminRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var bodyReq updateAdminBody
	if err := ctx.ShouldBindJSON(&bodyReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	params := dto.UpdateAdminParams{
		Nombre:          bodyReq.Name,
		Correo:          bodyReq.Email,
		Usuario:         bodyReq.User,
		Idadministrador: req.ID,
	}
	err := server.dbtx.UpdateAdmin(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Administrador modificado con éxito"})
}

type updateAdminPasswordParam struct {
	ID       int32  `json:"id" binding:"required,min=1"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) UpdateAdminPassword(ctx *gin.Context) {
	var req updateAdminRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var pswrdReq updateAdminPasswordParam
	if err := ctx.ShouldBindJSON(&pswrdReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	param := dto.UpdateAdminPasswordParams{
		Idadministrador: pswrdReq.ID,
		Contraseña:      pswrdReq.Password,
	}
	err := server.dbtx.UpdateAdminPassword(ctx, param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Contraseña modificada con éxito"})
}

type deleteAdminRequest struct {
	ID int32 `json:"id" binding:"required"`
}

func (server *Server) DeleteAdmin(ctx *gin.Context) {
	var req deleteAdminRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := server.dbtx.DeleteAdmin(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Administrador eliminado con éxito"})
}

type deleteAdminByNameRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) DeleteAdminByName(ctx *gin.Context) {
	var req deleteAdminByNameRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := server.dbtx.DeleteAdminByName(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Administrador eliminado con éxito"})
}
