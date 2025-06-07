package api

import (
	"database/sql"
	"net/http"
	"project/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateTravelRequest struct {
	Tipoviaje   string `json:"tipoviaje"`
	Descripcion string `json:"descripcion"`
}

func (server *Server) CreateTravel(ctx *gin.Context) {
	var request CreateTravelRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	params := dto.CreateTravelParams{
		Tipoviaje:   request.Tipoviaje,
		Descripcion: request.Descripcion,
	}
	travel, err := server.dbtx.CreateTravel(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, travel)
}

func (server *Server) DeleteTravel(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = server.dbtx.DeleteTravel(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Viaje no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar viaje"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Viaje eliminado con éxito"})
}

func (server *Server) GetAllTravels(ctx *gin.Context) {
	travels, err := server.dbtx.GetAllTravels(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, travels)
}

type getTravelByIdRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetTravelById(ctx *gin.Context) {
	var request getTravelByIdRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	travel, err := server.dbtx.GetTravelById(ctx, request.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, travel)
}

type updateTravelRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

type updateTravelBodyRequest struct {
	Tipoviaje   string `json:"tipoviaje"`
	Descripcion string `json:"descripcion"`
}

func (server *Server) UpdateTravel(ctx *gin.Context) {
	var request updateTravelRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var bodyReq updateTravelBodyRequest
	if err := ctx.ShouldBindJSON(&bodyReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	params := dto.UpdateTravelParams{
		Tipoviaje:   bodyReq.Tipoviaje,
		Descripcion: bodyReq.Descripcion,
	}
	err := server.dbtx.UpdateTravel(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Viaje fue actualizado con éxito"})
}
