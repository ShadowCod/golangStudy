package main

import "fmt"

func TypeJudge(items ...interface{}) {
	for index, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("第%v个参数是 bool类型，值是%v\n", index, v)
		case int, int32, int64:
			fmt.Printf("第%v个参数是 int类型，值是%v\n", index, v)
		case float32:
			fmt.Printf("第%v个参数是 float32类型，值是%v\n", index, v)
		case float64:
			fmt.Printf("第%v个参数是 float64类型，值是%v\n", index, v)
		case string:
			fmt.Printf("第%v个参数是 string类型，值是%v\n", index, v)
		default:
			fmt.Printf("第%v个参数是 非基本类型，值是%v\n", index, v)
		}
	}
}
func main() {
	// 类型断言实践
	var n1 float32 = 1.1
	var n2 float64 = 2.1
	var n3 int32 = 1
	var name string = "tom"
	address := "北京"
	TypeJudge(n1, n2, n3, name, address)

}
