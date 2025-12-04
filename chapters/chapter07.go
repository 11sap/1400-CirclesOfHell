package chapters

import "fmt"

func Run07() {
	tasks := []func(){
		ch7ex01, ch7ex02, ch7ex03, ch7ex04, ch7ex05, ch7ex06, ch7ex07, ch7ex08,
		ch7ex09, ch7ex10, ch7ex11, ch7ex12, ch7ex13, ch7ex14, ch7ex15, ch7ex16,
		ch7ex17, ch7ex18,
	}

	for {
		fmt.Print("\nГлава 7. Выберите задачу (1-181): ")
		var n int
		fmt.Scan(&n)

		if n == 0 {
			break
		}
		if n >= 1 && n <= 181 {
			tasks[n-1]()
		} else {
			fmt.Println("От 1 до 181! Для выхода введите 0.")
		}
	}
}

func ch7ex01() {
	fmt.Printf("Ex1. Сумма 10 чисел\n")
	sum := 0
	for i := 1; i <= 10; i++ {
		var num int
		fmt.Scan(&num)
		sum += num
	}
	fmt.Printf("Сумма: %d\n", sum)
}

func ch7ex02() {
	fmt.Printf("Ex2. Сумма n вещественных чисел\n")
	var n int
	fmt.Scan(&n)

	sum := 0.0
	for i := 1; i <= n; i++ {
		var num float64
		fmt.Scan(&num)
		sum += num
	}
	fmt.Printf("Сумма: %.2f\n", sum)
}

func ch7ex03() {
	fmt.Printf("Ex3. Периметр 12-угольника\n")
	sum := 0.0
	for i := 1; i <= 12; i++ {
		var side float64
		fmt.Scan(&side)
		sum += side
	}
	fmt.Printf("Периметр: %.2f\n", sum)
}

func ch7ex04() {
	fmt.Printf("Ex4. Общая масса груза\n")
	sum := 0.0
	fmt.Printf("Введите массы предметов (0 для окончания):\n")
	for {
		var mass float64
		fmt.Scan(&mass)
		if mass == 0 {
			break
		}
		sum += mass
	}
	fmt.Printf("Общая масса: %.2f\n", sum)
}

func ch7ex05() {
	fmt.Printf("Ex5. Общая сумма зарплат\n")
	sum := 0.0
	fmt.Printf("Введите зарплаты сотрудников (0 для окончания):\n")
	for {
		var salary float64
		fmt.Scan(&salary)
		if salary == 0 {
			break
		}
		sum += salary
	}
	fmt.Printf("Общая сумма выплат: %.2f\n", sum)
}

func ch7ex06() {
	fmt.Printf("Ex6. Последовательное соединение сопротивлений\n")
	sum := 0.0
	fmt.Printf("Введите сопротивления (0 для окончания):\n")
	for {
		var r float64
		fmt.Scan(&r)
		if r == 0 {
			break
		}
		sum += r
	}
	fmt.Printf("Общее сопротивление: %.2f\n", sum)
}

func ch7ex07() {
	fmt.Printf("Ex7. Параллельное соединение сопротивлений\n")
	sumReciprocal := 0.0
	fmt.Printf("Введите сопротивления (0 для окончания):\n")
	for {
		var r float64
		fmt.Scan(&r)
		if r == 0 {
			break
		}
		sumReciprocal += 1 / r
	}

	totalR := 1 / sumReciprocal
	fmt.Printf("Общее сопротивление: %.2f\n", totalR)
}

func ch7ex08() {
	fmt.Printf("Ex8. Последовательность до 0\n")
	sum := 0
	count := 0

	fmt.Printf("Введите числа (0 для окончания):\n")
	for {
		var num int
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		sum += num
		count++
	}

	fmt.Printf("а) Сумма: %d\n", sum)
	fmt.Printf("б) Количество: %d\n", count)
}

func ch7ex09() {
	fmt.Printf("Ex9. Сумма оценок двух учеников\n")
	sum1, sum2 := 0, 0

	fmt.Printf("Введите оценки первого ученика (4 предмета):\n")
	for i := 1; i <= 4; i++ {
		var grade int
		fmt.Scan(&grade)
		sum1 += grade
	}

	fmt.Printf("Введите оценки второго ученика (4 предмета):\n")
	for i := 1; i <= 4; i++ {
		var grade int
		fmt.Scan(&grade)
		sum2 += grade
	}

	fmt.Printf("Сумма оценок первого ученика: %d\n", sum1)
	fmt.Printf("Сумма оценок второго ученика: %d\n", sum2)
}

