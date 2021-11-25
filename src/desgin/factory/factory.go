package main

import "fmt"

//工厂模式核心概念：创建者和调用者分离

/**
工厂方法模式：
	1.定义一个工厂接口
	2.定义不同的工厂类
	3.不同工厂类去实现该接口的方法
	4.使用一个总控方法实现获取某个工厂
	5.使用获取的工厂获得产品
*/
//CarFactory 工厂接口
type CarFactory interface {
	CreateCar() Car
}

//Car 产品接口
type Car interface {
	Start()
	Close()
	Run()
}

//TeslaFactory 不同的工厂结构体
type TeslaFactory struct {
}

//TeslaCar 具体产品结构体
type TeslaCar struct {
}

//Start 实现接口方法
func (t *TeslaCar) Start() {
	fmt.Println("tesla start")
}

//Close 实现接口方法
func (t *TeslaCar) Close() {
	fmt.Println("tesla close")
}

//Run 实现接口方法
func (t *TeslaCar) Run() {
	fmt.Println("tesla run")
}

//CreateCar 不同的工厂结构体实现接口方法
func (t *TeslaFactory) CreateCar() Car {
	return &TeslaCar{}
}

//DaZhongFactory 不同的工厂结构体
type DaZhongFactory struct {
}

// DaZhongCar 具体产品结构体
type DaZhongCar struct {
}

//Start 实现接口方法
func (d *DaZhongCar) Start() {
	fmt.Println("dazhong start")
}

//Close 实现接口方法
func (d *DaZhongCar) Close() {
	fmt.Println("dazhong close")
}

//Run 实现接口方法
func (d *DaZhongCar) Run() {
	fmt.Println("dazhong run")
}

//CreateCar 不同的工厂结构体实现接口方法
func (d *DaZhongFactory) CreateCar() Car {
	return &DaZhongCar{}
}

//GetFactory 总控方法实现获取某个工厂
func GetFactory(name string) (f CarFactory) {
	switch name {
	case "tesla":
		f = &TeslaFactory{}
		return
	case "dazhong":
		f = &DaZhongFactory{}
		return
	}
	return nil
}

//工厂方法模式使用
// func main() {
// 	car1 := GetFactory("dazhong").CreateCar()
// 	car1.Start()
// 	car1.Close()
// 	car1.Run()
// 	car2 := GetFactory("tesla").CreateCar()
// 	car2.Start()
// 	car2.Close()
// 	car2.Run()
// }

/**
简单工厂模式|静态工厂模式：
	1.声明一个产品接口
	2.不同的产品实现对应的接口方法
	3.使用一个总控方法实现获取某个具体的产品
缺点：
	每增加一种类型的产品都需要修改总控方法
*/
//GetACar 总控方法实现获取某个具体的产品
func GetACar(name string) Car {
	switch name {
	case "dazhong":
		return &DaZhongCar{}
	case "teals":
		return &TeslaCar{}
	}
	return nil
}

//简单工厂方法使用
func main() {
	car1 := GetACar("dazhong")
	car1.Start()
	car2 := GetACar("teals")
	car2.Start()
}

/**
抽象工厂模式：
	也就是比工厂方式多了一种可以生成的产品
	比如：工厂方法模式中，小米工厂只生成手机
	而在抽象工厂模式中，小米可以生产手机也可以生产路由器
*/
