package config


import (
	"database/sql"
	"github.com/jfeng45/k8sdemo/dataservice"
	"github.com/jfeng45/k8sdemo/dataservice/userdata"
	"github.com/jfeng45/k8sdemo/usecase"
	"github.com/jfeng45/k8sdemo/usecase/registration"
	"github.com/jfeng45/k8sdemo/tool"
	"os"
)

type dbConfig struct {
	dbHost     string
	dbPort     string
	dbDatabase string
	dbUser string
	dbPassword string
}
func BuildRegistrationInterface(dataServiceName string)(usecase.RegistrationUseCaseInterface, error) {
	//uds := user
	RegisterLogrusLog()
	//RegisterZapLog()
	uds, err := BuildUserDataInterface(dataServiceName)
	if err != nil {
		return nil, err
	}
	ucsi := registration.UseCase{uds}
	return &ucsi, nil
}


func BuildUserDataInterface(dataServiceName string) (dataservice.UserDataInterface, error) {
	//var ds dataStoreUserI
	//var err error
	switch dataServiceName {
	case "mysql":
		return buildMysql()
	default:
		return buildMysql()
	}
	//return ds, err
}


func buildMysql() (dataservice.UserDataInterface, error) {
	tool.Log.Debug("connect to database ")
	dc :=  buildDbConfig ()
	dataSourceName := dc.dbUser + ":"+ dc.dbPassword + "@tcp(" +dc.dbHost +":" +dc.dbPort +")/" + dc.dbDatabase + "?charset=utf8";
	tool.Log.Debug("dataSourceName:", dataSourceName)
	//db, err := sql.Open("mysql", "dbuser:dbuser@tcp(localhost:3306)/service_config?charset=utf8")
	db, err := sql.Open("mysql", dataSourceName)
	checkErr(err)
	dataService := userdata.UserDataMysql{DB: db}
	return &dataService, err
}

func buildDbConfig () dbConfig{
	dc :=dbConfig{}
	dc.dbHost = os.Getenv("MYSQL_HOST")
	dc.dbPort = os.Getenv("MYSQL_PORT")
	dc.dbDatabase = os.Getenv("MYSQL_DATABASE")
	dc.dbUser = os.Getenv("MYSQL_USER_NAME")
	dc.dbPassword = os.Getenv("MYSQL_USER_PASSWORD")
	return dc
}

func checkErr(err error) {
	if err != nil {
		tool.Log.Debug("Error is ", err)
		os.Exit(-1)
	}
}


