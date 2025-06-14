package api

import (
	"database/sql"
	"net/http"
	"project/dto"

	"github.com/gin-gonic/gin"
)

// =====================================================

func (server *Server) GetAllReservations(ctx *gin.Context) {

	reservations, err := server.dbtx.GetAllReservations(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, reservations)
}

// =====================================================

type createReservationRequest struct {
	UsuarioID int32  `json:"idUsuario" binding:"required"`
	DetailID  int32  `json:"idDetail" binding:"required"`
	Estado    string `json:"estado" binding:"required"`
}

func (server *Server) CreateReservation(ctx *gin.Context) {

	var request createReservationRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	params := dto.CreateReservationParams{
		Idusuario: request.UsuarioID,
		Iddetalle: request.DetailID,
		Estado:    request.Estado,
	}

	reservation, err := server.dbtx.CreateReservation(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, reservation)

}

// =====================================================

type updateReservationRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

type updateReservationBodyRequest struct {
	UsuarioID int32 `json:"idusuario" binding:"required"`
	DetailID  int32 `json:"idDetail" binding:"required"`
}

func (server *Server) UpdateReservations(ctx *gin.Context) {

	var request updateReservationRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var bodyReq updateReservationBodyRequest
	if err := ctx.ShouldBindJSON(&bodyReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	params := dto.UpdateReservationParams{
		Idusuario:  bodyReq.UsuarioID,
		Iddetalle:  bodyReq.DetailID,
		Idreservas: request.ID,
	}
	err := server.dbtx.UpdateReservation(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": " La reserva  se ha actualizado con éxito"})
}

type updateStatusRequest struct {
	ID int32 `uri:"id" binding:"required"`
}
type updateStatusBodyRequest struct {
	Estado string `json:"estado" binding:"required"`
}

func (server *Server) UpdateStatus(ctx *gin.Context) {

	var request updateStatusRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var bodyReq updateStatusBodyRequest
	if err := ctx.ShouldBindJSON(&bodyReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	params := dto.UpdateStatusParams{
		Estado:     bodyReq.Estado,
		Idreservas: request.ID,
	}
	err := server.dbtx.UpdateStatus(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": " El estado de la reserva se ha actualizado con éxito"})
}

// =====================================================
type deleteReservationRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) DeleteReservation(ctx *gin.Context) {

	var request deleteReservationRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := server.dbtx.DeleteReservation(ctx, request.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Reserva eliminada con éxito"})
}

// =====================================================

// ===========================Gets=======================

type getReservationByIdRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetReservationsById(ctx *gin.Context) {

	var request getReservationByIdRequest

	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	reservations, err := server.dbtx.GetReservationsById(ctx, request.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, reservations)

}

//=====================================================
