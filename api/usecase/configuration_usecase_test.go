package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/muhammadhidayah/configuration-service/api/mocks"
	ucase "github.com/muhammadhidayah/configuration-service/api/usecase"
	pb "github.com/muhammadhidayah/configuration-service/proto/configuration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetConfigurationClient(t *testing.T) {
	mockConfigRepo := new(mocks.Repository)
	mockConfigClient := &pb.ConfigurationClient{
		ConfigClientId:     1,
		ConfigClientUuid:   "a6e2745e-c930-4717-a9d1-d1cfb2a64aa4",
		MultipleLanguageId: 2,
		Appname:            "client1.inactsoft.com",
		ReportTitle:        "Client Satu",
		CompanySubsId:      "012-031-234-542",
		IsConfigDeleted:    0,
	}

	mockListConfigClient := make([]*pb.ConfigurationClient, 0)
	mockListConfigClient = append(mockListConfigClient, mockConfigClient)

	t.Run("success", func(t *testing.T) {
		mockConfigRepo.On("GetConfigurationClient", mock.Anything).Return(mockListConfigClient, nil).Once()

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)
		list, err := uc.GetConfigurationClient(context.TODO())
		assert.NoError(t, err)
		assert.Len(t, list.Configclients, 1)

		mockConfigRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockConfigRepo.On("GetConfigurationClient", mock.Anything).Return(nil, errors.New("Unexpected Error")).Once()

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)
		list, err := uc.GetConfigurationClient(context.TODO())
		assert.Error(t, err)
		assert.Nil(t, list)

		mockConfigRepo.AssertExpectations(t)
	})
}

func TestGetConfigurationClientBySubs(t *testing.T) {
	mockConfigRepo := new(mocks.Repository)
	mockConfigClient := pb.ConfigurationClient{
		ConfigClientId:     1,
		ConfigClientUuid:   "a6e2745e-c930-4717-a9d1-d1cfb2a64aa4",
		MultipleLanguageId: 2,
		Appname:            "client1.inactsoft.com",
		ReportTitle:        "Client Satu",
		CompanySubsId:      "012-031-234-542",
		IsConfigDeleted:    0,
	}

	mockListConfigClient := make([]pb.ConfigurationClient, 0)
	mockListConfigClient = append(mockListConfigClient, mockConfigClient)

	t.Run("Success", func(t *testing.T) {
		mockConfigRepo.On("GetConfigurationClientBySubs", mock.Anything, mock.AnythingOfType("string")).Return(&mockConfigClient, nil).Once()

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)
		configClient, err := uc.GetConfigurationClientBySubs(context.TODO(), mockConfigClient.CompanySubsId)
		assert.NoError(t, err)
		assert.NotNil(t, configClient)

		mockConfigRepo.AssertExpectations(t)
	})

	t.Run("Success when get data", func(t *testing.T) {
		mockConfigRepo.On("GetConfigurationClientBySubs", mock.Anything, mock.AnythingOfType("string")).Return(&mockConfigClient, nil).Once()

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)
		configClient, err := uc.GetConfigurationClientBySubs(context.TODO(), mockConfigClient.CompanySubsId)
		assert.NoError(t, err)
		assert.Equal(t, mockConfigClient.ConfigClientUuid, configClient.Configclient.ConfigClientUuid)

		mockConfigRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		mockConfigRepo.On("GetConfigurationClientBySubs", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("UnExpected Error")).Once()

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)
		configClient, err := uc.GetConfigurationClientBySubs(context.TODO(), mockConfigClient.CompanySubsId)

		assert.Error(t, err)
		assert.Nil(t, configClient)

		mockConfigRepo.AssertExpectations(t)
	})
}

func TestAddConfigurationClient(t *testing.T) {
	mockConfigRepo := new(mocks.Repository)
	mockConfigClient := &pb.ConfigurationClient{
		ConfigClientId:     1,
		MultipleLanguageId: 2,
		Appname:            "client1.inactsoft.com",
		ReportTitle:        "Client Satu",
		CompanySubsId:      "012-031-234-542",
		IsConfigDeleted:    0,
	}

	t.Run("Success Add ConfigurationClient", func(t *testing.T) {
		mockConfigRepo.On("AddConfigurationClient", mock.Anything, mock.AnythingOfType("*configuration.ConfigurationClient")).Return(true, nil).Once()

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)
		inserted, err := uc.AddConfigurationClient(context.TODO(), mockConfigClient)

		assert.NoError(t, err)
		assert.True(t, inserted.Status.Created)
	})

	t.Run("Error Add Configuration", func(t *testing.T) {
		mockConfigRepo.On("AddConfigurationClient", mock.Anything, mock.AnythingOfType("*configuration.ConfigurationClient")).Return(false, errors.New("Unexpected syntax error")).Once()

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)
		inserted, err := uc.AddConfigurationClient(context.TODO(), mockConfigClient)

		assert.Error(t, err)
		assert.False(t, inserted.Status.Created)
	})
}

