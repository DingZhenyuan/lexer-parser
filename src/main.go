package main

import (
	"fmt"
)

/*栈区*/
//符号栈
var symbolStack StackRune
//状态栈
var stateStack StackInt

var str string

func main() {
	//parser
	for {
		var actHow rune
		var ch rune

		//输入字符串
		fmt.Println("***************************")
		fmt.Println("请输入待测试的字符串：")
		fmt.Scanln(&str)
		str = str + "#"

		//状态0入栈
		stateStack.Push(0)
		for index := 0; index < len(str); {
			sub := 0; actNum := 0; actHow = 0; ch = 0
			state, _:= stateStack.Top()
			//查询action表
			actHow, actNum, ch, sub = Action(state, index)

			if actHow == 's' {
				//移进
				fmt.Println("移进")
				symbolStack.Push(rune(str[index]))
				stateStack.Push(actNum)
				index++
			} else if actHow == 'r' {
				//规约
				for i := 0; i < sub; i++ {
					if !stateStack.IsEmpty() {
						stateStack.Pop()
					}
					if !symbolStack.IsEmpty() {
						symbolStack.Pop()
					}
				}
				t, _ := stateStack.Top()
				symbolStack.Push(ch)
				stateStack.Push(Go2(t, ch))
			} else if actHow == 'a' {
				fmt.Println("接收成功")
				break
			} else {
				fmt.Println("接收错误")
				index = Error(state, actNum, index)
				break
			}
		}
	}
}
