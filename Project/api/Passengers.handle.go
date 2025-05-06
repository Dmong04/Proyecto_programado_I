package api

import (
	"database/sql"
	"net/http"
	"project/dto"

	"github.com/gin-gonic/gin"
)

//=================================================================

func (server *Server) GetAllPassengers(ctx *gin.Context) {
	passengers, err := server.dbtx.GetAllPassengers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, passengers)
}

//=================================================================

type createPassengersRequest struct {
	Name     string `json:"name" binding:"required"`
	Age      int32  `json:"age" binding:"required"`
	DetailID int32  `json:"idDetail" binding:"required"`
}

func (server *Server) CreatePassenger(ctx *gin.Context) {
	var request createPassengersRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	params := dto.CreatePassengerParams{
		Nombre:    request.Name,
		Edad:      request.Age,
		Iddetalle: request.DetailID,
	}
	passenger, err := server.dbtx.CreatePassenger(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, passenger)
}

//=================================================================

type updatePassengerRequest struct {
	ID int32 `json:"id" binding:"required"`
}

type updatePassengerBodyRequest struct {
	Name     string `json:"name" binding:"required"`
	Age      int32  `json:"age" binding:"required"`
	DetailID int32  `json:"idDetail" binding:"required"`
}

func (server *Server) UpdatePassengers(ctx *gin.Context) {
	var request updatePassengerRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var bodyReq updatePassengerBodyRequest
	if err := ctx.ShouldBindJSON(&bodyReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	params := dto.UpdatePassengerParams{
		Nombre:      bodyReq.Name,
		Edad:        bodyReq.Age,
		Iddetalle:   bodyReq.DetailID,
		Idpasajeros: request.ID,
	}
	err := server.dbtx.UpdatePassenger(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Pasajero actualizado con éxito"})
}

//=================================================================

type deletePassengerRequest struct {
	ID int32 `json:"id" binding:"required"`
}

func (server *Server) DeletePassenger(ctx *gin.Context) {
	var request deletePassengerRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := server.dbtx.DeletePassenger(ctx, request.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Pasajero eliminado con éxito"})
}

//=================================================================

// ===========================Gets==================================

type getPassengersByDetailIDRequest struct {
	DetailID int32 `json:"idDetail" binding:"required"`
}

func (server *Server) GetPassengersByDetailID(ctx *gin.Context) {
	var request getPassengersByDetailIDRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	passengers, err := server.dbtx.GetPassengersByDetailID(ctx, request.DetailID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, passengers)
}

// =================================================================

type getPassengersByIdRequest struct {
	ID int32 `json:"id" binding:"required,min=1"`
}

func (server *Server) GetPassengersById(ctx *gin.Context) {
	var request getPassengersByIdRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	passenger, err := server.dbtx.GetPassengersById(ctx, request.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, passenger)
}

//=================================================================

type getPassengersByNameRequest struct {
	Name string `json:"name" binding:"required,min=1"`
}

func (server *Server) GetPassengersByName(ctx *gin.Context) {
	var request getPassengersByNameRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	passenger, err := server.dbtx.GetPassengersByName(ctx, request.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, passenger)
}
