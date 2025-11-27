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
		ch5ex25, ch5ex26, ch5ex27, ch5ex28, ch5ex29, ch5ex30, ch5ex31, ch5ex32,
		ch5ex33, ch5ex34, ch5ex35, ch5ex36, ch5ex37, ch5ex38, ch5ex39, ch5ex40,
		ch5ex41, ch5ex42, ch5ex43, ch5ex44, ch5ex45, ch5ex46, ch5ex47, ch5ex48,
		ch5ex49, ch5ex50, ch5ex51, ch5ex52, ch5ex53, ch5ex54, ch5ex55, ch5ex56,
		ch5ex57, ch5ex58, ch5ex59, ch5ex60, ch5ex61, ch5ex62, ch5ex63, ch5ex64,
		ch5ex65, ch5ex66, ch5ex67, ch5ex68, ch5ex69, ch5ex70, ch5ex71, ch5ex72,
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

func ch5ex31() {
	fmt.Printf("Ex31. Сумма n^2 + (n+1)^2 + ... + (2n)^2\n")
	var n int
	fmt.Scan(&n)
	sum := 0
	for i := n; i <= 2*n; i++ {
		sum += i * i
	}
	fmt.Printf("Сумма: %d\n", sum)
}

func ch5ex32() {
	fmt.Printf("Ex32. Сумма 1 + 1/2 + 1/3 + ... + 1/n\n")
	var n int
	fmt.Scan(&n)
	sum := 0.0
	for i := 1; i <= n; i++ {
		sum += 1.0 / float64(i)
	}
	fmt.Printf("Сумма: %.6f\n", sum)
}

func ch5ex33() {
	fmt.Printf("Ex33. Сумма 2/3 + 3/4 + 4/5 + ... + 10/11\n")
	sum := 0.0
	for i := 2; i <= 10; i++ {
		sum += float64(i) / float64(i+1)
	}
	fmt.Printf("Сумма: %.6f\n", sum)
}

func ch5ex34() {
	fmt.Printf("Ex34. Сумма 1^2 + 2^2 + 3^2 + ... + n^2\n")
	var n int
	fmt.Scan(&n)
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	fmt.Printf("Сумма: %d\n", sum)
}

func ch5ex35() {
	fmt.Printf("Ex35. Сумма 1*2 + 2*3 + ... + (n-1)*n\n")
	var n int
	fmt.Scan(&n)
	sum := 0
	for i := 1; i < n; i++ {
		sum += i * (i + 1)
	}
	fmt.Printf("Сумма: %d\n", sum)
}

func ch5ex36() {
	fmt.Printf("Ex36. Сумма 1 + 1/3 + 1/3^2 + ... + 1/3^8\n")
	sum, term := 1.0, 1.0
	for i := 1; i <= 8; i++ {
		term /= 3.0
		sum += term
	}
	fmt.Printf("Сумма: %.6f\n", sum)
}

func ch5ex37() {
	fmt.Printf("Ex37. Сумма 1 - 1/2 + 1/3 + ... + (-1)^n+1/n\n")
	var n int
	fmt.Scan(&n)
	sum, sign := 0.0, 1.0
	for i := 1; i <= n; i++ {
		sum += sign / float64(i)
		sign = -sign
	}
	fmt.Printf("Сумма: %.6f\n", sum)
}

func ch5ex38() {
	fmt.Printf("Ex38. Сумма x + x^3/3 + x^5/5 + ... + x^11/11 при x=2\n")
	x := 2.0
	sum, power := 0.0, x
	for i := 1; i <= 11; i += 2 {
		sum += power / float64(i)
		power *= x * x
	}
	fmt.Printf("Сумма: %.6f\n", sum)
}

func ch5ex39() {
	fmt.Printf("Ex39. Сумма 1 - 2/3 x + 3/4 x^2 - ... + 11/12 x^10 при x=2\n")
	x := 2.0
	sum, power, sign := 0.0, 1.0, 1.0
	for i := 1; i <= 11; i++ {
		sum += sign * float64(i) / float64(i+1) * power
		power *= x
		sign = -sign
	}
	fmt.Printf("Сумма: %.6f\n", sum)
}

