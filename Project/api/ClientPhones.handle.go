package api

import (
	"database/sql"
	"net/http"
	"project/dto"

	"github.com/gin-gonic/gin"
)

// =====================================================

func (server *Server) GetAllClientPhones(ctx *gin.Context) {

	ClientPhones, err := server.dbtx.GetAllClientPhones(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, ClientPhones)
}

// =====================================================

type createClientPhonesRequest struct {
	Number     string `json:"number" binding:"required"`
	NumberType string `json:"numberType" binding:"required"`
	Idclient   int32  `json:"idclient" binding:"required"`
}

func (server *Server) CreateClientPhones(ctx *gin.Context) {

	var request createClientPhonesRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	params := dto.CreateClientPhonesParams{
		Numero:    request.Number,
		Tipo:      request.NumberType,
		Idcliente: request.Idclient,
	}

	ClientPhone, err := server.dbtx.CreateClientPhones(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, ClientPhone)

}

// =====================================================

type updateClientPhonesRequest struct {
	ID int32 `json:"id" binding:"required"`
}

type updateClientPhonesRequestBody struct {
	Number     string `json:"number" binding:"required"`
	NumberType string `json:"numberType" binding:"required"`
	IdClient   int32  `json:"idclient" binding:"required"`
}

func (server *Server) UpdateClientPhones(ctx *gin.Context) {

	var request updateClientPhonesRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var bodyReq updateClientPhonesRequestBody
	if err := ctx.ShouldBindJSON(&bodyReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	params := dto.UpdateClientPhonesParams{
		Numero:             bodyReq.Number,
		Tipo:               bodyReq.NumberType,
		Idtelefonoclientes: bodyReq.IdClient,
	}
	err := server.dbtx.UpdateClientPhones(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Telefono del cliente actualizado con éxito"})
}

// =====================================================

type deleteClientPhonesRequest struct {
	ID int32 `json:"id" binding:"required"`
}

func (server *Server) DeleteClientPhones(ctx *gin.Context) {

	var request deleteClientPhonesRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := server.dbtx.DeleteClientPhones(ctx, request.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Telefono del cliente  eliminado con éxito"})
}

// =====================================================

// ===========================Gets=======================

type getClientPhonesByIdRequest struct {
	ID int32 `json:"id" binding:"required,min=1"`
}

func (server *Server) GetClientPhonesById(ctx *gin.Context) {

	var request getClientPhonesByIdRequest

	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ClientPhone, err := server.dbtx.GetClientPhonesById(ctx, request.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, ClientPhone)

}

// =====================================================

type getClientPhonesByClientIDRequest struct {
	IdClient int32 `json:"idclient" binding:"required"`
}

func (server *Server) GetClientPhonesByClientID(ctx *gin.Context) {

	var req getClientPhonesByClientIDRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ClientPhone, err := server.dbtx.GetClientPhonesByClientID(ctx, req.IdClient)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, ClientPhone)
}

// =====================================================
