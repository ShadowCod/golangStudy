package main

import "fmt"

// interface:定义一直方法，不需要实现，interface不能包含任何变量
// 实现接口：是要把接口里的所有方法都实现才算

// 声明|定义一个接口
//type Usb interface {
//	// 声明了两个没有实现的方法
//	Start()
//	Stop()
//}
//
//type Phone struct {
//}
//
//// Phone实现Usb接口方法
//func (p Phone) Start() {
//	fmt.Println("手机开始工作")
//}
//
//func (p Phone) Stop() {
//	fmt.Println("手机结束工作")
//}
//
//type Camera struct {
//}
//
//// Phone实现Usb接口方法
//func (c Camera) Start() {
//	fmt.Println("相机开始工作")
//}
//
//func (c Camera) Stop() {
//	fmt.Println("相机结束工作")
//}
//
//type Computer struct {
//}
//
//func (c Computer) working(u Usb) {
//	u.Start()
//	u.Stop()
//}

func main() {
	fmt.Println("111")
	//computer := Computer{}
	//phone := Phone{}
	//computer.working(phone)
	//
	//camera := Camera{}
	//computer.working(camera)
}
