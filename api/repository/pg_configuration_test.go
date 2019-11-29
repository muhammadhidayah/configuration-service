package repository_test

import (
	"context"
	"fmt"
	"testing"

	sqlMock "github.com/DATA-DOG/go-sqlmock"
	repo "github.com/muhammadhidayah/configuration-service/api/repository"
	pb "github.com/muhammadhidayah/configuration-service/proto/configuration"
	"github.com/stretchr/testify/assert"
)

// Testing Store Configuration to Table
func TestAddConfigurationClient(t *testing.T) {
	cc := &pb.ConfigurationClient{
		ConfigClientUuid:   "111-111-111-111",
		MultipleLanguageId: 3,
		Appname:            "client1.inactsoft.com",
		ReportTitle:        "Client 1",
		CompanySubsId:      "180-000-123-0321",
		IsConfigDeleted:    0,
	}

	db, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database conncection", err)
	}

	defer db.Close()

	prep := mock.ExpectPrepare("INSERT configuration_client")
	prep.ExpectExec().WithArgs(cc.ConfigClientUuid, cc.MultipleLanguageId, cc.Appname, cc.ReportTitle, cc.CompanySubsId, cc.IsConfigDeleted).WillReturnResult(sqlMock.NewResult(1, 1))

	clientRepo := repo.NewPgConfiguration(db)
	created, err := clientRepo.AddConfigurationClient(context.TODO(), cc)

	assert.NoError(t, err)
	assert.True(t, created)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// Testing success to update table configuration_client
func TestUpdateConfigurationClientByConfigClientIdSubs(t *testing.T) {
	cc := &pb.ConfigurationClient{
		ConfigClientUuid:   "a6e2745e-c930-4717-a9d1-d1cfb2a64aa4",
		MultipleLanguageId: 3,
		Appname:            "client1.inactsoft.com",
		ReportTitle:        "Client 1",
		CompanySubsId:      "180-000-123-0321",
		IsConfigDeleted:    0,
	}

	db, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	prep := mock.ExpectPrepare("UPDATE configuration_client")
	prep.ExpectExec().WithArgs(cc.MultipleLanguageId, cc.Appname, cc.ReportTitle, cc.CompanySubsId, cc.ConfigClientUuid).WillReturnResult(sqlMock.NewResult(0, 1))

	clientRepo := repo.NewPgConfiguration(db)
	updated, err := clientRepo.UpdateConfigurationClientBySubs(context.TODO(), cc)
	if err != nil {
		t.Fatalf("error cannot update to db %s", err.Error())
	}

	assert.True(t, updated)
}

// Testing failed to update table configuration_client
func TestFailUpdateConfigurationClientBySubs(t *testing.T) {
	cc := &pb.ConfigurationClient{
		ConfigClientUuid:   "",
		MultipleLanguageId: 3,
		Appname:            "client1.inactsoft.com",
		ReportTitle:        "Client 1",
		CompanySubsId:      "180-000-123-0321",
		IsConfigDeleted:    0,
	}

	db, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	prep := mock.ExpectPrepare("UPDATE configuration_client")
	prep.ExpectExec().WithArgs(cc.MultipleLanguageId, cc.Appname, cc.ReportTitle, cc.CompanySubsId, cc.ConfigClientUuid).WillReturnResult(sqlMock.NewResult(0, 0))

	clientRepo := repo.NewPgConfiguration(db)
	updated, err := clientRepo.UpdateConfigurationClientBySubs(context.TODO(), cc)

	assert.Error(t, err)
	assert.False(t, updated)
}

// Testing success to Delete (actually update flags is_deleted)
func TestDeleteConfigurationClientBySubs(t *testing.T) {
	cc := &pb.ConfigurationClient{
		CompanySubsId: "180-000-123-0321",
	}

	db, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	prepare := mock.ExpectPrepare("UPDATE configuration_client")
	prepare.ExpectExec().WithArgs(cc.CompanySubsId).WillReturnResult(sqlMock.NewResult(0, 1))

	clientRepo := repo.NewPgConfiguration(db)
	deleted, err := clientRepo.DeleteConfigurationClientBySubs(context.TODO(), cc)
	assert.NoError(t, err)
	assert.True(t, deleted)
}

