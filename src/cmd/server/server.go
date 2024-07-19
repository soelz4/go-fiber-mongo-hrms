package server

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/mongo"

	"go-fiber-mongo-hrms/src/services/employee"
)

// Server
type Server struct {
	address  string
	client   *mongo.Client
	database *mongo.Database
}

// Return New Server
func NewServer(address string, client *mongo.Client, database *mongo.Database) *Server {
	return &Server{address: address, client: client, database: database}
}

// Server RUN
func (s *Server) Run() error {
	// Fiber New App
	app := fiber.New()

	// Employee Handler
	employeeStore := employee.NewStore(s.client, s.database)
	employeeHandler := employee.NewHandler(employeeStore)

	// Routes
	employeeHandler.RegisterRoutes(app)

	// Logs
	log.Println("Listening On", s.address)

	// Create Server
	return app.Listen(s.address)
}