func TestUpdateConfigurationClientBySubs(t *testing.T) {
	mockConfigRepo := new(mocks.Repository)
	mockConfigClient := &pb.ConfigurationClient{
		ConfigClientId:     1,
		MultipleLanguageId: 2,
		Appname:            "client1.inactsoft.com",
		ReportTitle:        "Client Satu",
		CompanySubsId:      "012-031-234-542",
		IsConfigDeleted:    0,
	}

	t.Run("Success Update Configuration Client", func(t *testing.T) {
		mockConfigRepo.On("UpdateConfigurationClientBySubs", mock.Anything, mock.AnythingOfType("*configuration.ConfigurationClient")).Return(true, nil).Once()

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)
		reslt, err := uc.UpdateConfigurationClientBySubs(context.TODO(), mockConfigClient)

		assert.NoError(t, err)
		assert.True(t, reslt.Status.Updated)
	})

	t.Run("Failed Update Configuration Client", func(t *testing.T) {
		mockConfigRepo.On("UpdateConfigurationClientBySubs", mock.Anything, mock.AnythingOfType("*configuration.ConfigurationClient")).Return(false, errors.New("Unexpected Error")).Once()

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)
		reslt, err := uc.UpdateConfigurationClientBySubs(context.TODO(), mockConfigClient)

		assert.Error(t, err)
		assert.False(t, reslt.Status.Updated)
	})
}

func TestDeleteConfigurationClientBySubs(t *testing.T) {
	mockConfigRepo := new(mocks.Repository)
	mockConfigClient := &pb.ConfigurationClient{
		ConfigClientId:     1,
		MultipleLanguageId: 2,
		Appname:            "client1.inactsoft.com",
		ReportTitle:        "Client Satu",
		CompanySubsId:      "",
		IsConfigDeleted:    0,
	}

	t.Run("Success to delete configuration", func(t *testing.T) {
		mockConfigRepo.On("DeleteConfigurationClientBySubs", mock.Anything, mock.AnythingOfType("*configuration.ConfigurationClient")).Return(true, nil).Once()

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)
		res, err := uc.DeleteConfigurationClientBySubs(context.TODO(), mockConfigClient)

		assert.NoError(t, err)
		assert.True(t, res.Status.Deleted)
	})

	t.Run("Failed to delete configuration company_subs_id not found", func(t *testing.T) {
		mockConfigRepo.On("DeleteConfigurationClientBySubs", mock.Anything, mock.AnythingOfType("*configuration.ConfigurationClient")).Return(false, errors.New("Company subs id not found")).Once()

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)
		res, err := uc.DeleteConfigurationClientBySubs(context.TODO(), mockConfigClient)

		assert.Error(t, err)
		assert.False(t, res.Status.Deleted)
	})
}

func TestAddConfigurationGlobal(t *testing.T) {
	mockConfigRepo := new(mocks.Repository)
	mockConfigGlobal := &pb.ConfigurationGlobal{
		ConfigGlobalId: 1,
		Footertext:     "Technical support : +62-21-7509077 ext. 109 | Email: support@inactsoft.com | <a href='http://wiki.inactsoft.com' target='_blank' class='footer-link'>online help</a>",
		ServerSmpt:     "mail.google.com",
		Ssl:            false,
		Port:           5432,
		IsAuth:         true,
		Username:       "notification1@inactsoft.com",
		Password:       "123456789087654",
		IsActive:       true,
	}

	t.Run("Success Add Configuration Global", func(t *testing.T) {
		mockConfigRepo.On("AddConfigurationGlobal", mock.Anything, mock.AnythingOfType("*configuration.ConfigurationGlobal")).Return(true, nil).Once()

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)
		res, err := uc.AddConfigurationGlobal(context.TODO(), mockConfigGlobal)

		assert.NoError(t, err)
		assert.True(t, res.Configstatus.Created)
	})

	t.Run("Failed Add Configuration Global", func(t *testing.T) {
		mockConfigRepo.On("AddConfigurationGlobal", mock.Anything, mock.AnythingOfType("*configuration.ConfigurationGlobal")).Return(false, errors.New("Unexpected syntax error")).Once()

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)
		res, err := uc.AddConfigurationGlobal(context.TODO(), mockConfigGlobal)

		assert.Error(t, err)
		assert.False(t, res.Configstatus.Created)
	})
}