// Testing success to Delete (actually update flags is_deleted)
func TestFailedDeleteConfigurationClientBySubs(t *testing.T) {
	cc := &pb.ConfigurationClient{
		CompanySubsId: "",
	}

	db, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	prepare := mock.ExpectPrepare("UPDATE configuration_client")
	prepare.ExpectExec().WithArgs(cc.CompanySubsId).WillReturnResult(sqlMock.NewResult(0, 0))

	clientRepo := repo.NewPgConfiguration(db)
	deleted, err := clientRepo.DeleteConfigurationClientBySubs(context.TODO(), cc)
	assert.Error(t, err)
	assert.False(t, deleted)
}

func TestGetConfigurationClientBySubs(t *testing.T) {
	mockConfigurationClient := []pb.ConfigurationClient{
		pb.ConfigurationClient{
			ConfigClientId:     1,
			ConfigClientUuid:   "a6e2745e-c930-4717-a9d1-d1cfb2a64aa4",
			MultipleLanguageId: 2,
			Appname:            "client1.inactsoft.com",
			ReportTitle:        "Client Satu",
			CompanySubsId:      "012-031-234-542",
			IsConfigDeleted:    0,
		},
		pb.ConfigurationClient{
			ConfigClientId:     2,
			ConfigClientUuid:   "bf8bd542-a347-4cd1-838e-8b0debecb0f3",
			MultipleLanguageId: 3,
			Appname:            "client2.inactsoft.com",
			ReportTitle:        "Client Dua",
			CompanySubsId:      "011-021-234-542",
			IsConfigDeleted:    0,
		},
	}

	db, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlMock.NewRows([]string{"config_client_id", "config_client_uuid", "multiple_language_id", "appname", "report_title", "company_subs_id", "is_config_deleted"}).AddRow(mockConfigurationClient[1].ConfigClientId, mockConfigurationClient[1].ConfigClientUuid, mockConfigurationClient[1].MultipleLanguageId, mockConfigurationClient[1].Appname, mockConfigurationClient[1].ReportTitle, mockConfigurationClient[1].CompanySubsId, mockConfigurationClient[1].IsConfigDeleted)

	mock.ExpectQuery("SELECT config_client_id, config_client_uuid, multiple_language_id, appname, report_title, company_subs_id, is_config_deleted FROM configuration_client WHERE company_subs_id = \\? AND is_config_deleted = 0").WillReturnRows(rows)

	clientRepo := repo.NewPgConfiguration(db)

	companySubsId := "011-021-234-542"
	res, err := clientRepo.GetConfigurationClientBySubs(context.TODO(), companySubsId)
	assert.NoError(t, err)
	assert.Equal(t, "bf8bd542-a347-4cd1-838e-8b0debecb0f3", res.ConfigClientUuid)
}

