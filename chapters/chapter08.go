package chapters

import (
	"fmt"
	"math"
)

func Run08() {
	tasks := []func(){
		ch8ex01, ch8ex02, ch8ex03, ch8ex04, ch8ex05, ch8ex06, ch8ex07, ch8ex08,
		ch8ex09, ch8ex10, ch8ex11, ch8ex12, ch8ex13, ch8ex14, ch8ex15, ch8ex16,
		ch8ex17, ch8ex18, ch8ex19,
	}

	for {
		fmt.Print("\nГлава 8. Выберите задачу (1-19): ")
		var n int
		fmt.Scan(&n)

		if n == 0 {
			break
		}
		if n >= 1 && n <= 19 {
			tasks[n-1]()
		} else {
			fmt.Println("От 1 до 19! Для выхода введите 0.")
		}
	}
}

func ch8ex01() {
	fmt.Printf("Ex1. Квадраты чисел, не превышающие n\n")

	var n int
	fmt.Print("Введите число n: ")

	if _, err := fmt.Scan(&n); err != nil || n < 1 {
		fmt.Printf("Ошибка: требуется натуральное число\n")
		return
	}

	fmt.Printf("Квадраты чисел, не превышающие %d:\n", n)

	i := 1
	square := 1

	for square <= n {
		fmt.Printf("%d ", square)
		i++
		square = i * i
	}

	fmt.Printf("\n")
}

func ch8ex02() {
	fmt.Printf("Ex2. Первый квадрат, больший n\n")

	var n int
	fmt.Print("Введите число n: ")

	if _, err := fmt.Scan(&n); err != nil || n < 0 {
		fmt.Printf("Ошибка: требуется неотрицательное число\n")
		return
	}

	fmt.Printf("Способ 1 (с циклом while):\n")
	i := 1
	for {
		square := i * i
		if square > n {
			fmt.Printf("Первое число, большее %d: %d\n", n, square)
			break
		}
		i++
	}

	fmt.Printf("Способ 2 (без цикла while):\n")
	i = 1
	square := i * i
	for square <= n {
		i++
		square = i * i
	}
	fmt.Printf("Первое число, большее %d: %d\n", n, square)
}

func ch8ex03() {
	fmt.Printf("Ex3. Дроби i/(i+1), не меньшие a\n")

	var a float64
	fmt.Print("Введите число a (0 < a < 1): ")

	if _, err := fmt.Scan(&a); err != nil || a <= 0 || a >= 1 {
		fmt.Printf("Ошибка: число должно быть в интервале (0, 1)\n")
		return
	}

	fmt.Printf("Дроби i/(i+1), не меньшие %.3f:\n", a)

	fmt.Printf("Вариант 1:\n")
	i := 1
	fraction := float64(i) / float64(i+1)

	for fraction >= a {
		fmt.Printf("%d/%d = %.3f\n", i, i+1, fraction)
		i++
		fraction = float64(i) / float64(i+1)
	}

	fmt.Printf("\nВариант 2:\n")
	i = 1
	for {
		fraction = float64(i) / float64(i+1)
		if fraction < a {
			break
		}
		fmt.Printf("%d/%d = %.3f\n", i, i+1, fraction)
		i++
	}
}

func ch8ex04() {
	fmt.Printf("Ex4. Первая дробь i/(i+1), меньшая a\n")

	var a float64
	fmt.Print("Введите число a (0 < a < 1): ")

	if _, err := fmt.Scan(&a); err != nil || a <= 0 || a >= 1 {
		fmt.Printf("Ошибка: число должно быть в интервале (0, 1)\n")
		return
	}

	i := 1
	fraction := float64(i) / float64(i+1)

	for fraction >= a {
		i++
		fraction = float64(i) / float64(i+1)
	}

	fmt.Printf("Первая дробь, меньшая %.3f: %d/%d = %.3f\n",
		a, i, i+1, fraction)
}

func ch8ex05() {
	fmt.Printf("Ex5. Числа 1+1/i, не меньшие a\n")

	var a float64
	fmt.Print("Введите число a (1 < a < 1.5): ")

	if _, err := fmt.Scan(&a); err != nil || a <= 1 || a >= 1.5 {
		fmt.Printf("Ошибка: число должно быть в интервале (1, 1.5)\n")
		return
	}

	fmt.Printf("Числа 1+1/i, не меньшие %.3f:\n", a)

	fmt.Printf("Вариант 1:\n")
	i := 2
	value := 1.0 + 1.0/float64(i)

	for value >= a {
		fmt.Printf("1 + 1/%d = %.3f\n", i, value)
		i++
		value = 1.0 + 1.0/float64(i)
	}

	fmt.Printf("\nВариант 2:\n")
	i = 2
	for {
		value = 1.0 + 1.0/float64(i)
		if value < a {
			break
		}
		fmt.Printf("1 + 1/%d = %.3f\n", i, value)
		i++
	}
}

