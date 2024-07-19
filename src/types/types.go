package types

import (
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDB Config
type MongoDBConfig struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

// Employee
type Employee struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    float64 `json:"age"`
}

// Employee Store
type EmployeeStore interface {
	FindAllEmployees(c fiber.Ctx) (*mongo.Cursor, error)
	FindEmployeeByID(c fiber.Ctx, id string) (Employee, error)
	InsertEmployee(c fiber.Ctx, employee *Employee) (*Employee, error)
	DeleteEmployeeByID(c fiber.Ctx, id string) error
	UpdateEmployeeByID(c fiber.Ctx, id string, employee *Employee) (*Employee, error)
}
