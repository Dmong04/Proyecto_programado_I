package api

import (
	"project/dto"
	"project/security"

	"github.com/gin-gonic/gin"
)

type Server struct {
	dbtx         *dto.DbTransaction
	tokenBuilder security.Builder
	router       *gin.Engine
}

func NewServer(dbtx *dto.DbTransaction) (*Server, error) {
	tokenBuilder, err := security.NewPasetoBuilder("87654321876543218765432187654321")
	if err != nil {
		return nil, err
	}
	server := &Server{
		dbtx:         dbtx,
		tokenBuilder: tokenBuilder,
	}
	router := gin.Default()
	// Rutas (Endpoints) De la API
	// CRUD Aministrador
	router.GET("/api/v1/Admin/all", server.GetAllAdmins)
	router.POST("/api/v1/Admin", server.CreateAdmin)
	router.GET("/api/v1/Admin/:id", server.GetAdminByID)
	router.GET("/api/v1/Admin/name/:name", server.GetAdminByName)
	router.PATCH("/api/v1/Admin/update/:id", server.UpdateAdmin)
	router.PATCH("/api/v1/Admin/update/password/:id", server.UpdateAdminPassword)
	router.DELETE("/api/v1/Admin/delete/:id", server.DeleteAdmin)
	router.DELETE("/api/v1/Admin/delete/name/:name", server.DeleteAdminByName)
	// CRUD Pasajeros
	router.GET("/api/v1/Passengers/all", server.GetAllPassengers)
	router.POST("/api/v1/Passengers", server.CreatePassenger)
	router.GET("/api/v1/Passengers/:id", server.GetPassengersById)
	router.GET("/api/v1/Passengers/detail/:detail_id", server.GetPassengersByDetailID)
	router.GET("/api/v1/Passengers/name/:name", server.GetPassengersByName)
	router.PATCH("/api/v1/Passengers/update/:id", server.UpdatePassengers)
	router.DELETE("/api/v1/Passengers/delete/:id", server.DeletePassenger)
	// CRUD DetalleViajes
	router.GET("api/v1/Details/all", server.getAllDetails)
	router.GET("api/v1/Details/:id", server.getDetailsByID)
	router.POST("api/v1/Details", server.CreateDetail)
	router.PATCH("api/v1/Details/update/:id", server.UpdateDetail)
	router.DELETE("api/v1/Details/delete/:id", server.DeleteDetail)
	// CRUD Client
	router.POST("api/v1/Client", server.CreateClient)
	router.GET("api/v1/Client/:id", server.GetClientByID)
	router.GET("api/v1/Client", server.GetAllClients)
	router.GET("api/v1/Client/:name", server.GetClientByName)
	router.PATCH("api/v1/Client/:id", server.UpdateClient)
	router.PATCH("api/v1/Client/:id", server.UpdateClientPassword)
	router.DELETE("api/v1/Client/:id", server.DeleteClient)
	router.DELETE("api/v1/Client/:name", server.DeleteClientByName)
	// CRUD Travel
	router.POST("api/v1/Travel", server.CreateTravel)
	router.GET("api/v1/Travel/all", server.GetAllTravels)
	router.GET("api/v1/Travel/:id", server.GetTravelById)
	router.DELETE("api/v1/Travel:id", server.DeleteTravel)
	router.PATCH("api/v1/Travel:id", server.UpdateTravel)
	router.GET("/api/v1/Details/all", server.getAllDetails)
	router.POST("/api/v1/Details", server.CreateDetail)
	router.GET("/api/v1/Details/:id", server.getDetailsByID)
	// Rutas (Endpoints) De la API
	server.router = router
	return server, nil
}

func (server *Server) Start(url string) error {
	return server.router.Run(url)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
