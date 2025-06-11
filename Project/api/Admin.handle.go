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
	Name string `json:"name" binding:"required"`
}

func (server *Server) CreateAdmin(ctx *gin.Context) {
	var req createAdminRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	params := req.Name
	result, err := server.dbtx.CreateAdmin(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var lastId, _ = result.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{"generated_id": lastId})
}

type getAdminByIDRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetAdminByID(ctx *gin.Context) {
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

type getAdminByNameRequest struct {
	Nombre string `uri:"name" binding:"required,min=1"`
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
	ID int32 `uri:"id" binding:"required"`
}

type updateAdminBodyRequest struct {
	Name string `json:"name" binding:"required,min=1"`
}

func (server *Server) UpdateAdmin(ctx *gin.Context) {
	var req updateAdminRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var bodyReq updateAdminBodyRequest
	if err := ctx.ShouldBindJSON(&bodyReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	params := dto.UpdateAdminParams{
		Nombre:          bodyReq.Name,
		Idadministrador: req.ID,
	}
	_, err := server.dbtx.UpdateAdmin(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Administrador modificado con éxito"})
}

type deleteAdminRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) DeleteAdmin(ctx *gin.Context) {
	var req deleteAdminRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	_, err := server.dbtx.DeleteAdmin(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Administrador eliminado con éxito"})
}

type deleteAdminByNameRequest struct {
	Name string `uri:"name" binding:"required"`
}

func (server *Server) DeleteAdminByName(ctx *gin.Context) {
	var req deleteAdminByNameRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	_, err := server.dbtx.DeleteAdminByName(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Administrador eliminado con éxito"})
}
