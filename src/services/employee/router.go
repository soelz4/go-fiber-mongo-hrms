package employee

import (
	"github.com/gofiber/fiber/v3"

	"go-fiber-mongo-hrms/src/types"
)

// Employee Handler
type Handler struct {
	store types.EmployeeStore
}

// Return New Employee Handler
func NewHandler(store types.EmployeeStore) *Handler {
	return &Handler{store: store}
}

// Routes
func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Get("/api/v1/employees", h.handleGetEmployees)
	app.Get("/api/v1/employee/:id", h.handleGetEmployeeByID)
	app.Post("/api/v1/newemployee", h.handleNewEmployee)
	app.Put("/api/v1/employee/:id", h.handleUpdateEmployeeByID)
	app.Delete("/api/v1/employee/:id", h.handleDeleteEmployeeByID)
}