func ch8ex06() {
	fmt.Printf("Ex6. Числа 1+1/i, не большие a\n")

	var a float64
	fmt.Print("Введите число a (1 < a < 1.5): ")

	if _, err := fmt.Scan(&a); err != nil || a <= 1 || a >= 1.5 {
		fmt.Printf("Ошибка: число должно быть в интервале (1, 1.5)\n")
		return
	}

	fmt.Printf("Числа 1+1/i, не большие %.3f:\n", a)

	i := 2
	for {
		value := 1.0 + 1.0/float64(i)
		if value <= a {
			fmt.Printf("1 + 1/%d = %.3f\n", i, value)
		} else {
			break
		}
		i++
	}
}

func ch8ex07() {
	fmt.Printf("Ex7. Первое число 1+1/i, меньшее a\n")

	var a float64
	fmt.Print("Введите число a (1 < a < 1.5): ")

	if _, err := fmt.Scan(&a); err != nil || a <= 1 || a >= 1.5 {
		fmt.Printf("Ошибка: число должно быть в интервале (1, 1.5)\n")
		return
	}

	i := 2
	for {
		value := 1.0 + 1.0/float64(i)
		if value < a {
			fmt.Printf("Первое число, меньшее %.3f: 1 + 1/%d = %.3f\n",
				a, i, value)
			break
		}
		i++
	}
}

func ch8ex08() {
	fmt.Printf("Ex8. Все n, при которых 1+1/i ≥ a\n")

	var a float64
	fmt.Print("Введите число a (1 < a < 1.5): ")

	if _, err := fmt.Scan(&a); err != nil || a <= 1 || a >= 1.5 {
		fmt.Printf("Ошибка: число должно быть в интервале (1, 1.5)\n")
		return
	}

	fmt.Printf("Все значения n, при которых 1+1/n ≥ %.3f:\n", a)

	i := 2
	count := 0
	for {
		value := 1.0 + 1.0/float64(i)
		if value < a {
			break
		}
		fmt.Printf("n = %d (1 + 1/%d = %.3f)\n", i, i, value)
		i++
		count++
	}

	fmt.Printf("Всего: %d значений\n", count)
}

func ch8ex09() {
	fmt.Printf("Ex9. Наименьшее n, при котором 1+1/n < a\n")

	var a float64
	fmt.Print("Введите число a (1 < a < 1.5): ")

	if _, err := fmt.Scan(&a); err != nil || a <= 1 || a >= 1.5 {
		fmt.Printf("Ошибка: число должно быть в интервале (1, 1.5)\n")
		return
	}

	n := 2
	for {
		value := 1.0 + 1.0/float64(n)
		if value < a {
			fmt.Printf("Наименьшее n: %d (1 + 1/%d = %.3f < %.3f)\n",
				n, n, value, a)
			break
		}
		n++
	}
}

func ch8ex10() {
	fmt.Printf("Ex10. Числа 1+1/i, меньшие a\n")

	var a float64
	fmt.Print("Введите вещественное число a: ")

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Printf("Ошибка ввода\n")
		return
	}

	fmt.Printf("Числа 1+1/i, меньшие %.3f:\n", a)

	i := 1
	for {
		value := 1.0 + 1.0/float64(i)
		if value >= a {
			break
		}
		fmt.Printf("1 + 1/%d = %.3f\n", i, value)
		i++
	}
}

func ch8ex11() {
	fmt.Printf("Ex11. Первое число 1+1/i, большее a\n")

	var a float64
	fmt.Print("Введите вещественное число a: ")

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Printf("Ошибка ввода\n")
		return
	}

	i := 1
	for {
		value := 1.0 + 1.0/float64(i)
		if value > a {
			fmt.Printf("Первое число, большее %.3f: 1 + 1/%d = %.3f\n",
				a, i, value)
			break
		}
		i++
	}
}

