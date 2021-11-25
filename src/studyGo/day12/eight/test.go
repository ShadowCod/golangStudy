package eight

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// monster结构体
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

//序列化方法
func (m *Monster) Store() {
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("序列化失败,", err)
		return
	}
	filePath := "f:/testFun.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(string(data))
	if err != nil {
		fmt.Println("写入失败")
	}
	writer.Flush()
}

// 反序列化
func (m *Monster) Restore() {
	filePath := "f:/testFun.txt"
	str, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read file err", err)
		return
	}
	// fmt.Println(string(str))
	err = json.Unmarshal(str, m)
	if err != nil {
		fmt.Println("unmarshal err", err)
		return
	}
	fmt.Println(m)
}
