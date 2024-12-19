package main

// isBracketsRight() Проверяет правильность расстановки скобок
func isBracketsRight(str string) bool {
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
