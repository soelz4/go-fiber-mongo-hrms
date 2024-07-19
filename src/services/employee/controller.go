package employee

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"

	"go-fiber-mongo-hrms/src/types"
)

func (h *Handler) handleGetEmployees(c fiber.Ctx) error {
	// Find All Employees
	cursor, err := h.store.FindAllEmployees(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// List of Employees
	employees := make([]types.Employee, 0)

	// Insert into ~> employees
	err = cursor.All(c.Context(), &employees)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Return
	return c.JSON(employees)
}

func (h *Handler) handleGetEmployeeByID(c fiber.Ctx) error {
	// Find Employee By ID
	id := c.Params("id")
	var employee types.Employee

	employee, err := h.store.FindEmployeeByID(c, id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Return
	return c.JSON(employee)
}

func (h *Handler) handleNewEmployee(c fiber.Ctx) error {
	employee := new(types.Employee)

	// Request Body
	body := c.Body()
	// Unmarshal Parses the JSON-Encoded Data and Stores the Result
	err := json.Unmarshal(body, employee)
	if err != nil {
		return c.Status(503).SendString("JSON Decoding Failed")
	}

	// Insert Data into DataBase and Return OutPut
	createdEmployee, err := h.store.InsertEmployee(c, employee)
	if err != nil {
		c.Status(500).SendString(err.Error())
	}

	// Return
	return c.JSON(createdEmployee)
}

func (h *Handler) handleUpdateEmployeeByID(c fiber.Ctx) error {
	// Find Employee By ID
	id := c.Params("id")

	employee := new(types.Employee)

	// Request Body
	body := c.Body()
	// Unmarshal Parses the JSON-Encoded Data and Stores the Result
	err := json.Unmarshal(body, employee)
	if err != nil {
		return c.Status(503).SendString("JSON Decoding Failed")
	}

	updatedEmployee, err := h.store.UpdateEmployeeByID(c, id, employee)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Return
	return c.JSON(updatedEmployee)
}

func (h *Handler) handleDeleteEmployeeByID(c fiber.Ctx) error {
	// Find Employee By ID
	id := c.Params("id")

	// Delete Data from DataBase
	err := h.store.DeleteEmployeeByID(c, id)
	if err != nil {
		c.Status(500).SendString(err.Error())
	}

	return c.SendString("Record with ID " + id + " Deleted")
}
