package main

import (
	"fmt"
	"strconv"
)

// Интерфейс с методами стека
type PushPopper[T comparable] interface {
	Push(n T)
	Pop() T
	Len() int
	GetArray() []T
}

// infixExprToPostfixString() преобразует выражение инфиксной записи в выражение постфиксной записи и
// заполняет стек математических операций
func infixExprToPostfixString(infixExpr string, operationsStack PushPopper[string]) (string, error) {
	postfixExpr := ""
	for i := 0; i < len(infixExpr); i++ {
		s := string(infixExpr[i])
		if s == "(" {
			operationsStack.Push(s)
			continue
		}
		if s == "-" || s == "+" {
			if operationsStack.Len() == 0 {
				operationsStack.Push(s)
				continue
			}
			element := operationsStack.Pop()
			if element == "*" || element == "/" {
				postfixExpr += element
				if operationsStack.Len() != 0 {
					for i := operationsStack.Len() - 1; i >= 0; i-- {
						element := operationsStack.Pop()
						if element == "(" || element == "+" || element == "-" {
							break
						}
						if element == "*" || element == "/" {
							postfixExpr += element
							continue
						}
					}
				}
			} else {
				operationsStack.Push(element)
			}
			operationsStack.Push(s)
			continue
		}
		if s == "*" || s == "/" {
			if operationsStack.Len() == 0 {
				operationsStack.Push(s)
				continue
			}
			element := operationsStack.Pop()
			if element == "+" || element == "-" {
				operationsStack.Push(element)
				operationsStack.Push(s)
				continue
			}
			operationsStack.Push(element)
			operationsStack.Push(s)
			continue
		}
		if s == ")" {
			element := operationsStack.Pop()
			postfixExpr += element
			operationsStack.Pop()
			continue
		}
		a, err := strconv.Atoi(s)
		if err == nil {
			if a >= 0 && a <= 9 {
				postfixExpr += s
			}
		} else {
			return "", fmt.Errorf("Invalid expression")
		}
	}

	for i := operationsStack.Len() - 1; i >= 0; i-- {
		postfixExpr += operationsStack.GetArray()[i]
	}
	return postfixExpr, nil
}

// Проводит математические операции с полученным на вход выражением в постфиксной записи,
// попутно заполняя и убирая элементы из стека с числами
func stackCalc(postfixExpr string, numbersStack PushPopper[float64]) (float64, error) {
	for i := 0; i < len(postfixExpr); i++ {
		element := string(postfixExpr[i])
		n, err := strconv.Atoi(element)
		if err == nil {
			numbersStack.Push(float64(n))
			continue
		}
		if numbersStack.Len() < 2 {
			return 0, fmt.Errorf("Enter at least 2 numbers")
		}
		if element == "+" {
			n1 := numbersStack.Pop()
			n2 := numbersStack.Pop()
			oper := n2 + n1
			numbersStack.Push(oper)
		}
		if element == "-" {
			n1 := numbersStack.Pop()
			n2 := numbersStack.Pop()
			oper := n2 - n1
			numbersStack.Push(oper)
		}
		if element == "/" {
			n1 := numbersStack.Pop()
			n2 := numbersStack.Pop()
			oper := n2 / n1
			numbersStack.Push(oper)
		}
		if element == "*" {
			n1 := numbersStack.Pop()
			n2 := numbersStack.Pop()
			oper := n2 * n1
			numbersStack.Push(oper)
		}
	}
	if numbersStack.Len() != 1 {
		return 0, fmt.Errorf("Enter at least 1 mathematical operation")
	} else {
		return numbersStack.Pop(), nil
	}
}

// Calc() вызывает проверочные и вычислительные функции
func Calc(infixExpr string) (float64, error) {
	if !isBracketsRight(infixExpr) {
		return 0, fmt.Errorf("Invalid expression")
	}

	str, err := infixExprToPostfixString(infixExpr, NewStack[string]())
	if err != nil {
		return 0, err
	}

	return stackCalc(str, NewStack[float64]())
}

func main() {
	res, err := Calc("2+(1+3)*4")
	if err != nil {
		fmt.Println("ERROR: ", err)
		fmt.Println("ERROR: ", err)
		fmt.Println("ERROR: ", err)
		fmt.Println("ERROR: ", err)
	} else {
		fmt.Println(res)
	}
}
