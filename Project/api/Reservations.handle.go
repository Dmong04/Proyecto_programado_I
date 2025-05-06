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
	ClientID        int32 `json:"idclient" binding:"required"`
	AdministratorID int32 `json:"idadministrator" binding:"required"`
	DetailID        int32 `json:"idDetail" binding:"required"`
}

func (server *Server) CreateReservation(ctx *gin.Context) {

	var request createReservationRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	params := dto.CreateReservationParams{
		Idcliente:       request.ClientID,
		Idadministrador: request.AdministratorID,
		Iddetalle:       request.DetailID,
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
	ID int32 `json:"id" binding:"required"`
}

type updateReservationBodyRequest struct {
	ClientID        int32 `json:"idclient" binding:"required"`
	AdministratorID int32 `json:"idadministrator" binding:"required"`
	DetailID        int32 `json:"idDetail" binding:"required"`
}

func (server *Server) UpdateReservations(ctx *gin.Context) {

	var request updateReservationRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var bodyReq updateReservationBodyRequest
	if err := ctx.ShouldBindJSON(&bodyReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	params := dto.UpdateReservationParams{
		Idcliente:       bodyReq.ClientID,
		Idadministrador: bodyReq.AdministratorID,
		Iddetalle:       bodyReq.DetailID,
	}
	err := server.dbtx.UpdateReservation(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Reserva actualizado con éxito"})
}

// =====================================================
type deleteReservationRequest struct {
	ID int32 `json:"id" binding:"required"`
}

func (server *Server) DeleteReservation(ctx *gin.Context) {

	var request deleteReservationRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
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
	ctx.JSON(http.StatusOK, gin.H{"message": "Reserva eliminado con éxito"})
}

// =====================================================

// ===========================Gets=======================

type getReservationByIdRequest struct {
	ID int32 `json:"id" binding:"required,min=1"`
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
