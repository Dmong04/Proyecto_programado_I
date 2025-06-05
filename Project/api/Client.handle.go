package api

import (
	"database/sql"
	"net/http"
	"project/dto"

	"github.com/gin-gonic/gin"
)

func (server *Server) GetAllClients(ctx *gin.Context) {
	clients, err := server.dbtx.GetAllClients(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, clients)
}

type CreateClientRequest struct {
	Nombre string `json:"name"`
}

func (server *Server) CreateClient(ctx *gin.Context) {
	var req CreateClientRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := req.Nombre
	result, err := server.dbtx.CreateClient(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var lastId, _ = result.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{"generated_id": lastId})
}

type getClientByIDRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetClientByID(ctx *gin.Context) {
	var req getClientByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	client, err := server.dbtx.GetClientById(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, client)
}

type getClientByNameRequest struct {
	Nombre string `uri:"name" binding:"required,min=1"`
}

func (server *Server) GetClientByName(ctx *gin.Context) {
	var req getClientByNameRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	client, err := server.dbtx.GetClientByName(ctx, req.Nombre)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, client)
}

type updateClientRequest struct {
	ID int32 `uri:"id" binding:"required"`
}
type updateClientBody struct {
	Name string `json:"name" binding:"required,min=1"`
}

func (server *Server) UpdateClient(ctx *gin.Context) {
	var req updateClientRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var bodyReq updateClientBody
	if err := ctx.ShouldBindJSON(&bodyReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	params := dto.UpdateClientParams{
		Nombre:    bodyReq.Name,
		Idcliente: req.ID,
	}
	_, err := server.dbtx.UpdateClient(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Cliente se modificado con éxito"})
}

type deleteClientRequest struct {
	ID int32 `json:"id" binding:"required"`
}

func (server *Server) DeleteClient(ctx *gin.Context) {
	var req deleteClientRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	_, err := server.dbtx.DeleteClient(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Cliente fue eliminado con éxito"})
}

type deleteClientByNameRequest struct {
	Name string `uri:"name" binding:"required"`
}

func (server *Server) DeleteClientByName(ctx *gin.Context) {
	var req deleteClientByNameRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	_, err := server.dbtx.DeleteClientByName(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Cliente eliminado con éxito"})
}
