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
	// CRUD Aministrador (Funciona)
	router.GET("/api/v1/Admin/all", server.GetAllAdmins)
	router.POST("/api/v1/Admin", server.CreateAdmin)
	router.GET("/api/v1/Admin/:id", server.GetAdminByID)
	router.GET("/api/v1/Admin/name/:name", server.GetAdminByName)
	router.PATCH("/api/v1/Admin/update/:id", server.UpdateAdmin)
	router.PATCH("/api/v1/Admin/update/password/:id", server.UpdateAdminPassword)
	router.DELETE("/api/v1/Admin/delete/:id", server.DeleteAdmin)
	router.DELETE("/api/v1/Admin/delete/name/:name", server.DeleteAdminByName)
	// CRUD Pasajeros (Funciona)
	router.GET("/api/v1/Passengers/all", server.GetAllPassengers)
	router.POST("/api/v1/Passengers", server.CreatePassenger)
	router.GET("/api/v1/Passengers/:id", server.GetPassengersById)
	router.GET("/api/v1/Passengers/detail/:detail_id", server.GetPassengersByDetailID)
	router.GET("/api/v1/Passengers/name/:name", server.GetPassengersByName)
	router.PATCH("/api/v1/Passengers/update/:id", server.UpdatePassengers)
	router.DELETE("/api/v1/Passengers/delete/:id", server.DeletePassenger)
	// CRUD DetalleViajes (Funciona)
	router.GET("/api/v1/Details/all", server.getAllDetails)
	router.POST("/api/v1/Details", server.CreateDetail)
	router.GET("/api/v1/Details/:id", server.getDetailsByID)
	router.PATCH("/api/v1/Details/update/:id", server.UpdateDetail)
	router.DELETE("/api/v1/Details/delete/:id", server.DeleteDetail)
	//CRUD Historial (Funciona)
	router.GET("/api/v1/History/all", server.GetAllHistories)
	router.POST("/api/v1/History", server.CreateHistory)
	router.GET("/api/v1/History/:id", server.GetHistoryByID)
	router.PATCH("/api/v1/History/update/:id", server.UpdateHistory)
	router.DELETE("/api/v1/History/delete/:id", server.DeleteHistory)
	//CRUD Proveedor (Funciona)
	router.GET("/api/v1/Provider/all", server.GetAllProviders)
	router.POST("/api/v1/Provider", server.CreateProvider)
	router.GET("/api/v1/Provider/:id", server.GetProviderByID)
	router.GET("/api/v1/Provider/name/:name", server.GetProviderByName)
	router.PATCH("/api/v1/Provider/update/:id", server.UpdateProvider)
	router.PATCH("/api/v1/Provider/update/name/:name", server.UpdateProviderByName)
	router.DELETE("/api/v1/Provider/delete/:id", server.DeleteProvider)
	router.DELETE("/api/v1/Provider/delete/name/:name", server.DeleteProviderByName)
	// CRUD Client (Funciona)
	router.POST("/api/v1/Client", server.CreateClient)
	router.GET("/api/v1/Client/id/:id", server.GetClientByID)
	router.GET("/api/v1/Client/all", server.GetAllClients)
	router.GET("/api/v1/Client/name/:name", server.GetClientByName)
	router.PATCH("/api/v1/Client/update/:id", server.UpdateClient)
	router.PATCH("/api/v1/Client/password/:id", server.UpdateClientPassword)
	router.DELETE("/api/v1/Client/delete/:id", server.DeleteClient)
	router.DELETE("/api/v1/Client/delete/name/:name", server.DeleteClientByName)
	// CRUD ClientPhones (Funciona)
	router.POST("/api/v1/Client/Phones", server.CreateClientPhones)
	router.GET("/api/v1/Client/Phones/id/:id", server.GetClientPhonesById)
	router.GET("/api/v1/Client/Phones/client/:idclient", server.GetClientPhonesByClientID)
	router.GET("/api/v1/Client/Phones/all", server.GetAllClientPhones)
	router.PATCH("/api/v1/Client/Phones/update/:id", server.UpdateClientPhones)
	router.DELETE("/api/v1/Client/Phones/delete/:id", server.DeleteClientPhones)
	// CRUD Reservaciones
	router.POST("/api/v1/Reservations", server.CreateReservation)
	router.GET("/api/v1/Reservations/all", server.GetAllReservations)
	router.GET("/api/v1/Reservations/:id", server.GetReservationsById)
	router.PATCH("/api/v1/Reservations/update/:id", server.UpdateReservations)
	router.DELETE("/api/v1/Reservations/delete/:id", server.DeleteReservation)
	// CRUD Travel (Funciona)
	router.POST("/api/v1/Travel", server.CreateTravel)
	router.GET("/api/v1/Travel/all", server.GetAllTravels)
	router.GET("/api/v1/Travel/:id", server.GetTravelById)
	router.DELETE("/api/v1/Travel/delete/:id", server.DeleteTravel)
	router.PATCH("/api/v1/Travel/update/:id", server.UpdateTravel)
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
