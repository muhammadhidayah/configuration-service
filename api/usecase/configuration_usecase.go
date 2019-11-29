package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/muhammadhidayah/configuration-service/api"
	pb "github.com/muhammadhidayah/configuration-service/proto/configuration"
)

type configurationUseCase struct {
	configRepo     api.Repository
	contextTimeout time.Duration
}

func NewConfigurationUsecase(repo api.Repository, timeout time.Duration) api.Usecase {
	return &configurationUseCase{repo, timeout}
}

// this function will return pointer of ResponseConfigClient and Error. this function will call GetConfigurationClient method of Repository to get all data in table configuration_client
func (ucase *configurationUseCase) GetConfigurationClient(ctx context.Context) (*pb.ResponseConfigClient, error) {
	// created context time out to cancel process database
	c, cancel := context.WithTimeout(ctx, ucase.contextTimeout)

	// cancel will called if time morethan field timeout
	defer cancel()

	// call GetConfigurationClient method of Repository
	listConfigClient, err := ucase.configRepo.GetConfigurationClient(c)
	if err != nil {
		return nil, err
	}

	// store listConfiguration in respConfigClient
	respConfigClient := &pb.ResponseConfigClient{
		Configclients: listConfigClient,
	}

	return respConfigClient, nil
}

// this function will return pointer of ResponseConfigClient and Error. this function will call GetConfigurationClientBySubs method of Repository to get one data in table configuration_client with condition company_subs_id equal subsID
func (ucase *configurationUseCase) GetConfigurationClientBySubs(c context.Context, subsID string) (*pb.ResponseConfigClient, error) {
	// created context time out to cancel process database
	ctx, cancel := context.WithTimeout(c, ucase.contextTimeout)

	// cancel and terminated all process will called if time morethan field timeout
	defer cancel()

	// call GetConfigurationClientBySubs method of Repository
	configClient, err := ucase.configRepo.GetConfigurationClientBySubs(ctx, subsID)
	if err != nil {
		return nil, err
	}

	// store configClient in respConfigClient
	respConfigClient := &pb.ResponseConfigClient{
		Configclient: configClient,
	}

	return respConfigClient, nil
}

func (ucase *configurationUseCase) AddConfigurationClient(c context.Context, cc *pb.ConfigurationClient) (*pb.ResponseConfigClient, error) {

	// create variable to contain struct responseConfigClient. for first initiate will set status.Created is false
	respConfigC := &pb.ResponseConfigClient{
		Status: &pb.ConfigurationStatus{Created: false},
	}

	// generate uuid for configClientUuid
	configClientUuid, err := uuid.NewV4()
	if err != nil {
		return respConfigC, err
	}

	// store uuid to ConfigClientId field, uuid is result of generated before
	cc.ConfigClientUuid = configClientUuid.String()

	// create context timeout to cancel process database
	ctx, cancel := context.WithTimeout(c, ucase.contextTimeout)

	defer cancel()

	// call function
	resp, err := ucase.configRepo.AddConfigurationClient(ctx, cc)
	if err != nil {
		return respConfigC, err
	}

	// store configClient in respConfigClient
	respConfigC.Status.Created = resp
	respConfigC.Configclient = cc

	return respConfigC, nil
}

func (ucase *configurationUseCase) UpdateConfigurationClientBySubs(c context.Context, cc *pb.ConfigurationClient) (*pb.ResponseConfigClient, error) {
	// create variable to contain struct responseConfigClient. for first initiate will set status.Updated is false
	responseConfigC := &pb.ResponseConfigClient{
		Status: &pb.ConfigurationStatus{
			Updated: false,
		},
	}

	// create context timeout to cancel process database when process to long
	ctx, cancel := context.WithTimeout(c, ucase.contextTimeout)

	defer cancel()

	// call UpdateConfigurationClientBySubs method of configRepo to update data in table configuration_client
	resp, err := ucase.configRepo.UpdateConfigurationClientBySubs(ctx, cc)
	if err != nil {
		return responseConfigC, err
	}

	// update value status.updated
	responseConfigC.Status.Updated = resp

	return responseConfigC, nil
}

