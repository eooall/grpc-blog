package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB = &sql.DB{}
var err error

func init() {
	//DB, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1)/eooall_blog?charset=utf8&parseTime=True&loc=Local")
	//if err != nil {
	//	panic(err)
	//}
	//DB.DB().SetMaxOpenConns(50)
	//DB.DB().SetConnMaxLifetime(10)
	DB, err = sql.Open("mysql", "root:root@tcp(127.0.0.1)/eooall_blog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	DB.SetMaxOpenConns(50)
	DB.SetConnMaxLifetime(20)
}
