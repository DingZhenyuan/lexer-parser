package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

var signlist = make(map[string] int, 100)

//预处理函数，将文件中的空格，换行等无关字符处理掉
func Pretreatment(fileName string) {
	//打开文件
	fpRead, err1 := os.Open(fileName)
	if !check(err1, "open") {
		return
	}
	fpWrite, err2 := os.Create("test.temp")
	if !check(err2, "open") {
		return
	}

	//关闭文件
	defer fpRead.Close()
	defer fpWrite.Close()


	sign := 0
	r := bufio.NewReader(fpRead)
	//不断读取，进行预处理
	for {
		read, err := r.ReadBytes('\n')
		//fmt.Println(err)
		if err == io.EOF {  //读文件结束
			fmt.Println("结束")
			break
		} else {
			check(err, "read")
		}

		length := len(read)
		for index := 0; index < length - 1; index++ {
			//fmt.Println(string(read[index]))
			if sign == 0 {
				if read[index] == ' ' {
					continue
				}
			}
			if read[index] == '#' { //后面均为注释，直接结束此行分析
				break
			} else if read[index] == ' ' || read[index] == '\t' || read[index] == '\n' || read[index] == '\r'{
				if sign == 1 {
					continue
				} else {
					sign = 1
					_, err := fpWrite.WriteString(" ")
					check(err, "write")
				}
			} else if read[index] == '"' || read[index] == '\'' {
				_, err1 := fpWrite.WriteString(string(read[index]))
				check(err1, "write")
				temp := read[index]
				index++
				for index < length && read[index] != temp {
					_, err := fpWrite.WriteString(string(read[index]))
					check(err, "write")
					index++
				}
				if index >= length {
					break
				}
				_, err2 := fpWrite.WriteString(string(read[index]))
				check(err2, "write")
			} else {
				sign = 3
				_, err2 := fpWrite.WriteString(string(read[index]))
				check(err2, "write")
			}
		}
	}
	//最后加一个回车，方便下次读取
	_, err := fpWrite.WriteString("\n")
	check(err, "write")
}

func save(str string) {
	//判定str是否在keywords中
	if _, ok := keywords[str]; ok {
		if _, ok := signlist[str]; !ok {
			signlist[str] = keywords[str]
		}
	} else {
		saveConst(str)
	}
}

func save_var(str string) {
	if _, ok := signlist[str]; !ok {
		if len(strings.TrimSpace(str)) >= 1 {
			if isSignal(str) == 1 {
				signlist[str] = 301
			} else {
				signlist[str] = 501
			}
		}
	}
}

func saveConst(str string) {
	if _, ok := signlist[str]; !ok {
		signlist[str] = 401
	}
}

func saveError(str string) {
	if _, ok := signlist[str]; !ok {
		signlist[str] = 501
	}
}


func isSignal(s string) int {
	if s[0] == '_' || unicode.Is(unicode.ASCII_Hex_Digit, rune(s[0])) {
		for i := 0; i < len(s); i++ {
			if unicode.Is(unicode.ASCII_Hex_Digit, rune(s[i])) || s[i] == '_' || unicode.IsDigit(rune(s[i])) {

			} else {
				return 0
			}
		}
		return 1
	} else {
		return 0
	}
}

func recognition(fileName string) {
	fpRead, err := os.Open(fileName)
	if !check(err, "open") {
		return
	}
	str := ""
	sign := 0

	r := bufio.NewReader(fpRead)
	read, err := r.ReadBytes('\n')
	//fmt.Println(err)
	if err == io.EOF {  //读文件结束
		fmt.Println("结束")
	} else {
		check(err, "read")
	}
	//对读出的一行进行处理
	for index := 0; index < len(read) - 1; index++ {
		ch := read[index]
		if ch == ' ' {
			if len(strings.TrimSpace(str)) < 1 {
				sign = 0
			} else {
				if sign == 1 || sign == 2 {
					str += string(ch)
				} else {
					save(str)
					str = ""
					sign = 0
				}
			}
		} else if ch == '(' || ch == ')' || ch == '[' || ch == ']' || ch == '{' || ch == '}' || ch == ':' {
			if sign == 1 || sign == 2 {
				str += string(ch)
			} else {
				save(str)
				str = ""
				save(string(ch))
			}
		} else if ch == '<' || ch == '>' || ch == ',' || ch == '+' || ch == '='{
			save(str)
			str = ""
			save(string(ch))
		} else if ch == '\'' {
			str += string(ch)
			if sign == 1 {
				sign = 0
				saveConst(str)
				str = ""
			} else {
				if sign != 2 {
					sign = 1
				}
			}
		} else if ch == '"' {
			str += string(ch)
			if sign == 2 {
				sign = 0
				saveConst(str)
				str = ""
			} else {
				if sign != 1 {
					sign = 2
				}
			}
		} else {
			str += string(ch)
		}
	}
}