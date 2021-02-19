filepath = filepath + "\\"
	file, err := os.Create(filepath + filename)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	file.Write([]byte(content)) //写入字节切片数据