func ch7ex10() {
	fmt.Printf("Ex10. Сумма баллов спортсменов-пятиборцев\n")
	sum1, sum2 := 0.0, 0.0

	fmt.Printf("Введите результаты первого спортсмена (5 видов):\n")
	for i := 1; i <= 5; i++ {
		var score float64
		fmt.Scan(&score)
		sum1 += score
	}

	fmt.Printf("Введите результаты второго спортсмена (5 видов):\n")
	for i := 1; i <= 5; i++ {
		var score float64
		fmt.Scan(&score)
		sum2 += score
	}

	fmt.Printf("Сумма баллов первого спортсмена: %.2f\n", sum1)
	fmt.Printf("Сумма баллов второго спортсмена: %.2f\n", sum2)
}

func ch7ex11() {
	fmt.Printf("Ex11. Произведение 5 чисел\n")
	product := 1
	for i := 1; i <= 5; i++ {
		var num int
		fmt.Scan(&num)
		product *= num
	}
	fmt.Printf("Произведение: %d\n", product)
}

func ch7ex12() {
	fmt.Printf("Ex12. Последовательность до 0 (вывод a1, a2)\n")
	var a1, a2 int
	count := 0

	fmt.Printf("Введите числа (0 для окончания):\n")
	for {
		var num int
		fmt.Scan(&num)
		if num == 0 {
			break
		}

		count++
		if count == 1 {
			a1 = num
		} else if count == 2 {
			a2 = num
		}
	}

	if count >= 2 {
		fmt.Printf("a1 = %d, a2 = %d\n", a1, a2)
	} else {
		fmt.Printf("Недостаточно чисел в последовательности\n")
	}
}

func ch7ex13() {
	fmt.Printf("Ex13. Сумма квадратов 10 чисел\n")
	sum := 0
	for i := 1; i <= 10; i++ {
		var num int
		fmt.Scan(&num)
		sum += num * num
	}
	fmt.Printf("Сумма квадратов: %d\n", sum)
}

func ch7ex14() {
	fmt.Printf("Ex14. Сумма квадратов n вещественных чисел\n")
	var n int
	fmt.Scan(&n)

	sum := 0.0
	for i := 1; i <= n; i++ {
		var num float64
		fmt.Scan(&num)
		sum += num * num
	}
	fmt.Printf("Сумма квадратов: %.2f\n", sum)
}

func ch7ex15() {
	fmt.Printf("Ex15. Сумма 10 вещественных > 100.78?\n")
	sum := 0.0
	for i := 1; i <= 10; i++ {
		var num float64
		fmt.Scan(&num)
		sum += num
	}

	fmt.Printf("Сумма: %.2f\n", sum)
	fmt.Printf("Превышает 100.78: %v\n", sum > 100.78)
}

func ch7ex16() {
	fmt.Printf("Ex16. Сумма n чисел < p?\n")
	var n, p int
	fmt.Scan(&n, &p)

	sum := 0
	for i := 1; i <= n; i++ {
		var num int
		fmt.Scan(&num)
		sum += num
	}

	fmt.Printf("Сумма: %d\n", sum)
	fmt.Printf("Меньше %d: %v\n", p, sum < p)
}

func ch7ex17() {
	fmt.Printf("Ex17. Сумма 5 чисел четна?\n")
	sum := 0
	for i := 1; i <= 5; i++ {
		var num int
		fmt.Scan(&num)
		sum += num
	}

	fmt.Printf("Сумма: %d\n", sum)
	fmt.Printf("Четная: %v\n", sum%2 == 0)
}

func ch7ex18() {
	fmt.Printf("Ex18. Сумма n чисел кратна b?\n")
	var n, b int
	fmt.Scan(&n, &b)

	sum := 0
	for i := 1; i <= n; i++ {
		var num int
		fmt.Scan(&num)
		sum += num
	}

	fmt.Printf("Сумма: %d\n", sum)
	fmt.Printf("Кратна %d: %v\n", b, sum%b == 0)
}
