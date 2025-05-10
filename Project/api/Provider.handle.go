package api

import (
	"database/sql"
	"net/http"
	"project/dto"

	"github.com/gin-gonic/gin"
)

func (server *Server) GetAllProviders(ctx *gin.Context) {
	provider, err := server.dbtx.GetAllProviders(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, provider)
}

type createProviderRequest struct {
	Name      string `json:"name" binding:"required"`
	Descript  string `json:"descript" binding:"required"`
}

func (server *Server) CreateProvider(ctx *gin.Context) {
	var req createProviderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := dto.CreateProviderParams{
		Nombre: req.Name,
		Descrip: req.Descript,
		
	}
	provider, err := server.dbtx.CreateProvider(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, provider)
}

type getProviderByIDRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetProviderByID(ctx *gin.Context) {
	var req getProviderByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	provider, err := server.dbtx.GetProviderById(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, provider)
}

type getProviderByNameRequest struct {
	Nombre string `uri:"name" binding:"required,min=1"`
}

func (server *Server) GetProviderByName(ctx *gin.Context) {
	var req getProviderByNameRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	provider, err := server.dbtx.GetProviderByName(ctx, req.Nombre)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, provider)
}

type updateProviderRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

type updateProviderBody struct {
	Name     string `json:"name" binding:"required,min=1"`
	Descript string `json:"descript" binding:"required"`
}

func (server *Server) UpdateProvider(ctx *gin.Context) {
	var req updateProviderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var bodyReq updateProviderBody
	if err := ctx.ShouldBindJSON(&bodyReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	params := dto.UpdateProviderParams{
		Nombre:          bodyReq.Name,
		Descrip:         bodyReq.Descript,
		Idproveedor:     req.ID,
	}
	err := server.dbtx.UpdateProvider(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Proveedor modificado con éxito"})
}

type updateProviderByNameRequest struct {
	Name string `uri:"name" binding:"required"`
}

type updateProviderByNameBody struct {
	Name_2   string `json:"new_name" binding:"required,min=1"`
	Descript string `json:"descript" binding:"required"`
}

func (server *Server) UpdateProviderByName(ctx *gin.Context) {
	var req updateProviderByNameRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var bodyReq updateProviderByNameBody
	if err := ctx.ShouldBindJSON(&bodyReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	params := dto.UpdateProviderByNameParams{
		Nombre_2:        bodyReq.Name_2,
		Descrip:         bodyReq.Descript,
		Nombre:          req.Name,
	}
	err := server.dbtx.UpdateProviderByName(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Proveedor modificado con éxito"})
}

type deleteProviderRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) DeleteProvider(ctx *gin.Context) {
	var req deleteProviderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := server.dbtx.DeleteProvider(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Proveedor eliminado con éxito"})
}

type deleteProviderByNameRequest struct {
	Name string `uri:"name" binding:"required"`
}

func (server *Server) DeleteProviderByName(ctx *gin.Context) {
	var req deleteProviderByNameRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := server.dbtx.DeleteProviderByName(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Proveedor eliminado con éxito"})
}