func ch5ex40() {
	fmt.Printf("Ex40. Сумма цифр 9-значного числа\n")
	var num int
	fmt.Scan(&num)
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	fmt.Printf("Сумма цифр: %d\n", sum)
}

func ch5ex41() {
	fmt.Printf("Ex41. Произведение и сумма последних n цифр числа\n")
	var num, n int
	fmt.Scan(&num, &n)
	sum, product := 0, 1

	for i := 0; i < n; i++ {
		digit := num % 10
		sum += digit
		product *= digit
		num /= 10
	}
	fmt.Printf("Сумма: %d, Произведение: %d\n", sum, product)
}

func ch5ex42() {
	fmt.Printf("Ex42. Странный муж\n")
	n := 100
	distance := 0.0
	totalPath := 0.0
	sign := 1.0

	for i := 1; i <= n; i++ {
		distance += sign * (1.0 / float64(i))
		totalPath += 1.0 / float64(i)
		sign = -sign
	}

	fmt.Printf("а) Расстояние от дома: %.4f км\n", distance)
	fmt.Printf("б) Общий путь: %.4f км\n", totalPath)
}

func ch5ex43() {
	fmt.Printf("Ex43. Последовательность a0=1, ak=k*a(k-1)+1/k\n")
	var n int
	fmt.Scan(&n)
	a := 1.0
	fmt.Printf("a0 = %.6f\n", a)
	for k := 1; k <= n; k++ {
		a = float64(k)*a + 1.0/float64(k)
		fmt.Printf("a%d = %.4f\n", k, a)
	}
}

func ch5ex44() {
	fmt.Printf("Ex44. Последовательность Фибоначчи\n")
	var n int
	fmt.Scan(&n)

	a, b := 1, 1
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	fmt.Printf("а) %d-й член: %d\n", n, b)

	fmt.Printf("б) Первые %d членов: ", n)
	a, b = 1, 1
	if n >= 1 {
		fmt.Printf("1 ")
	}
	if n >= 2 {
		fmt.Printf("1 ")
	}
	for i := 3; i <= n; i++ {
		a, b = b, a+b
		fmt.Printf("%d ", b)
	}
	fmt.Println()
}

func ch5ex45() {
	fmt.Printf("Ex45. 3-й, 4-й, ..., 10-й член Фибоначчи\n")
	a, b := 1, 1
	for i := 3; i <= 10; i++ {
		a, b = b, a+b
		fmt.Printf("%d-й член: %d\n", i, b)
	}
}

func ch5ex46() {
	fmt.Printf("Ex46. Последовательность дробей\n")

	var k int
	fmt.Scan(&k)
	num1, num2 := 1, 2
	den1, den2 := 1, 1
	for i := 3; i <= k; i++ {
		num1, num2 = num2, num1+num2
		den1, den2 = den2, den1+den2
	}
	fmt.Printf("а) %d-й член: %d/%d = %.6f\n", k, num2, den2, float64(num2)/float64(den2))

	var n int
	fmt.Scan(&n)
	fmt.Printf("б) Первые %d членов:\n", n)
	num1, num2 = 1, 2
	den1, den2 = 1, 1
	if n >= 1 {
		fmt.Printf("1: %d/%d = %.6f\n", num1, den1, float64(num1)/float64(den1))
	}
	if n >= 2 {
		fmt.Printf("2: %d/%d = %.6f\n", num2, den2, float64(num2)/float64(den2))
	}
	for i := 3; i <= n; i++ {
		num1, num2 = num2, num1+num2
		den1, den2 = den2, den1+den2
		fmt.Printf("%d: %d/%d = %.6f\n", i, num2, den2, float64(num2)/float64(den2))
	}
}

