package main

import "errors"

//stack int
type StackInt []int

func (stack StackInt) Len() int {
	return len(stack)
}

func (stack StackInt) IsEmpty() bool  {
	return len(stack) == 0
}

func (stack StackInt) Cap() int {
	return cap(stack)
}

func (stack StackInt) Top() (int, error)  {
	if len(stack) == 0 {
		return 0, errors.New("Out of index, len is 0")
	}
	return stack[len(stack) - 1], nil
}

func (stack *StackInt) Push(value int)  {
	*stack = append(*stack, value)
}

func (stack *StackInt) Pop() (int, error)  {
	theStack := *stack
	if len(theStack) == 0 {
		return 0, errors.New("Out of index, len is 0")
	}
	value := theStack[len(theStack) - 1]
	*stack = theStack[:len(theStack) - 1]
	return value, nil
}


//stack rune
type StackRune []rune

func (stack StackRune) Len() int {
	return len(stack)
}

func (stack StackRune) IsEmpty() bool  {
	return len(stack) == 0
}

func (stack StackRune) Cap() int {
	return cap(stack)
}

func (stack StackRune) Top() (rune, error)  {
	if len(stack) == 0 {
		return ' ', errors.New("Out of index, len is 0")
	}
	return stack[len(stack) - 1], nil
}

func (stack *StackRune) Push(value rune)  {
	*stack = append(*stack, value)
}

func (stack *StackRune) Pop() (rune, error)  {
	theStack := *stack
	if len(theStack) == 0 {
		return ' ', errors.New("Out of index, len is 0")
	}
	value := theStack[len(theStack) - 1]
	*stack = theStack[:len(theStack) - 1]
	return value, nil
}