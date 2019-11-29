package repository

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
	"github.com/muhammadhidayah/configuration-service/api"
	pb "github.com/muhammadhidayah/configuration-service/proto/configuration"
)

type pgConfiguration struct {
	conn *sql.DB
}

func NewPgConfiguration(conn *sql.DB) api.Repository {
	return &pgConfiguration{conn}
}

// this function will be used to add configuration client
func (repo *pgConfiguration) AddConfigurationClient(ctx context.Context, cc *pb.ConfigurationClient) (bool, error) {
	query := "INSERT configuration_client INTO (config_client_uuid,multiple_language_id, appname, report_title, company_subs_id, is_config_deleted) VALUES(?,?,?,?,?)"

	// using function handlingStoreQuery to inserting in table configuration_client
	_, err := repo.handlingStoreQuery(ctx, query, cc.ConfigClientUuid, cc.MultipleLanguageId, cc.Appname, cc.ReportTitle, cc.CompanySubsId, cc.IsConfigDeleted)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (repo *pgConfiguration) UpdateConfigurationClientBySubs(ctx context.Context, cc *pb.ConfigurationClient) (bool, error) {
	query := "UPDATE configuration_client SET multiple_language_id = ?, appname = ?, report_title = ?, company_subs_id = ? WHERE config_client_uuid = ?"

	res, err := repo.handlingStoreQuery(ctx, query, cc.MultipleLanguageId, cc.Appname, cc.ReportTitle, cc.CompanySubsId, cc.ConfigClientUuid)

	if err != nil {
		return false, err
	}

	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
		return false, errors.New("Data Not Found to Update")
	}

	return true, nil
}

func (repo *pgConfiguration) DeleteConfigurationClientBySubs(ctx context.Context, cc *pb.ConfigurationClient) (bool, error) {
	query := "UPDATE configuration_client SET is_config_deleted = 1 WHERE company_subs_id = ?"

	res, err := repo.handlingStoreQuery(ctx, query, cc.CompanySubsId)
	if err != nil {
		return false, err
	}

	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
		return false, errors.New("Data Not Found to Delete")
	}

	return true, nil
}

func (repo *pgConfiguration) handlingStoreQuery(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	stmt, err := repo.conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	return stmt.ExecContext(ctx, args...)
}

// this function will return array pointer of configurationClient and error
// in params query, query must follow column name as sequentially : config_client_id, multiple_language_id, appname, report_title, company_subs_id, is_config_deleted
func (repo *pgConfiguration) fetchDataConfigClient(ctx context.Context, query string, args ...interface{}) ([]*pb.ConfigurationClient, error) {
	// execute query using querycontext
	rows, err := repo.conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	// rows will be closed when this function ended
	defer rows.Close()

	// make array of configurationClient with size 0 for the first
	configClient := make([]*pb.ConfigurationClient, 0)

	// iteration to get data all rows
	for rows.Next() {
		// this variable is temporary, will be use to containt value of row
		temp := &pb.ConfigurationClient{}

		// mapping data of row to field in variable temp, and if error will be store in variable err. and function will be exit
		err = rows.Scan(
			&temp.ConfigClientId,
			&temp.ConfigClientUuid,
			&temp.MultipleLanguageId,
			&temp.Appname,
			&temp.ReportTitle,
			&temp.CompanySubsId,
			&temp.IsConfigDeleted,
		)

		if err != nil {
			return configClient, err
		}

		// append value from temp to array of configClient
		configClient = append(configClient, temp)
	}

	return configClient, nil
}

// this function will fetch data of configurationclient with have condition company_subs_id. then this function return pointer of configurationClient and error
func (repo *pgConfiguration) GetConfigurationClientBySubs(ctx context.Context, clientSubsID string) (*pb.ConfigurationClient, error) {
	query := "SELECT config_client_id, config_client_uuid, multiple_language_id, appname, report_title, company_subs_id, is_config_deleted FROM configuration_client WHERE company_subs_id = ? AND is_config_deleted = 0"

	// for quering, will using function fetchDataConfigClient
	res, err := repo.fetchDataConfigClient(ctx, query, clientSubsID)
	if err != nil {
		return nil, err
	}

	// checking variable res, result of quering. if len res more than zero, will return only 1 data index 0
	if len(res) > 0 {
		return res[0], nil
	}

	// return error, if res has no data
	return nil, errors.New("Data Not Found")
}

// this function will fetch all of data configurationclient, and will return pointer configurationclient in array, and error
func (repo *pgConfiguration) GetConfigurationClient(ctx context.Context) ([]*pb.ConfigurationClient, error) {
	query := "SELECT config_client_id, config_client_uuid, multiple_language_id, appname, report_title, company_subs_id, is_config_deleted FROM configuration_client"

	// for quering, will using function fetchDataConfigClient
	res, err := repo.fetchDataConfigClient(ctx, query)
	if err != nil {
		return nil, err
	}

	// checking variable res, result of quering. if len res more than zero, will return all data
	if len(res) > 0 {
		return res, nil
	}

	// return error, if res has no data
	return nil, errors.New("Data Not Found")
}

