package main

import "fmt"

func Action(state int, index int) (actHow rune, actNum int, ch rune, sub int) {
	var j int
	switch str[index] {
	case 'i':
		j = 0
	case '+':
		j = 1
	case '*':
		j = 2
	case '(':
		j = 3
	case ')':
		j = 4
	case '#':
		j = 5
	default:
		j = -1
	}
	if j != -1 {
		actHow = sym[state][j]
		actNum = snum[state][j]
		if actHow == 'r' {
			switch actNum {
			case 1:
				ch = 'E'
				sub = 3
				fmt.Println("按照E->E+T规约")
			case 2:
				ch = 'E'
				sub = 1
				fmt.Println("按E->T规约")
			case 3:
				ch = 'T'
				sub = 3
				fmt.Println("按T->T*F规约")
			case 4:
				ch = 'T'
				sub = 1
				fmt.Println("按T->F规约")
			case 5:
				ch = 'F'
				sub = 3
				fmt.Println("按F->(E)规约")
			case 6:
				ch = 'F'
				sub = 1
				fmt.Println("按F->id规约")
			}
		}
	}
	return
}

//goto[t,A]
func Go2(t int,ch rune) (r int) {
	switch ch {
	case 'E':
		r = go2[t][0]
	case 'T':
		r = go2[t][1]
	case 'F':
		r = go2[t][2]
	}
	return
}