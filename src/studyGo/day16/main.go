package main

import "fmt"

func main() {
	// buf := new(bytes.Buffer)
	// fmt.Println(buf)
	// tw := tar.NewWriter(buf)
	// fmt.Println(tw)
	// var a, b []byte
	// if bytes.Compare(a, b) < 0 {
	// 	fmt.Println("<")
	// }
	// if bytes.Compare(a, b) == 0 {
	// 	fmt.Println("=")
	// }
	// var a interface{}
	// b := 1
	// a = b

	// fmt.Println(a.(int))
	// var arr = [6]int{1, 2, 0, 3, 0, 8}
	// i, j := 0, 0
	// for ; j < len(arr); j++ {
	// 	if arr[j] != 0 {
	// 		arr[i], arr[j] = arr[j], arr[i]
	// 		i++
	// 	}
	// }
	// fmt.Println(arr)
	arr := [6]int{1, 2, 0, 3, 0, 8}
	i, j := 0, len(arr)-1
	for {
		if i < j {
			if arr[i] != 0 {
				i++
			} else if arr[j] != 0 {
				arr[i], arr[j] = arr[j], arr[i]
				j--
			} else {
				j--
			}
		} else {
			break
		}
	}
	fmt.Println(arr)
}
