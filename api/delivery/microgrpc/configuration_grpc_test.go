package microgrpc_test

import (
	"context"
	"errors"
	"testing"

	micro "github.com/muhammadhidayah/configuration-service/api/delivery/microgrpc"
	"github.com/muhammadhidayah/configuration-service/api/mocks"
	pb "github.com/muhammadhidayah/configuration-service/proto/configuration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetConfigurationClient(t *testing.T) {
	mockUseCaseConf := new(mocks.Usecase)
	mockRespConfigClient := &pb.ResponseConfigClient{
		Configclients: []*pb.ConfigurationClient{
			{
				ConfigClientId:     1,
				ConfigClientUuid:   "a6e2745e-c930-4717-a9d1-d1cfb2a64aa4",
				MultipleLanguageId: 2,
				Appname:            "client1.inactsoft.com",
				ReportTitle:        "Client Satu",
				CompanySubsId:      "012-031-234-542",
				IsConfigDeleted:    0,
			},
			{
				ConfigClientId:     2,
				ConfigClientUuid:   "a6e2745e-c930-4717-a9d1-d1cfb2a64ab2",
				MultipleLanguageId: 2,
				Appname:            "client1.inactsoft.com",
				ReportTitle:        "Client Satu",
				CompanySubsId:      "012-031-234-542",
				IsConfigDeleted:    0,
			},
		},
	}

	mockRespConfigClientRes := &pb.ResponseConfigClient{}

	mockReqConfigClient := &pb.RequestConfigCient{}

	t.Run("Get Configuration Client", func(t *testing.T) {
		mockUseCaseConf.On("GetConfigurationClient", mock.Anything).Return(mockRespConfigClient, nil).Once()

		handler := micro.NewMicroGrpc(mockUseCaseConf)
		err := handler.GetConfigurationClient(context.TODO(), mockReqConfigClient, mockRespConfigClientRes)

		assert.NoError(t, err)
		assert.Len(t, mockRespConfigClientRes.Configclients, 2)
	})

	t.Run("Failed Get Config Client", func(t *testing.T) {
		mockUseCaseConf.On("GetConfigurationClient", mock.Anything).Return(nil, errors.New("Unexpected syntax error")).Once()

		handler := micro.NewMicroGrpc(mockUseCaseConf)
		err := handler.GetConfigurationClient(context.TODO(), mockReqConfigClient, mockRespConfigClientRes)

		assert.Error(t, err)
	})
}

func TestGetConfigurationClientBySubs(t *testing.T) {
	mockUseCaseConf := new(mocks.Usecase)
	mockRespConfigClient := &pb.ResponseConfigClient{
		Configclients: []*pb.ConfigurationClient{
			{
				ConfigClientId:     2,
				ConfigClientUuid:   "a6e2745e-c930-4717-a9d1-d1cfb2a64ab2",
				MultipleLanguageId: 2,
				Appname:            "client1.inactsoft.com",
				ReportTitle:        "Client Satu",
				CompanySubsId:      "012-031-234-542",
				IsConfigDeleted:    0,
			},
		},
	}

	mockRespConfigClientRes := &pb.ResponseConfigClient{}

	mockReqConfigClient := &pb.RequestConfigCient{
		Configclient: &pb.ConfigurationClient{
			CompanySubsId: "012-031-234-542",
		},
	}

	t.Run("Get Configuration Clients By Company Subs ID", func(t *testing.T) {
		mockUseCaseConf.On("GetConfigurationClientBySubs", mock.Anything, mock.AnythingOfType("string")).Return(mockRespConfigClient, nil).Once()

		handler := micro.NewMicroGrpc(mockUseCaseConf)
		err := handler.GetConfigurationClientBySubs(context.TODO(), mockReqConfigClient, mockRespConfigClientRes)

		assert.NoError(t, err)
	})

	t.Run("Failed Get Configuration By Company Subs ID", func(t *testing.T) {
		mockUseCaseConf.On("GetConfigurationClientBySubs", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("Data Not Found")).Once()

		handler := micro.NewMicroGrpc(mockUseCaseConf)
		err := handler.GetConfigurationClientBySubs(context.TODO(), mockReqConfigClient, mockRespConfigClientRes)

		assert.Error(t, err)
	})
}

