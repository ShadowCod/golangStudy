package main
/*
	gorm使用
*/
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type UserInfo struct{
	ID int
	Name string
	Age uint
	Gender string
}

func main() {
//	连接mysql数据库
	db,err:=gorm.Open("mysql","root:123456@tcp(127.0.0.1:3306)/gin_project?charset=utf8mb4&parseTime=True&loc=Local")
	if err!=nil{
		fmt.Printf("open fail,err:%v\n",err)
		return
	}
	//连接使用完毕后需要关闭
	defer db.Close()
//建立迁移	迁移会自动创建对应结构体字段的数据表
	db.AutoMigrate(&UserInfo{})
//	插入一条数据(重复创建会报主键重复)
	db.Create(&UserInfo{1,"tom",18,"男"})
//	查询第一条数据
	u:=&UserInfo{}
	db.First(u)
	fmt.Println("first u:",u)
//	更新一条数据
	db.Model(u).Update("name","tom")
//	删除一条数据
	db.Delete(u)
}
