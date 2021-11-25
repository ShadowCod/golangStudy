package main

import (
	//"crypto/rand" //真随机
	"fmt"
	"math/rand" //伪随机
	"time"
)

func main() {
	Pseudorandom()
}

/*
	伪随机实现
*/
func Pseudorandom() {
	// 生成随机种子
	rand.Seed(time.Now().Unix())
	// 生成20个0-100伪随机数
	for i := 1; i <= 20; i++ {
		result := rand.Intn(100)
		fmt.Println(result)
	}
}

/*
	Reallyrandom 真随机实现
*/
// func Reallyrandom() {
// 	// 生成20个0-100的真随机数
// 	for i := 1; i <= 20; i++ {
// 		result, _ := rand.Int(rand.Reader, big.NewInt(100))
// 		fmt.Println(result)
// 	}
// }