func ch8ex12() {
	fmt.Printf("Ex12. Все n, при которых 1+1/n < a\n")

	var a float64
	fmt.Print("Введите вещественное число a: ")

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Printf("Ошибка ввода\n")
		return
	}

	fmt.Printf("Все значения n, при которых 1+1/n < %.3f:\n", a)

	i := 1
	count := 0
	for {
		value := 1.0 + 1.0/float64(i)
		if value >= a {
			break
		}
		fmt.Printf("n = %d (1 + 1/%d = %.3f)\n", i, i, value)
		i++
		count++
	}

	if count == 0 {
		fmt.Printf("Нет таких значений n\n")
	} else {
		fmt.Printf("Всего: %d значений\n", count)
	}
}

func ch8ex13() {
	fmt.Printf("Ex13. Наименьшее n, при котором 1+1/n < a\n")

	var a float64
	fmt.Print("Введите вещественное число a: ")

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Printf("Ошибка ввода\n")
		return
	}

	if a > 2.0 {
		fmt.Printf("Наименьшее n: 1 (1 + 1/1 = 2.000 < %.3f)\n", a)
		return
	}

	n := 1
	for {
		value := 1.0 + 1.0/float64(n)
		if value < a {
			fmt.Printf("Наименьшее n: %d (1 + 1/%d = %.3f < %.3f)\n",
				n, n, value, a)
			break
		}
		n++

		if n > 1000000 {
			fmt.Printf("Не удалось найти подходящее n\n")
			break
		}
	}
}

func ch8ex14() {
	fmt.Printf("Ex14. Первая дробь i/(i+1), меньшая a\n")

	var a float64
	fmt.Print("Введите число a (0 < a < 1): ")

	if _, err := fmt.Scan(&a); err != nil || a <= 0 || a >= 1 {
		fmt.Printf("Ошибка: число должно быть в интервале (0, 1)\n")
		return
	}

	fmt.Printf("Вариант 1:\n")
	i := 1
	for float64(i)/float64(i+1) >= a {
		i++
	}
	fmt.Printf("Первая дробь, меньшая %.3f: %d/%d = %.3f\n",
		a, i, i+1, float64(i)/float64(i+1))

	fmt.Printf("\nВариант 2:\n")

	minI := int(math.Ceil(a/(1.0-a))) + 1
	fmt.Printf("Первая дробь, меньшая %.3f: %d/%d = %.3f\n",
		a, minI, minI+1, float64(minI)/float64(minI+1))
}

func ch8ex15() {
	fmt.Printf("Ex15. Числа (x-100)/(x-200) при x=1..100, меньшие m\n")

	var m float64
	fmt.Print("Введите число m (0.52 < m < 53.7): ")

	if _, err := fmt.Scan(&m); err != nil || m <= 0.52 || m >= 53.7 {
		fmt.Printf("Ошибка: число должно быть в интервале (0.52, 53.7)\n")
		return
	}

	fmt.Printf("Числа (x-100)/(x-200) < %.3f:\n", m)

	fmt.Printf("Вариант 1:\n")
	count := 0
	for x := 1; x <= 100; x++ {
		value := float64(x-100) / float64(x-200)
		if value < m {
			fmt.Printf("x=%d: (x-100)/(x-200) = %.3f\n", x, value)
			count++
		}
	}
	fmt.Printf("Всего: %d значений\n", count)

	fmt.Printf("\nВариант 2:\n")
	count = 0
	x := 1
	for x <= 100 {
		value := float64(x-100) / float64(x-200)
		if value < m {
			fmt.Printf("x=%d: %.3f\n", x, value)
			count++
		}
		x++
	}
	fmt.Printf("Всего: %d значений\n", count)
}

func ch8ex16() {
	fmt.Printf("Ex16. Дроби 1/i, меньшие m\n")

	var m float64
	fmt.Print("Введите число m (0.5 < m < 1): ")

	if _, err := fmt.Scan(&m); err != nil || m <= 0.5 || m >= 1 {
		fmt.Printf("Ошибка: число должно быть в интервале (0.5, 1)\n")
		return
	}

	fmt.Printf("Дроби 1/i, меньшие %.3f:\n", m)

	fmt.Printf("Вариант 1:\n")
	i := 2
	for {
		value := 1.0 / float64(i)
		if value >= m {
			break
		}
		fmt.Printf("1/%d = %.4f\n", i, value)
		i++
	}

	fmt.Printf("\nВариант 2:\n")

	minI := int(math.Floor(1.0/m)) + 1
	fmt.Printf("Дроби, начиная с 1/%d:\n", minI)
	for i := minI; 1.0/float64(i) < m; i++ {
		fmt.Printf("1/%d = %.4f\n", i, 1.0/float64(i))
	}
}

