package models

import (
	"errors"

	"github.com/kasiss-liu/go-webserver/dbserver"

	_ "github.com/go-sql-driver/mysql"
)

const (
	tableName = "user"
	connName  = "mysqllocal"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Birthday string `json:"birthday"`
}

type userModel struct{}

//模型实例
//用于实际调用查询或者写入
var UserModel userModel

//获取单个用户
func (u *userModel) GetUserById(id int) (User, error) {
	newUser := User{}
	db, err := dbserver.GetMysql(connName)

	if err != nil {
		return newUser, err
	}
	err = db.QueryRow("select * from "+tableName+" where id = ?", id).Scan(&newUser.Id, &newUser.Name, &newUser.Age, &newUser.Birthday)
	if err != nil {
		return newUser, errors.New("user not found")
	}
	return newUser, nil
}

//获取所有用户
func (u *userModel) GetUserAll() (users []User) {
	users = make([]User, 0, 10)
	db, err := dbserver.GetMysql(connName)
	rows, err := db.Query("select * from " + tableName)
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Name, &user.Age, &user.Birthday)
		users = append(users, user)
	}

	return
}