func TestGetConfigurationClientBySubsNoData(t *testing.T) {

	db, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlMock.NewRows([]string{"config_client_id", "config_client_uuid", "multiple_language_id", "appname", "report_title", "company_subs_id", "is_config_deleted"})

	mock.ExpectQuery("SELECT config_client_id, config_client_uuid, multiple_language_id, appname, report_title, company_subs_id, is_config_deleted FROM configuration_client WHERE company_subs_id = \\? AND is_config_deleted = 0").WillReturnRows(rows)

	clientRepo := repo.NewPgConfiguration(db)

	companySubsId := "011-021-234-542"
	res, err := clientRepo.GetConfigurationClientBySubs(context.TODO(), companySubsId)
	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestGetConfigurationClient(t *testing.T) {
	mockConfigurationClient := []pb.ConfigurationClient{
		pb.ConfigurationClient{
			ConfigClientId:     1,
			ConfigClientUuid:   "a6e2745e-c930-4717-a9d1-d1cfb2a64aa4",
			MultipleLanguageId: 2,
			Appname:            "client1.inactsoft.com",
			ReportTitle:        "Client Satu",
			CompanySubsId:      "012-031-234-542",
			IsConfigDeleted:    0,
		},
		pb.ConfigurationClient{
			ConfigClientId:     2,
			ConfigClientUuid:   "bf8bd542-a347-4cd1-838e-8b0debecb0f3",
			MultipleLanguageId: 3,
			Appname:            "client2.inactsoft.com",
			ReportTitle:        "Client Dua",
			CompanySubsId:      "011-021-234-542",
			IsConfigDeleted:    0,
		},
	}

	db, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlMock.NewRows([]string{"config_client_id", "config_client_uuid", "multiple_language_id", "appname", "report_title", "company_subs_id", "is_config_deleted"}).AddRow(mockConfigurationClient[0].ConfigClientId, mockConfigurationClient[0].ConfigClientUuid, mockConfigurationClient[0].MultipleLanguageId, mockConfigurationClient[0].Appname, mockConfigurationClient[0].ReportTitle, mockConfigurationClient[0].CompanySubsId, mockConfigurationClient[0].IsConfigDeleted).AddRow(mockConfigurationClient[1].ConfigClientId, mockConfigurationClient[1].ConfigClientUuid, mockConfigurationClient[1].MultipleLanguageId, mockConfigurationClient[1].Appname, mockConfigurationClient[1].ReportTitle, mockConfigurationClient[1].CompanySubsId, mockConfigurationClient[1].IsConfigDeleted)

	mock.ExpectQuery("SELECT config_client_id, config_client_uuid, multiple_language_id, appname, report_title, company_subs_id, is_config_deleted FROM configuration_client").WillReturnRows(rows)

	clientRepo := repo.NewPgConfiguration(db)
	res, err := clientRepo.GetConfigurationClient(context.TODO())
	assert.NoError(t, err)
	assert.Len(t, res, 2)
}

func TestGetConfigurationClientNoData(t *testing.T) {
	db, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlMock.NewRows([]string{"config_client_id", "config_client_uuid", "multiple_language_id", "appname", "report_title", "company_subs_id", "is_config_deleted"})

	mock.ExpectQuery("SELECT config_client_id, config_client_uuid, multiple_language_id, appname, report_title, company_subs_id, is_config_deleted FROM configuration_client").WillReturnRows(rows)

	clientRepo := repo.NewPgConfiguration(db)
	res, err := clientRepo.GetConfigurationClient(context.TODO())
	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestAddConfigurationGlobal(t *testing.T) {
	cg := &pb.ConfigurationGlobal{
		Footertext: "Technical support : +62-21-7509077 ext. 109 | Email: support@inactsoft.com | <a href='http://wiki.inactsoft.com' target='_blank' class='footer-link'>online help</a>",
		ServerSmpt: "mail.google.com",
		Ssl:        true,
		Port:       1234,
		IsAuth:     true,
		Username:   "notification@inactsoft.com",
		Password:   "123456789087654",
	}

	db, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database conncection", err)
	}

	defer db.Close()

	prep := mock.ExpectPrepare("INSERT configuration_global")
	prep.ExpectExec().WithArgs(cg.Footertext, cg.ServerSmpt, cg.Ssl, cg.Port, cg.IsAuth, cg.Username, cg.Password).WillReturnResult(sqlMock.NewResult(1, 1))

	clientRepo := repo.NewPgConfiguration(db)
	created, err := clientRepo.AddConfigurationGlobal(context.TODO(), cg)
	assert.True(t, created)
	assert.NoError(t, err)
}

func TestUpdateConfigurationGlobal(t *testing.T) {
	cg := &pb.ConfigurationGlobal{
		ConfigGlobalId: 1,
		Footertext:     "Technical support : +62-21-7509077 ext. 109 | Email: support@inactsoft.com | <a href='http://wiki.inactsoft.com' target='_blank' class='footer-link'>online help</a>",
		ServerSmpt:     "mail.google.com",
		Ssl:            true,
		Port:           5431,
		IsAuth:         true,
		Username:       "notification@inactsoft.com",
		Password:       "123456789087654",
	}

	db, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database conncection", err)
	}

	defer db.Close()

	prep := mock.ExpectPrepare("UPDATE configuration_global")
	prep.ExpectExec().WithArgs(cg.Footertext, cg.ServerSmpt, cg.Ssl, cg.Port, cg.IsAuth, cg.Username, cg.Password, cg.ConfigGlobalId).WillReturnResult(sqlMock.NewResult(1, 1))

	configRepo := repo.NewPgConfiguration(db)
	updated, err := configRepo.UpdateConfigurationGlobal(context.TODO(), cg)
	assert.True(t, updated)
	assert.NoError(t, err)
}

