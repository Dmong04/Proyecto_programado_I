package api

import (
	"database/sql"
	"net/http"
	"project/dto"

	"github.com/gin-gonic/gin"
)

func (server *Server) GetAllHistories(ctx *gin.Context) {
	provider, err := server.dbtx.GetAllHistories(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, provider)
}

type createHistoryRequest struct {
	Descript        string `json:"description" binding:"required"`
	ID_Client       int32 `json:"idclient" binding:"required"`
	ID_Reservation  int32 `json:"idreservation" binding:"required"`
}

func (server *Server) CreateHistory(ctx *gin.Context) {
	var req createHistoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := dto.CreateHistoryParams{
		Descrip: req.Descript,
		Idcliente: req.ID_Client,
		Idreserva: req.ID_Reservation,
	}
	history, err := server.dbtx.CreateHistory(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, history)
}

type getHistoryByIDRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetHistoryByID(ctx *gin.Context) {
	var req getHistoryByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	history, err := server.dbtx.GetHistoryById(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, history)
}

type updateHistoryRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

type updateHistoryBody struct {
	Descript        string `json:"description" binding:"required"`
	ID_Client       int32 `json:"idclient" binding:"required"`
	ID_Reservation  int32 `json:"idreservation" binding:"required"`
}

func (server *Server) UpdateHistory(ctx *gin.Context) {
	var req updateHistoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var bodyReq updateHistoryBody
	if err := ctx.ShouldBindJSON(&bodyReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	params := dto.UpdateHistoryParams{
		Descrip:         bodyReq.Descript,
		Idcliente:       bodyReq.ID_Client,
		Idreserva:       bodyReq.ID_Reservation,
		Idhistorial:     req.ID,
	}
	err := server.dbtx.UpdateHistory(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Historial modificado con éxito"})
}

type deleteHistoryRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) DeleteHistory(ctx *gin.Context) {
	var req deleteHistoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := server.dbtx.DeleteHistory(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Historial eliminado con éxito"})
}
