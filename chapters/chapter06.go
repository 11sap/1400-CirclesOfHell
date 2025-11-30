package chapters

import (
	"fmt"
)

func Run06() {
	tasks := []func(){
		ch6ex01, ch6ex02, ch6ex03, ch6ex04, ch6ex05, ch6ex06, ch6ex07, ch6ex08,
		ch6ex09, ch6ex10, ch6ex11, ch6ex12, ch6ex13,
	}

	for {
		fmt.Print("\nГлава 6. Выберите задачу (1-83): ")
		var n int
		fmt.Scan(&n)

		if n == 0 {
			break
		}
		if n >= 1 && n <= 83 {
			tasks[n-1]()
		} else {
			fmt.Println("От 1 до 83! Для выхода введите 0.")
		}
	}
}

func ch6ex01() {
	fmt.Printf("Ex01. Деление без стандартных операций\n")
	var a, b, quotient int
	fmt.Scan(&a, &b)

	temp := a
	for temp >= b {
		temp -= b
		quotient++
	}
	fmt.Printf("а) Целая часть от деления %d на %d: %d\n", a, b, quotient)

	remainder := a
	for remainder >= b {
		remainder -= b
	}
	fmt.Printf("б) Остаток от деления %d на %d: %d\n", a, b, remainder)
}

func ch6ex02() {
	fmt.Printf("Ex02. Числа до n (без for с параметром)\n")
	var n int
	fmt.Scan(&n)

	i := 1
	for {
		if i > n {
			break
		}
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()
}

func ch6ex03() {
	fmt.Printf("Ex03. Нечетные 10-100 (без for с параметром)\n")
	i := 10
	for {
		if i > 100 {
			break
		}
		if i%2 != 0 {
			fmt.Printf("%d ", i)
		}
		i++
	}
	fmt.Println()
}

func ch6ex04() {
	fmt.Printf("Ex04. Минимальное >190, кратное 17\n")
	n := 191
	for {
		if n%17 == 0 {
			fmt.Printf("Число: %d\n", n)
			break
		}
		n++
	}
}

func ch6ex05() {
	fmt.Printf("Ex05. Максимальное ≤5000, кратное 139\n")
	n := 5000
	for {
		if n%139 == 0 {
			fmt.Printf("Число: %d\n", n)
			break
		}
		n--
	}
}

func ch6ex06() {
	fmt.Printf("Ex06. Число по факториалу\n")
	var factorial int
	fmt.Scan(&factorial)

	n := 1
	product := 1
	for {
		product *= n
		if product == factorial {
			fmt.Printf("Число: %d\n", n)
			break
		}
		if product > factorial {
			fmt.Printf("Такого факториала не существует\n")
			break
		}
		n++
	}
}

func ch6ex07() {
	fmt.Printf("Ex07. Числа, квадрат которых ≤ n\n")
	var n int
	fmt.Scan(&n)

	i := 1
	for {
		if i*i > n {
			break
		}
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()
}

func ch6ex08() {
	fmt.Printf("Ex08. Первое число, квадрат которого > n\n")
	var n int
	fmt.Scan(&n)

	fmt.Printf("Вариант 1: ")
	i := 1
	for i <= 100 {
		if i*i > n {
			fmt.Printf("%d\n", i)
			break
		}
		i++
	}

	fmt.Printf("Вариант 2: ")
	i = 1
	found := false
	for i <= 100 && !found {
		if i*i > n {
			fmt.Printf("%d\n", i)
			found = true
		}
		i++
	}

	fmt.Printf("Вариант 3: ")
	i = 1

	for i <= 100 {
		if i*i > n {
			fmt.Printf("%d\n", i)
			break
		}
		i++
	}
}

func ch6ex09() {
	fmt.Printf("Ex09. Ввод четного числа\n")
	var num int
	for {
		fmt.Printf("Введите четное число: ")
		fmt.Scan(&num)
		if num%2 == 0 {
			fmt.Printf("Спасибо! Вы ввели: %d\n", num)
			break
		}
		fmt.Printf("Ошибка! Число %d нечетное. Попробуйте еще раз.\n", num)
	}
}

func ch6ex10() {
	fmt.Printf("Ex10. Проверка пароля\n")
	password := 1234
	var input int

	for {
		fmt.Printf("Введите пароль: ")
		fmt.Scan(&input)
		if input == password {
			fmt.Printf("Добро пожаловать! Пароль верный.\n")
			break
		}
		fmt.Printf("Неверный пароль! Попробуйте еще раз.\n")
	}
}

func ch6ex11() {
	fmt.Printf("Ex11. Ввод 10 чисел или до 0\n")
	fmt.Printf("Введите 10 чисел (0 для досрочного завершения):\n")

	for i := 1; i <= 10; i++ {
		var num int
		fmt.Scan(&num)
		if num == 0 {
			fmt.Printf("Ввод прерван на числе 0\n")
			break
		}
		fmt.Printf("Введено число %d: %d\n", i, num)
	}
}

func ch6ex12() {
	fmt.Printf("Ex12. Ввод до 0\n")
	fmt.Printf("Введите числа (0 для завершения):\n")

	for {
		var num int
		fmt.Scan(&num)
		fmt.Printf("Вы ввели число: %d\n", num)
		if num == 0 {
			break
		}
	}
}

func ch6ex13() {
	fmt.Printf("Ex13. Цикл от 10 до 30\n")

	fmt.Printf("\nВариант с предусловием типа while:\n")
	i := 10
	for i <= 30 {
		fmt.Printf("%d\n", i)
		i++
	}

	fmt.Printf("\nВариант c пост условием типа do-while:\n")
	i = 10
	for {
		fmt.Printf("%d\n", i)
		i++
		if i > 30 {
			break
		}
	}
}