func TestUpdateConfigurationGlobal(t *testing.T) {
	mockConfigRepo := new(mocks.Repository)
	mockConfigGlobal := &pb.ConfigurationGlobal{
		ConfigGlobalId: 1,
		Footertext:     "Technical support : +62-21-7509077 ext. 109 | Email: support@inactsoft.com | <a href='http://wiki.inactsoft.com' target='_blank' class='footer-link'>online help</a>",
		ServerSmpt:     "mail.google.com",
		Ssl:            false,
		Port:           5432,
		IsAuth:         true,
		Username:       "notification1@inactsoft.com",
		Password:       "123456789087654",
		IsActive:       true,
	}

	t.Run("Success Update Configuration", func(t *testing.T) {
		mockConfigRepo.On("UpdateConfigurationGlobal", mock.Anything, mock.AnythingOfType("*configuration.ConfigurationGlobal")).Return(true, nil).Once()

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)
		res, err := uc.UpdateConfigurationGlobal(context.TODO(), mockConfigGlobal)

		assert.NoError(t, err)
		assert.True(t, res.Configstatus.Updated)
	})

	t.Run("Failed to update configuration global", func(t *testing.T) {
		mockConfigRepo.On("UpdateConfigurationGlobal", mock.Anything, mock.AnythingOfType("*configuration.ConfigurationGlobal")).Return(false, errors.New("Unexpected syntax error")).Once()

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)
		res, err := uc.UpdateConfigurationGlobal(context.TODO(), mockConfigGlobal)

		assert.Error(t, err)
		assert.False(t, res.Configstatus.Updated)
	})
}

func TestDeleteConfiguration(t *testing.T) {
	mockConfigRepo := new(mocks.Repository)
	mockConfigGlobal := &pb.ConfigurationGlobal{
		ConfigGlobalId: 1,
		Footertext:     "Technical support : +62-21-7509077 ext. 109 | Email: support@inactsoft.com | <a href='http://wiki.inactsoft.com' target='_blank' class='footer-link'>online help</a>",
		ServerSmpt:     "mail.google.com",
		Ssl:            false,
		Port:           5432,
		IsAuth:         true,
		Username:       "notification1@inactsoft.com",
		Password:       "123456789087654",
		IsActive:       true,
	}

	t.Run("Success Delete Configuration", func(t *testing.T) {
		mockConfigRepo.On("DeleteConfiguration", mock.Anything, mock.AnythingOfType("int32")).Return(true, nil).Once()

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)

		res, err := uc.DeleteConfiguration(context.TODO(), mockConfigGlobal.ConfigGlobalId)

		assert.NoError(t, err)
		assert.True(t, res.Configstatus.Deleted)
	})

	t.Run("Failed Deleted Configuration", func(t *testing.T) {
		mockConfigRepo.On("DeleteConfiguration", mock.Anything, mock.AnythingOfType("int32")).Return(false, errors.New("Unexpected syntax error")).Once()

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)

		res, err := uc.DeleteConfiguration(context.TODO(), mockConfigGlobal.ConfigGlobalId)

		assert.Error(t, err)
		assert.False(t, res.Configstatus.Deleted)
	})
}