// this function will change status is_delete to 1. actually not really remove from db. the function will return struct of ResponseConfigClient and error
func (ucase *configurationUseCase) DeleteConfigurationClientBySubs(c context.Context, cc *pb.ConfigurationClient) (*pb.ResponseConfigClient, error) {
	// create variable to contain struct responseConfigClient. for first initiate will set status.Deleted is false
	responseConfigC := &pb.ResponseConfigClient{
		Status: &pb.ConfigurationStatus{
			Deleted: false,
		},
	}

	// create context timeout to cancel process database when process to long
	ctx, cancel := context.WithTimeout(c, ucase.contextTimeout)

	defer cancel()

	// call DeleteConfigurationClientBySubs method of configRepo to change status is deleted to 1
	res, err := ucase.configRepo.DeleteConfigurationClientBySubs(ctx, cc)
	if err != nil {
		return responseConfigC, err
	}

	// update value status.deleted
	responseConfigC.Status.Deleted = res

	return responseConfigC, nil
}

func (ucase *configurationUseCase) AddConfigurationGlobal(c context.Context, cg *pb.ConfigurationGlobal) (*pb.ResponseConfigGlobal, error) {
	// create variable to contain struct responseConfigGlobal. for first initiate will set status.Created is false
	respConfigG := &pb.ResponseConfigGlobal{
		Configstatus: &pb.ConfigurationStatus{
			Created: false,
		},
	}

	// create context timeout to cancel process database when process to long
	ctx, cancel := context.WithTimeout(c, ucase.contextTimeout)

	defer cancel()

	// call AddConfigurationGlobal method of configRepo, to store data in table configuration_global
	res, err := ucase.configRepo.AddConfigurationGlobal(ctx, cg)
	if err != nil {
		return respConfigG, err
	}

	respConfigG.Configstatus.Created = res

	return respConfigG, nil

}

func (ucase *configurationUseCase) UpdateConfigurationGlobal(c context.Context, cg *pb.ConfigurationGlobal) (*pb.ResponseConfigGlobal, error) {
	// create variable to contain struct responseConfigGlobal. for first initiate will set status.Updated is false
	respConfigG := &pb.ResponseConfigGlobal{
		Configstatus: &pb.ConfigurationStatus{
			Updated: false,
		},
	}

	// create context timeout to cancel process database when process to long
	ctx, cancel := context.WithTimeout(c, ucase.contextTimeout)

	defer cancel()

	// call UpdateConfigurationGlobal method of configRepo, to update data exists by config_global_id in table configuration_global
	res, err := ucase.configRepo.UpdateConfigurationGlobal(ctx, cg)
	if err != nil {
		return respConfigG, err
	}

	respConfigG.Configstatus.Updated = res

	return respConfigG, nil
}

func (ucase *configurationUseCase) DeleteConfiguration(c context.Context, configGloalId int32) (*pb.ResponseConfigGlobal, error) {
	// create variable to contain struct responseConfigGlobal. for first initiate will set status.Deleted is false
	respConfigG := &pb.ResponseConfigGlobal{
		Configstatus: &pb.ConfigurationStatus{
			Deleted: false,
		},
	}

	// create context timeout to cancel process database when process to long
	ctx, cancel := context.WithTimeout(c, ucase.contextTimeout)

	defer cancel()

	// call DeleteConfiguration method of configRepo, to update data exists by config_global_id in table configuration_global
	res, err := ucase.configRepo.DeleteConfiguration(ctx, configGloalId)
	if err != nil {
		return respConfigG, err
	}

	respConfigG.Configstatus.Deleted = res

	return respConfigG, nil
}

