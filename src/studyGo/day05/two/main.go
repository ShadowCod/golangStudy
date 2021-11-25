package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//map
func main() {
	// var m1 map[string]int
	// fmt.Println(m1 == nil)        //没有初始化，未开辟内存空间
	// m2 := make(map[string]int, 3) //要估算好map的容量，避免在程序运行时再动听扩容
	// m2["北京"] = 18
	// m2["上海"] = 20
	// fmt.Println(m2)

	//取map的值
	// fmt.Println(m2["北京"])
	// fmt.Println(m2["广州"]) //如果不存在拿到对应值类型的零值
	//约定成俗用ok接收返回的布尔值
	// value, ok := m2["广州"]
	// if !ok {
	// 	fmt.Println("查无此key")
	// } else {
	// 	fmt.Println(value)
	// }

	//map的遍历
	// for k, v := range m2 {
	// 	fmt.Println(k, v)
	// }
	//只遍历key
	// for k := range m2 {
	// 	fmt.Println(k)
	// }
	//只遍历value
	// for _, v := range m2 {
	// 	fmt.Println(v)
	// }

	//删除map中的元素
	// delete(m2, "北京")
	// fmt.Println(m2)
	// delete(m2, "深圳")
	// fmt.Println(m2)

	rand.Seed(time.Now().UnixNano()) //初始化随机种子

	var scoreMap = make(map[string]int, 200)
	for i := 1; i <= 200; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}

	// fmt.Println(scoreMap)
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	// fmt.Println(keys)

	sort.Strings(keys)
	// fmt.Println(keys)
	for _, v := range keys {
		fmt.Println(v, scoreMap[v])
	}
}
