package main

import "fmt"

//错误处理
func check(err error, op string) (pass bool) {
	pass = true
	switch op {
	case "open":
		if err != nil {
			fmt.Println("打开文件操作失败")
			pass = false
		}
	case "read":
		if err != nil {
			fmt.Println("读取文件操作失败")
			pass = false
		}
	case "write":
		if err != nil {
			fmt.Println("写入文件操作失败")
			pass = false
		}
	}
	return pass

}