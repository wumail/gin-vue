package storage

import (
	"fmt"
	"os"
)

var path = "d:\\test\\"

//StorageCodeByCategory s
func StorageCodeByCategory(filename string, category string, content string) {
	filename = filename + ".cpp"
	filepath := path + category
	_, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(filepath, 0777) //  Everyone can read write and execute
		}
	} else {
		fmt.Println("create file folder failed, err:", err)
	}

	filepath = filepath + "\\"
	file, err := os.Create(filepath + filename)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	file.Write([]byte(content)) //写入字节切片数据
}

//StorageCode s
// func StorageCode(filename string, title string, content string) {
func StorageCode(content string) {
	// filename = filename + ".cpp"
	filename := "usercode.cpp"
	filepath := path

	// _, err := os.Stat(filepath)
	// if err != nil {
	// 	if os.IsNotExist(err) {
	// 		os.Mkdir(filepath, 0777) //  Everyone can read write and execute
	// 	}
	// } else {
	// 	fmt.Println("create file folder failed, err:", err)
	// }

	filepath = filepath + "\\"
	// file, err := os.Create(filepath + filename)
	file, err := os.Create(filepath + filename)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	file.Write([]byte(content)) //写入字节切片数据
}
