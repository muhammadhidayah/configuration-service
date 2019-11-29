package api

import (
	"context"

	pb "github.com/muhammadhidayah/configuration-service/proto/configuration"
)

type Usecase interface {
	GetConfigurationClient(context.Context) (*pb.ResponseConfigClient, error)
	GetConfigurationClientBySubs(context.Context, string) (*pb.ResponseConfigClient, error)
	AddConfigurationClient(context.Context, *pb.ConfigurationClient) (*pb.ResponseConfigClient, error)
	UpdateConfigurationClientBySubs(context.Context, *pb.ConfigurationClient) (*pb.ResponseConfigClient, error)
	DeleteConfigurationClientBySubs(context.Context, *pb.ConfigurationClient) (*pb.ResponseConfigClient, error)

	AddConfigurationGlobal(context.Context, *pb.ConfigurationGlobal) (*pb.ResponseConfigGlobal, error)
	UpdateConfigurationGlobal(context.Context, *pb.ConfigurationGlobal) (*pb.ResponseConfigGlobal, error)
	DeleteConfiguration(context.Context, int32) (*pb.ResponseConfigGlobal, error)
	GetConfigurationGlobal(context.Context) (*pb.ResponseConfigGlobal, error)
	GetConfigurationGlobalByID(context.Context, int32) (*pb.ResponseConfigGlobal, error)
	GetConfigurationGlobalActive(context.Context) (*pb.ResponseConfigGlobal, error)
	SetConfigurationGlobalActive(context.Context, *pb.ConfigurationGlobal) (*pb.ResponseConfigGlobal, error)
}
