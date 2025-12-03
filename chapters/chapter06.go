package chapters

import (
	"fmt"
	"math"
)

func Run06() {
	tasks := []func(){
		ch6ex01, ch6ex02, ch6ex03, ch6ex04, ch6ex05, ch6ex06, ch6ex07, ch6ex08,
		ch6ex09, ch6ex10, ch6ex11, ch6ex12, ch6ex13, ch6ex14, ch6ex15, ch6ex16,
		ch6ex17, ch6ex18, ch6ex19, ch6ex20, ch6ex21, ch6ex22, ch6ex23, ch6ex24,
		ch6ex25, ch6ex26, ch6ex27, ch6ex28, ch6ex29, ch6ex30, ch6ex31, ch6ex32,
		ch6ex33, ch6ex34, ch6ex35, ch6ex36, ch6ex37, ch6ex38, ch6ex39, ch6ex40,
		ch6ex41, ch6ex42, ch6ex43, ch6ex44, ch6ex45, ch6ex46, ch6ex47, ch6ex48,
		ch6ex49, ch6ex50, ch6ex51,
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

func ch6ex14() {
	fmt.Printf("Ex14. Числа 100-80 (без for с параметром)\n")
	i := 100
	for i >= 80 {
		fmt.Printf("%d\n", i)
		i--
	}
}

func ch6ex15() {
	fmt.Printf("Ex15. Цикл с параметром (аналог while)\n")
	for n := 21; n <= 151; n += 10 {
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}

func ch6ex16() {
	fmt.Printf("Ex16. Цикл с параметром (с вещественным шагом)\n")

	for n := 2.0; n <= 12.0; n += 0.5 {
		fmt.Printf("%.1f ", n)
	}
	fmt.Println()
}

func ch6ex17() {
	fmt.Printf("Ex17. Числа 1.0, 1.5, 2.0, ..., 13.5 (без for с параметром)\n")
	num := 1.0
	for {
		fmt.Printf("%.1f ", num)
		num += 0.5
		if num > 13.5 {
			break
		}
	}
	fmt.Println()
}

func ch6ex18() {
	fmt.Printf("Ex18. Квадратные корни нечетных a-b\n")
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Printf("а) С предусловием:\n")
	i := a
	for i >= b {
		if i%2 != 0 {
			root := math.Sqrt(float64(i))
			fmt.Printf("√%d = %.2f\n", i, root)
		}
		i--
	}

	fmt.Printf("\nб) С постусловием:\n")
	j := a
	for {
		if j%2 != 0 {
			root := math.Sqrt(float64(j))
			fmt.Printf("√%d = %.2f\n", j, root)
		}
		j--
		if j < b {
			break
		}
	}
}

func ch6ex19() {
	fmt.Printf("Ex19. Цифры числа в столбик\n")
	var n int
	fmt.Scan(&n)

	for n > 0 {
		digit := n % 10
		fmt.Printf("%d\n", digit)
		n /= 10
	}
}

func ch6ex20() {
	fmt.Printf("Ex20. Операции с цифрами числа\n")
	var n int
	fmt.Scan(&n)

	sum, count, sumSquares, sumCubes, firstDigit := 0, 0, 0, 0, 0
	product := 1
	lastDigit := n % 10

	temp := n
	for temp > 0 {
		digit := temp % 10
		sum += digit
		count++
		product *= digit
		sumSquares += digit * digit
		sumCubes += digit * digit * digit
		firstDigit = digit
		temp /= 10
	}

	sumFirstLast := firstDigit + lastDigit

	fmt.Printf("а) Сумма цифр: %d\n", sum)
	fmt.Printf("б) Количество цифр: %d\n", count)
	fmt.Printf("в) Произведение цифр: %d\n", product)
	fmt.Printf("г) Среднее арифметическое: %.2f\n", float64(sum)/float64(count))
	fmt.Printf("д) Сумма квадратов цифр: %d\n", sumSquares)
	fmt.Printf("е) Сумма кубов цифр: %d\n", sumCubes)
	fmt.Printf("ж) Первая цифра: %d\n", firstDigit)
	fmt.Printf("з) Сумма первой и последней цифры: %d\n", sumFirstLast)
}

func ch6ex21() {
	fmt.Printf("Ex21. Вторая цифра числа (n>9)\n")
	var n int
	fmt.Scan(&n)

	for n >= 10 {
		n /= 10
	}
	secondDigit := n % 10

	fmt.Printf("Вторая цифра: %d\n", secondDigit)
}

func ch6ex22() {
	fmt.Printf("Ex22. Третья цифра числа (n>99)\n")
	var n int
	fmt.Scan(&n)

	for n >= 100 {
		n /= 10
	}
	thirdDigit := n % 10

	fmt.Printf("Третья цифра: %d\n", thirdDigit)
}

func ch6ex23() {
	fmt.Printf("Ex23. Сумма последних m цифр\n")
	var n, m int
	fmt.Scan(&n, &m)

	sum := 0
	for i := 0; i < m; i++ {
		if n == 0 {
			break
		}
		sum += n % 10
		n /= 10
	}

	fmt.Printf("Сумма последних %d цифр: %d\n", m, sum)
}

func ch6ex24() {
	fmt.Printf("Ex24. Знакопеременные суммы цифр\n")
	var n int
	fmt.Scan(&n)

	sign, position := 1, 1
	sumA, sumB := 0, 0

	temp := n
	for temp > 0 {
		digit := temp % 10

		sumA += sign * digit
		sign = -sign

		if position%2 == 1 {
			sumB += digit
		} else {
			sumB -= digit
		}
		position++

		temp /= 10
	}

	fmt.Printf("а) Знакопеременная сумма (справа налево): %d\n", sumA)
	fmt.Printf("б) Знакопеременная сумма (с учетом позиции): %d\n", sumB)
}

func ch6ex25() {
	fmt.Printf("Ex25. Наименьший делитель, отличный от 1\n")
	var n int
	fmt.Scan(&n)

	divisor := 2
	for n%divisor != 0 {
		divisor++
	}

	fmt.Printf("Наименьший делитель, отличный от 1: %d\n", divisor)
}

func ch6ex26() {
	fmt.Printf("Ex26. Кратные тринадцати <100\n")

	fmt.Printf("а) Без while:\n")
	for i := 13; i < 100; i += 13 {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	fmt.Printf("б) С while:\n")
	i := 13
	for i < 100 {
		fmt.Printf("%d ", i)
		i += 13
	}
	fmt.Println()
}

func ch6ex27() {
	fmt.Printf("Ex27. 15 чисел >100, делящихся на 19\n")
	count := 0
	n := 100
	for count < 15 {
		n++
		if n%19 == 0 {
			fmt.Printf("%d ", n)
			count++
		}
	}
	fmt.Println()
}

func ch6ex28() {
	fmt.Printf("Ex28. 20 чисел >500, делящихся на 13 или 17\n")
	count := 0
	n := 500
	for count < 20 {
		n++
		if n%13 == 0 || n%17 == 0 {
			fmt.Printf("%d ", n)
			count++
		}
	}
	fmt.Println()
}

func ch6ex29() {
	fmt.Printf("Ex29. 10 чисел >100, оканчивающихся на 7, кратных 9\n")
	count := 0
	n := 100
	for count < 10 {
		n++
		if n%10 == 7 && n%9 == 0 {
			fmt.Printf("%d ", n)
			count++
		}
	}
	fmt.Println()
}

func ch6ex30() {
	fmt.Printf("Ex30. Целочисленное деление и остаток (любые числа)\n")
	var a, b float64
	fmt.Scan(&a, &b)

	quotient := 0
	temp := math.Abs(b)
	divisor := math.Abs(a)
	for temp >= divisor {
		temp -= divisor
		quotient++
	}
	if (a < 0 && b > 0) || (a > 0 && b < 0) {
		quotient = -quotient
	}

	remainder := b
	for math.Abs(remainder) >= math.Abs(a) {
		if remainder > 0 {
			remainder -= math.Abs(a)
		} else {
			remainder += math.Abs(a)
		}
	}

	fmt.Printf("а) Целая часть от деления %.2f на %.2f: %d\n", b, a, quotient)
	fmt.Printf("б) Остаток от деления %.2f на %.2f: %.2f\n", b, a, remainder)
}

func ch6ex31() {
	fmt.Printf("Ex31. Банковский вклад\n")
	deposit := 1000.0
	month := 0

	fmt.Printf("а) Прирост >30 руб:\n")
	current := deposit
	for {
		month++
		increase := current * 0.02
		current += increase
		if increase > 30 {
			fmt.Printf("В %d-й месяц прирост: %.2f руб\n", month, increase)
			break
		}
	}

	fmt.Printf("б) Вклад >1200 руб:\n")
	month = 0
	current = deposit
	for current <= 1200 {
		month++
		current *= 1.02
	}
	fmt.Printf("Через %d месяцев вклад: %.2f руб\n", month, current)
}

func ch6ex32() {
	fmt.Printf("Ex32. Тренировки лыжника\n")
	distance := 10.0
	day := 1

	fmt.Printf("а) Пробег >20 км:\n")
	current := distance
	for current <= 20 {
		day++
		current *= 1.1
	}
	fmt.Printf("В %d-й день пробег: %.2f км\n", day, current)

	fmt.Printf("б) Суммарный пробег >100 км:\n")
	day = 1
	current = distance
	total := current
	for total <= 100 {
		day++
		current *= 1.1
		total += current
	}
	fmt.Printf("В %d-й день суммарный пробег: %.2f км\n", day, total)
}

func ch6ex33() {
	fmt.Printf("Ex33. Урожайность ячменя\n")
	area := 100.0
	yield := 20.0
	year := 1

	fmt.Printf("а) Урожайность >22 ц/га:\n")
	currentYield := yield
	for currentYield <= 22 {
		year++
		currentYield *= 1.02
	}
	fmt.Printf("В %d-й год урожайность: %.2f ц/га\n", year, currentYield)

	fmt.Printf("б) Площадь >120 га:\n")
	year = 1
	currentArea := area
	for currentArea <= 120 {
		year++
		currentArea *= 1.05
	}
	fmt.Printf("В %d-й год площадь: %.2f га\n", year, currentArea)

	fmt.Printf("в) Общий урожай >800 ц:\n")
	year = 1
	currentArea = area
	currentYield = yield
	totalHarvest := 0.0
	for totalHarvest <= 800 {
		harvest := currentArea * currentYield
		totalHarvest += harvest
		year++
		currentArea *= 1.05
		currentYield *= 1.02
	}
	fmt.Printf("В %d-й год общий урожай: %.2f ц\n", year-1, totalHarvest)
}

func ch6ex34() {
	fmt.Printf("Ex34. Кратные взаимно простых чисел\n")
	var m, n int
	fmt.Scan(&m, &n)

	a, b := m, n
	for b != 0 {
		a, b = b, a%b
	}
	if a != 1 {
		fmt.Printf("Числа не взаимно простые\n")
		return
	}

	product := m * n
	for i := product; i <= m*n; i += product {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func ch6ex35() {
	fmt.Printf("Ex35. Статистика цифр числа\n")
	var n int
	fmt.Scan(&n)

	countThree := 0
	countLastDigit := 0
	lastDigit := n % 10
	countEven := 0
	sumGreaterFive := 0
	productGreaterSeven := 1
	countZeroOrFive := 0

	temp := n
	for temp > 0 {
		digit := temp % 10

		if digit == 3 {
			countThree++
		}

		if digit == lastDigit {
			countLastDigit++
		}

		countEven += 1 - (digit % 2)

		if digit > 5 {
			sumGreaterFive += digit
		}

		if digit > 7 {
			productGreaterSeven *= digit
		}

		if digit == 0 || digit == 5 {
			countZeroOrFive++
		}

		temp /= 10
	}

	fmt.Printf("а) Цифр 3: %d\n", countThree)
	fmt.Printf("б) Последняя цифра (%d) встречается: %d раз\n", lastDigit, countLastDigit)
	fmt.Printf("в) Четных цифр: %d\n", countEven)
	fmt.Printf("г) Сумма цифр >5: %d\n", sumGreaterFive)
	fmt.Printf("д) Произведение цифр >7: %d\n", productGreaterSeven)
	fmt.Printf("е) Цифр 0 или 5: %d\n", countZeroOrFive)
}

func ch6ex36() {
	fmt.Printf("Ex36. Цифры числа\n")
	var n, a int
	fmt.Scan(&n, &a)

	countA := 0
	sumGreaterA := 0
	sumEven := 0
	var x, y int
	fmt.Scan(&x, &y)
	countXY := 0

	temp := n
	for temp > 0 {
		digit := temp % 10

		if digit == a {
			countA++
		}

		if digit > a {
			sumGreaterA += digit
		}

		if digit%2 == 0 {
			sumEven += digit
		}

		if digit == x || digit == y {
			countXY++
		}

		temp /= 10
	}

	fmt.Printf("а) Цифра %d встречается: %d раз\n", a, countA)
	fmt.Printf("б) Сумма цифр >%d: %d\n", a, sumGreaterA)
	fmt.Printf("в) Сумма четных цифр: %d\n", sumEven)
	fmt.Printf("г) Цифры %d и %d встречаются: %d раз\n", x, y, countXY)
}

func ch6ex37() {
	fmt.Printf("Ex37. Номер цифры 8 от конца (самая левая)\n")
	var n int
	fmt.Scan(&n)

	position := 1
	result := 0
	temp := n

	for temp > 0 {
		if temp%10 == 8 {
			result = position
		}
		temp /= 10
		position++
	}

	fmt.Printf("Номер цифры 8 от конца: %d\n", result)
}

func ch6ex38() {
	fmt.Printf("Ex38. Номер цифры 3 от конца (самая правая)\n")
	var n int
	fmt.Scan(&n)

	position := 1
	result := 0
	temp := n

	for temp > 0 {
		digit := temp % 10
		if digit == 3 {
			result = position
			break
		}
		temp /= 10
		position++
	}

	fmt.Printf("Номер цифры 3 от конца: %d\n", result)
}

func ch6ex39() {
	fmt.Printf("Ex39. Цифры числа от первой\n")
	var n int
	fmt.Scan(&n)

	digits := []int{}
	temp := n
	for temp > 0 {
		digits = append(digits, temp%10)
		temp /= 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		fmt.Printf("%d\n", digits[i])
	}
}

func ch6ex40() {
	fmt.Printf("Ex40. Сколько раз встречается первая цифра\n")
	var n int
	fmt.Scan(&n)

	temp := n
	firstDigit := 0
	for temp > 0 {
		firstDigit = temp % 10
		temp /= 10
	}

	count := 0
	temp = n
	for temp > 0 {
		if temp%10 == firstDigit {
			count++
		}
		temp /= 10
	}

	fmt.Printf("Первая цифра %d встречается: %d раз\n", firstDigit, count)
}

func ch6ex41() {
	fmt.Printf("Ex41. Максимальная и минимальная цифры\n")
	var n int
	fmt.Scan(&n)

	maxDigit := -1
	minDigit := 10
	temp := n

	for temp > 0 {
		digit := temp % 10
		if digit > maxDigit {
			maxDigit = digit
		}
		if digit < minDigit {
			minDigit = digit
		}
		temp /= 10
	}

	fmt.Printf("а) Максимальная цифра: %d\n", maxDigit)
	fmt.Printf("б) Минимальная цифра: %d\n", minDigit)
}

func ch6ex42() {
	fmt.Printf("Ex42. Операции с max и min цифрами\n")
	var n int
	fmt.Scan(&n)

	maxDigit := -1
	minDigit := 10
	temp := n

	for temp > 0 {
		digit := temp % 10
		if digit > maxDigit {
			maxDigit = digit
		}
		if digit < minDigit {
			minDigit = digit
		}
		temp /= 10
	}

	difference := maxDigit - minDigit
	sum := maxDigit + minDigit

	fmt.Printf("а) Максимальная: %d, Минимальная: %d\n", maxDigit, minDigit)
	fmt.Printf("б) Разность max-min: %d\n", difference)
	fmt.Printf("в) Сумма max+min: %d\n", sum)
}

func ch6ex43() {
	fmt.Printf("Ex43. Сумма max+min кратна ли a\n")
	var n, a int
	fmt.Scan(&n, &a)

	maxDigit := -1
	minDigit := 10
	temp := n

	for temp > 0 {
		digit := temp % 10
		if digit > maxDigit {
			maxDigit = digit
		}
		if digit < minDigit {
			minDigit = digit
		}
		temp /= 10
	}

	sum := maxDigit + minDigit
	if sum%a == 0 {
		fmt.Printf("Сумма %d+%d=%d кратна %d\n", maxDigit, minDigit, sum, a)
	} else {
		fmt.Printf("Сумма %d+%d=%d не кратна %d\n", maxDigit, minDigit, sum, a)
	}
}

func ch6ex44() {
	fmt.Printf("Ex44. Разность max-min четна ли\n")
	var n int
	fmt.Scan(&n)

	maxDigit := -1
	minDigit := 10
	temp := n

	for temp > 0 {
		digit := temp % 10
		if digit > maxDigit {
			maxDigit = digit
		}
		if digit < minDigit {
			minDigit = digit
		}
		temp /= 10
	}

	difference := maxDigit - minDigit
	if difference%2 == 0 {
		fmt.Printf("Разность %d-%d=%d четная\n", maxDigit, minDigit, difference)
	} else {
		fmt.Printf("Разность %d-%d=%d нечетная\n", maxDigit, minDigit, difference)
	}
}

func ch6ex45() {
	fmt.Printf("Ex45. Порядковые номера max и min (все цифры различны)\n")
	var n int
	fmt.Scan(&n)

	maxDigit, minDigit := -1, 10
	maxPosFromEnd, minPosFromEnd := 0, 0
	maxPosFromStart, minPosFromStart := 0, 0

	temp := n
	positionFromEnd := 1
	totalDigits := 0

	temp2 := n
	for temp2 > 0 {
		totalDigits++
		temp2 /= 10
	}

	temp = n
	for temp > 0 {
		digit := temp % 10

		if digit > maxDigit {
			maxDigit = digit
			maxPosFromEnd = positionFromEnd
			maxPosFromStart = totalDigits - positionFromEnd + 1
		}

		if digit < minDigit {
			minDigit = digit
			minPosFromEnd = positionFromEnd
			minPosFromStart = totalDigits - positionFromEnd + 1
		}

		temp /= 10
		positionFromEnd++
	}

	fmt.Printf("а) Максимальная цифра %d:\n", maxDigit)
	fmt.Printf("   - от конца: позиция %d\n", maxPosFromEnd)
	fmt.Printf("   - от начала: позиция %d\n", maxPosFromStart)

	fmt.Printf("б) Минимальная цифра %d:\n", minDigit)
	fmt.Printf("   - от конца: позиция %d\n", minPosFromEnd)
	fmt.Printf("   - от начала: позиция %d\n", minPosFromStart)
}

func ch6ex46() {
	fmt.Printf("Ex46. Порядковые номера max и min\n")
	var n int
	fmt.Scan(&n)

	maxDigit, minDigit := -1, 10
	maxPosFromEnd, minPosFromEnd := 0, 0
	maxPosFromStart, minPosFromStart := 0, 0

	temp := n
	positionFromEnd := 1
	totalDigits := 0

	temp2 := n
	for temp2 > 0 {
		totalDigits++
		temp2 /= 10
	}

	temp = n
	for temp > 0 {
		digit := temp % 10

		if digit > maxDigit {
			maxDigit = digit
			maxPosFromEnd = positionFromEnd
			maxPosFromStart = totalDigits - positionFromEnd + 1
		}

		if digit < minDigit {
			minDigit = digit
			minPosFromEnd = positionFromEnd
			minPosFromStart = totalDigits - positionFromEnd + 1
		}

		temp /= 10
		positionFromEnd++
	}

	fmt.Printf("а) От конца числа:\n")
	fmt.Printf("   Максимальная цифра %d на позиции %d\n", maxDigit, maxPosFromEnd)
	fmt.Printf("   Минимальная цифра %d на позиции %d\n", minDigit, minPosFromEnd)

	fmt.Printf("б) От начала числа:\n")
	fmt.Printf("   Максимальная цифра %d на позиции %d\n", maxDigit, maxPosFromStart)
	fmt.Printf("   Минимальная цифра %d на позиции %d\n", minDigit, minPosFromStart)
}

func ch6ex47() {
	fmt.Printf("Ex47. Что левее: max или min\n")
	var n int
	fmt.Scan(&n)

	maxDigit, minDigit := -1, 10
	maxPosFromStart, minPosFromStart := 0, 0

	temp := n
	positionFromEnd := 1
	totalDigits := 0

	temp2 := n
	for temp2 > 0 {
		totalDigits++
		temp2 /= 10
	}

	temp = n
	for temp > 0 {
		digit := temp % 10
		positionFromStart := totalDigits - positionFromEnd + 1

		if digit > maxDigit {
			maxDigit = digit
			maxPosFromStart = positionFromStart
		}

		if digit < minDigit {
			minDigit = digit
			minPosFromStart = positionFromStart
		}

		temp /= 10
		positionFromEnd++
	}

	if maxPosFromStart < minPosFromStart {
		fmt.Printf("Максимальная цифра %d левее минимальной %d\n", maxDigit, minDigit)
	} else {
		fmt.Printf("Минимальная цифра %d левее максимальной %d\n", minDigit, maxDigit)
	}
}

func ch6ex48() {
	fmt.Printf("Ex48. Макс нечетная цифра и номер мин цифры слева\n")
	var n int
	fmt.Scan(&n)

	maxOdd := -1
	minDigit := 10
	minPosFromStart := 0

	temp := n
	positionFromEnd := 1
	totalDigits := 0

	temp2 := n
	for temp2 > 0 {
		totalDigits++
		temp2 /= 10
	}

	temp = n
	for temp > 0 {
		digit := temp % 10
		positionFromStart := totalDigits - positionFromEnd + 1

		if digit%2 == 1 && digit > maxOdd {
			maxOdd = digit
		}

		if digit < minDigit {
			minDigit = digit
			minPosFromStart = positionFromStart
		}

		temp /= 10
		positionFromEnd++
	}

	if maxOdd == -1 {
		fmt.Printf("а) Нет нечетных цифр\n")
	} else {
		fmt.Printf("а) Максимальная нечетная цифра: %d\n", maxOdd)
	}
	fmt.Printf("б) Минимальная цифра %d на позиции %d (слева направо)\n", minDigit, minPosFromStart)
}

func ch6ex49() {
	fmt.Printf("Ex49. Проверки числа\n")
	var n, b int
	fmt.Scan(&n, &b)

	sum := 0
	product := 1
	count := 0
	firstDigit := 0
	lastDigit := n % 10

	temp := n
	for temp > 0 {
		digit := temp % 10

		sum += digit
		product *= digit
		count++

		if temp < 10 {
			firstDigit = temp
		}

		temp /= 10
	}

	fmt.Printf("а) Сумма цифр > %d: ", b)
	fmt.Println(sum > b)

	fmt.Printf("б) Произведение цифр < 50: ")
	fmt.Println(product < 50)

	fmt.Printf("в) Количество цифр четное: ")
	fmt.Println(count%2 == 0)

	fmt.Printf("г) Число четырехзначное: ")
	fmt.Println(n >= 1000 && n <= 9999)

	fmt.Printf("д) Первая цифра ≤ 5: ")
	fmt.Println(firstDigit <= 5)

	fmt.Printf("е) Начинается и заканчивается одной цифрой: ")
	fmt.Println(firstDigit == lastDigit)

	fmt.Printf("ж) Первая цифра > последней: ")
	fmt.Println(firstDigit > lastDigit)
}

func ch6ex50() {
	fmt.Printf("Ex50. Проверки числа 2\n")
	var n, a, b, k int
	fmt.Scan(&n, &a, &b, &k)

	sum := 0
	product := 1
	count := 0
	firstDigit := 0

	temp := n
	for temp > 0 {
		digit := temp % 10

		sum += digit
		product *= digit
		count++

		if temp < 10 {
			firstDigit = temp
		}

		temp /= 10
	}

	fmt.Printf("а) Сумма цифр < %d: ", a)
	fmt.Println(sum < a)

	fmt.Printf("б) Произведение цифр > %d: ", b)
	fmt.Println(product > b)

	fmt.Printf("в) Число %d-значное: ", k)
	lower := 1
	for i := 1; i < k; i++ {
		lower *= 10
	}
	upper := lower * 10
	fmt.Println(n >= lower && n < upper)

	fmt.Printf("г) Первая цифра > %d: ", k)
	fmt.Println(firstDigit > k)
}

func ch6ex51() {
	fmt.Printf("Ex51. Комбинированные проверки\n")
	var n, k, b, x, y, a, m int
	fmt.Scan(&n, &k, &b, &x, &y, &a, &m)

	sum := 0
	product := 1
	count := 0
	firstDigit := 0
	lastDigit := n % 10

	temp := n
	for temp > 0 {
		digit := temp % 10

		sum += digit
		product *= digit
		count++

		if temp < 10 {
			firstDigit = temp
		}

		temp /= 10
	}

	resultA := sum > k && n%2 == 0
	fmt.Printf("а) Сумма цифр > %d и число четное: ", k)
	fmt.Println(resultA)

	resultB := count%2 == 0 && n <= b
	fmt.Printf("б) Количество цифр четное и число ≤ %d: ", b)
	fmt.Println(resultB)

	resultC := firstDigit == x && lastDigit == y
	fmt.Printf("в) Начинается на %d и заканчивается на %d: ", x, y)
	fmt.Println(resultC)

	resultD := product < a && n%b == 0
	fmt.Printf("г) Произведение цифр < %d и число делится на %d: ", a, b)
	fmt.Println(resultD)

	resultE := sum > m && n%k == 0
	fmt.Printf("д) Сумма цифр > %d и число делится на %d: ", m, k)
	fmt.Println(resultE)
}
