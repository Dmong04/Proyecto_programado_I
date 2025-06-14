package api

import (
	"database/sql"
	"net/http"
	"project/dto"

	"github.com/gin-gonic/gin"
)

// GET /clients
func (server *Server) GetAllClients(ctx *gin.Context) {
	clients, err := server.dbtx.GetAllClients(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, clients)
}

// POST /clients
type createClientRequest struct {
	Name     string `json:"name" binding:"required"`
	Telefono string `json:"telefono" binding:"required"`
}

func (server *Server) CreateClient(ctx *gin.Context) {
	var req createClientRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	params := dto.CreateClientParams{
		Nombre:   req.Name,
		Telefono: req.Telefono,
	}
	result, err := server.dbtx.CreateClient(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var lastId, _ = result.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{"generated_id": lastId})
}

// GET /clients/:id
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
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Cliente no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, client)
}

// GET /clients/name/:name
type getClientByNameRequest struct {
	Nombre string `uri:"name" binding:"required"`
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
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Cliente no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, client)
}

// PUT /clients/:id
type updateClientURI struct {
	ID int32 `uri:"id" binding:"required"`
}
type updateClientBody struct {
	Nombre   string `json:"nombre" binding:"required"`
	Telefono string `json:"telefono" binding:"required"`
}

func (server *Server) UpdateClient(ctx *gin.Context) {
	var uri updateClientURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var body updateClientBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	params := dto.UpdateClientParams{
		Nombre:    body.Nombre,
		Telefono:  body.Telefono,
		Idcliente: uri.ID,
	}
	if _, err := server.dbtx.UpdateClient(ctx, params); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Cliente actualizado con éxito"})
}

// DELETE /clients (body: { "id": 1 })
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
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Cliente no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Cliente eliminado con éxito"})
}

// DELETE /clients/name/:name
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