func ch5ex47() {
	fmt.Printf("Ex47. Последовательность v1, v2, ...\n")
	var n int
	fmt.Scan(&n)
	v := make([]float64, n+1)
	v[1], v[2], v[3] = 0, 0, 1.5

	for i := 4; i <= n; i++ {
		v[i] = (float64(i-1)/(float64(i*i)+1))*v[i-1] - v[i-2] + v[i-3]
	}

	fmt.Printf("v%d = %.6f\n", n, v[n])
}

func ch5ex48() {
	fmt.Printf("Ex48. Деление амебы\n")
	cells := 1
	fmt.Printf("Начально: %d клеток\n", cells)
	for hours := 3; hours <= 24; hours += 3 {
		cells *= 2
		fmt.Printf("Через %d часов: %d клеток\n", hours, cells)
	}
}

func ch5ex49() {
	fmt.Printf("Ex49. Банковский вклад\n")
	deposit := 1000.0

	fmt.Printf("а) Прирост суммы вклада:\n")
	current := deposit
	for month := 1; month <= 10; month++ {
		increase := current * 0.02
		current += increase
		fmt.Printf("За %d-й месяц: %.2f руб.\n", month, increase)
	}

	fmt.Printf("б) Сумма вклада:\n")
	current = deposit
	for month := 1; month <= 12; month++ {
		current *= 1.02
		if month >= 3 {
			fmt.Printf("Через %d месяцев: %.2f руб.\n", month, current)
		}
	}
}

func ch5ex50() {
	fmt.Printf("Ex50. Тренировки лыжника\n")
	distance := 10.0

	fmt.Printf("а) Пробег по дням:\n")
	fmt.Printf("1-й день: %.2f км\n", distance)
	for day := 2; day <= 10; day++ {
		distance *= 1.1
		fmt.Printf("%d-й день: %.2f км\n", day, distance)
	}

	fmt.Printf("б) Суммарный путь за 7 дней:\n")
	distance = 10.0
	total := distance
	for day := 2; day <= 7; day++ {
		distance *= 1.1
		total += distance
	}
	fmt.Printf("За 7 дней: %.2f км\n", total)
}

func ch5ex51() {
	fmt.Printf("Ex51. Урожайность ячменя\n")
	area := 100.0
	yield := 20.0

	fmt.Printf("а) Урожайность:\n")
	fmt.Printf("1-й год: %.2f ц/га\n", yield)
	for year := 2; year <= 8; year++ {
		yield *= 1.02
		fmt.Printf("%d-й год: %.2f ц/га\n", year, yield)
	}

	fmt.Printf("б) Площадь участка:\n")
	area = 100.0
	for year := 1; year <= 7; year++ {
		area *= 1.05
		if year >= 4 {
			fmt.Printf("%d-й год: %.2f га\n", year, area)
		}
	}

	fmt.Printf("в) Урожай за 6 лет:\n")
	area = 100.0
	yield = 20.0
	totalHarvest := 0.0
	for year := 1; year <= 6; year++ {
		harvest := area * yield
		totalHarvest += harvest
		area *= 1.05
		yield *= 1.02
	}
	fmt.Printf("За 6 лет: %.2f ц\n", totalHarvest)
}

func ch5ex52() {
	fmt.Printf("Ex52. Объем шаров\n")
	thickness := 0.5
	innerDiameter := 10.0
	totalVolume := 0.0

	currentDiameter := innerDiameter
	for i := 1; i <= 12; i++ {
		radius := currentDiameter / 2
		volume := (4.0 / 3.0) * math.Pi * math.Pow(radius, 3)
		totalVolume += volume
		currentDiameter += 2 * thickness
	}

	totalVolumeLiters := totalVolume / 1000
	fmt.Printf("Суммарный объем: %.6f л\n", totalVolumeLiters)
}

func ch5ex53() {
	fmt.Printf("Ex53. Сумма 2^2 + 3^2 + ... + 2^10\n")
	sum, power := 0, 1
	for i := 1; i <= 10; i++ {
		power *= 2
		sum += power
	}
	fmt.Printf("Сумма: %d\n", sum)
}