func TestGetConfigurationGlobal(t *testing.T) {
	mockConfigRepo := new(mocks.Repository)
	mockConfigGlobal := &pb.ConfigurationGlobal{
		ConfigGlobalId: 1,
		Footertext:     "Technical support : +62-21-7509077 ext. 109 | Email: support@inactsoft.com | <a href='http://wiki.inactsoft.com' target='_blank' class='footer-link'>online help</a>",
		ServerSmpt:     "mail.google.com",
		Ssl:            false,
		Port:           5432,
		IsAuth:         true,
		Username:       "notification1@inactsoft.com",
		Password:       "123456789087654",
		IsActive:       true,
	}

	mockListConfigGlobal := make([]*pb.ConfigurationGlobal, 0)
	mockListConfigGlobal = append(mockListConfigGlobal, mockConfigGlobal)

	t.Run("Get All Configuration Global.", func(t *testing.T) {
		mockConfigRepo.On("GetConfigurationGlobal", mock.Anything).Return(mockListConfigGlobal, nil).Once()

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)
		res, err := uc.GetConfigurationGlobal(context.TODO())

		assert.NoError(t, err)
		assert.Len(t, res.Configglobals, 1)
	})

	t.Run("Failed Get All Configuration Global.", func(t *testing.T) {
		mockConfigRepo.On("GetConfigurationGlobal", mock.Anything).Return(nil, errors.New("Unexpected syntax error")).Once()

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)
		res, err := uc.GetConfigurationGlobal(context.TODO())

		assert.Error(t, err)
		assert.Len(t, res.Configglobals, 0)
	})
}

func TestGetConfigurationGlobalByID(t *testing.T) {
	mockConfigRepo := new(mocks.Repository)
	mockConfigGlobal := &pb.ConfigurationGlobal{
		ConfigGlobalId: 1,
		Footertext:     "Technical support : +62-21-7509077 ext. 109 | Email: support@inactsoft.com | <a href='http://wiki.inactsoft.com' target='_blank' class='footer-link'>online help</a>",
		ServerSmpt:     "mail.google.com",
		Ssl:            false,
		Port:           5432,
		IsAuth:         true,
		Username:       "notification1@inactsoft.com",
		Password:       "123456789087654",
		IsActive:       true,
	}

	t.Run("Get Configuration Global with condition id equal params", func(t *testing.T) {
		mockConfigRepo.On("GetConfigurationGlobalByID", mock.Anything, mock.AnythingOfType("int32")).Return(mockConfigGlobal, nil).Once()

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)
		res, err := uc.GetConfigurationGlobalByID(context.TODO(), mockConfigGlobal.ConfigGlobalId)

		assert.NoError(t, err)
		assert.Equal(t, mockConfigGlobal.ConfigGlobalId, res.Configglobal.ConfigGlobalId)
	})

	t.Run("Failed Get Configuration Global with condition id equal params", func(t *testing.T) {
		mockConfigRepo.On("GetConfigurationGlobalByID", mock.Anything, mock.AnythingOfType("int32")).Return(nil, errors.New("Unexpected syntax error")).Once()

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)
		res, err := uc.GetConfigurationGlobalByID(context.TODO(), mockConfigGlobal.ConfigGlobalId)

		assert.Error(t, err)
		assert.NotNil(t, res)
	})
}

