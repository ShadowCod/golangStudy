package service

import "studyGo/GoProject/project02/model"

type CustomerService struct {
	CustomerNum   int
	CustomerSlice []model.Customer
}

//创建
func CreateCustomerServict() *CustomerService {
	return &CustomerService{}
}

// 返回slice切片
func (c *CustomerService) Show() []model.Customer {
	return c.CustomerSlice
}

// 添加客户
func (c *CustomerService) Add(name string, gender string, age int, phone string, email string) {
	c.CustomerNum += 1
	customer := model.CreateCustomer(c.CustomerNum, name, gender, age, phone, email)
	c.CustomerSlice = append(c.CustomerSlice, customer)
}

// 根据编号查找index索引
func (c *CustomerService) FindIndex(num int) int {
	var index = -1
	for i, v := range c.CustomerSlice {
		if v.Id == num {
			index = i
			break
		}
	}
	return index
}

// 根据index索引删除slice中的元素
func (c *CustomerService) Delete(index int) {
	c.CustomerSlice = append(c.CustomerSlice[:index], c.CustomerSlice[index+1:]...)
}

// 根据编号修改元素中的值
func (c *CustomerService) UpData(index int, name string, gender string, age int, phone string, email string) {
	c.CustomerSlice[index].Name = name
	c.CustomerSlice[index].Gender = gender
	c.CustomerSlice[index].Age = age
	c.CustomerSlice[index].Phone = phone
	c.CustomerSlice[index].Email = email
}
