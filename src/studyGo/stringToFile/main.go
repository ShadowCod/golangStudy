package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

//将单行string写入文件
// func main() {
// 	f, err := os.Create("test.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer f.Close()
// 	n, err := f.WriteString("string to file")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(n)
// }
// 将多行string写入文件
// func main() {
// 	f, err := os.Create("lines")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer f.Close()
// 	str := []string{"Welcome to the world of golang.", "Go is a complied language.", "It is easy to learn Go."}
// 	for _, v := range str {
// 		_, err = fmt.Fprintln(f, v)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 	}
// 	fmt.Println("file written successfully")
// }

// 追加string到文件
// func main() {
// 	f, err := os.OpenFile("lines", os.O_APPEND|os.O_WRONLY, 0666)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer f.Close()
// 	str := "File handling is easy."
// 	_, err = fmt.Fprintln(f, str)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("file appended successfully")
// }
// 并发写入文件
// func produce(data chan int, wg *sync.WaitGroup) {
// 	n := rand.Intn(999)
// 	data <- n
// 	wg.Done()
// }
// func consume(data chan int, done chan bool) {
// 	f, err := os.Create("./concurrent")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer f.Close()
// 	for d := range data {
// 		_, err = fmt.Fprintln(f, d)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 	}
// 	done <- true

// }
// func main() {
// 	data := make(chan int)
// 	done := make(chan bool)
// 	wg := sync.WaitGroup{}
// 	for i := 0; i <= 100; i++ {
// 		wg.Add(1)
// 		go produce(data, &wg)
// 	}
// 	go consume(data, done)

// 	go func() {
// 		wg.Wait()
// 		close(data)
// 	}()
// 	d := <-done
// 	if d == true {
// 		fmt.Println("File writring successfully")
// 	} else {
// 		fmt.Println("File writing failed")
// 	}
// }
// 读取文件内容
// func main() {
// 	b, err := ioutil.ReadFile("./concurrent")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(string(b))
// }
// 使用 命令行标记传递文件路径
// func main() {
// 	f := flag.String("fpath", "test.txt", "file path to read from")
// 	flag.Parse()
// 	fmt.Println(f)
// 	b, err := ioutil.ReadFile(*f)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(string(b))
// }
// 分块读取文件
// func main() {
// 	f, err := os.Open("./test.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer f.Close()
// 	r := bufio.NewReader(f)
// 	b := make([]byte, 3)
// 	for {
// 		_, err = r.Read(b)
// 		if err != nil {
// 			fmt.Println(err)
// 			break
// 		}
// 		fmt.Println(string(b))
// 	}
// }
// 逐行读取文件
// func main() {
// 	f := flag.String("fpath", "concurrent", "file")
// 	file, err := os.Open(*f)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer file.Close()
// 	s := bufio.NewScanner(file)
// 	for s.Scan() {
// 		fmt.Println(s.Text())
// 	}
// 	err = s.Err()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

func main() {
	f := flag.String("fpath", "lines", "test")
	file, err := os.Open(*f)
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(file)
	line, err := reader.ReadBytes('.')
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(line))
}
