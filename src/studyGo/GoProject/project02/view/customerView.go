package main

import (
	"fmt"
	"studyGo/GoProject/project02/service"
)

type customerView struct {
	key             string
	loop            bool
	name            string
	gender          string
	age             int
	phone           string
	email           string
	customerService *service.CustomerService
}

func (c *customerView) menu() {
	for {
		fmt.Println("\n-----------客户信息管理----------")
		fmt.Println("             1 添加客户")
		fmt.Println("             2 修改客户")
		fmt.Println("             3 删除客户")
		fmt.Println("             4 客户列表")
		fmt.Println("             5 退出界面")
		fmt.Println()
		fmt.Print("请选择1-5:")
		fmt.Scanln(&c.key)
		switch c.key {
		case "1":
			c.add()
		case "2":
			c.updata()
		case "3":
			c.delete()
		case "4":
			c.show()
		case "5":
			c.loop = true
		default:
			fmt.Println("输入不正确...")
		}
		if c.loop {
			fmt.Println("即将退出")
			break
		}
	}
}

func (c *customerView) show() {
	if c.customerService == nil {
		fmt.Println("请先添加客户")
	} else {
		customerSlice := c.customerService.Show()
		fmt.Println("---------------客户列表-------------------")
		fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
		for _, v := range customerSlice {
			fmt.Println(v.GetInfo())
		}
		fmt.Println("------------------------------------------")
	}
}

func (c *customerView) add() {
	fmt.Print("请输入名称:")
	fmt.Scanln(&c.name)
	fmt.Print("请输入性别:")
	fmt.Scanln(&c.gender)
	fmt.Print("请输入年龄:")
	fmt.Scanln(&c.age)
	fmt.Print("请输入电话:")
	fmt.Scanln(&c.phone)
	fmt.Print("请输入邮箱:")
	fmt.Scanln(&c.email)
	c.customerService.Add(c.name, c.gender, c.age, c.phone, c.email)
}

func (c *customerView) delete() {
	fmt.Print("请选择需要删除的客户编号(-1退出):")
	index := -1
	fmt.Scanln(&index)
	if index != -1 {
		findIndex := c.customerService.FindIndex(index)
		if findIndex == -1 {
			fmt.Println("客户不存在")
		} else {
			fmt.Print("确认删除(y/n):")
			flag := "n"
			fmt.Scanln(&flag)
			if flag == "y" {
				c.customerService.Delete(findIndex)
			}
		}
	}
}

func (c *customerView) updata() {
	fmt.Print("请选择需要修改的客户编号(-1退出):")
	index := -1
	fmt.Scanln(&index)
	if index == -1 {
		return
	} else {
		findIndex := c.customerService.FindIndex(index)
		if findIndex == -1 {
			fmt.Println("客户不存在")
		} else {
			fmt.Print("请输入名称:")
			fmt.Scanln(&c.name)
			fmt.Print("请输入性别:")
			fmt.Scanln(&c.gender)
			fmt.Print("请输入年龄:")
			fmt.Scanln(&c.age)
			fmt.Print("请输入电话:")
			fmt.Scanln(&c.phone)
			fmt.Print("请输入邮箱:")
			fmt.Scanln(&c.email)
			c.customerService.UpData(findIndex, c.name, c.gender, c.age, c.phone, c.email)
		}
	}
}

func main() {
	v := customerView{
		key:  "",
		loop: false,
	}
	v.customerService = service.CreateCustomerServict()
	v.menu()
}
