package api

import (
	"project/dto"

	"github.com/gin-gonic/gin"
)

type Server struct {
	dbtx   *dto.DbTransaction
	router *gin.Engine
}

func NewServer(dbtx *dto.DbTransaction) *Server {
	server := &Server{dbtx: dbtx}
	router := gin.Default()
	// Rutas (Endpoints) De la API
	// CRUD Aministrador
	router.POST("api/v1/Admin", server.CreateAdmin)
	router.GET("api/v1/Admin", server.GetAllAdmins)
	router.GET("api/v1/Admin:id", server.GetCategoryByID)
	router.GET("api/v1/Admin:name", server.GetAdminByName)
	router.PATCH("api/v1/Admin:id", server.UpdateAdmin)
	router.PATCH("api/v1/Admin:id", server.UpdateAdminPassword)
	router.DELETE("api/v1/Admin:id", server.DeleteAdmin)
	router.DELETE("api/v1/Admin:name", server.DeleteAdminByName)
	// CRUD Pasajeros
	router.POST("api/v1/Passengers", server.CreatePassenger)
	router.GET("api/v1/Passengers", server.GetAllPassengers)
	router.GET("api/v1/Passengers:id", server.GetPassengersById)
	router.GET("api/v1/Passengers:detailId", server.GetPassengersByDetailID)
	router.GET("api/v1/Passengers:name", server.GetPassengersByName)
	router.PATCH("api/v1/Passengers:id", server.UpdatePassengers)
	router.DELETE("api/v1/Passengers:id", server.DeletePassenger)
	// CRUD DetalleViajes
	router.GET("api/v1/Details", server.getAllDetails)
	router.GET("api/v1/Details:id", server.getDetailsByID)
	router.POST("api/v1/Details", server.CreateDetail)
	router.PATCH("api/v1/Details:id", server.UpdateDetail)
	router.DELETE("api/v1/Details:id", server.DeleteDetail)
	// CRUD Client
	router.POST("api/v1/Client", server.CreateClient)
	router.GET("api/v1/Client:id", server.GetClientByID)
	router.GET("api/v1/Client", server.GetAllClients)
	router.GET("api/v1/Client:name", server.GetClientByName)
	router.PATCH("api/v1/Client:id", server.UpdateClient)
	router.PATCH("api/v1/Client:id", server.UpdateClientPassword)
	router.DELETE("api/v1/Client:id", server.DeleteClient)
	router.DELETE("api/v1/Client:name", server.DeleteClientByName)
	// CRUD Travel
	router.POST("api/v1/Travel", server.CreateTravel)
	router.GET("api/v1/Travel", server.GetAllTravels)
	router.GET("api/v1/Travel:id", server.GetTravelById)
	router.DELETE("api/v1/Travel:id", server.DeleteTravel)
	router.PATCH("api/v1/Travel:id", server.UpdateTravel)
	// Rutas (Endpoints) De la API
	server.router = router
	return server
}

func (server *Server) Start(url string) error {
	return server.router.Run(url)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
