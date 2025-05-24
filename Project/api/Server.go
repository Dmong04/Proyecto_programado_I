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
	auth := authMiddleware(tokenBuilder)
	router.POST("api/v1/login", server.login)
	// Rutas (Endpoints) De la API
	adminRoutes := router.Group("/")
	clientRoutes := router.Group("/")
	sharedRoutes := router.Group("/")
	router.POST("api/v1/Client", server.CreateClient)
	sharedRoutes.Use(auth, roleMiddleware("Admin", "Client"))
	{
		// Rutas de consulta de clientes
		sharedRoutes.GET("api/v1/Client/all", server.GetAllClients)
		sharedRoutes.GET("api/v1/Client/name/:name", server.GetClientByName)
		sharedRoutes.GET("api/v1/Client/id/:id", server.GetClientByID)
		// Rutas de consulta a teléfonos de clientes
		sharedRoutes.GET("api/v1/Client/Phones/:id", server.GetClientPhonesById)
		sharedRoutes.GET("api/v1/Client/Phones/client/:idclient", server.GetClientPhonesByClientID)
		sharedRoutes.GET("api/v1/Client/Phones/all", server.GetAllClientPhones)
		// Rutas de gestión de reservas
		sharedRoutes.GET("api/v1/Reservations/all", server.GetAllReservations)
		sharedRoutes.GET("api/v1/Reservations/:id", server.GetReservationsById)
		sharedRoutes.PATCH("api/v1/Reservations/update/:id", server.UpdateReservations)
		sharedRoutes.DELETE("api/v1/Reservations/delete/:id", server.DeleteReservation)
		// Rutas de gestión de pasajeros
		sharedRoutes.GET("api/v1/Passengers/all", server.GetAllPassengers)
		sharedRoutes.GET("api/v1/Passengers/:id", server.GetPassengersById)
		sharedRoutes.GET("api/v1/Passengers/detail/:detail_id", server.GetPassengersByDetailID)
		sharedRoutes.GET("api/v1/Passengers/name/:name", server.GetPassengersByName)
		// Gestión en los detalles del viaje
		sharedRoutes.GET("api/v1/Details/all", server.getAllDetails)
		sharedRoutes.GET("api/v1/Details/:id", server.getDetailsByID)
	}
	adminRoutes.Use(auth, roleMiddleware("Admin"))
	{
		// CRUD Aministrador (Funciona)
		adminRoutes.GET("api/v1/Admin/all", server.GetAllAdmins)
		adminRoutes.POST("api/v1/Admin", server.CreateAdmin)
		adminRoutes.GET("api/v1/Admin/:id", server.GetAdminByID)
		adminRoutes.GET("api/v1/Admin/name/:name", server.GetAdminByName)
		adminRoutes.PATCH("api/v1/Admin/update/:id", server.UpdateAdmin)
		adminRoutes.PATCH("api/v1/Admin/update/password/:id", server.UpdateAdminPassword)
		adminRoutes.DELETE("api/v1/Admin/delete/:id", server.DeleteAdmin)
		adminRoutes.DELETE("api/v1/Admin/delete/name/:name", server.DeleteAdminByName)
		//CRUD Proveedor (Funciona)
		adminRoutes.GET("api/v1/Provider/all", server.GetAllProviders)
		adminRoutes.POST("api/v1/Provider", server.CreateProvider)
		adminRoutes.GET("api/v1/Provider/:id", server.GetProviderByID)
		adminRoutes.GET("api/v1/Provider/name/:name", server.GetProviderByName)
		adminRoutes.PATCH("api/v1/Provider/update/:id", server.UpdateProvider)
		adminRoutes.PATCH("api/v1/Provider/update/name/:name", server.UpdateProviderByName)
		adminRoutes.DELETE("api/v1/Provider/delete/:id", server.DeleteProvider)
		adminRoutes.DELETE("api/v1/Provider/delete/name/:name", server.DeleteProviderByName)
		// CRUD Travel (Funciona)
		adminRoutes.POST("api/v1/Travel", server.CreateTravel)
		adminRoutes.GET("api/v1/Travel/all", server.GetAllTravels)
		adminRoutes.GET("api/v1/Travel/:id", server.GetTravelById)
		adminRoutes.DELETE("api/v1/Travel/delete/:id", server.DeleteTravel)
		adminRoutes.PATCH("api/v1/Travel/update/:id", server.UpdateTravel)
	}
	clientRoutes.Use(auth, roleMiddleware("Client"))
	{
		// CRUD Client (Funciona)
		clientRoutes.PATCH("api/v1/Client/update/:id", server.UpdateClient)
		clientRoutes.PATCH("api/v1/Client/password/:id", server.UpdateClientPassword)
		clientRoutes.DELETE("api/v1/Client/delete/:id", server.DeleteClient)
		clientRoutes.DELETE("api/v1/Client/delete/name/:name", server.DeleteClientByName)
		// CRUD Pasajeros (Funciona)
		clientRoutes.POST("api/v1/Passengers", server.CreatePassenger)
		clientRoutes.PATCH("api/v1/Passengers/update/:id", server.UpdatePassengers)
		clientRoutes.DELETE("api/v1/Passengers/delete/:id", server.DeletePassenger)
		// CRUD DetalleViajes (Funciona)
		clientRoutes.POST("api/v1/Details", server.CreateDetail)
		clientRoutes.PATCH("api/v1/Details/update/:id", server.UpdateDetail)
		clientRoutes.DELETE("api/v1/Details/delete/:id", server.DeleteDetail)
		// CRUD ClientPhones (Funciona)
		clientRoutes.POST("api/v1/Client/Phones", server.CreateClientPhones)
		clientRoutes.PATCH("api/v1/Client/Phones/update/:id", server.UpdateClientPhones)
		clientRoutes.DELETE("api/v1/Client/Phones/delete/:id", server.DeleteClientPhones)
		// CRUD Reservaciones (Funciona)
		clientRoutes.POST("api/v1/Reservations", server.CreateReservation)
	}
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
