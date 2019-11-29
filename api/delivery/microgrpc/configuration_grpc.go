package microgrpc

import (
	"context"

	"github.com/muhammadhidayah/configuration-service/api"
	pb "github.com/muhammadhidayah/configuration-service/proto/configuration"
)

type microgrpc struct {
	uscase api.Usecase
}

func NewMicroGrpc(ucase api.Usecase) *microgrpc {
	return &microgrpc{ucase}
}

/**
 *
 *
 */
func (micro *microgrpc) GetConfigurationClient(ctx context.Context, req *pb.RequestConfigCient, res *pb.ResponseConfigClient) error {
	resp, err := micro.uscase.GetConfigurationClient(ctx)
	if err != nil {
		return err
	}

	res.Configclients = resp.GetConfigclients()
	return nil
}

func (micro *microgrpc) GetConfigurationClientBySubs(ctx context.Context, req *pb.RequestConfigCient, res *pb.ResponseConfigClient) error {
	companySubsID := req.Configclient.CompanySubsId

	resp, err := micro.uscase.GetConfigurationClientBySubs(ctx, companySubsID)
	if err != nil {
		return err
	}

	res.Configclient = resp.Configclient

	return nil
}

func (micro *microgrpc) AddConfigurationClient(ctx context.Context, req *pb.RequestConfigCient, res *pb.ResponseConfigClient) error {
	configClient := req.Configclient

	resp, err := micro.uscase.AddConfigurationClient(ctx, configClient)
	if err != nil {
		return err
	}

	res.Status = resp.GetStatus()
	res.Configclient = configClient

	return nil
}

func (micro *microgrpc) UpdateConfigurationClientBySubs(ctx context.Context, req *pb.RequestConfigCient, res *pb.ResponseConfigClient) error {
	configClient := req.Configclient

	resp, err := micro.uscase.UpdateConfigurationClientBySubs(ctx, configClient)
	if err != nil {
		return err
	}

	res.Status = resp.GetStatus()

	return nil
}

func (micro *microgrpc) DeleteConfigurationClientBySubs(ctx context.Context, req *pb.RequestConfigCient, res *pb.ResponseConfigClient) error {
	configClient := req.Configclient

	resp, err := micro.uscase.DeleteConfigurationClientBySubs(ctx, configClient)
	if err != nil {
		return err
	}

	res.Status = resp.GetStatus()

	return nil
}

func (micro *microgrpc) AddConfigurationGlobal(ctx context.Context, req *pb.RequestConfigGlobal, res *pb.ResponseConfigGlobal) error {
	configGlobal := req.Configglobal

	resp, err := micro.uscase.AddConfigurationGlobal(ctx, configGlobal)
	if err != nil {
		return err
	}

	res.Configstatus = resp.GetConfigstatus()
	res.Configglobal = configGlobal

	return nil
}

func (micro *microgrpc) UpdateConfigurationGlobal(ctx context.Context, req *pb.RequestConfigGlobal, res *pb.ResponseConfigGlobal) error {
	configGlobal := req.Configglobal

	resp, err := micro.uscase.UpdateConfigurationGlobal(ctx, configGlobal)
	if err != nil {
		return err
	}

	res.Configstatus = resp.GetConfigstatus()
	return nil
}

func (micro *microgrpc) DeleteConfiguration(ctx context.Context, req *pb.RequestConfigGlobal, res *pb.ResponseConfigGlobal) error {
	configGlobalID := req.Configglobal.GetConfigGlobalId()

	resp, err := micro.uscase.DeleteConfiguration(ctx, configGlobalID)
	if err != nil {
		return err
	}

	res.Configstatus = resp.GetConfigstatus()

	return nil
}

func (micro *microgrpc) GetConfigurationGlobal(ctx context.Context, req *pb.RequestConfigGlobal, res *pb.ResponseConfigGlobal) error {
	resp, err := micro.uscase.GetConfigurationGlobal(ctx)
	if err != nil {
		return err
	}

	res.Configglobals = resp.GetConfigglobals()
	return nil
}

func (micro *microgrpc) GetConfigurationGlobalByID(ctx context.Context, req *pb.RequestConfigGlobal, res *pb.ResponseConfigGlobal) error {
	configGlobalID := req.Configglobal.GetConfigGlobalId()

	resp, err := micro.uscase.GetConfigurationGlobalByID(ctx, configGlobalID)
	if err != nil {
		return err
	}

	res.Configglobal = resp.GetConfigglobal()
	return nil
}

func (micro *microgrpc) GetConfigurationGlobalActive(ctx context.Context, req *pb.RequestConfigGlobal, res *pb.ResponseConfigGlobal) error {
	resp, err := micro.uscase.GetConfigurationGlobalActive(ctx)
	if err != nil {
		return err
	}

	res.Configglobal = resp.GetConfigglobal()
	return nil
}

func (micro *microgrpc) SetConfigurationGlobalActive(ctx context.Context, req *pb.RequestConfigGlobal, res *pb.ResponseConfigGlobal) error {
	configGlobal := req.GetConfigglobal()

	resp, err := micro.uscase.SetConfigurationGlobalActive(ctx, configGlobal)
	if err != nil {
		res.Configstatus.Updated = false
		return err
	}

	res.Configstatus = resp.GetConfigstatus()
	return nil
}