func TestAddConfigurationClient(t *testing.T) {
	mockUseCaseConf := new(mocks.Usecase)
	mockRespConfigClient := &pb.ResponseConfigClient{
		Status: &pb.ConfigurationStatus{Created: true},
	}

	mockRespConfigClientRes := &pb.ResponseConfigClient{}

	mockReqConfigClient := &pb.RequestConfigCient{
		Configclient: &pb.ConfigurationClient{
			ConfigClientId:     2,
			ConfigClientUuid:   "a6e2745e-c930-4717-a9d1-d1cfb2a64ab2",
			MultipleLanguageId: 2,
			Appname:            "client1.inactsoft.com",
			ReportTitle:        "Client Satu",
			CompanySubsId:      "012-031-234-542",
			IsConfigDeleted:    0,
		},
	}

	t.Run("Add configuration client", func(t *testing.T) {
		mockUseCaseConf.On("AddConfigurationClient", mock.Anything, mock.AnythingOfType("*configuration.ConfigurationClient")).Return(mockRespConfigClient, nil).Once()

		handler := micro.NewMicroGrpc(mockUseCaseConf)
		err := handler.AddConfigurationClient(context.TODO(), mockReqConfigClient, mockRespConfigClientRes)

		assert.True(t, mockRespConfigClientRes.Status.Created)
		assert.NoError(t, err)
	})

	t.Run("Failed Add configuration client", func(t *testing.T) {
		mockRespConfigClient.Status.Created = false
		mockUseCaseConf.On("AddConfigurationClient", mock.Anything, mock.AnythingOfType("*configuration.ConfigurationClient")).Return(mockRespConfigClient, errors.New("Unexpected syntax error")).Once()

		handler := micro.NewMicroGrpc(mockUseCaseConf)
		err := handler.AddConfigurationClient(context.TODO(), mockReqConfigClient, mockRespConfigClientRes)

		assert.False(t, mockRespConfigClientRes.Status.Created)
		assert.Error(t, err)
	})
}

func TestUpdateConfigurationClientBySubs(t *testing.T) {
	mockUseCaseConf := new(mocks.Usecase)
	mockRespConfigClient := &pb.ResponseConfigClient{
		Status: &pb.ConfigurationStatus{Updated: true},
	}

	mockRespConfigClientRes := &pb.ResponseConfigClient{}

	mockReqConfigClient := &pb.RequestConfigCient{
		Configclient: &pb.ConfigurationClient{
			ConfigClientId:     2,
			ConfigClientUuid:   "a6e2745e-c930-4717-a9d1-d1cfb2a64ab2",
			MultipleLanguageId: 2,
			Appname:            "client1.inactsoft.com",
			ReportTitle:        "Client Satu",
			CompanySubsId:      "012-031-234-542",
			IsConfigDeleted:    0,
		},
	}

	t.Run("Update configuration client", func(t *testing.T) {
		mockUseCaseConf.On("UpdateConfigurationClientBySubs", mock.Anything, mock.AnythingOfType("*configuration.ConfigurationClient")).Return(mockRespConfigClient, nil).Once()

		handler := micro.NewMicroGrpc(mockUseCaseConf)
		err := handler.UpdateConfigurationClientBySubs(context.TODO(), mockReqConfigClient, mockRespConfigClientRes)

		assert.NoError(t, err)
		assert.True(t, mockRespConfigClientRes.Status.Updated)
	})

	t.Run("Failed Update configuration client", func(t *testing.T) {
		mockRespConfigClient.Status.Updated = false
		mockUseCaseConf.On("UpdateConfigurationClientBySubs", mock.Anything, mock.AnythingOfType("*configuration.ConfigurationClient")).Return(mockRespConfigClient, errors.New("Unexpected Error")).Once()

		handler := micro.NewMicroGrpc(mockUseCaseConf)
		err := handler.UpdateConfigurationClientBySubs(context.TODO(), mockReqConfigClient, mockRespConfigClientRes)

		assert.Error(t, err)
		assert.False(t, mockRespConfigClientRes.Status.Updated)
	})
}

