package chapters

import (
	"fmt"
	"math"
)

func Run01() {
	tasks := []func(){
		ch1ex01, ch1ex02, ch1ex03, ch1ex04, ch1ex05, ch1ex06, ch1ex07, ch1ex08, ch1ex09,
		ch1ex10, ch1ex11, ch1ex12, ch1ex13, ch1ex14, ch1ex15, ch1ex16, ch1ex17,
	}

	for {
		fmt.Print("\nГлава 1. Выберите задачу (1-17): ")
		var n int
		fmt.Scan(&n)

		if n == 0 {
			break
		}
		if n >= 1 && n <= 17 {
			tasks[n-1]()
		} else {
			fmt.Println("От 1 до 17! Для выхода введите 0.")
		}
	}
}

func ch1ex01() {
	fmt.Printf("Ex01.  %d %d %d\n", 31, 18, 79)
}

func ch1ex02() {
	fmt.Printf("Ex02. %d  %d  %d\n", 47, 52, 150)
}

func ch1ex03() {
	fmt.Printf("Ex03. \t %d\n\t %d\n", 50, 10)
}

func ch1ex04() {
	fmt.Printf("Ex04. \t %d\n\t %d\n\t %d\n", 5, 10, 21)
}

func ch1ex05() {
	fmt.Printf("Ex05. \t %d\n\n\t %d\n", 1, 2)
}

func ch1ex06() {
	fmt.Printf("Ex06. %.3f\n", math.Pi)
}

func ch1ex07() {
	fmt.Printf("Ex07. %.1f\n", math.E)
}

func ch1ex08() {
	var n int
	fmt.Printf("Ex08. Введите число: ")
	fmt.Scan(&n)
	fmt.Printf("Ex08. Вы ввели число: %d\n", n)
}
func ch1ex09() {
	var n int
	fmt.Printf("Ex09. Введите число: ")
	fmt.Scan(&n)
	fmt.Printf("Ex09. %d - вот такое число Вы ввели\n", n)
}
func ch1ex10() {
	var s string
	fmt.Printf("Ex10. Введите своё имя: ")
	fmt.Scan(&s)
	fmt.Printf("Ex10. Ваше имя: %s\n", s)
}
func ch1ex11() {
	var s string
	fmt.Printf("Ex11. Введите название футбольной команды: ")
	fmt.Scan(&s)
	fmt.Printf("Ex11. %s - это чемпион!\n", s)
}
func ch1ex12() {
	var s string
	fmt.Printf("Ex12. Введите ваше имя: ")
	fmt.Scan(&s)
	fmt.Printf("Ex12. Привет, %s!\n", s)
}
func ch1ex13() {
	var n int
	fmt.Printf("Ex13. Введите число: ")
	fmt.Scan(&n)
	fmt.Printf("Ex13. Следующее за числом %d число - %d.\n", n, n+1)
	fmt.Printf("Ex13. Для числа %d предыдущее число - %d.\n", n, n-1)
}
func ch1ex14() {
	var n1, n2, n3 int
	fmt.Printf("Ex14. Введите три числа: ")
	fmt.Scan(&n1, &n2, &n3)
	fmt.Printf("Ex14. %d  %d  %d\n", n1, n2, n3)
}

func ch1ex15() {
	var n1, n2, n3, n4 int
	fmt.Printf("Ex15. Введите четыре числа: ")
	fmt.Scan(&n1, &n2, &n3, &n4)
	fmt.Printf("Ex15. %d %d %d %d\n", n1, n2, n3, n4)
}

func ch1ex16() {
	var t, v, x, y int
	fmt.Printf("Ex16. Введите четыре числа: ")
	fmt.Scan(&t, &v, &x, &y)
	fmt.Printf("Ex16. a) 5 10 б) 100 %d в) %d 25\n", t, x)
	fmt.Printf("Ex16.    7 см   1949 %d    %d %d\n", v, x, y)
}

func ch1ex17() {
	var x, y, a, b int
	fmt.Printf("Ex17. Введите четыре числа: ")
	fmt.Scan(&a, &b, &x, &y)
	fmt.Printf("Ex17. a) 2  кг б) %d 1  в) %d %d\n", a, x, y)
	fmt.Printf("Ex17.    13 17    19 %d    5 %d\n", b, y)
}
