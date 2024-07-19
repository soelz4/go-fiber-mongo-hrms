package employee

import (
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"go-fiber-mongo-hrms/src/types"
)

func (s *Store) FindAllEmployees(c fiber.Ctx) (*mongo.Cursor, error) {
	// Query
	query := bson.D{{}}

	// Choose Your Database and Collection
	collection := s.database.Collection("employee")

	// Cursor
	cursor, err := collection.Find(c.Context(), query)

	// Return
	return cursor, err
}

func (s *Store) FindEmployeeByID(c fiber.Ctx, id string) (types.Employee, error) {
	var employee types.Employee

	// ObjectIDFromHex Creates a New ObjectID from a Hex String. It Returns an Error if the Hex String is not a Valid ObjectID.
	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return employee, c.SendStatus(400)
	}

	// Query
	query := bson.M{"_id": employeeID}

	// Choose Your Database and Collection
	collection := s.database.Collection("employee")

	// Find Document by ID - Filter by ID
	err = collection.FindOne(c.Context(), query).Decode(&employee)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return employee, err
		} else {
			return employee, err
		}
	}

	return employee, err
}

func (s *Store) InsertEmployee(c fiber.Ctx, employee *types.Employee) (*types.Employee, error) {
	// Choose Your Database and Collection
	collection := s.database.Collection("employee")

	// Created Employee
	createdEmployee := &types.Employee{}

	employee.ID = ""

	// Insert Record
	insertResult, err := collection.InsertOne(c.Context(), employee)
	if err != nil {
		return nil, c.Status(500).SendString(err.Error())
	}

	// Filter
	filter := bson.D{{Key: "_id", Value: insertResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	// Decode into createdEmployee Var
	err = createdRecord.Decode(createdEmployee)
	if err != nil {
		return nil, c.Status(400).SendString(err.Error())
	}

	// Return
	return createdEmployee, err
}

func (s *Store) DeleteEmployeeByID(c fiber.Ctx, id string) error {
	// ObjectIDFromHex Creates a New ObjectID from a Hex String. It Returns an Error if the Hex String is not a Valid ObjectID.
	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.SendStatus(400)
	}

	// Query
	query := bson.D{{Key: "_id", Value: employeeID}}

	// Choose Your Database and Collection
	collection := s.database.Collection("employee")

	result, err := collection.DeleteOne(c.Context(), &query)
	if err != nil {
		return c.SendStatus(500)
	}
	if result.DeletedCount < 1 {
		return c.SendStatus(404)
	}

	return nil
}

func (s *Store) UpdateEmployeeByID(
	c fiber.Ctx,
	id string,
	employee *types.Employee,
) (*types.Employee, error) {
	updatedEmployee := &types.Employee{}

	// ObjectIDFromHex Creates a New ObjectID from a Hex String. It Returns an Error if the Hex String is not a Valid ObjectID.
	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, c.SendStatus(400)
	}

	// Choose Your Database and Collection
	collection := s.database.Collection("employee")

	// Query
	query := bson.D{{Key: "_id", Value: employeeID}}
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "name", Value: employee.Name},
				{Key: "age", Value: employee.Age},
				{Key: "salary", Value: employee.Salary},
			},
		},
	}

	result := collection.FindOneAndUpdate(c.Context(), query, update)

	err = result.Decode(updatedEmployee)
	if err != nil {
		return nil, c.Status(400).SendString(err.Error())
	}

	err = result.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, c.SendStatus(400)
		}
		return nil, c.SendStatus(500)
	}

	updatedEmployee.ID = id

	return updatedEmployee, err
}