func TestDeleteConfigurationClientBySubs(t *testing.T) {
	mockUseCaseConf := new(mocks.Usecase)
	mockRespConfigClient := &pb.ResponseConfigClient{
		Status: &pb.ConfigurationStatus{Deleted: true},
	}

	mockRespConfigClientRes := &pb.ResponseConfigClient{}

	mockReqConfigClient := &pb.RequestConfigCient{
		Configclient: &pb.ConfigurationClient{
			ConfigClientId:     2,
			ConfigClientUuid:   "a6e2745e-c930-4717-a9d1-d1cfb2a64ab2",
			MultipleLanguageId: 2,
			Appname:            "client1.inactsoft.com",
			ReportTitle:        "Client Satu",
			CompanySubsId:      "012-031-234-542",
			IsConfigDeleted:    0,
		},
	}

	t.Run("Delete configuration client", func(t *testing.T) {
		mockUseCaseConf.On("DeleteConfigurationClientBySubs", mock.Anything, mock.AnythingOfType("*configuration.ConfigurationClient")).Return(mockRespConfigClient, nil).Once()

		handler := micro.NewMicroGrpc(mockUseCaseConf)
		err := handler.DeleteConfigurationClientBySubs(context.TODO(), mockReqConfigClient, mockRespConfigClientRes)

		assert.NoError(t, err)
		assert.True(t, mockRespConfigClientRes.Status.GetDeleted())
	})

	t.Run("Delete configuration client", func(t *testing.T) {
		mockRespConfigClient.Status.Deleted = false
		mockUseCaseConf.On("DeleteConfigurationClientBySubs", mock.Anything, mock.AnythingOfType("*configuration.ConfigurationClient")).Return(mockRespConfigClient, errors.New("Unexpected Error")).Once()

		handler := micro.NewMicroGrpc(mockUseCaseConf)
		err := handler.DeleteConfigurationClientBySubs(context.TODO(), mockReqConfigClient, mockRespConfigClientRes)

		assert.Error(t, err)
		assert.False(t, mockRespConfigClientRes.Status.GetDeleted())
	})
}

