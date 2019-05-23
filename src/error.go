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

//Error处理函数
func Error(state, actNum int, index int) (i int) {
	//fmt.Println("error")
	switch actNum {
	case 1:
		stateStack.Push(5)
		symbolStack.Push('i')
		fmt.Println("缺少运算对象")
	case 2:
		//从输入中删除右括号
		index++
		fmt.Println("不配对的右括号")
	case 3:
		//期望碰到+，但是输入id或者左括号，假设已经输入算符
		stateStack.Push(6)
		symbolStack.Push('+')
		fmt.Println("缺少运算符")
	case 4:
		stateStack.Push(11)
		symbolStack.Push(')')
		fmt.Println("缺少运算符")
	case 5:
		index++
		fmt.Println("*号无效，应该输入+号")
	case 6:
		index++
	}
	return index
}