func ch8ex17() {
	fmt.Printf("Ex17. Последовательность с разностью ≤ 0.001\n")

	num1, den1 := 1.0, 1.0
	num2, den2 := 2.0, 1.0

	prevValue := num1 / den1
	currentValue := num2 / den2
	iteration := 2

	fmt.Printf("Последовательность дробей:\n")
	fmt.Printf("1: %d/%d = %.4f\n", int(num1), int(den1), prevValue)
	fmt.Printf("2: %d/%d = %.4f\n", int(num2), int(den2), currentValue)

	for math.Abs(currentValue-prevValue) > 0.001 {

		nextNum := num1 + num2
		nextDen := den1 + den2

		num1, den1 = num2, den2
		num2, den2 = nextNum, nextDen

		prevValue = currentValue
		currentValue = num2 / den2
		iteration++

		fmt.Printf("%d: %d/%d = %.4f (разность: %.4f)\n",
			iteration, int(num2), int(den2), currentValue,
			math.Abs(currentValue-prevValue))
	}

	fmt.Printf("\nПервый член с разностью ≤ 0.001: %d/%d = %.4f\n",
		int(num2), int(den2), currentValue)
	fmt.Printf("Разность с предыдущим: %.4f\n", math.Abs(currentValue-prevValue))
	fmt.Printf("Номер итерации: %d\n", iteration)
}

func ch8ex18() {
	fmt.Printf("Ex18. Итерационная последовательность\n")

	var a, x, eps float64
	fmt.Print("Введите a (>0): ")

	if _, err := fmt.Scan(&a); err != nil || a <= 0 {
		fmt.Printf("Ошибка: a должно быть положительным\n")
		return
	}

	fmt.Print("Введите x (>0): ")

	if _, err := fmt.Scan(&x); err != nil || x <= 0 {
		fmt.Printf("Ошибка: x должно быть положительным\n")
		return
	}

	fmt.Print("Введите ε (точность, >0): ")

	if _, err := fmt.Scan(&eps); err != nil || eps <= 0 {
		fmt.Printf("Ошибка: ε должно быть положительным\n")
		return
	}

	y := 0.5 * (x + a/x)
	iteration := 1

	fmt.Printf("Итерации:\n")
	fmt.Printf("y1 = 0.5*(%.2f + %.2f/%.2f) = %.6f\n", x, a, x, y)

	for math.Abs(y*y-a) >= eps {
		x = y
		y = 0.5 * (x + a/x)
		iteration++

		fmt.Printf("y%d = 0.5*(%.6f + %.2f/%.6f) = %.6f\n",
			iteration, x, a, x, y)

		if iteration > 1000 {
			fmt.Printf("Превышено максимальное количество итераций\n")
			break
		}
	}

	fmt.Printf("\nРезультат: y = %.6f (после %d итераций)\n", y, iteration)
	fmt.Printf("y² = %.6f, разность с a: %.6f\n", y*y, math.Abs(y*y-a))
}

func ch8ex19() {
	fmt.Printf("Ex19. Последовательность Фибоначчи\n")

	fmt.Printf("\nа) Сумма чисел Фибоначчи ≤ 1000:\n")

	a, b := 1, 1
	sum := a + b

	fmt.Printf("Числа Фибоначчи: %d, %d", a, b)

	for {
		next := a + b
		if next > 1000 {
			break
		}
		fmt.Printf(", %d", next)
		sum += next
		a, b = b, next
	}

	fmt.Printf("\nСумма: %d\n", sum)

	fmt.Printf("\nб) Первое число Фибоначчи, большее n:\n")

	var n int
	fmt.Print("Введите n (>1): ")

	if _, err := fmt.Scan(&n); err != nil || n <= 1 {
		fmt.Printf("Ошибка: n должно быть больше 1\n")
		return
	}

	a, b = 1, 1

	fmt.Print("Числа Фибоначчи: ")
	for i := 1; ; i++ {
		if i == 1 || i == 2 {
			fmt.Printf("%d ", a)
			if a > n {
				fmt.Printf("\nПервое число, большее %d: %d\n", n, a)
				break
			}
		} else {
			next := a + b
			fmt.Printf("%d ", next)
			if next > n {
				fmt.Printf("\nПервое число, большее %d: %d\n", n, next)
				break
			}
			a, b = b, next
		}
	}
}
