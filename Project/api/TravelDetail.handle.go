package api

import (
	"database/sql"
	"net/http"
	"project/dto"

	"github.com/gin-gonic/gin"
)

func (server *Server) getAllDetails(ctx *gin.Context) {
	details, err := server.dbtx.GetAllTravelDetails(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, details)
}

type getDetailsByIDRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getDetailsByID(ctx *gin.Context) {
	var request getDetailsByIDRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	detail, err := server.dbtx.GetTravelDetailById(ctx, request.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, detail)
}

type createDetailsRequest struct {
	Fecha       string `json:"fecha" binding:"required"`
	Hora        string `json:"hora" binding:"required"`
	Idproveedor int32  `json:"idproveedor" binding:"omitempty"`
	Idviaje     int32  `json:"idviaje" binding:"required"`
}

func (server *Server) CreateDetail(ctx *gin.Context) {
	var request createDetailsRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := dto.CreateTravelDetailParams{
		Fecha:       request.Fecha,
		Hora:        request.Hora,
		Idproveedor: request.Idproveedor,
		Idviaje:     request.Idviaje,
	}
	detail, err := server.dbtx.CreateTravelDetail(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var lastId, _ = detail.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{"generated_id": lastId})
}

type updateDetailrequest struct {
	ID int32 `uri:"id" binding:"required"`
}
type updateDetailBodyrequest struct {
	Fecha       string `json:"fecha" binding:"required"`
	Hora        string `json:"hora" binding:"required"`
	Idproveedor int32  `json:"idproveedor" binding:"omitempty"`
	Idviaje     int32  `json:"idviaje" binding:"required"`
}

func (server *Server) UpdateDetail(ctx *gin.Context) {
	var request updateDetailrequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var bodyReq updateDetailBodyrequest
	if err := ctx.ShouldBindJSON(&bodyReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	params := dto.UpdateTravelDetailParams{
		Fecha:          bodyReq.Fecha,
		Hora:           bodyReq.Hora,
		Idproveedor:    bodyReq.Idproveedor,
		Idviaje:        bodyReq.Idviaje,
		Iddetalleviaje: request.ID,
	}
	err := server.dbtx.UpdateTravelDetail(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "detalles del viaje actualizados con éxito"})
}

type deleteDetailRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) DeleteDetail(ctx *gin.Context) {
	var request deleteDetailRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := server.dbtx.DeleteTravelDetail(ctx, request.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Detalles del viaje eliminados con éxito"})
}
