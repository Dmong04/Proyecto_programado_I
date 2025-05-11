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
	Nombre     string `json:"name"`
	Correo     string `json:"email"`
	Usuario    string `json:"user"`
	Contraseña string `json:"password"`
}

func (server *Server) CreateClient(ctx *gin.Context) {
	var req CreateClientRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := dto.CreateClientParams{
		Nombre:     req.Nombre,
		Correo:     req.Correo,
		Usuario:    req.Usuario,
		Contraseña: req.Contraseña,
	}
	client, err := server.dbtx.CreateClient(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, client)
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
	Name  string `json:"name" binding:"required,min=1"`
	Email string `json:"email" binding:"required,email"`
	User  string `json:"user" binding:"required,alphanum"`
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
		Correo:    bodyReq.Email,
		Usuario:   bodyReq.User,
		Idcliente: req.ID,
	}
	err := server.dbtx.UpdateClient(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Cliente se modificado con éxito"})
}

type updateClientPasswordParam struct {
	ID       int32  `json:"id" binding:"required,min=1"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) UpdateClientPassword(ctx *gin.Context) {
	var req updateClientRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var pswrdReq updateClientPasswordParam
	if err := ctx.ShouldBindJSON(&pswrdReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	param := dto.UpdateClientPasswordParams{
		Idcliente:  pswrdReq.ID,
		Contraseña: pswrdReq.Password,
	}
	err := server.dbtx.UpdateClientPassword(ctx, param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Contraseña de cliente modificada con éxito"})
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
	err := server.dbtx.DeleteClient(ctx, req.ID)
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
	err := server.dbtx.DeleteClientByName(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Cliente eliminado con éxito"})
}