func ch5ex54() {
	fmt.Printf("Ex54. Степени числа a\n")
	var a float64
	var n int
	fmt.Scan(&a, &n)

	p := 1.0
	for i := 1; i <= n; i++ {
		p *= a
		fmt.Printf("a^%d = %.6f\n", i, p)
	}
}

func ch5ex55() {
	fmt.Printf("Ex55. Сумма -1^2 + 2^2 - 3^2 + ... + 10^2\n")
	sum, sign := 0, -1
	for i := 1; i <= 10; i++ {
		sum += sign * i * i
		sign = -sign
	}
	fmt.Printf("Сумма: %d\n", sum)
}

func ch5ex56() {
	fmt.Printf("Ex56. Площадь арки синусоиды\n")

	a, b := 0.0, math.Pi
	n := 1000
	h := (b - a) / float64(n)
	area := 0.0

	for i := 0; i < n; i++ {
		x := a + float64(i)*h
		area += math.Sin(x) * h
	}

	fmt.Printf("Площадь арки синусоиды: %.6f\n", area)
}

func ch5ex57() {
	fmt.Printf("Ex57. Площадь фигуры y=0.3(x-1)^2+3\n")

	x1 := 1.0 - math.Sqrt(1.0/0.3)
	x2 := 1.0 + math.Sqrt(1.0/0.3)

	n := 1000
	h := (x2 - x1) / float64(n)
	area := 0.0

	for i := 0; i < n; i++ {
		x := x1 + float64(i)*h
		y := 0.3*math.Pow(x-1, 2) + 3
		area += y * h
	}

	fmt.Printf("Площадь фигуры: %.6f\n", area)
}

func ch5ex58() {
	fmt.Printf("Ex58. Площадь фигуры y=0.4(x+2)^2+1\n")

	x1 := -2.0 - math.Sqrt(2.5)
	x2 := -2.0 + math.Sqrt(2.5)

	n := 1000
	h := (x2 - x1) / float64(n)
	area := 0.0

	for i := 0; i < n; i++ {
		x := x1 + float64(i)*h
		y := 0.4*math.Pow(x+2, 2) + 1
		area += y * h
	}

	fmt.Printf("Площадь фигуры: %.6f\n", area)
}

func ch5ex59() {
	fmt.Printf("Ex59. Факториал числа\n")
	var n int
	fmt.Scan(&n)

	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}

	fmt.Printf("%d! = %d\n", n, factorial)
}

func ch5ex60() {
	fmt.Printf("Ex60. Произведение без умножения\n")
	var n int
	var a float64
	fmt.Scan(&n, &a)

	product := 0.0
	for i := 0; i < n; i++ {
		product += a
	}

	fmt.Printf("%d * %.2f = %.2f\n", n, a, product)
}

func ch5ex61() {
	fmt.Printf("Ex61. Произведение через сложение\n")
	var x, y int
	fmt.Scan(&x, &y)

	product1 := 0
	for i := 0; i < y; i++ {
		product1 += x
	}

	product2 := 0
	for i := 0; i < x; i++ {
		product2 += y
	}

	fmt.Printf("Способ 1: %d * %d = %d\n", x, y, product1)
	fmt.Printf("Способ 2: %d * %d = %d\n", x, y, product2)
}

func ch5ex62() {
	fmt.Printf("Ex62. Степень без возведения в степень\n")
	var a float64
	var n int
	fmt.Scan(&a, &n)

	power := 1.0
	for i := 0; i < n; i++ {
		power *= a
	}

	fmt.Printf("%.2f^%d = %.6f\n", a, n, power)
}

func ch5ex63() {
	fmt.Printf("Ex63. Сложное выражение\n")
	result := 20.0*20.0 - 19.0*19.0
	result = result * result / 2.0

	for i := 18; i >= 1; i-- {
		result = result*result - float64(i*i)
	}

	fmt.Printf("Результат: %.6f\n", result)
}

