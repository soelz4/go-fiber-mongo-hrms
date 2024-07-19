package employee

import "go.mongodb.org/mongo-driver/mongo"

// Employee Store
type Store struct {
	client   *mongo.Client
	database *mongo.Database
}

// Return New Employee Store
func NewStore(client *mongo.Client, database *mongo.Database) *Store {
	return &Store{client: client, database: database}
}
