package db

import (
	"go-fiber-mongo-hrms/src/envs"
	"go-fiber-mongo-hrms/src/types"
)

func GetConfig() *types.MongoDBConfig {
	return &types.MongoDBConfig{
		PublicHost: envs.Envs.PublicHost,
		Port:       envs.Envs.Port,
		DBUser:     envs.Envs.DBUser,
		DBPassword: envs.Envs.DBPassword,
		DBAddress:  envs.Envs.DBAddress,
		DBName:     envs.Envs.DBName,
	}
}
