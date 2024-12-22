package calculation

import (
	"calculator/pkg/stack"
	"fmt"
	"strconv"
	"unicode"
)

// Интерфейс с методами стека
type PushPopper[T comparable] interface {
	Push(n T)
	Pop() T
	Len() int
	GetArray() []T
}

// InfixExprToPostfixString() преобразует выражение инфиксной записи в выражение постфиксной записи и
// заполняет стек математических операций
func InfixExprToPostfixString(infixExpr string, operationsStack PushPopper[string], postfixExpr PushPopper[string]) (PushPopper[string], error) {
	for i := 0; i < len(infixExpr); i++ {
		s := string(infixExpr[i])
		s_rune := []rune(s)[0]
		if unicode.IsDigit(s_rune) {
			number := ""
			for j := i; j < len(infixExpr); j++ {
				s = string(infixExpr[j])
				s_rune := []rune(s)[0]
				if unicode.IsDigit(s_rune) != true {
					postfixExpr.Push(number)
					i = j - 1
					break
				}
				if j == len(infixExpr)-1 {
					number += s
					postfixExpr.Push(number)
					i = j
					break
				}
				number += s
			}
			continue
		}

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
				postfixExpr.Push(element)
				if operationsStack.Len() != 0 {
					for i := operationsStack.Len() - 1; i >= 0; i-- {
						element := operationsStack.Pop()
						if element == "(" || element == "+" || element == "-" {
							postfixExpr.Push(element)
							break
						}
						if element == "*" || element == "/" {
							postfixExpr.Push(element)
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
			if element != "+" && element != "-" && element != "/" && element != "*" {
				return stack.NewStack[string](), fmt.Errorf("500")
			}
			postfixExpr.Push(element)
			operationsStack.Pop()
			continue
		}
	}

	for i := operationsStack.Len() - 1; i >= 0; i-- {
		postfixExpr.Push(operationsStack.GetArray()[i])
	}
	return postfixExpr, nil
}

// Проводит математические операции с полученным на вход выражением в постфиксной записи,
// попутно заполняя и убирая элементы из стека с числами
func StackCalc(postfixExpr PushPopper[string], numbersStack PushPopper[float64]) (float64, error) {
	for i := 0; i < postfixExpr.Len(); i++ {
		element := postfixExpr.GetArray()[i]
		n, err := strconv.Atoi(element)
		if err == nil {
			numbersStack.Push(float64(n))
			continue
		}

		if numbersStack.Len() < 2 {
			return 0, fmt.Errorf("Invalid expression")
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
		return 0, fmt.Errorf("Invalid expression")
	} else {
		return numbersStack.Pop(), nil
	}
}

// IsBracketsRight() Проверяет правильность расстановки скобок
func IsBracketsRight(str string) bool {
	num := 0
	for i := 0; i < len(str); i++ {
		if string(str[i]) == "(" {
			num += 1
		} else if string(str[i]) == ")" {
			num -= 1
		}
		if num < 0 {
			return false
		}
	}
	if num == 0 {
		return true
	} else {
		return false
	}
}

// Calc() вызывает проверочные и вычислительные функции
func Calc(infixExpr string) (float64, error) {
	for i := 0; i < len(infixExpr); i++ {
		if infixExpr[i] == '+' || infixExpr[i] == '-' || infixExpr[i] == '/' || infixExpr[i] == '*' ||
			infixExpr[i] == '(' || infixExpr[i] == ')' {
			continue
		} else {
			_, err := strconv.Atoi(string(infixExpr[i]))
			if err != nil {
				return 0, fmt.Errorf("Invalid expression")
			}
		}
	}

	if !IsBracketsRight(infixExpr) {
		return 0, fmt.Errorf("Invalid expression")
	}

	str, err := InfixExprToPostfixString(infixExpr, stack.NewStack[string](), stack.NewStack[string]())
	if err != nil {
		return 0, err
	}

	return StackCalc(str, stack.NewStack[float64]())
}