func ch5ex64() {
	fmt.Printf("Ex64. Число наоборот (7-значное)\n")
	var num int
	fmt.Scan(&num)

	reversed := 0
	temp := num
	for temp > 0 {
		digit := temp % 10
		reversed = reversed*10 + digit
		temp /= 10
	}

	fmt.Printf("Исходное: %d, Наоборот: %d\n", num, reversed)
}

func ch5ex65() {
	fmt.Printf("Ex65. Квадрат через сумму нечетных\n")
	var n int
	fmt.Scan(&n)

	sum := 0
	for i := 1; i <= 2*n-1; i += 2 {
		sum += i
	}

	fmt.Printf("%d^2 = %d\n", n, sum)
}

func ch5ex66() {
	fmt.Printf("Ex66. Сумма квадратов 1^2 + 2^2 + ... + 12^2\n")
	totalSum := 0
	for n := 1; n <= 12; n++ {
		sum := 0
		for i := 1; i <= 2*n-1; i += 2 {
			sum += i
		}
		totalSum += sum
	}

	fmt.Printf("Сумма квадратов: %d\n", totalSum)
}

func ch5ex67() {
	fmt.Printf("Ex67. Куб через сумму последовательных нечетных\n")
	var n int
	fmt.Scan(&n)

	first := n*n - n + 1
	sum := 0
	for i := 0; i < n; i++ {
		sum += first + 2*i
	}

	fmt.Printf("%d^3 = %d\n", n, sum)
}

func ch5ex68() {
	fmt.Printf("Ex68. Сумма факториалов 1! + 2! + ... + n!\n")
	var n int
	fmt.Scan(&n)

	totalSum := 0
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
		totalSum += factorial
	}

	fmt.Printf("Сумма факториалов: %d\n", totalSum)
}

func ch5ex69() {
	fmt.Printf("Ex69. Сумма 1 + 1/1! + 1/2! + ... + 1/n!\n")
	var n int
	fmt.Scan(&n)

	sum := 1.0
	factorial := 1.0
	for i := 1; i <= n; i++ {
		factorial *= float64(i)
		sum += 1.0 / factorial
	}

	fmt.Printf("Сумма: %.6f\n", sum)
}

func ch5ex70() {
	fmt.Printf("Ex70. Сумма 1 + x^1/1! + x^2/2! + ... + x^n/n!\n")
	var n int
	var x float64
	fmt.Scan(&n, &x)

	sum := 1.0
	term := 1.0
	factorial := 1.0
	for i := 1; i <= n; i++ {
		term *= x
		factorial *= float64(i)
		sum += term / factorial
	}

	fmt.Printf("Сумма: %.6f\n", sum)
}

func ch5ex71() {
	fmt.Printf("Ex71. Сумма вложенных корней\n")
	result := math.Sqrt(50.0)
	for i := 49; i >= 1; i-- {
		result = math.Sqrt(float64(i) + result)
	}

	fmt.Printf("Результат: %.6f\n", result)
}

func ch5ex72() {
	fmt.Printf("Ex72. Разные суммы\n")
	var n int
	fmt.Scan(&n)

	sumA := 0.0
	sinSum := 0.0
	for i := 1; i <= n; i++ {
		sinSum += math.Sin(float64(i))
		sumA += 1.0 / sinSum
	}

	resultB := math.Sqrt(2.0)
	for i := 2; i <= n; i++ {
		resultB = math.Sqrt(2.0 + resultB)
	}

	sumC := 0.0
	cosSum := 0.0
	for i := 1; i <= n; i++ {
		cosSum += math.Cos(float64(i))
		sumC += cosSum
	}

	resultD := math.Sqrt(3.0 * float64(n))
	for i := n - 1; i >= 1; i-- {
		resultD = math.Sqrt(3.0*float64(i) + resultD)
	}

	fmt.Printf("а) Сумма обратных синусов: %.6f\n", sumA)
	fmt.Printf("б) Вложенные корни из 2: %.6f\n", resultB)
	fmt.Printf("в) Сумма сумм косинусов: %.6f\n", sumC)
	fmt.Printf("г) Вложенные корни с 3: %.6f\n", resultD)
}
