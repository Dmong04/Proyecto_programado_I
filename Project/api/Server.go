package api

import (
	"project/dto"
	"project/security"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
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
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE, PATCH",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))
	auth := authMiddleware(tokenBuilder)
	router.POST("api/v1/login", server.login)
	router.POST("api/v1/User", server.createUser)
	router.POST("api/v1/Client", server.CreateClient) //(funciona la v2)
	// Rutas (Endpoints) De la API
	adminRoutes := router.Group("/")
	clientRoutes := router.Group("/")
	sharedRoutes := router.Group("/")
	sharedRoutes.Use(auth, roleMiddleware("Admin", "Client"))
	{
		// Rutas de consulta de clientes
		sharedRoutes.PATCH("api/v1/Client/update/:id", server.UpdateClient)
		sharedRoutes.DELETE("api/v1/Client/delete/:id", server.DeleteClient)
		sharedRoutes.GET("api/v1/Client/name/:name", server.GetClientByName)
		sharedRoutes.GET("api/v1/Client/id/:id", server.GetClientByID)
		// Rutas de gestión de reservas
		sharedRoutes.GET("api/v1/Reservations/all", server.GetAllReservations)          // (funciona)
		sharedRoutes.GET("api/v1/Reservations/:id", server.GetReservationsById)         // (funciona)
		sharedRoutes.PATCH("api/v1/Reservations/update/:id", server.UpdateReservations) // (funciona)
		sharedRoutes.DELETE("api/v1/Reservations/delete/:id", server.DeleteReservation) // (funciona)
		sharedRoutes.POST("api/v1/Reservations", server.CreateReservation)              // (Funciona)
		// Gestión en los detalles del viaje
		sharedRoutes.GET("api/v1/Details/all", server.getAllDetails)
		sharedRoutes.GET("api/v1/Details/:id", server.getDetailsByID)
		sharedRoutes.POST("api/v1/Details", server.CreateDetail)
		sharedRoutes.PATCH("api/v1/Details/update/:id", server.UpdateDetail)
		sharedRoutes.DELETE("api/v1/Details/delete/:id", server.DeleteDetail)
		// Gestion de viaje
		sharedRoutes.GET("api/v1/Travel/all", server.GetAllTravels)
		sharedRoutes.GET("api/v1/Travel/:id", server.GetTravelById)
		// Gestion de proveedor
		sharedRoutes.GET("api/v1/Provider/all", server.GetAllProviders)
		sharedRoutes.GET("api/v1/Provider/:id", server.GetProviderByID)
		sharedRoutes.GET("api/v1/Provider/name/:name", server.GetProviderByName)
		// Gestión de usuario
		sharedRoutes.PATCH("api/v1/User/update/:id", server.updateUser) // (Funciona)
		sharedRoutes.PATCH("api/v1/User/password/:id", server.updatePassword)
		sharedRoutes.DELETE("api/v1/User/delete/:id", server.deleteUser) // (Funciona)
	}
	adminRoutes.Use(auth, roleMiddleware("Admin"))
	{
		// All clients
		adminRoutes.GET("api/v1/Client/all", server.GetAllClients)
		// CRUD Aministrador (Funciona)
		adminRoutes.GET("api/v1/Admin/all", server.GetAllAdmins)
		adminRoutes.POST("api/v1/Admin", server.CreateAdmin)
		adminRoutes.GET("api/v1/Admin/:id", server.GetAdminByID)
		adminRoutes.GET("api/v1/Admin/name/:name", server.GetAdminByName)
		adminRoutes.PATCH("api/v1/Admin/update/:id", server.UpdateAdmin)
		adminRoutes.DELETE("api/v1/Admin/delete/:id", server.DeleteAdmin)
		adminRoutes.DELETE("api/v1/Admin/delete/name/:name", server.DeleteAdminByName)
		//CRUD Proveedor (Funciona)
		adminRoutes.POST("api/v1/Provider", server.CreateProvider)
		adminRoutes.PATCH("api/v1/Provider/update/:id", server.UpdateProvider)
		adminRoutes.PATCH("api/v1/Provider/update/name/:name", server.UpdateProviderByName)
		adminRoutes.DELETE("api/v1/Provider/delete/:id", server.DeleteProvider)
		adminRoutes.DELETE("api/v1/Provider/delete/name/:name", server.DeleteProviderByName)
		// CRUD Travel (Funciona)
		adminRoutes.POST("api/v1/Travel", server.CreateTravel)
		adminRoutes.DELETE("api/v1/Travel/delete/:id", server.DeleteTravel)
		adminRoutes.PATCH("api/v1/Travel/update/:id", server.UpdateTravel)
		// GESTION USUARIO
		adminRoutes.GET("api/v1/User/all", server.getAllUsers)                  // (Funciona)
		adminRoutes.GET("api/v1/User/:id", server.getUserById)                  // (Funciona)
		adminRoutes.GET("api/v1/User/UserName/:user", server.getUserByUserName) // (Funciona)
		// Update status de reserva
		adminRoutes.PATCH("api/v1/Reservations/status/:id", server.UpdateStatus) // ()
	}
	clientRoutes.Use(auth, roleMiddleware("Client"))
	{
		// CRUD Client (Funciona)
		clientRoutes.DELETE("api/v1/Client/delete/name/:name", server.DeleteClientByName)
		// CRUD DetalleViajes (Funciona)
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
