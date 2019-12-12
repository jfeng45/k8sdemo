package main

import (
	"github.com/jfeng45/k8sdemo/config"
	"github.com/jfeng45/k8sdemo/tool"
)

func main(){
	//testDataService()
	testRegistration()

}

func testRegistration() {
	dataServiceName :="mysql"

	ruci, err:= config.BuildRegistrationInterface(dataServiceName)
	if err != nil {
		tool.Log.Debug("registration interface build failed:", err)
	}
	//user :=model.User{Name:"jinefng", Department:"ddevelopment", Created:"2018-12-09"}
	//b, err :=ruci.Register(&user)
	//if b {
	//	tool.Log.Debug("user registered:", user)
	//} else {
	//	tool.Log.Debug("user registere failed:", err)
	//}
	users, err :=ruci.ListUser()
	if err != nil {

		tool.Log.Debug("user registere failed:", err)
	} else {
		tool.Log.Debug("user lst skaffold:", users)
	}

}

//func testDataService() {
//	id := 15
//	err := runSql(id)
//	checkErr(err)
//	err = runCouchdb(id)
//	checkErr(err)
//}

//func runCouchdb(id int) error {
//	dataServiceName :="couchdb"
//	userDataService, err := dataservice.BuildUserDataInterface(dataServiceName)
//	defer userDataService.Close()
//	if err != nil {
//		return err
//	}
//
//	err = userDataService.Find(id)
//	checkErr(err)
//	return nil
//}
//func runSql(id int) error{
//	dataServiceName :="mysql"
//	userDataService, err := dataservice.BuildUserDataInterface(dataServiceName)
//	if err != nil {
//		return err
//	}
//	//userDataService, err := buildDataStore()
//	defer userDataService.Close()
//	user :=model.User{Name:"jinefng", Department:"ddevelopment", Created:"2018-12-09"}
//	userDataService.Insert(&user)
//	user1 :=model.User{Id: user.Id, Name:"yeyang", Department:"sales", Created:"2019-1-21"}
//	err, affect := userDataService.Update(&user1)
//
//	err = userDataService.Find(id)
//	_, err = userDataService.FindAll()
//
//	userDataService.Remove(err, int64(id), affect)
//	return nil
//}
//
//func checkErr(err error) {
//	if err != nil {
//		tool.Log.Debug("Error is ", err)
//		os.Exit(-1)
//	}
//}