func (ucase *configurationUseCase) GetConfigurationGlobal(c context.Context) (*pb.ResponseConfigGlobal, error) {
	// create variable to contain struct responseConfigGlobal.
	respConfigG := &pb.ResponseConfigGlobal{}

	// create context timeout to cancel process database when process to long
	ctx, cancel := context.WithTimeout(c, ucase.contextTimeout)

	defer cancel()

	// call GetConfigurationGlobal method of configRepo, to get All data in table configuration_global
	res, err := ucase.configRepo.GetConfigurationGlobal(ctx)
	if err != nil {
		return respConfigG, err
	}

	respConfigG.Configglobals = res
	return respConfigG, nil
}

func (ucase *configurationUseCase) GetConfigurationGlobalByID(c context.Context, configGlobalID int32) (*pb.ResponseConfigGlobal, error) {
	// create variable to contain struct responseConfigGlobal.
	respConfigG := &pb.ResponseConfigGlobal{}

	// create context timeout to cancel process database when process to long
	ctx, cancel := context.WithTimeout(c, ucase.contextTimeout)

	defer cancel()

	// call GetConfigurationGlobal method of configRepo, to get data by id in table configuration_global
	res, err := ucase.configRepo.GetConfigurationGlobalByID(ctx, configGlobalID)
	if err != nil {
		return respConfigG, err
	}

	respConfigG.Configglobal = res

	return respConfigG, nil
}

func (ucase *configurationUseCase) GetConfigurationGlobalActive(c context.Context) (*pb.ResponseConfigGlobal, error) {
	// create context timeout to cancel process database when process to long
	ctx, cancel := context.WithTimeout(c, ucase.contextTimeout)

	defer cancel()

	// call GetConfigurationGlobal method of configRepo, to get data by id in table configuration_global
	res, err := ucase.configRepo.GetConfigurationGlobalActive(ctx)
	if err != nil {
		return nil, err
	}

	// checking is res nil or not. if nil will will set default first data in configuration_global
	if res == nil {
		listConfgiGlobal, err := ucase.configRepo.GetConfigurationGlobal(ctx)
		if err != nil {
			return nil, errors.New("Cannot set default configuration global")
		}

		if len(listConfgiGlobal) > 0 {
			res = listConfgiGlobal[0]
		} else {
			return nil, errors.New("Cannot set default configuration global")
		}

	}

	// create variable to contain struct responseConfigGlobal.
	respConfigG := &pb.ResponseConfigGlobal{}

	respConfigG.Configglobal = res

	return respConfigG, nil
}

func (ucase *configurationUseCase) SetConfigurationGlobalActive(c context.Context, cg *pb.ConfigurationGlobal) (*pb.ResponseConfigGlobal, error) {
	// set isActive field of cg param to be true
	cg.IsActive = true

	// create context timeout to cancel process database when process to long
	ctx, cancel := context.WithTimeout(c, ucase.contextTimeout)

	defer cancel()

	// Get All Configuration Global
	listConfigGlobal, _ := ucase.configRepo.GetConfigurationGlobal(ctx)

	// checking configuration active, then deactive
	for _, configGlobal := range listConfigGlobal {
		if configGlobal.IsActive {
			configGlobal.IsActive = false

			// because we wont waitting, and let the syntax run as asyncronous without blocking other process
			go ucase.configRepo.UpdateConfigurationGlobal(ctx, configGlobal)
		}
	}

	// call UpdateConfigurationGlobal method of configRepo, to update data by id in table configuration_global
	updated, err := ucase.configRepo.UpdateConfigurationGlobal(ctx, cg)
	if err != nil {
		return nil, err
	}

	// create variable to contain struct responseConfigGlobal.
	respConfigG := &pb.ResponseConfigGlobal{}
	respConfigG.Configglobal = cg
	respConfigG.Configstatus = &pb.ConfigurationStatus{Updated: updated}

	return respConfigG, nil
}