func TestUpdateConfigurationGlobalNoData(t *testing.T) {
	cg := &pb.ConfigurationGlobal{
		ConfigGlobalId: 1,
		Footertext:     "Technical support : +62-21-7509077 ext. 109 | Email: support@inactsoft.com | <a href='http://wiki.inactsoft.com' target='_blank' class='footer-link'>online help</a>",
		ServerSmpt:     "mail.google.com",
		Ssl:            true,
		Port:           5431,
		IsAuth:         true,
		Username:       "notification@inactsoft.com",
		Password:       "123456789087654",
	}

	db, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database conncection", err)
	}

	defer db.Close()

	prep := mock.ExpectPrepare("UPDATE configuration_global")
	prep.ExpectExec().WithArgs(cg.Footertext, cg.ServerSmpt, cg.Ssl, cg.Port, cg.IsAuth, cg.Username, cg.Password, cg.ConfigGlobalId).WillReturnResult(sqlMock.NewResult(0, 0))

	configRepo := repo.NewPgConfiguration(db)
	updated, err := configRepo.UpdateConfigurationGlobal(context.TODO(), cg)
	assert.False(t, updated)
	assert.Error(t, err)
}

func TestDeleteConfiguration(t *testing.T) {
	db, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database conncection", err)
	}

	defer db.Close()

	cg := &pb.ConfigurationGlobal{ConfigGlobalId: 1}

	prep := mock.ExpectPrepare("DELETE FROM configuration_global")
	prep.ExpectExec().WithArgs(cg.ConfigGlobalId).WillReturnResult(sqlMock.NewResult(0, 1))

	configRepo := repo.NewPgConfiguration(db)
	deleted, err := configRepo.DeleteConfiguration(context.TODO(), cg.ConfigGlobalId)

	assert.True(t, deleted)
	assert.NoError(t, err)
}

func TestFailDeleteConfigurationNoData(t *testing.T) {
	db, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database conncection", err)
	}

	defer db.Close()

	cg := &pb.ConfigurationGlobal{ConfigGlobalId: 1}

	prep := mock.ExpectPrepare("DELETE FROM configuration_global")
	prep.ExpectExec().WillReturnResult(sqlMock.NewResult(0, 0))

	configRepo := repo.NewPgConfiguration(db)
	deleted, err := configRepo.DeleteConfiguration(context.TODO(), cg.ConfigGlobalId)

	assert.False(t, deleted)
	assert.Error(t, err)
}

func TestFailDeleteConfigurationSyntaxErr(t *testing.T) {
	db, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database conncection", err)
	}

	defer db.Close()

	cg := &pb.ConfigurationGlobal{ConfigGlobalId: 1}

	prep := mock.ExpectPrepare("DELETE FROM configuration_global")
	prep.ExpectExec().WillReturnError(fmt.Errorf("configuration_global_id not exists"))

	configRepo := repo.NewPgConfiguration(db)
	deleted, err := configRepo.DeleteConfiguration(context.TODO(), cg.ConfigGlobalId)

	assert.False(t, deleted)
	assert.Error(t, err)
}

