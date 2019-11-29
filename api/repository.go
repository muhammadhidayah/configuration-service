package api

import (
	"context"

	pb "github.com/muhammadhidayah/configuration-service/proto/configuration"
)

type Repository interface {
	GetConfigurationClient(context.Context) ([]*pb.ConfigurationClient, error)
	GetConfigurationClientBySubs(context.Context, string) (*pb.ConfigurationClient, error)
	AddConfigurationClient(context.Context, *pb.ConfigurationClient) (bool, error)
	UpdateConfigurationClientBySubs(context.Context, *pb.ConfigurationClient) (bool, error)
	DeleteConfigurationClientBySubs(context.Context, *pb.ConfigurationClient) (bool, error)

	AddConfigurationGlobal(context.Context, *pb.ConfigurationGlobal) (bool, error)
	UpdateConfigurationGlobal(context.Context, *pb.ConfigurationGlobal) (bool, error)
	DeleteConfiguration(context.Context, int32) (bool, error)
	GetConfigurationGlobal(context.Context) ([]*pb.ConfigurationGlobal, error)
	GetConfigurationGlobalByID(context.Context, int32) (*pb.ConfigurationGlobal, error)
	GetConfigurationGlobalActive(context.Context) (*pb.ConfigurationGlobal, error)
}
