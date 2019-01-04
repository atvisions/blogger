package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() error {
	var err error
	DB,err = sqlx.Open("mysql","root:123456@tcp(127.0.0.1:3306)/blogger?charset=utf8&parseTime=true")
	//检查是否连接成功
	err = DB.Ping()
	if err != nil{
		fmt.Println(err)
	}
	return nil
}