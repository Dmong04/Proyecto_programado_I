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
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))
	// Rutas (Endpoints) De la API
	// CRUD Aministrador
	router.GET("/api/v1/Admin/all", server.GetAllAdmins)
	router.POST("/api/v1/Admin", server.CreateAdmin)
	router.GET("/api/v1/Admin/:id", server.GetAdminByID)
	router.GET("/api/v1/Admin/name/:name", server.GetAdminByName)
	router.PATCH("/api/v1/Admin/update/:id", server.UpdateAdmin)
	router.DELETE("/api/v1/Admin/delete/:id", server.DeleteAdmin)
	router.DELETE("/api/v1/Admin/delete/name/:name", server.DeleteAdminByName)
	// CRUD Pasajeros
	router.GET("/api/v1/Passengers/all", server.GetAllPassengers)
	router.POST("/api/v1/Passengers", server.CreatePassenger)
	router.GET("/api/v1/Passengers/:id", server.GetPassengersById)
	router.GET("/api/v1/Passengers/detail/:detailId", server.GetPassengersByDetailID)
	router.GET("/api/v1/Passengers/name/:name", server.GetPassengersByName)
	router.PATCH("/api/v1/Passengers/update/:id", server.UpdatePassengers)
	router.DELETE("/api/v1/Passengers/delete/:id", server.DeletePassenger)
	// CRUD DetalleViajes
	router.GET("api/v1/Details/all", server.getAllDetails)
	router.GET("api/v1/Details:id", server.getDetailsByID)
	router.POST("api/v1/Details", server.CreateDetail)
	router.PATCH("api/v1/Details:id", server.UpdateDetail)
	router.DELETE("api/v1/Details:id", server.DeleteDetail)
	// CRUD Client
	router.POST("/api/v1/Client", server.CreateClient)
	router.GET("/api/v1/Client/id/:id", server.GetClientByID)
	router.GET("/api/v1/Client/all", server.GetAllClients)
	router.GET("/api/v1/Client/name/:name", server.GetClientByName)
	router.PATCH("/api/v1/Client/update/id/:id", server.UpdateClient)
	router.DELETE("/api/v1/Client/delete/id/:id", server.DeleteClient)
	router.DELETE("/api/v1/Client/delete/name/:name", server.DeleteClientByName)

	// CRUD Travel
	router.GET("/api/v1/Travel/all", server.GetAllTravels)
	router.GET("/api/v1/Travel/id/:id", server.GetTravelById)
	router.DELETE("/api/v1/Travel/delete/:id", server.DeleteTravel)
	router.PATCH("/api/v1/Travel/update/:id", server.UpdateTravel)

	//RUTAS CON MIDDLEWARE
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenBuilder))
	authRoutes.POST("api/v1/Travel", server.CreateTravel)
	authRoutes.PUT("api/v1/User/:id", server.updateUser)
	authRoutes.PUT("api/v1/User", server.updateUser)
	authRoutes.PUT("api/v1/Userpass", server.updatePassword)
	authRoutes.PUT("api/v1/Userrole", server.updateRole)

	// FIN RUTAS
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
