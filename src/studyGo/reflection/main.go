package main

import (
	"fmt"
	"reflect"
)

type order struct {
	OrderID    int64
	CustomerID int64
}
type student struct {
	ID   int64
	Name string
	Age  byte
	Meg  string
}

func createQuery(q interface{}) string {
	// fmt.Println(reflect.ValueOf(q).Kind())
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		str := fmt.Sprintf("insert into %s values(", t)
		fmt.Println(str)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			fmt.Println(v.Field(i))
			// fmt.Println(v.Field(i).Int())
		}

	}
	return ""
}

func main() {
	order := order{
		OrderID:    1,
		CustomerID: 2,
	}
	createQuery(order)
}
