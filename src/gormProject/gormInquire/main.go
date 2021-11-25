package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

/*
	练习gorm中mysql的查询语句
*/
//定义一个学生模型
type Stu struct{
	gorm.Model
	No uint
	Name string
	Score uint
}

func main() {
	db,err :=gorm.Open("mysql","root:123456@tcp(1278.0.0.1:3306)/ginProject?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil{
		fmt.Printf("open err:%v",err)
		return
	}
	defer db.Close()
//根据模型创建数据表（默认表名为模版名称后面加个S） 自定义表名
	db.AutoMigrate(&Stu{})
//	插入数据
	db.Create(&Stu{})
}
