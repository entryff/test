package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sum(x int, y int) int {
	return x + y
}

func subtr(x int, y int) int {
	return x - y
}

func mult(x int, y int) int {
	return x * y
}

func div(x int, y int) int {
	return x / y
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	a, _ := reader.ReadString('\n')
	a = strings.Trim(a, "\r\n")
	var b []string
	var oper string
	var str string
	var res int
	if strings.Contains(a, ".") || strings.Contains(a, ",") {
		panic("Введены нецелые числа")
	}
	for _, k := range a {
		switch {
		case k == '+':
			b = strings.Split(a, "+")
			if (len(b)) > 2 {
				panic("не удовлетворяет заданию — ошибка")
			} else {
				oper = "sum"
			}
		case k == '-':
			b = strings.Split(a, "-")
			if (len(b)) > 2 {
				panic("не удовлетворяет заданию — ошибка")
			} else {
				oper = "subtr"
			}
		case k == '*':
			b = strings.Split(a, "*")
			if (len(b)) > 2 {
				panic("не удовлетворяет заданию — ошибка")
			} else {
				oper = "mult"
			}
		case k == '/':
			b = strings.Split(a, "/")
			if (len(b)) > 2 {
				panic("не удовлетворяет заданию — ошибка")
			} else {
				oper = "div"
			}
		}
	}
	if oper == "" {
		panic("Введеный символ не сответсвует существующей арифметичской операции")
	}
	slovar := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	slovar1 := map[int]string{1: "I", 4: "IV", 5: "V", 9: "IV", 10: "X", 40: "XL", 50: "L", 90: "XC", 100: "C"}
	srez := [9]int{1, 4, 5, 9, 10, 40, 50, 90, 100}
	c1, err1 := strconv.Atoi(b[0])
	c2, err2 := strconv.Atoi(b[1])
	if (err1 != nil && err2 == nil) || (err1 == nil && err2 != nil) {
		panic("не удовлетворяет заданию — ошибка")
	} else if err1 != nil && err2 != nil {
		c1 = slovar[b[0]]
		c2 = slovar[b[1]]
	}
	if c1 < 1 || c1 > 10 || c2 < 1 || c2 > 10 {
		panic("Значения не входят в диапазон допустимого")
	}
	if oper == "sum" {
		res = sum(c1, c2)
	} else if oper == "subtr" {
		res = subtr(c1, c2)
	} else if oper == "mult" {
		res = mult(c1, c2)
	} else if oper == "div" {
		res = div(c1, c2)
	}
	if err1 != nil && err2 != nil && res > 0 {
		N := 8
		for res > 0 {
			for srez[N] > res {
				N = N - 1
			}
			str += slovar1[srez[N]]
			res -= srez[N]
			N = 8
		}
		fmt.Println(str)
	} else if err1 != nil && err2 != nil && res < 1 {
		panic("Результатом работы калькулятора с римскими числами могут быть только положительные числа")
	} else {
		fmt.Println(res)
	}
}