// this function will store data to configuration_global, return bool and error
func (repo *pgConfiguration) AddConfigurationGlobal(ctx context.Context, cg *pb.ConfigurationGlobal) (bool, error) {
	query := "INSERT configuration_global INTO (footer_text, server_smpt, ssl, port, is_auth, username, password) VALUES(?,?,?,?,?,?,?)"

	// insert data to table configuration_global use handlingStoreQuhandlingStoreQueryery function of pgRepository
	_, err := repo.handlingStoreQuery(ctx, query, cg.Footertext, cg.ServerSmpt, cg.Ssl, cg.Port, cg.IsAuth, cg.Username, cg.Password)
	if err != nil {
		return false, err
	}

	return true, nil
}

// this function will update data to configuration_global with condition config_global_id, return bool and error
func (repo *pgConfiguration) UpdateConfigurationGlobal(ctx context.Context, cg *pb.ConfigurationGlobal) (bool, error) {
	query := "UPDATE configuration_global SET footer_text = ?, server_smtp = ?, ssl = ?, port = ?, is_auth = ?, username = ?, password = ? WHERE config_global_id = ?"

	// to execute query to update data in table configuration_global use handlingStoreQuery function of pgRepository
	res, err := repo.handlingStoreQuery(ctx, query, cg.Footertext, cg.ServerSmpt, cg.Ssl, cg.Port, cg.IsAuth, cg.Username, cg.Password, cg.ConfigGlobalId)
	if err != nil {
		return false, err
	}

	// check is row affected or not. if not will return false, and error because no data to updated
	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
		return false, errors.New("No Data to Update")
	}

	return true, nil
}

// this function will delete row by id in table configuration_global. return bool and error
func (repo *pgConfiguration) DeleteConfiguration(ctx context.Context, configGlobalID int32) (bool, error) {
	query := "DELETE FROM configuration_global WHERE configuration_global_id = ?"

	// to execute query delete in table configuration_global will use handlingStoreQuery function of pgRepository
	res, err := repo.handlingStoreQuery(ctx, query, configGlobalID)
	if err != nil {
		return false, err
	}

	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
		return false, errors.New("No Data to Delete From DB")
	}

	return true, nil
}

func (repo *pgConfiguration) fetchConfigurationGlobal(ctx context.Context, query string, args ...interface{}) ([]*pb.ConfigurationGlobal, error) {
	// execute query, and get all data in rows. if error will store in variable err
	rows, err := repo.conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// initiate map of pointer configurationGlobal with size 0
	dataConfigGlobals := make([]*pb.ConfigurationGlobal, 0)

	// iteration for data rows
	for rows.Next() {
		// variable to containt data temporary. and will be used to append in dataConfigGlobals
		temp := &pb.ConfigurationGlobal{}

		// mapping data to field of temp
		err = rows.Scan(
			&temp.ConfigGlobalId,
			&temp.Footertext,
			&temp.ServerSmpt,
			&temp.Ssl,
			&temp.Port,
			&temp.IsAuth,
			&temp.Username,
			&temp.Password,
			&temp.IsActive,
		)

		if err != nil {
			return nil, err
		}

		// append data config temporary to dataConfigGlobals
		dataConfigGlobals = append(dataConfigGlobals, temp)
	}

	return dataConfigGlobals, nil
}

// this function will fetch all data rows, return array of pointer configurationGlobal and error
func (repo *pgConfiguration) GetConfigurationGlobal(ctx context.Context) ([]*pb.ConfigurationGlobal, error) {
	query := "SELECT config_global_id, footertext, server_smpt, ssl, port, is_auth, username, password, is_active FROM configuration_global"

	// execute query, and get all data in rows. if error will store in variable err
	dataConfigGlobals, err := repo.fetchConfigurationGlobal(ctx, query)

	if err != nil {
		return nil, err
	}

	return dataConfigGlobals, nil
}

// this function will return data pointer ConfigurationGlobal and error. this function will query to table configuration_global with condition configuration_global_id must equal configGlobalID (param 2)
func (repo *pgConfiguration) GetConfigurationGlobalByID(ctx context.Context, configGlobalID int32) (*pb.ConfigurationGlobal, error) {
	query := "SELECT config_global_id, footertext, server_smpt, ssl, port, is_auth, username, password, is_active FROM configuration_global WHERE config_global_id = ?"

	// execute query, and get all data in rows using condition config_global_id must equal. if error will store in variable err
	data, err := repo.fetchConfigurationGlobal(ctx, query, configGlobalID)

	if err != nil {
		return nil, err
	}

	// checking if data more than 0, will return only one data. index 0
	if len(data) > 0 {
		return data[0], nil
	}

	return nil, nil
}

// this function will return data pointer ConfigurationGlobal and error. this function will query to table configuration_global with condition configration is active
func (repo *pgConfiguration) GetConfigurationGlobalActive(ctx context.Context) (*pb.ConfigurationGlobal, error) {
	query := "SELECT config_global_id, footertext, server_smpt, ssl, port, is_auth, username, password, is_active FROM configuration_global WHERE is_active = ?"

	// execute query to get data configuration_global is active
	res, err := repo.fetchConfigurationGlobal(ctx, query, true)
	if err != nil {
		return nil, err
	}

	if len(res) > 0 {
		return res[0], nil
	}

	return nil, nil
}
