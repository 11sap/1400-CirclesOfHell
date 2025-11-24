package chapters

import (
	"fmt"
	"math"
)

func Run05() {
	tasks := []func(){
		ch5ex01, ch5ex02, ch5ex03, ch5ex04, ch5ex05, ch5ex06, ch5ex07, ch5ex08,
		ch5ex09, ch5ex10, ch5ex11, ch5ex12, ch5ex13, ch5ex14, ch5ex15, ch5ex16,
		ch5ex17, ch5ex18, ch5ex19, ch5ex20, ch5ex21, ch5ex22, ch5ex23, ch5ex24,
		ch5ex25, ch5ex26, ch5ex27, ch5ex28, ch5ex29, ch5ex30,
	}

	for {
		fmt.Print("\nГлава 5. Выберите задачу (1-94): ")
		var n int
		fmt.Scan(&n)

		if n == 0 {
			break
		}
		if n >= 1 && n <= 94 {
			tasks[n-1]()
		} else {
			fmt.Println("От 1 до 94! Для выхода введите 0.")
		}
	}
}

func ch5ex01() {
	fmt.Printf("Ex01. Ряд чисел 20\n")
	for i := 0; i < 8; i++ {
		fmt.Printf("20 ")
	}
	fmt.Println()
}

func ch5ex02() {
	var num, count int
	fmt.Printf("Ex02. Число повторенное N раз\n")
	fmt.Printf("Ex02. Введите число и количество повторений: ")
	fmt.Scan(&num, &count)

	for i := 0; i < count; i++ {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

func ch5ex03() {
	fmt.Printf("Ex03. Числа столбиком\n")

	fmt.Printf("Ex03a. От 20 до 55:\n")
	for i := 20; i <= 55; i++ {
		fmt.Printf("%d\n", i)
	}

	var a int
	fmt.Printf("Ex03b. Введите a (a ≤ 50): ")
	fmt.Scan(&a)
	fmt.Printf("Квадраты от %d до 50:\n", a)
	for i := a; i <= 50; i++ {
		fmt.Printf("%d\n", i*i)
	}

	var b int
	fmt.Printf("Ex03c. Введите b (b > 10): ")
	fmt.Scan(&b)
	fmt.Printf("Кубы от 10 до %d:\n", b)
	for i := 10; i <= b; i++ {
		fmt.Printf("%d\n", i*i*i)
	}

	var a2, b2 int
	fmt.Printf("Ex03d. Введите a и b (b > a): ")
	fmt.Scan(&a2, &b2)
	fmt.Printf("Числа от %d до %d:\n", a2, b2)
	for i := a2; i <= b2; i++ {
		fmt.Printf("%d\n", i)
	}
}

func ch5ex04() {
	fmt.Printf("Ex04. Таблицы чисел\n")

	fmt.Printf("Ex04a:\n")
	for i := 10; i <= 25; i++ {
		fmt.Printf("%d %.1f\n", i, float64(i)+0.4)
	}

	fmt.Printf("Ex04b:\n")
	for i := 25; i <= 35; i++ {
		fmt.Printf("%d %.1f %.1f\n", i, float64(i)+0.5, float64(i)-0.2)
	}
}

func ch5ex05() {
	fmt.Printf("Ex05. Обратные таблицы чисел\n")

	fmt.Printf("Ex05a:\n")
	for i := 21; i >= 10; i-- {
		fmt.Printf("%d %.1f\n", i, float64(i)-1.8)
	}

	fmt.Printf("Ex05b:\n")
	for i := 45; i >= 25; i-- {
		fmt.Printf("%d %.1f %.1f\n", i, float64(i)-0.5, float64(i)-0.8)
	}
}

func ch5ex06() {
	fmt.Printf("Ex06. Возрастающие таблицы чисел\n")

	fmt.Printf("Ex06a:\n")
	for i := 21; i <= 35; i++ {
		fmt.Printf("%d %.1f\n", i, float64(i)-0.6)
	}

	fmt.Printf("Ex06b:\n")
	for i := 16; i <= 24; i++ {
		fmt.Printf("%d %.1f %.1f\n", i, float64(i)-0.5, float64(i)+0.8)
	}
}

func ch5ex07() {
	price := 20.4
	fmt.Printf("Ex07. Стоимость товара\n")
	fmt.Printf("Количество\tСтоимость\n")
	for i := 2; i <= 20; i++ {
		cost := price * float64(i)
		fmt.Printf("%d\t\t%.1f руб.\n", i, cost)
	}
}

func ch5ex08() {
	fmt.Printf("Ex08. Таблица фунты-килограммы\n")
	fmt.Printf("Фунты\tкг\n")
	for pounds := 1; pounds <= 10; pounds++ {
		kg := float64(pounds) * 0.453
		fmt.Printf("%d\t%.3f\n", pounds, kg)
	}
}

func ch5ex09() {
	fmt.Printf("Ex09. Таблица дюймы-сантиметры\n")
	fmt.Printf("Дюймы\tСантиметры\n")
	for inches := 10; inches <= 22; inches++ {
		cm := float64(inches) * 2.54
		fmt.Printf("%d\t%.2f\n", inches, cm)
	}
}

func ch5ex10() {
	var rate float64
	fmt.Printf("Ex10. Таблица доллары-рубли\n")
	fmt.Printf("Введите курс доллара: ")
	fmt.Scan(&rate)

	fmt.Printf("Доллары\tРубли\n")
	for dollars := 1; dollars <= 20; dollars++ {
		rubles := float64(dollars) * rate
		fmt.Printf("%d\t%.2f\n", dollars, rubles)
	}
}

func ch5ex11() {
	R := 6350.0
	fmt.Printf("Ex11. Расстояние до горизонта\n")
	fmt.Printf("Высота (км)\tРасстояние (км)\n")
	for h := 1; h <= 10; h++ {
		distance := math.Sqrt(math.Pow(R+float64(h), 2) - math.Pow(R, 2))
		fmt.Printf("%d\t\t%.2f\n", h, distance)
	}
}

func ch5ex12() {
	p := 1.29
	z := 1.25e-4
	fmt.Printf("Ex12. Плотность воздуха от высоты\n")
	fmt.Printf("Высота (м)\tПлотность (кг/м³)\n")
	for h := 0; h <= 1000; h += 100 {
		p := p * math.Exp(-float64(h)/z)
		fmt.Printf("%d\t\t%.4f\n", h, p)
	}
}

func ch5ex13() {
	fmt.Printf("Ex13. Таблица умножения на 7\n")
	for i := 1; i <= 9; i++ {
		fmt.Printf("%d × 7 = %d\n", i, i*7)
	}
}

func ch5ex14() {
	fmt.Printf("Ex14. Таблица умножения на 9\n")
	for i := 1; i <= 9; i++ {
		fmt.Printf("9 × %d = %d\n", i, 9*i)
	}
}

func ch5ex15() {
	var n int
	fmt.Printf("Ex15. Таблица умножения на n\n")
	fmt.Printf("Введите число n (1-9): ")
	fmt.Scan(&n)

	for i := 1; i <= 9; i++ {
		fmt.Printf("%d × %d = %d\n", i, n, i*n)
	}
}

func ch5ex16() {
	fmt.Printf("Ex16. Значения sin 2..15\n")
	for i := 2; i <= 15; i++ {
		fmt.Printf("sin %d = %.4f\n", i, math.Sin(float64(i)))
	}
}

func ch5ex17() {
	fmt.Printf("Ex17. Значения функции y\n")
	fmt.Printf("x\ty\n")
	for x := 4; x <= 28; x++ {
		t := float64(x) + 3
		y := 3*math.Pow(t, 2) + 4.87*t - 3
		fmt.Printf("%d\t%.2f\n", x, y)
	}
}

func ch5ex18() {
	fmt.Printf("Ex18. Значения функции z\n")
	fmt.Printf("a\tz\n")
	for a := 2; a <= 17; a++ {
		t := 3 * float64(a)
		z := 4.3*math.Pow(t, 2) - 8*t + 13
		fmt.Printf("%d\t%.2f\n", a, z)
	}
}

func ch5ex19() {
	fmt.Printf("Ex19. Значения sin 0.1..1.5\n")
	for i := 1; i <= 15; i++ {
		x := float64(i) / 10.0
		fmt.Printf("sin %.1f = %.4f\n", x, math.Sin(x))
	}
}

func ch5ex20() {
	fmt.Printf("Ex20. Значения sqrt 0.1..0.9\n")
	for i := 1; i <= 9; i++ {
		x := float64(i) / 10.0
		fmt.Printf("sqrt %.1f = %.4f\n", x, math.Sqrt(x))
	}
}

func ch5ex21() {
	var pricePerKg float64
	fmt.Printf("Ex21. Стоимость сыра\n")
	fmt.Printf("Введите стоимость 1 кг сыра: ")
	fmt.Scan(&pricePerKg)

	fmt.Printf("Вес (г)\tСтоимость (руб)\n")
	for weight := 50; weight <= 1000; weight += 50 {
		cost := pricePerKg * float64(weight) / 1000
		fmt.Printf("%d\t%.2f\n", weight, cost)
	}
}

func ch5ex22() {
	var pricePerKg float64
	fmt.Printf("Ex22. Стоимость конфет\n")
	fmt.Printf("Введите стоимость 1 кг конфет: ")
	fmt.Scan(&pricePerKg)

	fmt.Printf("Вес (г)\tСтоимость (руб)\n")
	for weight := 100; weight <= 2000; weight += 100 {
		cost := pricePerKg * float64(weight) / 1000
		fmt.Printf("%d\t%.2f\n", weight, cost)
	}
}

func ch5ex23() {
	fmt.Printf("Ex23. Числа 2.1-2.8 столбиком\n")
	for i := 21; i <= 28; i++ {
		fmt.Printf("%.1f\n", float64(i)/10)
	}
}

func ch5ex24() {
	fmt.Printf("Ex24. Числа 3.2-3.9 столбиком\n")
	for i := 32; i <= 39; i++ {
		fmt.Printf("%.1f\n", float64(i)/10)
	}
}

func ch5ex25() {
	fmt.Printf("Ex25. Числа 2.2-4.2 с шагом 0.2\n")
	for i := 22; i <= 42; i += 2 {
		fmt.Printf("%.1f\n", float64(i)/10)
	}
}

func ch5ex26() {
	fmt.Printf("Ex26. Числа 4.4-6.4 с шагом 0.2\n")
	for i := 44; i <= 64; i += 2 {
		fmt.Printf("%.1f\n", float64(i)/10)
	}
}

func ch5ex27() {
	fmt.Printf("Ex27. Суммы целых чисел\n")

	sum := 0
	for i := 200; i <= 600; i++ {
		sum += i
	}
	fmt.Printf("Ex27a. Сумма от 200 до 600: %d\n", sum)

	var a int
	fmt.Printf("Ex27b. Введите a (a ≤ 400): ")
	fmt.Scan(&a)
	sum = 0
	for i := a; i <= 400; i++ {
		sum += i
	}
	fmt.Printf("Сумма от %d до 400: %d\n", a, sum)

	var b int
	fmt.Printf("Ex27c. Введите b (b > -15): ")
	fmt.Scan(&b)
	sum = 0
	for i := -15; i <= b; i++ {
		sum += i
	}
	fmt.Printf("Сумма от -15 до %d: %d\n", b, sum)

	var a2, b2 int
	fmt.Printf("Ex27d. Введите a и b (b > a): ")
	fmt.Scan(&a2, &b2)
	sum = 0
	for i := a2; i <= b2; i++ {
		sum += i
	}
	fmt.Printf("Сумма от %d до %d: %d\n", a2, b2, sum)
}

func ch5ex28() {
	fmt.Printf("Ex28. Произведения целых чисел\n")

	product := 1
	for i := 7; i <= 14; i++ {
		product *= i
	}
	fmt.Printf("Ex28a. Произведение от 7 до 14: %d\n", product)

	var a int
	fmt.Printf("Ex28b. Введите a (1 ≤ a < 15): ")
	fmt.Scan(&a)
	product = 1
	for i := a; i <= 15; i++ {
		product *= i
	}
	fmt.Printf("Произведение от %d до 15: %d\n", a, product)

	var b int
	fmt.Printf("Ex28c. Введите b (1 ≤ b < 10): ")
	fmt.Scan(&b)
	product = 1
	for i := 1; i <= b; i++ {
		product *= i
	}
	fmt.Printf("Произведение от 1 до %d: %d\n", b, product)

	var a2, b2 int
	fmt.Printf("Ex28d. Введите a и b (b > a): ")
	fmt.Scan(&a2, &b2)
	product = 1
	for i := a2; i <= b2; i++ {
		product *= i
	}
	fmt.Printf("Произведение от %d до %d: %d\n", a2, b2, product)
}

func ch5ex29() {
	fmt.Printf("Ex29. Средние арифметические\n")

	sum, count := 0, 0
	for i := 3; i <= 750; i++ {
		sum += i
		count++
	}
	average := float64(sum) / float64(count)
	fmt.Printf("Ex29a. Среднее от 3 до 750: %.2f\n", average)

	var b int
	fmt.Printf("Ex29b. Введите b (b > 150): ")
	fmt.Scan(&b)
	sum, count = 0, 0
	for i := 150; i <= b; i++ {
		sum += i
		count++
	}
	average = float64(sum) / float64(count)
	fmt.Printf("Среднее от 150 до %d: %.2f\n", b, average)

	var a int
	fmt.Printf("Ex29c. Введите a (a ≤ 200): ")
	fmt.Scan(&a)
	sum, count = 0, 0
	for i := a; i <= 200; i++ {
		sum += i
		count++
	}
	average = float64(sum) / float64(count)
	fmt.Printf("Среднее от %d до 200: %.2f\n", a, average)

	var a2, b2 int
	fmt.Printf("Ex29d. Введите a и b (b > a): ")
	fmt.Scan(&a2, &b2)
	sum, count = 0, 0
	for i := a2; i <= b2; i++ {
		sum += i
		count++
	}
	average = float64(sum) / float64(count)
	fmt.Printf("Среднее от %d до %d: %.2f\n", a2, b2, average)
}

func ch5ex30() {
	fmt.Printf("Ex30. Суммы степеней\n")

	sum := 0
	for i := 25; i <= 40; i++ {
		sum += i * i * i
	}
	fmt.Printf("Ex30a. Сумма кубов от 25 до 40: %d\n", sum)

	var a int
	fmt.Printf("Ex30b. Введите a (0 ≤ a ≤ 50): ")
	fmt.Scan(&a)
	sum = 0
	for i := a; i <= 50; i++ {
		sum += i * i
	}
	fmt.Printf("Сумма квадратов от %d до 50: %d\n", a, sum)

	var n int
	fmt.Printf("Ex30c. Введите n (1 ≤ n ≤ 100): ")
	fmt.Scan(&n)
	sum = 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	fmt.Printf("Сумма квадратов от 1 до %d: %d\n", n, sum)

	var a2, b2 int
	fmt.Printf("Ex30d. Введите a и b (b > a): ")
	fmt.Scan(&a2, &b2)
	sum = 0
	for i := a2; i <= b2; i++ {
		sum += i * i
	}
	fmt.Printf("Сумма квадратов от %d до %d: %d\n", a2, b2, sum)
}