func TestAddConfigurationGlobal(t *testing.T) {
	mockUseCaseConf := new(mocks.Usecase)
	mockRespConfGlobal := &pb.ResponseConfigGlobal{
		Configstatus: &pb.ConfigurationStatus{Created: true},
	}

	mockReqConfGlobal := &pb.RequestConfigGlobal{
		Configglobal: &pb.ConfigurationGlobal{
			ConfigGlobalId: 1,
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

	mockRespConfGlobalRes := &pb.ResponseConfigGlobal{}

	t.Run("Add Configuration Global", func(t *testing.T) {
		mockUseCaseConf.On("AddConfigurationGlobal", mock.Anything, mock.AnythingOfType("*configuration.ConfigurationGlobal")).Return(mockRespConfGlobal, nil).Once()

		handler := micro.NewMicroGrpc(mockUseCaseConf)
		err := handler.AddConfigurationGlobal(context.TODO(), mockReqConfGlobal, mockRespConfGlobalRes)

		assert.NoError(t, err)
		assert.True(t, mockRespConfGlobalRes.Configstatus.GetCreated())
		assert.Equal(t, mockReqConfGlobal.GetConfigglobal().GetConfigGlobalId(), mockRespConfGlobalRes.GetConfigglobal().GetConfigGlobalId())
	})

	t.Run("Add Configuration Global", func(t *testing.T) {
		mockRespConfGlobal.Configstatus.Created = false
		mockUseCaseConf.On("AddConfigurationGlobal", mock.Anything, mock.AnythingOfType("*configuration.ConfigurationGlobal")).Return(mockRespConfGlobal, errors.New("Unexpected Error")).Once()

		handler := micro.NewMicroGrpc(mockUseCaseConf)
		err := handler.AddConfigurationGlobal(context.TODO(), mockReqConfGlobal, mockRespConfGlobalRes)

		assert.Error(t, err)
		assert.False(t, mockRespConfGlobalRes.Configstatus.GetCreated())
	})
}

func TestUpdateConfigurationGlobal(t *testing.T) {
	mockUseCaseConf := new(mocks.Usecase)
	mockRespConfGlobal := &pb.ResponseConfigGlobal{
		Configstatus: &pb.ConfigurationStatus{Updated: true},
	}

	mockReqConfGlobal := &pb.RequestConfigGlobal{
		Configglobal: &pb.ConfigurationGlobal{
			ConfigGlobalId: 1,
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

	mockRespConfGlobalRes := &pb.ResponseConfigGlobal{}

	t.Run("Update Configuration Global", func(t *testing.T) {
		mockUseCaseConf.On("UpdateConfigurationGlobal", mock.Anything, mock.AnythingOfType("*configuration.ConfigurationGlobal")).Return(mockRespConfGlobal, nil).Once()

		handler := micro.NewMicroGrpc(mockUseCaseConf)
		err := handler.UpdateConfigurationGlobal(context.TODO(), mockReqConfGlobal, mockRespConfGlobalRes)

		assert.NoError(t, err)
		assert.True(t, mockRespConfGlobalRes.Configstatus.GetUpdated())
	})

	t.Run("Failed Update Configuration Global", func(t *testing.T) {
		mockRespConfGlobal.Configstatus.Updated = false
		mockUseCaseConf.On("UpdateConfigurationGlobal", mock.Anything, mock.AnythingOfType("*configuration.ConfigurationGlobal")).Return(mockRespConfGlobal, errors.New("Unexpected syntax error")).Once()

		handler := micro.NewMicroGrpc(mockUseCaseConf)
		err := handler.UpdateConfigurationGlobal(context.TODO(), mockReqConfGlobal, mockRespConfGlobalRes)

		assert.Error(t, err)
		assert.False(t, mockRespConfGlobalRes.Configstatus.GetUpdated())
	})
}

func TestDeleteConfiguration(t *testing.T) {
	mockUseCaseConf := new(mocks.Usecase)
	mockRespConfGlobal := &pb.ResponseConfigGlobal{
		Configstatus: &pb.ConfigurationStatus{Deleted: true},
	}

	mockReqConfGlobal := &pb.RequestConfigGlobal{
		Configglobal: &pb.ConfigurationGlobal{
			ConfigGlobalId: 1,
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

	mockRespConfGlobalRes := &pb.ResponseConfigGlobal{}

	t.Run("Delete Configuration Global", func(t *testing.T) {
		mockUseCaseConf.On("DeleteConfiguration", mock.Anything, mock.AnythingOfType("int32")).Return(mockRespConfGlobal, nil).Once()

		handler := micro.NewMicroGrpc(mockUseCaseConf)
		err := handler.DeleteConfiguration(context.TODO(), mockReqConfGlobal, mockRespConfGlobalRes)

		assert.NoError(t, err)
		assert.True(t, mockRespConfGlobalRes.Configstatus.GetDeleted())
	})

	t.Run("Failed Delete Configuration Global", func(t *testing.T) {
		mockRespConfGlobal.Configstatus.Deleted = false
		mockUseCaseConf.On("DeleteConfiguration", mock.Anything, mock.AnythingOfType("int32")).Return(mockRespConfGlobal, errors.New("Unexpected syntax error")).Once()

		handler := micro.NewMicroGrpc(mockUseCaseConf)
		err := handler.DeleteConfiguration(context.TODO(), mockReqConfGlobal, mockRespConfGlobalRes)

		assert.Error(t, err)
		assert.False(t, mockRespConfGlobalRes.Configstatus.GetDeleted())
	})
}

func TestGetConfigurationGlobal(t *testing.T) {
	mockUseCaseConf := new(mocks.Usecase)
	mockRespConfGlobal := &pb.ResponseConfigGlobal{
		Configglobals: []*pb.ConfigurationGlobal{
			{
				ConfigGlobalId: 1,
				Footertext:     "Technical support : +62-21-7509077 ext. 109 | Email: support@inactsoft.com | <a href='http://wiki.inactsoft.com' target='_blank' class='footer-link'>online help</a>",
				ServerSmpt:     "mail.google.com",
				Ssl:            false,
				Port:           5432,
				IsAuth:         true,
				Username:       "notification1@inactsoft.com",
				Password:       "123456789087654",
				IsActive:       false,
			}, {
				ConfigGlobalId: 2,
				Footertext:     "Technical support : +62-21-7509077 ext. 109 | Email: support@inactsoft.com | <a href='http://wiki.inactsoft.com' target='_blank' class='footer-link'>online help</a>",
				ServerSmpt:     "mail.google.com",
				Ssl:            false,
				Port:           5432,
				IsAuth:         true,
				Username:       "notification1@inactsoft.com",
				Password:       "123456789087654",
				IsActive:       false,
			},
		},
	}

	mockReqConfGlobal := &pb.RequestConfigGlobal{}

	mockRespConfGlobalRes := &pb.ResponseConfigGlobal{}

	t.Run("Get Configuration Global", func(t *testing.T) {
		mockUseCaseConf.On("GetConfigurationGlobal", mock.Anything).Return(mockRespConfGlobal, nil).Once()

		handler := micro.NewMicroGrpc(mockUseCaseConf)

		err := handler.GetConfigurationGlobal(context.TODO(), mockReqConfGlobal, mockRespConfGlobalRes)

		assert.NoError(t, err)
		assert.Len(t, mockRespConfGlobalRes.GetConfigglobals(), 2)
	})

	t.Run("Failed Get Configuration Global", func(t *testing.T) {
		mockUseCaseConf.On("GetConfigurationGlobal", mock.Anything).Return(nil, errors.New("Unexpected syntax error")).Once()

		handler := micro.NewMicroGrpc(mockUseCaseConf)

		err := handler.GetConfigurationGlobal(context.TODO(), mockReqConfGlobal, mockRespConfGlobalRes)

		assert.Error(t, err)
	})
}

func TestGetConfigurationGlobalByID(t *testing.T) {
	mockUseCaseConf := new(mocks.Usecase)
	mockRespConfGlobal := &pb.ResponseConfigGlobal{
		Configglobal: &pb.ConfigurationGlobal{
			ConfigGlobalId: 1,
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

	mockReqConfGlobal := &pb.RequestConfigGlobal{
		Configglobal: &pb.ConfigurationGlobal{ConfigGlobalId: 1},
	}

	mockRespConfGlobalRes := &pb.ResponseConfigGlobal{}

	t.Run("Get Configuration By ID", func(t *testing.T) {
		mockUseCaseConf.On("GetConfigurationGlobalByID", mock.Anything, mock.AnythingOfType("int32")).Return(mockRespConfGlobal, nil).Once()

		handler := micro.NewMicroGrpc(mockUseCaseConf)

		err := handler.GetConfigurationGlobalByID(context.TODO(), mockReqConfGlobal, mockRespConfGlobalRes)

		assert.NoError(t, err)
		assert.Equal(t, mockRespConfGlobal.Configglobal.GetFootertext(), mockRespConfGlobalRes.Configglobal.GetFootertext())
	})

	t.Run("Get Configuration By ID", func(t *testing.T) {
		mockUseCaseConf.On("GetConfigurationGlobalByID", mock.Anything, mock.AnythingOfType("int32")).Return(nil, errors.New("Unexpected Syntax Error")).Once()

		handler := micro.NewMicroGrpc(mockUseCaseConf)

		err := handler.GetConfigurationGlobalByID(context.TODO(), mockReqConfGlobal, mockRespConfGlobalRes)

		assert.Error(t, err)
	})
}

func TestGetConfigurationGlobalActive(t *testing.T) {
	mockUseCaseConf := new(mocks.Usecase)
	mockRespConfGlobal := &pb.ResponseConfigGlobal{
		Configglobal: &pb.ConfigurationGlobal{
			ConfigGlobalId: 1,
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

	mockReqConfGlobal := &pb.RequestConfigGlobal{}

	mockRespConfGlobalRes := &pb.ResponseConfigGlobal{}

	t.Run("Get Configuration Global Active", func(t *testing.T) {
		mockUseCaseConf.On("GetConfigurationGlobalActive", mock.Anything).Return(mockRespConfGlobal, nil).Once()

		handler := micro.NewMicroGrpc(mockUseCaseConf)
		err := handler.GetConfigurationGlobalActive(context.TODO(), mockReqConfGlobal, mockRespConfGlobalRes)

		assert.NoError(t, err)
		assert.Equal(t, "notification1@inactsoft.com", mockRespConfGlobalRes.Configglobal.GetUsername())
	})

	t.Run("Failed Get Configuration Global Active", func(t *testing.T) {
		mockUseCaseConf.On("GetConfigurationGlobalActive", mock.Anything).Return(nil, errors.New("Unexpected Syntax Error")).Once()

		handler := micro.NewMicroGrpc(mockUseCaseConf)
		err := handler.GetConfigurationGlobalActive(context.TODO(), mockReqConfGlobal, mockRespConfGlobalRes)

		assert.Error(t, err)
	})
}

func TestSetConfigurationGlobalActive(t *testing.T) {
	mockUseCaseConf := new(mocks.Usecase)
	mockRespConfGlobal := &pb.ResponseConfigGlobal{
		Configstatus: &pb.ConfigurationStatus{Updated: true},
	}

	mockReqConfGlobal := &pb.RequestConfigGlobal{
		Configglobal: &pb.ConfigurationGlobal{ConfigGlobalId: 1},
	}

	mockRespConfGlobalRes := &pb.ResponseConfigGlobal{}

	t.Run("Set Configuration Global Active", func(t *testing.T) {
		mockUseCaseConf.On("SetConfigurationGlobalActive", mock.Anything, mock.AnythingOfType("*configuration.ConfigurationGlobal")).Return(mockRespConfGlobal, nil).Once()

		handler := micro.NewMicroGrpc(mockUseCaseConf)
		err := handler.SetConfigurationGlobalActive(context.TODO(), mockReqConfGlobal, mockRespConfGlobalRes)

		assert.NoError(t, err)
		assert.True(t, mockRespConfGlobalRes.Configstatus.GetUpdated())
	})

	t.Run("Set Configuration Global Active", func(t *testing.T) {
		mockUseCaseConf.On("SetConfigurationGlobalActive", mock.Anything, mock.AnythingOfType("*configuration.ConfigurationGlobal")).Return(nil, errors.New("Unexpected Error")).Once()

		handler := micro.NewMicroGrpc(mockUseCaseConf)
		err := handler.SetConfigurationGlobalActive(context.TODO(), mockReqConfGlobal, mockRespConfGlobalRes)

		assert.Error(t, err)
		assert.False(t, mockRespConfGlobalRes.Configstatus.GetUpdated())
	})
}