func TestGetConfigurationGlobal(t *testing.T) {
	mockConfigurationGlobal := []pb.ConfigurationGlobal{
		{
			ConfigGlobalId: 1,
			Footertext:     "Technical support : +62-21-7509077 ext. 109 | Email: support@inactsoft.com | <a href='http://wiki.inactsoft.com' target='_blank' class='footer-link'>online help</a>",
			ServerSmpt:     "mail.google.com",
			Ssl:            true,
			Port:           5431,
			IsAuth:         true,
			Username:       "notification3@inactsoft.com",
			Password:       "123456789087654",
			IsActive:       false,
		}, {
			ConfigGlobalId: 2,
			Footertext:     "Technical support : +62-21-7509077 ext. 109 | Email: support@inactsoft.com | <a href='http://wiki.inactsoft.com' target='_blank' class='footer-link'>online help</a>",
			ServerSmpt:     "mail.google.com",
			Ssl:            true,
			Port:           5431,
			IsAuth:         true,
			Username:       "notification2@inactsoft.com",
			Password:       "123456789087654",
			IsActive:       false,
		}, {
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

	db, mock, err := sqlMock.New()
	query := "SELECT config_global_id, footertext, server_smpt, ssl, port, is_auth, username, password, is_active FROM configuration_global"

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database conncection", err)
	}

	defer db.Close()

	t.Run("Get Configuration Global Return all data", func(t *testing.T) {
		rows := sqlMock.NewRows([]string{"config_global_id", "footertext", "server_smpt", "ssl", "port", "is_auth", "username", "password", "is_active"}).AddRow(mockConfigurationGlobal[0].ConfigGlobalId, mockConfigurationGlobal[0].Footertext, mockConfigurationGlobal[0].ServerSmpt, mockConfigurationGlobal[0].Ssl, mockConfigurationGlobal[0].Port, mockConfigurationGlobal[0].IsAuth, mockConfigurationGlobal[0].Username, mockConfigurationGlobal[0].Password, mockConfigurationGlobal[0].IsActive).AddRow(mockConfigurationGlobal[1].ConfigGlobalId, mockConfigurationGlobal[1].Footertext, mockConfigurationGlobal[1].ServerSmpt, mockConfigurationGlobal[1].Ssl, mockConfigurationGlobal[1].Port, mockConfigurationGlobal[1].IsAuth, mockConfigurationGlobal[1].Username, mockConfigurationGlobal[1].Password, mockConfigurationGlobal[1].IsActive).AddRow(mockConfigurationGlobal[2].ConfigGlobalId, mockConfigurationGlobal[2].Footertext, mockConfigurationGlobal[2].ServerSmpt, mockConfigurationGlobal[2].Ssl, mockConfigurationGlobal[2].Port, mockConfigurationGlobal[2].IsAuth, mockConfigurationGlobal[2].Username, mockConfigurationGlobal[2].Password, mockConfigurationGlobal[2].IsActive)

		sqlrows := mock.ExpectQuery(query)
		sqlrows.WillReturnRows(rows)

		configRepo := repo.NewPgConfiguration(db)
		res, err := configRepo.GetConfigurationGlobal(context.TODO())
		assert.NoError(t, err)
		assert.Len(t, res, 3)
	})

	t.Run("Get Configuration Global, but error failed query", func(t *testing.T) {
		sqlRows := mock.ExpectQuery(query)
		sqlRows.WillReturnError(fmt.Errorf("table or column doesnt exists"))

		clientRepo := repo.NewPgConfiguration(db)
		res, err := clientRepo.GetConfigurationGlobal(context.TODO())
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("Get Configuration, but error scan because nil data", func(t *testing.T) {
		rows := sqlMock.NewRows([]string{"config_global_id", "footertext", "server_smpt", "ssl", "port", "is_auth", "username", "password", "is_active"}).AddRow(mockConfigurationGlobal[0].ConfigGlobalId, nil, mockConfigurationGlobal[0].ServerSmpt, mockConfigurationGlobal[0].Ssl, mockConfigurationGlobal[0].Port, mockConfigurationGlobal[0].IsAuth, mockConfigurationGlobal[0].Username, mockConfigurationGlobal[0].Password, mockConfigurationGlobal[0].IsActive)

		sqlRows := mock.ExpectQuery(query)
		sqlRows.WillReturnRows(rows)

		clientRepo := repo.NewPgConfiguration(db)
		res, err := clientRepo.GetConfigurationGlobal(context.TODO())
		assert.Error(t, err)
		assert.Nil(t, res)

	})
}

func TestGetConfigurationGlobalByID(t *testing.T) {
	mockConfigurationGlobal := []pb.ConfigurationGlobal{
		{
			ConfigGlobalId: 1,
			Footertext:     "Technical support : +62-21-7509077 ext. 109 | Email: support@inactsoft.com | <a href='http://wiki.inactsoft.com' target='_blank' class='footer-link'>online help</a>",
			ServerSmpt:     "mail.google.com",
			Ssl:            true,
			Port:           5431,
			IsAuth:         true,
			Username:       "notification3@inactsoft.com",
			Password:       "123456789087654",
			IsActive:       false,
		}, {
			ConfigGlobalId: 2,
			Footertext:     "Technical support : +62-21-7509077 ext. 109 | Email: support@inactsoft.com | <a href='http://wiki.inactsoft.com' target='_blank' class='footer-link'>online help</a>",
			ServerSmpt:     "mail.google.com",
			Ssl:            true,
			Port:           5431,
			IsAuth:         true,
			Username:       "notification2@inactsoft.com",
			Password:       "123456789087654",
			IsActive:       false,
		}, {
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

	db, mock, err := sqlMock.New()
	query := "SELECT config_global_id, footertext, server_smpt, ssl, port, is_auth, username, password, is_active FROM configuration_global WHERE config_global_id = \\?"

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database conncection", err)
	}

	defer db.Close()

	t.Run("Get Configuration Global using condition configuration id", func(t *testing.T) {
		rows := sqlMock.NewRows([]string{"config_global_id", "footertext", "server_smpt", "ssl", "port", "is_auth", "username", "password", "is_active"}).AddRow(mockConfigurationGlobal[0].ConfigGlobalId, mockConfigurationGlobal[0].Footertext, mockConfigurationGlobal[0].ServerSmpt, mockConfigurationGlobal[0].Ssl, mockConfigurationGlobal[0].Port, mockConfigurationGlobal[0].IsAuth, mockConfigurationGlobal[0].Username, mockConfigurationGlobal[0].Password, mockConfigurationGlobal[0].IsActive)

		sqlRows := mock.ExpectQuery(query)
		sqlRows.WithArgs(mockConfigurationGlobal[0].ConfigGlobalId).WillReturnRows(rows)

		clientRepo := repo.NewPgConfiguration(db)
		data, err := clientRepo.GetConfigurationGlobalByID(context.TODO(), mockConfigurationGlobal[0].ConfigGlobalId)
		assert.NoError(t, err)
		assert.Equal(t, mockConfigurationGlobal[0].ConfigGlobalId, data.ConfigGlobalId)
	})

	t.Run("Get Configuration Global using condition configuration id, when no data", func(t *testing.T) {
		rows := sqlMock.NewRows([]string{"config_global_id", "footertext", "server_smpt", "ssl", "port", "is_auth", "username", "password", "is_active"})

		sqlRows := mock.ExpectQuery(query)
		sqlRows.WithArgs(mockConfigurationGlobal[0].ConfigGlobalId).WillReturnRows(rows)

		clientRepo := repo.NewPgConfiguration(db)
		data, err := clientRepo.GetConfigurationGlobalByID(context.TODO(), mockConfigurationGlobal[0].ConfigGlobalId)
		assert.NoError(t, err)
		assert.Nil(t, data)
	})

	t.Run("Get Configuration Global using condition configuration id, then syntax error", func(t *testing.T) {
		sqlRows := mock.ExpectQuery(query)
		sqlRows.WithArgs(mockConfigurationGlobal[0].ConfigGlobalId).WillReturnError(fmt.Errorf("Sql syntax error"))

		clientRepo := repo.NewPgConfiguration(db)
		data, err := clientRepo.GetConfigurationGlobalByID(context.TODO(), mockConfigurationGlobal[0].ConfigGlobalId)
		assert.Error(t, err)
		assert.Nil(t, data)
	})
}

func TestGetConfigurationGlobalActive(t *testing.T) {
	mockConfigurationGlobal := []pb.ConfigurationGlobal{
		{
			ConfigGlobalId: 1,
			Footertext:     "Technical support : +62-21-7509077 ext. 109 | Email: support@inactsoft.com | <a href='http://wiki.inactsoft.com' target='_blank' class='footer-link'>online help</a>",
			ServerSmpt:     "mail.google.com",
			Ssl:            true,
			Port:           5431,
			IsAuth:         true,
			Username:       "notification3@inactsoft.com",
			Password:       "123456789087654",
			IsActive:       false,
		}, {
			ConfigGlobalId: 2,
			Footertext:     "Technical support : +62-21-7509077 ext. 109 | Email: support@inactsoft.com | <a href='http://wiki.inactsoft.com' target='_blank' class='footer-link'>online help</a>",
			ServerSmpt:     "mail.google.com",
			Ssl:            true,
			Port:           5431,
			IsAuth:         true,
			Username:       "notification2@inactsoft.com",
			Password:       "123456789087654",
			IsActive:       false,
		}, {
			ConfigGlobalId: 3,
			Footertext:     "Technical support : +62-21-7509077 ext. 109 | Email: support@inactsoft.com | <a href='http://wiki.inactsoft.com' target='_blank' class='footer-link'>online help</a>",
			ServerSmpt:     "mail.google.com",
			Ssl:            false,
			Port:           5432,
			IsAuth:         true,
			Username:       "notification1@inactsoft.com",
			Password:       "123456789087654",
			IsActive:       true,
		},
	}

	db, mock, err := sqlMock.New()
	query := "SELECT config_global_id, footertext, server_smpt, ssl, port, is_auth, username, password, is_active FROM configuration_global WHERE is_active = \\?"

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database conncection", err)
	}

	defer db.Close()

	t.Run("Get Configuration Global Active (Success)", func(t *testing.T) {
		rows := sqlMock.NewRows([]string{"config_global_id", "footertext", "server_smpt", "ssl", "port", "is_auth", "username", "password", "is_active"}).AddRow(mockConfigurationGlobal[2].ConfigGlobalId, mockConfigurationGlobal[2].Footertext, mockConfigurationGlobal[2].ServerSmpt, mockConfigurationGlobal[2].Ssl, mockConfigurationGlobal[2].Port, mockConfigurationGlobal[2].IsAuth, mockConfigurationGlobal[2].Username, mockConfigurationGlobal[2].Password, mockConfigurationGlobal[2].IsActive)

		queryExpect := mock.ExpectQuery(query)
		queryExpect.WithArgs(mockConfigurationGlobal[2].IsActive).WillReturnRows(rows)

		clientRepo := repo.NewPgConfiguration(db)
		data, err := clientRepo.GetConfigurationGlobalActive(context.TODO())
		assert.NoError(t, err)
		assert.Equal(t, mockConfigurationGlobal[2].ConfigGlobalId, data.ConfigGlobalId)
	})

	t.Run("Get Configuration Global Active. SQL syntax error", func(t *testing.T) {
		queryExpect := mock.ExpectQuery(query)
		queryExpect.WithArgs(mockConfigurationGlobal[2].IsActive).WillReturnError(fmt.Errorf("SQL syntax error"))

		clientRepo := repo.NewPgConfiguration(db)
		data, err := clientRepo.GetConfigurationGlobalActive(context.TODO())
		assert.Error(t, err)
		assert.Nil(t, data)
	})

	t.Run("Get Configuration Global Active. Data Nil Fom DB", func(t *testing.T) {
		rows := sqlMock.NewRows([]string{"config_global_id", "footertext", "server_smpt", "ssl", "port", "is_auth", "username", "password", "is_active"}).AddRow(mockConfigurationGlobal[2].ConfigGlobalId, mockConfigurationGlobal[2].Footertext, mockConfigurationGlobal[2].ServerSmpt, mockConfigurationGlobal[2].Ssl, mockConfigurationGlobal[2].Port, mockConfigurationGlobal[2].IsAuth, nil, mockConfigurationGlobal[2].Password, mockConfigurationGlobal[2].IsActive)

		queryExpect := mock.ExpectQuery(query)
		queryExpect.WithArgs(mockConfigurationGlobal[2].IsActive).WillReturnRows(rows)

		clientRepo := repo.NewPgConfiguration(db)
		data, err := clientRepo.GetConfigurationGlobalActive(context.TODO())
		assert.Error(t, err)
		assert.Nil(t, data)
	})

	t.Run("Get Configuration Global Active. No DataData", func(t *testing.T) {
		rows := sqlMock.NewRows([]string{"config_global_id", "footertext", "server_smpt", "ssl", "port", "is_auth", "username", "password", "is_active"})

		queryExpect := mock.ExpectQuery(query)
		queryExpect.WithArgs(mockConfigurationGlobal[2].IsActive).WillReturnRows(rows)

		clientRepo := repo.NewPgConfiguration(db)
		data, err := clientRepo.GetConfigurationGlobalActive(context.TODO())
		assert.NoError(t, err)
		assert.Nil(t, data)
	})

}