func TestGetConfigurationGlobalActive(t *testing.T) {
	mockConfigRepo := new(mocks.Repository)
	mockConfigGlobal := &pb.ConfigurationGlobal{
		ConfigGlobalId: 1,
		Footertext:     "Technical support : +62-21-7509077 ext. 109 | Email: support@inactsoft.com | <a href='http://wiki.inactsoft.com' target='_blank' class='footer-link'>online help</a>",
		ServerSmpt:     "mail.google.com",
		Ssl:            false,
		Port:           5432,
		IsAuth:         true,
		Username:       "notification1@inactsoft.com",
		Password:       "123456789087654",
		IsActive:       true,
	}

	mockListConfigGlobal := make([]*pb.ConfigurationGlobal, 0)
	mockListConfigGlobalZLen := make([]*pb.ConfigurationGlobal, 0)
	mockListConfigGlobal = append(mockListConfigGlobal, mockConfigGlobal)

	t.Run("Get Configuration Global Active", func(t *testing.T) {
		mockConfigRepo.On("GetConfigurationGlobalActive", mock.Anything).Return(mockConfigGlobal, nil).Once()

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)
		res, err := uc.GetConfigurationGlobalActive(context.TODO())

		assert.NoError(t, err)
		assert.Equal(t, mockConfigGlobal.ConfigGlobalId, res.Configglobal.ConfigGlobalId)
	})

	t.Run("Get Configuration Global Active but Nil, will set first data in configuration_global", func(t *testing.T) {
		mockConfigRepo.On("GetConfigurationGlobalActive", mock.Anything).Return(nil, nil).Once()
		mockConfigRepo.On("GetConfigurationGlobal", mock.Anything).Return(mockListConfigGlobal, nil).Once()

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)
		res, err := uc.GetConfigurationGlobalActive(context.TODO())

		assert.NoError(t, err)
		assert.Equal(t, mockConfigGlobal.ConfigGlobalId, res.Configglobal.ConfigGlobalId)
	})

	t.Run("Failed Set Default Get Configuration Global Active but Nil, will set first data in configuration_global", func(t *testing.T) {
		mockConfigRepo.On("GetConfigurationGlobalActive", mock.Anything).Return(nil, nil).Once()
		mockConfigRepo.On("GetConfigurationGlobal", mock.Anything).Return(nil, errors.New("Unexpected syntax error")).Once()

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)
		res, err := uc.GetConfigurationGlobalActive(context.TODO())

		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("Failed Set Default, becasue data not found but no error", func(t *testing.T) {
		mockConfigRepo.On("GetConfigurationGlobalActive", mock.Anything).Return(nil, nil).Once()
		mockConfigRepo.On("GetConfigurationGlobal", mock.Anything).Return(mockListConfigGlobalZLen, nil).Once()

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)
		res, err := uc.GetConfigurationGlobalActive(context.TODO())

		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("Failed Get Configuration Global Active", func(t *testing.T) {
		mockConfigRepo.On("GetConfigurationGlobalActive", mock.Anything).Return(nil, errors.New("Unexpected syntax error")).Once()

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)
		res, err := uc.GetConfigurationGlobalActive(context.TODO())

		assert.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestSetConfigurationGlobalActive(t *testing.T) {
	mockConfigRepo := new(mocks.Repository)
	mockListConfigGlobal := []*pb.ConfigurationGlobal{
		{
			ConfigGlobalId: 1,
			Footertext:     "Technical support : +62-21-7509077 ext. 109 | Email: support@inactsoft.com | <a href='http://wiki.inactsoft.com' target='_blank' class='footer-link'>online help</a>",
			ServerSmpt:     "mail.google.com",
			Ssl:            false,
			Port:           5432,
			IsAuth:         true,
			Username:       "notification1@inactsoft.com",
			Password:       "123456789087654",
			IsActive:       true,
		},
		{
			ConfigGlobalId: 2,
			Footertext:     "Technical support : +62-21-7509077 ext. 109 | Email: support@inactsoft.com | <a href='http://wiki.inactsoft.com' target='_blank' class='footer-link'>online help</a>",
			ServerSmpt:     "mail.google.com",
			Ssl:            false,
			Port:           5432,
			IsAuth:         true,
			Username:       "notification1@inactsoft.com",
			Password:       "123456789087654",
			IsActive:       true,
		},
		{
			ConfigGlobalId: 3,
			Footertext:     "Technical support : +62-21-7509077 ext. 109 | Email: support@inactsoft.com | <a href='http://wiki.inactsoft.com' target='_blank' class='footer-link'>online help</a>",
			ServerSmpt:     "mail.google.com",
			Ssl:            false,
			Port:           5432,
			IsAuth:         true,
			Username:       "notification1@inactsoft.com",
			Password:       "123456789087654",
			IsActive:       false,
		},
	}

	t.Run("Set Configuration Global Active", func(t *testing.T) {
		mockConfigRepo.On("GetConfigurationGlobal", mock.Anything).Return(mockListConfigGlobal, nil).Once()
		mockConfigRepo.On("UpdateConfigurationGlobal", mock.Anything, mock.AnythingOfType("*configuration.ConfigurationGlobal")).Return(true, nil)

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)
		res, err := uc.SetConfigurationGlobalActive(context.TODO(), mockListConfigGlobal[2])

		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("Set Configuration Global Active and Inactive Config Gloabl", func(t *testing.T) {
		mockConfigRepo.On("GetConfigurationGlobal", mock.Anything).Return(mockListConfigGlobal, nil).Once()
		mockConfigRepo.On("UpdateConfigurationGlobal", mock.Anything, mock.AnythingOfType("*configuration.ConfigurationGlobal")).Return(true, nil)

		uc := ucase.NewConfigurationUsecase(mockConfigRepo, time.Second*2)
		res, err := uc.SetConfigurationGlobalActive(context.TODO(), mockListConfigGlobal[2])

		assert.NoError(t, err)
		assert.NotNil(t, res)
	})
}
