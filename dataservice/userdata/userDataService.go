package userdata

import (
	"database/sql"
	"github.com/jfeng45/k8sdemo/model"
	"github.com/jfeng45/k8sdemo/tool"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type UserDataMysql struct {
	DB *sql.DB
}

func checkErr(err error) {
	if err != nil {
		tool.Log.Debug("Error is ", err)
		os.Exit(-1)
	}

}

func (userData *UserDataMysql ) Close() error{
	return userData.DB.Close()
}
func (userData *UserDataMysql) Remove(err error, id int64, affect int64) {
	stmt, err := userData.DB.Prepare("delete from userinfo where uid=?")
	checkErr(err)
	res, err := stmt.Exec(id)
	checkErr(err)
	affect, err = res.RowsAffected()
	checkErr(err)
	tool.Log.Debug("remove:row affected", affect)
}

func (userData *UserDataMysql) Find(id int) error {
	rows, err := userData.DB.Query("SELECT * FROM userinfo where uid =?", id)
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		user := model.User{uid, username, department, created}

		tool.Log.Debug("find user:", user)
		//fmt.Print("find: uid=",uid)
		//fmt.Print("username=",username)
		//fmt.Print("department=", department)
		//tool.Log.Debug("created=", created)
	}
	return err
}
func (userData *UserDataMysql) FindByName(name string) (bool, error) {
	rows, err := userData.DB.Query("SELECT * FROM userinfo where username =?", name)
	if err != nil {
		return false, err
	}
	if rows.Next() {
		return true, nil
		//var uid int
		//var username string
		//var department string
		//var created string
		//err = rows.Scan(&uid, &username, &department, &created)
		//user := model.User{uid, username, department, created}
		//checkErr(err)
		//tool.Log.Debug("find user:", user)
		//fmt.Print("find: uid=",uid)
		//fmt.Print("username=",username)
		//fmt.Print("department=", department)
		//tool.Log.Debug("created=", created)
	}

	return false, nil
}

func (userData *UserDataMysql) FindAll() ([]model.User, error) {
	tool.Log.Debug("FindAll()")

	rows, err := userData.DB.Query("SELECT * FROM userinfo ")
	if err != nil {
		return nil, err
	}
	var users = []model.User{}
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		tool.Log.Debug("created=", created)
		user := model.User{uid, username, department, created}
		tool.Log.Debug("find user:", user)

		users = append(users, user)


		//fmt.Print("find: uid=",uid)
		//fmt.Print("username=",username)
		//fmt.Print("department=", department)

	}
	tool.Log.Debug("find user list:", users)
	return users, nil
}

func (userData *UserDataMysql) Update(user *model.User) (error, int64) {
	stmt, err  := userData.DB.Prepare("update userinfo set username=?, department=?, created=? where uid=?")
	defer stmt.Close()
	checkErr(err)
	res, err := stmt.Exec(user.Name, user.Department, user.Created, user.Id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	tool.Log.Debug("update: rows affected: ", affect)
	return err, affect
}

func (userData *UserDataMysql) Insert(user *model.User) int64 {
	stmt, err := userData.DB.Prepare("INSERT userinfo SET username=?,department=?,created=?")
	defer stmt.Close()
	checkErr(err)
	res, err := stmt.Exec(user.Name, user.Department, user.Created)
	checkErr(err)
	id, err := res.LastInsertId()
	user.Id = int(id)
	checkErr(err)
	tool.Log.Debug("user inserted:", user)
	return id
}

