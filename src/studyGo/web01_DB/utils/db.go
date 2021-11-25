package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err := sql.Open("mysql", "账号:密码@tcp(地址:端口)/库名")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(Db)
}
