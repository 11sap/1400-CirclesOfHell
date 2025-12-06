package chapters

import (
	"fmt"
	"math"
)

func Run07() {
	tasks := []func(){
		ch7ex01, ch7ex02, ch7ex03, ch7ex04, ch7ex05, ch7ex06, ch7ex07, ch7ex08,
		ch7ex09, ch7ex10, ch7ex11, ch7ex12, ch7ex13, ch7ex14, ch7ex15, ch7ex16,
		ch7ex17, ch7ex18, ch7ex19, ch7ex20, ch7ex21, ch7ex22, ch7ex23, ch7ex24,
		ch7ex25, ch7ex26, ch7ex27, ch7ex28, ch7ex29, ch7ex30, ch7ex31, ch7ex32,
		ch7ex33, ch7ex34, ch7ex35, ch7ex36, ch7ex37, ch7ex38, ch7ex39, ch7ex40,
		ch7ex41, ch7ex42, ch7ex43, ch7ex44, ch7ex45, ch7ex46, ch7ex47, ch7ex48,
		ch7ex49, ch7ex50, ch7ex51, ch7ex52, ch7ex53, ch7ex54, ch7ex55, ch7ex56,
		ch7ex57, ch7ex58, ch7ex59, ch7ex60, ch7ex61, ch7ex62, ch7ex63, ch7ex64,
		ch7ex65, ch7ex66, ch7ex67, ch7ex68, ch7ex69, ch7ex70, ch7ex71, ch7ex72,
		ch7ex73, ch7ex74, ch7ex75, ch7ex76, ch7ex77, ch7ex78, ch7ex79, ch7ex80,
		ch7ex81, ch7ex82, ch7ex83, ch7ex84, ch7ex85, ch7ex86, ch7ex87, ch7ex88,
		ch7ex89, ch7ex90, ch7ex91, ch7ex92, ch7ex93, ch7ex94, ch7ex95, ch7ex96,
		ch7ex97, ch7ex98, ch7ex99, ch7ex100, ch7ex101, ch7ex102, ch7ex103, ch7ex104,
		ch7ex105, ch7ex106, ch7ex107, ch7ex108, ch7ex109, ch7ex110, ch7ex111, ch7ex112,
		ch7ex113, ch7ex114, ch7ex115, ch7ex116, ch7ex117, ch7ex118, ch7ex119, ch7ex120,
		ch7ex121, ch7ex122, ch7ex123, ch7ex124, ch7ex125, ch7ex126, ch7ex127, ch7ex128,
		ch7ex129, ch7ex130, ch7ex131, ch7ex132, ch7ex132, ch7ex133, ch7ex134, ch7ex134,
		ch7ex135, ch7ex136, ch7ex137, ch7ex138, ch7ex139, ch7ex140, ch7ex141, ch7ex142,
		ch7ex143, ch7ex144, ch7ex145, ch7ex146, ch7ex147, ch7ex148, ch7ex149, ch7ex150,
		ch7ex151, ch7ex152, ch7ex153, ch7ex154, ch7ex155,
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

func ch7ex19() {
	fmt.Printf("Ex19. Осадки за февраль превысили прошлогодние?\n")
	days := 28
	sumCurrent, sumLast := 0.0, 0.0

	fmt.Printf("Введите осадки за каждый день февраля (текущий год):\n")
	for i := 1; i <= days; i++ {
		var precip float64
		fmt.Scan(&precip)
		sumCurrent += precip
	}

	fmt.Printf("Введите осадки за каждый день февраля (прошлый год):\n")
	for i := 1; i <= days; i++ {
		var precip float64
		fmt.Scan(&precip)
		sumLast += precip
	}

	fmt.Printf("Текущий год: %.1f, Прошлый год: %.1f\n", sumCurrent, sumLast)
	fmt.Printf("Текущий год превысил прошлогодний: %v\n", sumCurrent > sumLast)
}

func ch7ex20() {
	fmt.Printf("Ex20. Не превышена ли грузоподъемность?\n")
	var maxLoad float64
	fmt.Scan(&maxLoad)

	totalLoad := 0.0
	fmt.Printf("Введите массы грузов (0 для окончания):\n")
	for {
		var mass float64
		fmt.Scan(&mass)
		if mass == 0 {
			break
		}
		totalLoad += mass
	}

	fmt.Printf("Общая масса: %.1f, Грузоподъемность: %.1f\n", totalLoad, maxLoad)
	fmt.Printf("Не превышена: %v\n", totalLoad <= maxLoad)
}

func ch7ex21() {
	fmt.Printf("Ex21. Лучший результат десятиборцев\n")
	sum1, sum2 := 0.0, 0.0

	fmt.Printf("Введите результаты первого спортсмена (10 видов):\n")
	for i := 1; i <= 10; i++ {
		var score float64
		fmt.Scan(&score)
		sum1 += score
	}

	fmt.Printf("Введите результаты второго спортсмена (10 видов):\n")
	for i := 1; i <= 10; i++ {
		var score float64
		fmt.Scan(&score)
		sum2 += score
	}

	fmt.Printf("Первый спортсмен: %.1f, Второй спортсмен: %.1f\n", sum1, sum2)
	if sum1 > sum2 {
		fmt.Printf("Лучший результат у первого спортсмена\n")
	} else if sum2 > sum1 {
		fmt.Printf("Лучший результат у второго спортсмена\n")
	} else {
		fmt.Printf("Результаты равны\n")
	}
}

func ch7ex22() {
	fmt.Printf("Ex22. Какой набор предметов дешевле?\n")
	sum1, sum2 := 0.0, 0.0

	fmt.Printf("Введите стоимость 8 предметов первого набора:\n")
	for i := 1; i <= 8; i++ {
		var price float64
		fmt.Scan(&price)
		sum1 += price
	}

	fmt.Printf("Введите стоимость 8 предметов второго набора:\n")
	for i := 1; i <= 8; i++ {
		var price float64
		fmt.Scan(&price)
		sum2 += price
	}

	fmt.Printf("Первый набор: %.2f, Второй набор: %.2f\n", sum1, sum2)
	if sum1 < sum2 {
		fmt.Printf("Первый набор дешевле\n")
	} else if sum2 < sum1 {
		fmt.Printf("Второй набор дешевле\n")
	} else {
		fmt.Printf("Наборы стоят одинаково\n")
	}
}

func ch7ex23() {
	fmt.Printf("Ex23. Произведение 8 чисел < 10000?\n")
	product := 1
	for i := 1; i <= 8; i++ {
		var num int
		fmt.Scan(&num)
		product *= num
	}

	fmt.Printf("Произведение: %d\n", product)
	fmt.Printf("Меньше 10000: %v\n", product < 10000)
}

func ch7ex24() {
	fmt.Printf("Ex24. Произведение n вещественных > 5?\n")
	var n int
	fmt.Scan(&n)

	product := 1.0
	for i := 1; i <= n; i++ {
		var num float64
		fmt.Scan(&num)
		product *= num
	}

	fmt.Printf("Произведение: %.2f\n", product)
	fmt.Printf("Больше 5: %v\n", product > 5)
}

func ch7ex25() {
	fmt.Printf("Ex25. Среднее арифметическое 10 чисел\n")
	sum := 0.0
	for i := 1; i <= 10; i++ {
		var num float64
		fmt.Scan(&num)
		sum += num
	}

	average := sum / 10
	fmt.Printf("Среднее арифметическое: %.2f\n", average)
}

func ch7ex26() {
	fmt.Printf("Ex26. Среднее арифметическое n вещественных\n")
	var n int
	fmt.Scan(&n)

	sum := 0.0
	for i := 1; i <= n; i++ {
		var num float64
		fmt.Scan(&num)
		sum += num
	}

	average := sum / float64(n)
	fmt.Printf("Среднее арифметическое: %.2f\n", average)
}

func ch7ex27() {
	fmt.Printf("Ex27. Средняя оценка по физике в классе (20 учеников)\n")
	sum := 0
	for i := 1; i <= 20; i++ {
		var grade int
		fmt.Scan(&grade)
		sum += grade
	}

	average := float64(sum) / 20
	fmt.Printf("Средняя оценка: %.2f\n", average)
}

func ch7ex28() {
	fmt.Printf("Ex28. Средняя оценка ученика по 10 предметам\n")
	sum := 0
	for i := 1; i <= 10; i++ {
		var grade int
		fmt.Scan(&grade)
		sum += grade
	}

	average := float64(sum) / 10
	fmt.Printf("Средняя оценка: %.2f\n", average)
}

func ch7ex29() {
	fmt.Printf("Ex29. Средняя оценка по алгебре в классе\n")
	var count int
	fmt.Scan(&count)

	sum := 0
	for i := 1; i <= count; i++ {
		var grade int
		fmt.Scan(&grade)
		sum += grade
	}

	average := float64(sum) / float64(count)
	fmt.Printf("Средняя оценка: %.2f\n", average)
}

func ch7ex30() {
	fmt.Printf("Ex30. Средняя масса предметов в наборе\n")
	var count int
	fmt.Scan(&count)

	sum := 0.0
	for i := 1; i <= count; i++ {
		var mass float64
		fmt.Scan(&mass)
		sum += mass
	}

	average := sum / float64(count)
	fmt.Printf("Средняя масса: %.2f\n", average)
}

func ch7ex31() {
	fmt.Printf("Ex31. Среднее арифметическое неотрицательных чисел до отрицательного\n")
	sum := 0.0
	count := 0

	fmt.Printf("Введите числа (отрицательное для окончания):\n")
	for {
		var num float64
		fmt.Scan(&num)
		if num < 0 {
			break
		}
		sum += num
		count++
	}

	if count > 0 {
		average := sum / float64(count)
		fmt.Printf("Среднее арифметическое: %.2f\n", average)
	} else {
		fmt.Printf("Нет чисел для вычисления среднего\n")
	}
}

func ch7ex32() {
	fmt.Printf("Ex32. Средний возраст учеников двух классов\n")
	studentsPerClass := 20

	fmt.Printf("Введите возраст учеников первого класса:\n")
	sum1 := 0.0
	for i := 1; i <= studentsPerClass; i++ {
		var age float64
		fmt.Scan(&age)
		sum1 += age
	}
	avg1 := sum1 / float64(studentsPerClass)

	fmt.Printf("Введите возраст учеников второго класса:\n")
	sum2 := 0.0
	for i := 1; i <= studentsPerClass; i++ {
		var age float64
		fmt.Scan(&age)
		sum2 += age
	}
	avg2 := sum2 / float64(studentsPerClass)

	fmt.Printf("Средний возраст в первом классе: %.1f лет\n", avg1)
	fmt.Printf("Средний возраст во втором классе: %.1f лет\n", avg2)
}

func ch7ex33() {
	fmt.Printf("Ex33. Среднедневное количество осадков за январь и март\n")
	daysJan, daysMar := 31, 31

	fmt.Printf("Введите осадки за каждый день января:\n")
	sumJan := 0.0
	for i := 1; i <= daysJan; i++ {
		var precip float64
		fmt.Scan(&precip)
		sumJan += precip
	}
	avgJan := sumJan / float64(daysJan)

	fmt.Printf("Введите осадки за каждый день марта:\n")
	sumMar := 0.0
	for i := 1; i <= daysMar; i++ {
		var precip float64
		fmt.Scan(&precip)
		sumMar += precip
	}
	avgMar := sumMar / float64(daysMar)

	fmt.Printf("Среднедневные осадки в январе: %.1f\n", avgJan)
	fmt.Printf("Среднедневные осадки в марте: %.1f\n", avgMar)
}

func ch7ex34() {
	fmt.Printf("Ex34. Средний рост учеников двух классов\n")
	var studentsCount int
	fmt.Scan(&studentsCount)

	fmt.Printf("Введите рост учеников первого класса:\n")
	sum1 := 0.0
	for i := 1; i <= studentsCount; i++ {
		var height float64
		fmt.Scan(&height)
		sum1 += height
	}
	avg1 := sum1 / float64(studentsCount)

	fmt.Printf("Введите рост учеников второго класса:\n")
	sum2 := 0.0
	for i := 1; i <= studentsCount; i++ {
		var height float64
		fmt.Scan(&height)
		sum2 += height
	}
	avg2 := sum2 / float64(studentsCount)

	fmt.Printf("Средний рост в первом классе: %.1f см\n", avg1)
	fmt.Printf("Средний рост во втором классе: %.1f см\n", avg2)
}

func ch7ex35() {
	fmt.Printf("Ex35. Средняя оценка по физике в двух классах\n")
	var studentsCount int
	fmt.Scan(&studentsCount)

	fmt.Printf("Введите оценки по физике первого класса:\n")
	sum1 := 0
	for i := 1; i <= studentsCount; i++ {
		var grade int
		fmt.Scan(&grade)
		sum1 += grade
	}
	avg1 := float64(sum1) / float64(studentsCount)

	fmt.Printf("Введите оценки по физике второго класса:\n")
	sum2 := 0
	for i := 1; i <= studentsCount; i++ {
		var grade int
		fmt.Scan(&grade)
		sum2 += grade
	}
	avg2 := float64(sum2) / float64(studentsCount)

	fmt.Printf("Средняя оценка в первом классе: %.2f\n", avg1)
	fmt.Printf("Средняя оценка во втором классе: %.2f\n", avg2)
}

func ch7ex36() {
	fmt.Printf("Ex36. Пшеница в области (10 районов)\n")
	districts := 10
	totalArea := 0.0
	totalHarvest := 0.0

	for i := 1; i <= districts; i++ {
		var area, yield float64
		fmt.Scan(&area, &yield)

		totalArea += area
		totalHarvest += area * yield
	}

	averageYield := totalHarvest / totalArea
	fmt.Printf("Общее количество пшеницы: %.1f ц\n", totalHarvest)
	fmt.Printf("Средняя урожайность по области: %.1f ц/га\n", averageYield)
}

func ch7ex37() {
	fmt.Printf("Ex37. Средняя плотность населения области (12 районов)\n")
	districts := 12
	totalPopulation := 0.0
	totalArea := 0.0

	for i := 1; i <= districts; i++ {
		var population, area float64
		fmt.Scan(&population, &area)

		totalPopulation += population
		totalArea += area
	}

	averageDensity := totalPopulation / totalArea
	fmt.Printf("Средняя плотность населения: %.1f тыс. чел/км²\n", averageDensity)
}

func ch7ex38() {
	fmt.Printf("Ex38. Общая площадь области (12 районов)\n")
	districts := 12
	totalArea := 0.0

	for i := 1; i <= districts; i++ {
		var population, density float64
		fmt.Scan(&population, &density)

		area := population / density
		totalArea += area
	}

	fmt.Printf("Общая площадь области: %.1f км²\n", totalArea)
}

func ch7ex39() {
	fmt.Printf("Ex39. Операции с n целыми числами\n")
	var n int
	fmt.Scan(&n)

	sumAbs := 0

	productAbs := 1

	sumOddPos := 0

	alternatingSum := 0
	sign := 1

	numbers := make([]int, n)

	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		numbers[i] = num

		if num < 0 {
			sumAbs -= num
		} else {
			sumAbs += num
		}

		if num < 0 {
			productAbs *= -num
		} else {
			productAbs *= num
		}

		alternatingSum += sign * num
		sign = -sign
	}

	for i := 0; i < n; i += 2 {
		sumOddPos += numbers[i]
	}

	fmt.Printf("а) Сумма модулей: %d\n", sumAbs)
	fmt.Printf("б) Произведение модулей: %d\n", productAbs)
	fmt.Printf("в) Сумма элементов на нечетных позициях: %d\n", sumOddPos)
	fmt.Printf("г) Знакопеременная сумма: %d\n", alternatingSum)
}

func ch7ex40() {
	fmt.Printf("Ex40. Сумма вещественных > 10.75\n")
	var n int
	fmt.Scan(&n)

	sum := 0.0
	for i := 1; i <= n; i++ {
		var num float64
		fmt.Scan(&num)
		if num > 10.75 {
			sum += num
		}
	}

	fmt.Printf("Сумма чисел > 10.75: %.2f\n", sum)
}

func ch7ex41() {
	fmt.Printf("Ex41. Сумма вещественных > p\n")
	var n int
	var p float64
	fmt.Scan(&n, &p)

	sum := 0.0
	for i := 1; i <= n; i++ {
		var num float64
		fmt.Scan(&num)
		if num > p {
			sum += num
		}
	}

	fmt.Printf("Сумма чисел > %.2f: %.2f\n", p, sum)
}

func ch7ex42() {
	fmt.Printf("Ex42. Сумма четных чисел\n")
	var n int
	fmt.Scan(&n)

	sum := 0
	for i := 1; i <= n; i++ {
		var num int
		fmt.Scan(&num)
		if num%2 == 0 {
			sum += num
		}
	}

	fmt.Printf("Сумма четных чисел: %d\n", sum)
}

func ch7ex43() {
	fmt.Printf("Ex43. Сумма чисел, оканчивающихся на 0\n")
	sum := 0
	for i := 1; i <= 10; i++ {
		var num int
		fmt.Scan(&num)
		if num%10 == 0 {
			sum += num
		}
	}

	fmt.Printf("Сумма чисел, оканчивающихся на 0: %d\n", sum)
}

func ch7ex44() {
	fmt.Printf("Ex44. Сумма a1 + a2 + a4 + a5 + a7 + a8 + ...\n")
	var n int
	fmt.Scan(&n)

	sum := 0
	group := 1

	for i := 1; i <= n; i++ {
		var num int
		fmt.Scan(&num)

		if group == 1 {
			sum += num
		}

		if i%3 == 0 {
			group = 1
		} else if i%3 == 2 {
			group = 2
		}
	}

	fmt.Printf("Сумма: %d\n", sum)
}

func ch7ex45() {
	fmt.Printf("Ex45. Сумма -c1 - c2 - c3 - ...\n")
	var n int
	fmt.Scan(&n)

	sum := 0.0
	for i := 1; i <= n; i++ {
		var num float64
		fmt.Scan(&num)
		sum -= num
	}

	fmt.Printf("Сумма с противоположными знаками: %.2f\n", sum)
}

func ch7ex46() {
	fmt.Printf("Ex46. Сумма a1 + a3 и a2 - a4\n")
	var n int
	fmt.Scan(&n)

	var a1, a2, a3, a4 int
	for i := 1; i <= n && i <= 4; i++ {
		var num int
		fmt.Scan(&num)

		switch i {
		case 1:
			a1 = num
		case 2:
			a2 = num
		case 3:
			a3 = num
		case 4:
			a4 = num
		}
	}

	sum13 := a1 + a3
	diff24 := a2 - a4

	fmt.Printf("а) a1 + a3 = %d\n", sum13)
	fmt.Printf("б) a2 - a4 = %d\n", diff24)
}

func ch7ex47() {
	fmt.Printf("Ex47. Стоимость товаров дороже 1000 руб\n")
	totalCost := 0.0

	fmt.Printf("Введите стоимости товаров (0 для окончания):\n")
	for {
		var price float64
		fmt.Scan(&price)
		if price == 0 {
			break
		}
		if price > 1000 {
			totalCost += price
		}
	}

	fmt.Printf("Общая стоимость товаров дороже 1000 руб: %.2f\n", totalCost)
}

func ch7ex48() {
	fmt.Printf("Ex48. Общее количество страниц во всех журналах\n")
	totalPages := 0

	fmt.Printf("Введите количество страниц в журналах (0 для окончания):\n")
	for {
		var pages int
		fmt.Scan(&pages)
		if pages == 0 {
			break
		}
		if pages <= 16 {
			totalPages += pages
		}
	}

	fmt.Printf("Общее количество страниц во всех журналах: %d\n", totalPages)
}

func ch7ex49() {
	fmt.Printf("Ex49. Осадки за четные дни месяца\n")
	days := 30
	sumEvenDays := 0.0

	for day := 1; day <= days; day++ {
		var precip float64
		fmt.Scan(&precip)
		if day%2 == 0 {
			sumEvenDays += precip
		}
	}

	fmt.Printf("Сумма осадков за четные дни: %.1f\n", sumEvenDays)
}

func ch7ex50() {
	fmt.Printf("Ex50. Дети в нечетных классах\n")
	classes := 11
	totalChildren := 0

	for class := 1; class <= classes; class++ {
		var children int
		fmt.Scan(&children)
		if class%2 == 1 {
			totalChildren += children
		}
	}

	fmt.Printf("Общее количество детей в нечетных классах: %d\n", totalChildren)
}

func ch7ex51() {
	fmt.Printf("Ex51. Сумма нечетных чисел после первого нечетного\n")
	foundFirstOdd := false
	sum := 0

	fmt.Printf("Введите последовательность (0 для окончания):\n")
	for {
		var num int
		fmt.Scan(&num)
		if num == 0 {
			break
		}

		if num%2 == 1 {
			if !foundFirstOdd {
				foundFirstOdd = true
			} else {
				sum += num
			}
		}
	}

	fmt.Printf("Сумма нечетных чисел после первого нечетного: %d\n", sum)
}

func ch7ex52() {
	fmt.Printf("Ex52. Проверки сумм\n")
	var n int
	fmt.Scan(&n)

	sumGreater20 := 0
	sumLess50 := 0

	for i := 1; i <= n; i++ {
		var num int
		fmt.Scan(&num)

		if num > 20 {
			sumGreater20 += num
		}

		if num < 50 {
			sumLess50 += num
		}
	}

	fmt.Printf("а) Сумма чисел > 20 превышает 100: %v\n", sumGreater20 > 100)
	fmt.Printf("б) Сумма чисел < 50 четна: %v\n", sumLess50%2 == 0)
}

func ch7ex53() {
	fmt.Printf("Ex53. Проверки сумм вещественных чисел\n")
	var n int
	fmt.Scan(&n)

	sumLess25_5 := 0.0
	sumNotExceed20 := 0.0

	for i := 1; i <= n; i++ {
		var num float64
		fmt.Scan(&num)

		if num < 25.5 {
			sumLess25_5 += num
		}

		if num <= 20 {
			sumNotExceed20 += num
		}
	}

	fmt.Printf("а) Сумма чисел < 25.5 не превышает 50: %v\n", sumLess25_5 <= 50)
	fmt.Printf("б) Сумма чисел ≤ 20 кратна 3: %v\n", int(sumNotExceed20)%3 == 0)
}

func ch7ex54() {
	fmt.Printf("Ex54. Сумма вещественных > 20.5 меньше p?\n")
	var n int
	var p float64
	fmt.Scan(&n, &p)

	sumGreater20_5 := 0.0

	for i := 1; i <= n; i++ {
		var num float64
		fmt.Scan(&num)

		if num > 20.5 {
			sumGreater20_5 += num
		}
	}

	fmt.Printf("Сумма чисел > 20.5 = %.2f\n", sumGreater20_5)
	fmt.Printf("Меньше %.2f: %v\n", p, sumGreater20_5 < p)
}

func ch7ex55() {
	fmt.Printf("Ex55. Сумма чисел ≤ m превышает q?\n")
	var n, m, q int
	fmt.Scan(&n, &m, &q)

	sumNotGreaterM := 0

	for i := 1; i <= n; i++ {
		var num int
		fmt.Scan(&num)

		if num <= m {
			sumNotGreaterM += num
		}
	}

	fmt.Printf("Сумма чисел ≤ %d = %d\n", m, sumNotGreaterM)
	fmt.Printf("Превышает %d: %v\n", q, sumNotGreaterM > q)
}

func ch7ex56() {
	fmt.Printf("Ex56. Сумма чисел ≤ m кратна p?\n")
	var n, m, p int
	fmt.Scan(&n, &m, &p)

	sumNotExceedM := 0

	for i := 1; i <= n; i++ {
		var num int
		fmt.Scan(&num)

		if num <= m {
			sumNotExceedM += num
		}
	}

	fmt.Printf("Сумма чисел ≤ %d = %d\n", m, sumNotExceedM)
	fmt.Printf("Кратна %d: %v\n", p, sumNotExceedM%p == 0)
}

func ch7ex57() {
	fmt.Printf("Ex57. В четные дни больше осадков, чем в нечетные?\n")
	days := 28
	sumEven, sumOdd := 0.0, 0.0

	for day := 1; day <= days; day++ {
		var precip float64
		fmt.Scan(&precip)

		if day%2 == 0 {
			sumEven += precip
		} else {
			sumOdd += precip
		}
	}

	fmt.Printf("Осадки в четные дни: %.1f, в нечетные: %.1f\n", sumEven, sumOdd)
	fmt.Printf("В четные дни больше: %v\n", sumEven > sumOdd)
}

func ch7ex58() {
	fmt.Printf("Ex58. На какой стороне улицы живет больше людей?\n")
	var houses int
	fmt.Scan(&houses)

	evenSide, oddSide := 0, 0

	for house := 1; house <= houses; house++ {
		var residents int
		fmt.Scan(&residents)

		if house%2 == 0 {
			evenSide += residents
		} else {
			oddSide += residents
		}
	}

	fmt.Printf("Людей на четной стороне: %d, на нечетной: %d\n", evenSide, oddSide)
	if evenSide > oddSide {
		fmt.Printf("Больше людей живет на четной стороне\n")
	} else if oddSide > evenSide {
		fmt.Printf("Больше людей живет на нечетной стороне\n")
	} else {
		fmt.Printf("На обеих сторонах живет одинаково людей\n")
	}
}

func ch7ex59() {
	fmt.Printf("Ex59. Количество отличников по математике\n")
	var students int
	fmt.Scan(&students)

	excellentCount := 0

	for i := 1; i <= students; i++ {
		var grade int
		fmt.Scan(&grade)
		if grade == 5 {
			excellentCount++
		}
	}

	fmt.Printf("Количество отличников: %d\n", excellentCount)
}

func ch7ex60() {
	fmt.Printf("Ex60. Сколько раз температура опускалась ниже 0°C\n")
	days := 30
	belowZeroCount := 0

	for i := 1; i <= days; i++ {
		var temp float64
		fmt.Scan(&temp)
		if temp < 0 {
			belowZeroCount++
		}
	}

	fmt.Printf("Температура опускалась ниже 0°C: %d раз\n", belowZeroCount)
}

func ch7ex61() {
	fmt.Printf("Ex61. Юноши ростом менее 165 см\n")
	boys := 12
	count := 0

	for i := 1; i <= boys; i++ {
		var height float64
		fmt.Scan(&height)
		if height < 165 {
			count++
		}
	}

	fmt.Printf("Юношей ростом менее 165 см: %d\n", count)
}

func ch7ex62() {
	fmt.Printf("Ex62. Статистика по n целым числам\n")
	var n, p, k int
	fmt.Scan(&n, &p, &k)

	countGreaterP := 0
	countEndsWith5 := 0
	countDivisibleByK := 0

	for i := 1; i <= n; i++ {
		var num int
		fmt.Scan(&num)

		if num > p {
			countGreaterP++
		}

		if num%10 == 5 {
			countEndsWith5++
		}

		if num%k == 0 {
			countDivisibleByK++
		}
	}

	fmt.Printf("а) Числа > %d: %d\n", p, countGreaterP)
	fmt.Printf("б) Числа, оканчивающиеся на 5: %d\n", countEndsWith5)
	fmt.Printf("в) Числа, кратные %d: %d\n", k, countDivisibleByK)
}

func ch7ex63() {
	fmt.Printf("Ex63. Отличники и двоечники по химии\n")
	var students int
	fmt.Scan(&students)

	excellentCount := 0
	failingCount := 0

	for i := 1; i <= students; i++ {
		var grade int
		fmt.Scan(&grade)

		if grade == 5 {
			excellentCount++
		} else if grade == 2 {
			failingCount++
		}
	}

	fmt.Printf("Отличников: %d, Двоечников: %d\n", excellentCount, failingCount)
}

func ch7ex64() {
	fmt.Printf("Ex64. Люди по годам рождения\n")
	var people int
	fmt.Scan(&people)

	before1990 := 0
	after2000 := 0

	for i := 1; i <= people; i++ {
		var year int
		fmt.Scan(&year)

		if year < 1990 {
			before1990++
		} else if year > 2000 {
			after2000++
		}
	}

	fmt.Printf("Родившихся до 1990 года: %d\n", before1990)
	fmt.Printf("Родившихся после 2000 года: %d\n", after2000)
}

func ch7ex65() {
	fmt.Printf("Ex65. Команды с побед больше чем поражений\n")
	var teams int
	fmt.Scan(&teams)

	teamsWithMoreWins := 0

	for i := 1; i <= teams; i++ {
		var wins, losses int
		fmt.Scan(&wins, &losses)

		if wins > losses {
			teamsWithMoreWins++
		}
	}

	fmt.Printf("Команд с побед больше чем поражений: %d\n", teamsWithMoreWins)
}

func ch7ex66() {
	fmt.Printf("Ex66. Количество отрицательных чисел в начале последовательности\n")
	var n int
	fmt.Scan(&n)

	count := 0
	positiveFound := false

	for i := 1; i <= n; i++ {
		var num float64
		fmt.Scan(&num)

		if !positiveFound {
			if num < 0 {
				count++
			} else {
				positiveFound = true
			}
		}
	}

	fmt.Printf("Отрицательных чисел в начале: %d\n", count)
}

func ch7ex67() {
	fmt.Printf("Ex67. Количество чисел перед первым нулем\n")
	count := 0

	fmt.Printf("Введите последовательность (0 для окончания):\n")
	for {
		var num float64
		fmt.Scan(&num)

		if num == 0 {
			break
		}

		count++
	}

	fmt.Printf("Чисел перед первым нулем: %d\n", count)
}

func ch7ex68() {
	fmt.Printf("Ex68. Количество одинаковых элементов в начале последовательности\n")
	var n int
	fmt.Scan(&n)

	var firstNum int
	count := 1

	if n > 0 {
		fmt.Scan(&firstNum)

		for i := 2; i <= n; i++ {
			var num int
			fmt.Scan(&num)

			if num == firstNum {
				count++
			} else {
				break
			}
		}
	}

	fmt.Printf("Одинаковых элементов в начале: %d\n", count)
}

func ch7ex69() {
	fmt.Printf("Ex69. Подсчет пятерок по информатике (два случая)\n")

	fmt.Println("Случай 1: Пятерки есть не у всех")
	fmt.Println("Случай 2: Пятерки могут быть у всех")

	count5 := 0
	var grade int

	fmt.Println("Введите 20 оценок по информатике:")
	for i := 0; i < 20; i++ {
		fmt.Scan(&grade)
		if grade == 5 {
			count5++
		}
	}

	fmt.Printf("Количество учеников с оценкой '5': %d\n", count5)
}

func ch7ex70() {
	fmt.Printf("Ex70. Количество дней без осадков в мае\n")

	fmt.Println("Случай 1: В какие-то дни осадки были")
	fmt.Println("Случай 2: Осадков могло не быть вообще")

	maxDaysWithoutRain := 0
	currentStreak := 0

	fmt.Println("Введите количество осадков за каждый день мая (31 день):")
	for i := 1; i <= 31; i++ {
		var precipitation float64
		fmt.Scan(&precipitation)

		if precipitation == 0 {
			currentStreak++
		} else {
			if currentStreak > maxDaysWithoutRain {
				maxDaysWithoutRain = currentStreak
			}
			currentStreak = 0
		}
	}

	if currentStreak > maxDaysWithoutRain {
		maxDaysWithoutRain = currentStreak
	}

	fmt.Printf("Максимальное количество дней подряд без осадков: %d\n", maxDaysWithoutRain)
}

func ch7ex71() {
	fmt.Printf("Ex71. Студенты, получившие двойку на экзамене\n")

	var n int
	fmt.Print("Введите количество студентов: ")
	fmt.Scan(&n)

	count2 := 0
	fmt.Printf("Введите %d оценок студентов:\n", n)
	for i := 0; i < n; i++ {
		var grade int
		fmt.Scan(&grade)
		if grade == 2 {
			count2++
		}
	}

	fmt.Printf("Количество студентов с двойкой: %d\n", count2)
}

func ch7ex72() {
	fmt.Printf("Ex72. Подсчет положительных и отрицательных чисел\n")

	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	positiveCount := 0
	negativeCount := 0

	fmt.Printf("Введите %d вещественных чисел:\n", n)
	for i := 0; i < n; i++ {
		var num float64
		fmt.Scan(&num)

		if num > 0 {
			positiveCount++
		} else if num < 0 {
			negativeCount++
		}
	}

	fmt.Printf("Положительных чисел: %d\n", positiveCount)
	fmt.Printf("Отрицательных чисел: %d\n", negativeCount)
}

func ch7ex73() {
	fmt.Printf("Ex73. Числа, кратные трем и семи\n")

	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	countMultiple3 := 0
	countMultiple7 := 0

	fmt.Printf("Введите %d целых чисел:\n", n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)

		if num%3 == 0 {
			countMultiple3++
		}
		if num%7 == 0 {
			countMultiple7++
		}
	}

	fmt.Printf("Количество чисел, кратных 3: %d\n", countMultiple3)
	fmt.Printf("Количество чисел, кратных 7: %d\n", countMultiple7)
}

func ch7ex74() {
	fmt.Printf("Ex74. Проверка троек для построения треугольника\n")

	var n int
	fmt.Print("Введите количество троек: ")
	fmt.Scan(&n)

	validCount := 0

	fmt.Printf("Введите %d троек чисел (a < b < c):\n", n)
	for i := 0; i < n; i++ {
		var a, b, c float64
		fmt.Scan(&a, &b, &c)

		if a+b > c && a+c > b && b+c > a {
			validCount++
		}
	}

	fmt.Printf("Количество троек, подходящих для треугольника: %d\n", validCount)
}

func ch7ex75() {
	fmt.Printf("Ex75. Траектория снаряда и попадание в цель\n")

	g := 9.8

	var H, R float64
	fmt.Print("Введите высоту цели H (м): ")
	fmt.Scan(&H)
	fmt.Print("Введите расстояние до цели R (м): ")
	fmt.Scan(&R)

	var alpha, v0 float64
	fmt.Print("Введите угол выстрела α (в градусах): ")
	fmt.Scan(&alpha)
	fmt.Print("Введите начальную скорость v0 (м/с): ")
	fmt.Scan(&v0)

	alphaRad := alpha * math.Pi / 180

	t := R / (v0 * math.Cos(alphaRad))

	y := v0*t*math.Sin(alphaRad) - (g*t*t)/2

	if math.Abs(y-H) <= 0.1 {
		fmt.Println("Снаряд попал в цель!")
	} else {
		fmt.Printf("Снаряд не попал. Высота в точке R: %.2f м\n", y)
	}
}

func ch7ex76() {
	fmt.Printf("Ex76. Хоккейный матч - удаления по командам\n")

	totalRemovals1 := 0
	totalTime1 := 0
	totalRemovals2 := 0
	totalTime2 := 0

	fmt.Println("Введите данные об удалениях (24 раза):")
	fmt.Println("Формат: номер_команды продолжительность")
	fmt.Println("(продолжительность: 2, 5 или 10 минут)")

	for i := 0; i < 24; i++ {
		var team, duration int
		fmt.Scan(&team, &duration)

		if team == 1 {
			totalRemovals1++
			totalTime1 += duration
		} else if team == 2 {
			totalRemovals2++
			totalTime2 += duration
		}
	}

	fmt.Println("\nРезультаты:")
	fmt.Printf("Команда 1: удалений - %d, общее время - %d мин\n", totalRemovals1, totalTime1)
	fmt.Printf("Команда 2: удалений - %d, общее время - %d мин\n", totalRemovals2, totalTime2)
}

func ch7ex77() {
	fmt.Printf("Ex77. Хоккейный матч (расширенная версия)\n")

	var team, duration int
	totalRemovals1 := 0
	totalTime1 := 0
	totalRemovals2 := 0
	totalTime2 := 0

	fmt.Println("Вводите данные об удалениях (для окончания введите 0):")
	fmt.Println("Формат: номер_команды продолжительность")

	for {
		fmt.Scan(&team)
		if team == 0 {
			break
		}
		fmt.Scan(&duration)

		if team == 1 {
			totalRemovals1++
			totalTime1 += duration
		} else if team == 2 {
			totalRemovals2++
			totalTime2 += duration
		}
	}

	fmt.Println("\nИтоговые результаты:")
	fmt.Printf("Команда 1: удалений - %d, штрафное время - %d мин\n", totalRemovals1, totalTime1)
	fmt.Printf("Команда 2: удалений - %d, штрафное время - %d мин\n", totalRemovals2, totalTime2)
}

func ch7ex78() {
	fmt.Printf("Ex78. Подсчет оценок по физике\n")

	count5 := 0
	count4 := 0
	count3 := 0
	count2 := 0

	fmt.Println("Введите оценки учеников по физике:")
	for i := 0; i < 20; i++ {
		var grade int
		fmt.Scan(&grade)

		switch grade {
		case 5:
			count5++
		case 4:
			count4++
		case 3:
			count3++
		case 2:
			count2++
		}
	}

	fmt.Println("\nРезультаты подсчета:")
	fmt.Printf("Пятерок: %d\n", count5)
	fmt.Printf("Четверок: %d\n", count4)
	fmt.Printf("Троек: %d\n", count3)
	fmt.Printf("Двоек: %d\n", count2)
}

func ch7ex79() {
	fmt.Printf("Ex79. Футбольный чемпионат - подсчет очков\n")

	var n int
	fmt.Print("Введите количество игр: ")
	fmt.Scan(&n)

	wins := 0
	losses := 0
	draws := 0

	fmt.Printf("Введите результаты %d игр (3-выигрыш, 1-ничья, 0-проигрыш):\n", n)
	for i := 0; i < n; i++ {
		var result int
		fmt.Scan(&result)

		switch result {
		case 3:
			wins++
		case 1:
			draws++
		case 0:
			losses++
		}
	}

	fmt.Println("Результаты команды:")
	fmt.Printf("Выигрышей: %d\n", wins)
	fmt.Printf("Ничьих: %d\n", draws)
	fmt.Printf("Проигрышей: %d\n", losses)
}

func ch7ex80() {
	fmt.Printf("Ex80. Футбольные игры - анализ результатов\n")

	wins := 0
	draws := 0
	losses := 0
	totalPoints := 0
	gamesDiffGreater3 := 0

	fmt.Println("Введите 20 пар чисел (забитые-пропущенные):")
	for i := 1; i <= 20; i++ {
		var scored, conceded int
		fmt.Scan(&scored, &conceded)

		var result string
		if scored > conceded {
			result = "ВЫИГРЫШ"
			wins++
		} else if scored < conceded {
			result = "ПРОИГРЫШ"
			losses++
		} else {
			result = "НИЧЬЯ"
			draws++
		}
		fmt.Printf("Игра %d: %s (%d:%d)\n", i, result, scored, conceded)

		if scored > conceded {
			totalPoints += 3
		} else if scored == conceded {
			totalPoints += 1
		}

		if math.Abs(float64(scored-conceded)) >= 3 {
			gamesDiffGreater3++
		}
	}

	fmt.Println("\nИтоговая статистика:")
	fmt.Printf("a) Побед: %d, Ничьих: %d, Поражений: %d\n", wins, draws, losses)
	fmt.Printf("б) Всего побед: %d\n", wins)
	fmt.Printf("в) Побед: %d, Поражений: %d\n", wins, losses)
	fmt.Printf("г) Побед: %d, Ничьих: %d, Поражений: %d\n", wins, draws, losses)
	fmt.Printf("д) Игр с разницей голов ≥3: %d\n", gamesDiffGreater3)
	fmt.Printf("е) Всего очков: %d\n", totalPoints)
}

func ch7ex81() {
	fmt.Printf("Ex81. Анализ результатов игр (двузначный формат)\n")

	wins := 0
	draws := 0
	losses := 0

	fmt.Println("Введите 20 двузначных чисел (десятки - забитые, единицы - пропущенные):")
	for i := 1; i <= 20; i++ {
		var result int
		fmt.Scan(&result)

		scored := result / 10
		conceded := result % 10

		if scored > conceded {
			wins++
		} else if scored < conceded {
			losses++
		} else {
			draws++
		}
	}

	fmt.Printf("Побед: %d, Ничьих: %d, Поражений: %d\n", wins, draws, losses)
}

func ch7ex82() {
	fmt.Printf("Ex82. Анализ пар последовательных чисел\n")

	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	var prev int
	fmt.Scan(&prev)

	equalPairs := 0
	zeroPairs := 0
	evenPairs := 0
	endsWith5Pairs := 0

	for i := 1; i < n; i++ {
		var current int
		fmt.Scan(&current)

		if prev == current {
			equalPairs++
		}
		if prev == 0 && current == 0 {
			zeroPairs++
		}
		if prev%2 == 0 && current%2 == 0 {
			evenPairs++
		}
		if prev%10 == 5 && current%10 == 5 {
			endsWith5Pairs++
		}

		prev = current
	}

	fmt.Println("Результаты анализа:")
	fmt.Printf("a) Равных пар: %d\n", equalPairs)
	fmt.Printf("б) Пар нулей: %d\n", zeroPairs)
	fmt.Printf("в) Четных пар: %d\n", evenPairs)
	fmt.Printf("г) Пар, оканчивающихся на 5: %d\n", endsWith5Pairs)
}

func ch7ex83() {
	fmt.Printf("Ex83. Проверка наличия четного числа\n")

	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	hasEven := false

	fmt.Printf("Введите %d чисел:\n", n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)

		if num%2 == 0 {
			hasEven = true
		}
	}

	if hasEven {
		fmt.Println("В наборе есть хотя бы одно четное число")
	} else {
		fmt.Println("В наборе нет четных чисел")
	}
}

func ch7ex84() {
	fmt.Printf("Ex84. Проверка количества положительных чисел\n")

	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	positiveCount := 0

	fmt.Printf("Введите %d целых чисел:\n", n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)

		if num > 0 {
			positiveCount++
		}
	}

	if positiveCount <= 5 {
		fmt.Println("Верно: количество положительных чисел не превышает 5")
	} else {
		fmt.Println("Неверно: количество положительных чисел больше 5")
	}
}

func ch7ex85() {
	fmt.Printf("Ex85. Проверка чисел не больше 33,2\n")

	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	count := 0

	fmt.Printf("Введите %d вещественных чисел:\n", n)
	for i := 0; i < n; i++ {
		var x float64
		fmt.Scan(&x)

		if x <= 33.2 {
			count++
		}
	}

	if count%4 == 0 {
		fmt.Printf("Верно: количество чисел ≤33.2 (%d) кратно 4\n", count)
	} else {
		fmt.Printf("Неверно: количество чисел ≤33.2 (%d) не кратно 4\n", count)
	}
}

func ch7ex86() {
	fmt.Printf("Ex86. Проверка количества чисел меньше 20\n")

	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	count := 0

	fmt.Printf("Введите %d целых чисел:\n", n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)

		if num < 20 {
			count++
		}
	}

	if count%2 == 0 {
		fmt.Printf("Верно: количество чисел <20 (%d) четное\n", count)
	} else {
		fmt.Printf("Неверно: количество чисел <20 (%d) нечетное\n", count)
	}
}

func ch7ex87() {
	fmt.Printf("Ex87. Проверка количества положительных чисел\n")

	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	positiveCount := 0

	fmt.Printf("Введите %d целых чисел:\n", n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)

		if num > 0 {
			positiveCount++
		}
	}

	if positiveCount%3 == 0 {
		fmt.Printf("Верно: количество положительных чисел (%d) кратно 3\n", positiveCount)
	} else {
		fmt.Printf("Неверно: количество положительных чисел (%d) не кратно 3\n", positiveCount)
	}
}

func ch7ex88() {
	fmt.Printf("Ex88. Проверка количества отрицательных чисел\n")

	var n, k int
	fmt.Print("Введите количество чисел n и число k: ")
	fmt.Scan(&n, &k)

	negativeCount := 0

	fmt.Printf("Введите %d целых чисел:\n", n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)

		if num < 0 {
			negativeCount++
		}
	}

	if negativeCount > k {
		fmt.Printf("Верно: количество отрицательных чисел (%d) превышает %d\n", negativeCount, k)
	} else {
		fmt.Printf("Неверно: количество отрицательных чисел (%d) не превышает %d\n", negativeCount, k)
	}
}

func ch7ex89() {
	fmt.Printf("Ex89. Проверка чисел больше m\n")

	var n, m, p int
	fmt.Print("Введите n, m и p: ")
	fmt.Scan(&n, &m, &p)

	count := 0

	fmt.Printf("Введите %d целых чисел:\n", n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)

		if num > m {
			count++
		}
	}

	if count%p == 0 {
		fmt.Printf("Верно: количество чисел >%d (%d) кратно %d\n", m, count, p)
	} else {
		fmt.Printf("Неверно: количество чисел >%d (%d) не кратно %d\n", m, count, p)
	}
}

func ch7ex90() {
	fmt.Printf("Ex90. Проверка оценок ученика\n")

	hasGrade3 := false

	fmt.Println("Введите 12 оценок ученика:")
	for i := 0; i < 12; i++ {
		var grade int
		fmt.Scan(&grade)

		if grade == 3 {
			hasGrade3 = true
		}
	}

	if !hasGrade3 {
		fmt.Println("Верно: среди оценок нет троек")
	} else {
		fmt.Println("Неверно: среди оценок есть тройки")
	}

	fmt.Println("Оператор цикла с условием можно использовать для раннего выхода при обнаружении тройки.")
}

func ch7ex91() {
	fmt.Printf("Ex91. Проверка осадков в марте\n")

	daysWithoutRain := 0

	fmt.Println("Введите количество осадков за каждый день марта (31 день):")
	for i := 1; i <= 31; i++ {
		var precipitation float64
		fmt.Scan(&precipitation)

		if precipitation == 0 {
			daysWithoutRain++
		}
	}

	if daysWithoutRain >= 10 {
		fmt.Printf("Верно: дней без осадков (%d) было не менее 10\n", daysWithoutRain)
	} else {
		fmt.Printf("Неверно: дней без осадков (%d) было менее 10\n", daysWithoutRain)
	}

	fmt.Println("Оператор цикла с условием можно использовать для подсчета дней без осадков.")
}

func ch7ex92() {
	fmt.Printf("Ex92. Проверка наличия двоек по информатике\n")

	hasGrade2 := false

	fmt.Println("Введите 28 оценок по информатике:")
	for i := 0; i < 28; i++ {
		var grade int
		fmt.Scan(&grade)

		if grade == 2 {
			hasGrade2 = true
		}
	}

	if hasGrade2 {
		fmt.Println("Да, среди оценок есть двойки")
	} else {
		fmt.Println("Нет, среди оценок нет двоек")
	}
}

func ch7ex93() {
	fmt.Printf("Ex93. Автомобили с мощностью двигателя >200 л.с.\n")

	hasPowerfulCar := false

	fmt.Println("Введите мощность двигателя 30 автомобилей (в л.с.):")
	for i := 0; i < 30; i++ {
		var power float64
		fmt.Scan(&power)

		if power > 200 {
			hasPowerfulCar = true
		}
	}

	if hasPowerfulCar {
		fmt.Println("Да, есть автомобили с мощностью >200 л.с.")
	} else {
		fmt.Println("Нет автомобилей с мощностью >200 л.с.")
	}
}

func ch7ex94() {
	fmt.Printf("Ex94. Поиск строгих локальных максимумов\n")

	var n int
	fmt.Print("Введите количество элементов: ")
	fmt.Scan(&n)

	if n < 3 {
		fmt.Println("Для поиска локальных максимумов нужно минимум 3 элемента")
		return
	}

	count := 0

	var prev, current int
	fmt.Scan(&prev)
	fmt.Scan(&current)

	for i := 2; i < n; i++ {
		var next int
		fmt.Scan(&next)

		if current > prev && current > next {
			count++
		}

		prev = current
		current = next
	}

	fmt.Printf("Количество строгих локальных максимумов: %d\n", count)
}

func ch7ex95() {
	fmt.Printf("Ex95. Подсчет смен знаков в последовательности\n")

	var n int
	fmt.Print("Введите количество ненулевых чисел: ")
	fmt.Scan(&n)

	if n < 2 {
		fmt.Println("Для подсчета смен знака нужно минимум 2 числа")
		return
	}

	signChanges := 0

	var prev int
	fmt.Scan(&prev)

	for i := 1; i < n; i++ {
		var current int
		fmt.Scan(&current)

		if (prev > 0 && current < 0) || (prev < 0 && current > 0) {
			signChanges++
		}

		prev = current
	}

	fmt.Printf("Количество смен знака: %d\n", signChanges)
}

func ch7ex96() {
	fmt.Printf("Ex96. Поиск пары одинаковых соседних чисел\n")

	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	if n < 2 {
		fmt.Println("Для поиска пары нужно минимум 2 числа")
		return
	}

	found := false
	var firstNum, secondNum int

	var prev int
	fmt.Scan(&prev)

	for i := 2; i <= n; i++ {
		var current int
		fmt.Scan(&current)

		if !found && prev == current {
			found = true
			firstNum = i - 1
			secondNum = i
		}

		prev = current
	}

	if found {
		fmt.Printf("Да, найдена пара одинаковых чисел на позициях %d и %d\n", firstNum, secondNum)
	} else {
		fmt.Println("Нет пар одинаковых соседних чисел")
	}
}

func ch7ex97() {
	fmt.Printf("Ex97. Поиск пары одинаковых соседних чисел (окончание -1)\n")

	found := false
	var firstNum, secondNum int
	position := 1

	var prev, current int
	fmt.Print("Введите последовательность (оканчивается -1): ")
	fmt.Scan(&prev)

	if prev == -1 {
		fmt.Println("Последовательность пуста")
		return
	}

	for {
		fmt.Scan(&current)
		position++

		if current == -1 {
			break
		}

		if !found && prev == current {
			found = true
			firstNum = position - 1
			secondNum = position
		}

		prev = current
	}

	if found {
		fmt.Printf("Да, найдена пара одинаковых чисел на позициях %d и %d\n", firstNum, secondNum)
	} else {
		fmt.Println("Нет пар одинаковых соседних чисел")
	}
}

func ch7ex98() {
	fmt.Printf("Ex98. Поиск пары соседних нечетных чисел\n")

	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	if n < 2 {
		fmt.Println("Для поиска пары нужно минимум 2 числа")
		return
	}

	found := false
	var firstNum, secondNum int

	var prev int
	fmt.Scan(&prev)

	for i := 2; i <= n; i++ {
		var current int
		fmt.Scan(&current)

		if !found && prev%2 != 0 && current%2 != 0 {
			found = true
			firstNum = i - 1
			secondNum = i
		}

		prev = current
	}

	if found {
		fmt.Printf("Да, найдена пара нечетных чисел на позициях %d и %d\n", firstNum, secondNum)
	} else {
		fmt.Println("Нет пар соседних нечетных чисел")
	}
}

func ch7ex99() {
	fmt.Printf("Ex99. Поиск пары соседних четных чисел (окончание 9999)\n")

	found := false
	var firstNum, secondNum int
	position := 1

	var prev, current int
	fmt.Print("Введите последовательность (оканчивается 9999): ")
	fmt.Scan(&prev)

	if prev == 9999 {
		fmt.Println("Последовательность пуста")
		return
	}

	for {
		fmt.Scan(&current)
		position++

		if current == 9999 {
			break
		}

		if !found && prev%2 == 0 && current%2 == 0 {
			found = true
			firstNum = position - 1
			secondNum = position
		}

		prev = current
	}

	if found {
		fmt.Printf("Да, найдена пара четных чисел на позициях %d и %d\n", firstNum, secondNum)
	} else {
		fmt.Println("Нет пар соседних четных чисел")
	}
}

func ch7ex100() {
	fmt.Printf("Ex100. Проверка упорядоченности по возрастанию\n")

	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	if n < 2 {
		fmt.Println("Последовательность упорядочена")
		return
	}

	var prev, current float64
	fmt.Scan(&prev)

	for i := 2; i <= n; i++ {
		fmt.Scan(&current)

		if current <= prev {
			fmt.Printf("Не упорядочено. Первое нарушение на позиции %d\n", i)
			return
		}

		prev = current
	}

	fmt.Println("Последовательность упорядочена по возрастанию")
}

func ch7ex101() {
	fmt.Printf("Ex101. Проверка упорядоченности (окончание 10000)\n")

	position := 1

	var prev, current float64
	fmt.Print("Введите последовательность (оканчивается 10000): ")
	fmt.Scan(&prev)

	if prev == 10000 {
		fmt.Println("Последовательность пуста")
		return
	}

	for {
		fmt.Scan(&current)
		position++

		if current == 10000 {
			break
		}

		if current <= prev {
			fmt.Printf("Не упорядочено. Первое нарушение на позиции %d\n", position)
			return
		}

		prev = current
	}

	fmt.Println("Последовательность упорядочена по возрастанию")
}

func ch7ex102() {
	fmt.Printf("Ex102. Проверка упорядоченности списка по убыванию роста\n")

	var n int
	fmt.Print("Введите количество учеников: ")
	fmt.Scan(&n)

	if n < 2 {
		fmt.Println("Список упорядочен")
		return
	}

	var prevHeight, currentHeight float64
	fmt.Scan(&prevHeight)

	for i := 2; i <= n; i++ {
		fmt.Scan(&currentHeight)

		if currentHeight >= prevHeight {
			fmt.Println("Список не упорядочен по убыванию роста")
			return
		}

		prevHeight = currentHeight
	}

	fmt.Println("Список упорядочен по убыванию роста")
}

func ch7ex103() {
	fmt.Printf("Ex103. Проверка упорядоченности команд по набранным очкам\n")

	var n int
	fmt.Print("Введите количество команд: ")
	fmt.Scan(&n)

	if n < 2 {
		fmt.Println("Список упорядочен")
		return
	}

	var prevPoints, currentPoints int
	fmt.Scan(&prevPoints)

	for i := 2; i <= n; i++ {
		fmt.Scan(&currentPoints)

		if currentPoints > prevPoints {
			fmt.Println("Список не упорядочен по занятым местам")
			return
		}

		prevPoints = currentPoints
	}

	fmt.Println("Список упорядочен по занятым местам")
}

func ch7ex104() {
	fmt.Printf("Ex104. Проверка равенства всех элементов\n")

	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	if n <= 1 {
		fmt.Println("Все элементы равны")
		return
	}

	var first int
	fmt.Scan(&first)

	for i := 2; i <= n; i++ {
		var current int
		fmt.Scan(&current)

		if current != first {
			fmt.Println("Не все элементы равны")
			return
		}
	}

	fmt.Println("Все элементы равны")
}

func ch7ex105() {
	fmt.Printf("Ex105. Проверка равенства элементов (окончание отрицательным)\n")

	allEqual := true
	var first, current int

	fmt.Print("Введите последовательность (оканчивается отрицательным числом): ")
	fmt.Scan(&first)

	if first < 0 {
		fmt.Println("Последовательность пуста")
		return
	}

	for {
		fmt.Scan(&current)

		if current < 0 {
			break
		}

		if current != first {
			allEqual = false
		}
	}

	if allEqual {
		fmt.Println("Все элементы равны")
	} else {
		fmt.Println("Не все элементы равны")
	}
}

func ch7ex106() {
	fmt.Printf("Ex106. Среднее арифметическое чисел больше 20\n")

	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	sum := 0.0
	count := 0

	fmt.Printf("Введите %d вещественных чисел:\n", n)
	for i := 0; i < n; i++ {
		var num float64
		fmt.Scan(&num)

		if num > 20 {
			sum += num
			count++
		}
	}

	if count > 0 {
		average := sum / float64(count)
		fmt.Printf("Среднее арифметическое чисел >20: %.2f\n", average)
	} else {
		fmt.Println("Нет чисел больше 20")
	}
}

func ch7ex107() {
	fmt.Printf("Ex107. Среднее арифметическое чисел больше n\n")

	var m, n int
	fmt.Print("Введите количество чисел m и число n: ")
	fmt.Scan(&m, &n)

	sum := 0
	count := 0

	fmt.Printf("Введите %d целых чисел:\n", m)
	for i := 0; i < m; i++ {
		var num int
		fmt.Scan(&num)

		if num > n {
			sum += num
			count++
		}
	}

	if count > 0 {
		average := float64(sum) / float64(count)
		fmt.Printf("Среднее арифметическое чисел >%d: %.2f\n", n, average)
	} else {
		fmt.Printf("Нет чисел больше %d\n", n)
	}
}

func ch7ex108() {
	fmt.Printf("Ex108. Среднее арифметическое чисел, кратных N\n")

	var m, n int
	fmt.Print("Введите количество чисел m и число n: ")
	fmt.Scan(&m, &n)

	sum := 0
	count := 0

	fmt.Printf("Введите %d целых чисел:\n", m)
	for i := 0; i < m; i++ {
		var num int
		fmt.Scan(&num)

		if num%n == 0 {
			sum += num
			count++
		}
	}

	if count > 0 {
		average := float64(sum) / float64(count)
		fmt.Printf("Среднее арифметическое чисел, кратных %d: %.2f\n", n, average)
	} else {
		fmt.Printf("Нет чисел, кратных %d\n", n)
	}
}

func ch7ex109() {
	fmt.Printf("Ex109. Средние арифметические четных и нечетных чисел\n")

	evenSum := 0
	evenCount := 0
	oddSum := 0
	oddCount := 0

	fmt.Println("Введите 12 целых чисел:")
	for i := 0; i < 12; i++ {
		var num int
		fmt.Scan(&num)

		if num%2 == 0 {
			evenSum += num
			evenCount++
		} else {
			oddSum += num
			oddCount++
		}
	}

	if evenCount > 0 {
		evenAvg := float64(evenSum) / float64(evenCount)
		fmt.Printf("Среднее четных чисел: %.2f\n", evenAvg)
	} else {
		fmt.Println("Нет четных чисел")
	}

	if oddCount > 0 {
		oddAvg := float64(oddSum) / float64(oddCount)
		fmt.Printf("Среднее нечетных чисел: %.2f\n", oddAvg)
	} else {
		fmt.Println("Нет нечетных чисел")
	}
}

func ch7ex110() {
	fmt.Printf("Ex110. Средний вес полных и остальных людей\n")

	var n int
	fmt.Print("Введите количество людей: ")
	fmt.Scan(&n)

	fullSum := 0.0
	fullCount := 0
	otherSum := 0.0
	otherCount := 0

	fmt.Printf("Введите вес %d человек (в кг):\n", n)
	for i := 0; i < n; i++ {
		var weight float64
		fmt.Scan(&weight)

		if weight > 100 {
			fullSum += weight
			fullCount++
		} else {
			otherSum += weight
			otherCount++
		}
	}

	if fullCount > 0 {
		fullAvg := fullSum / float64(fullCount)
		fmt.Printf("Средний вес полных людей: %.2f кг\n", fullAvg)
	} else {
		fmt.Println("Нет полных людей")
	}

	if otherCount > 0 {
		otherAvg := otherSum / float64(otherCount)
		fmt.Printf("Средний вес остальных людей: %.2f кг\n", otherAvg)
	} else {
		fmt.Println("Все люди полные")
	}
}

func ch7ex111() {
	fmt.Printf("Ex111. Средний рост мальчиков и девочек\n")

	var n int
	fmt.Print("Введите количество учеников: ")
	fmt.Scan(&n)

	boysSum := 0.0
	boysCount := 0
	girlsSum := 0.0
	girlsCount := 0

	fmt.Printf("Введите рост %d учеников (рост мальчиков - отрицательные числа):\n", n)
	for i := 0; i < n; i++ {
		var height float64
		fmt.Scan(&height)

		if height < 0 {
			boysSum += math.Abs(height)
			boysCount++
		} else {
			girlsSum += height
			girlsCount++
		}
	}

	if boysCount > 0 {
		boysAvg := boysSum / float64(boysCount)
		fmt.Printf("Средний рост мальчиков: %.2f см\n", boysAvg)
	} else {
		fmt.Println("Нет мальчиков")
	}

	if girlsCount > 0 {
		girlsAvg := girlsSum / float64(girlsCount)
		fmt.Printf("Средний рост девочек: %.2f см\n", girlsAvg)
	} else {
		fmt.Println("Нет девочек")
	}
}

func ch7ex112() {
	fmt.Printf("Ex112. Среднее чисел >10 (возможно отсутствие)\n")

	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	sum := 0.0
	count := 0

	fmt.Printf("Введите %d вещественных чисел:\n", n)
	for i := 0; i < n; i++ {
		var num float64
		fmt.Scan(&num)

		if num > 10 {
			sum += num
			count++
		}
	}

	if count > 0 {
		average := sum / float64(count)
		fmt.Printf("Среднее арифметическое чисел >10: %.2f\n", average)
	} else {
		fmt.Println("Нет чисел больше 10")
	}
}

func ch7ex113() {
	fmt.Printf("Ex113. Среднее чисел >n (возможно отсутствие)\n")

	var m, n int
	fmt.Print("Введите количество чисел m и число n: ")
	fmt.Scan(&m, &n)

	sum := 0
	count := 0

	fmt.Printf("Введите %d целых чисел:\n", m)
	for i := 0; i < m; i++ {
		var num int
		fmt.Scan(&num)

		if num > n {
			sum += num
			count++
		}
	}

	if count > 0 {
		average := float64(sum) / float64(count)
		fmt.Printf("Среднее арифметическое чисел >%d: %.2f\n", n, average)
	} else {
		fmt.Printf("Нет чисел больше %d\n", n)
	}
}

func ch7ex114() {
	fmt.Printf("Ex114. Среднее четных чисел (возможно отсутствие)\n")

	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	sum := 0
	count := 0

	fmt.Printf("Введите %d целых чисел:\n", n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)

		if num%2 == 0 {
			sum += num
			count++
		}
	}

	if count > 0 {
		average := float64(sum) / float64(count)
		fmt.Printf("Среднее арифметическое четных чисел: %.2f\n", average)
	} else {
		fmt.Println("Нет четных чисел")
	}
}

func ch7ex115() {
	fmt.Printf("Ex115. Среднее чисел, кратных n (возможно отсутствие)\n")

	var m, n int
	fmt.Print("Введите количество чисел m и число n: ")
	fmt.Scan(&m, &n)

	sum := 0
	count := 0

	fmt.Printf("Введите %d целых чисел:\n", m)
	for i := 0; i < m; i++ {
		var num int
		fmt.Scan(&num)

		if num%n == 0 {
			sum += num
			count++
		}
	}

	if count > 0 {
		average := float64(sum) / float64(count)
		fmt.Printf("Среднее арифметическое чисел, кратных %d: %.2f\n", n, average)
	} else {
		fmt.Printf("Нет чисел, кратных %d\n", n)
	}
}

func ch7ex116() {
	fmt.Printf("Ex116. Сравнение средней стоимости автомобилей и мотоциклов\n")

	var carsCount, bikesCount int
	fmt.Print("Введите количество автомобилей и мотоциклов: ")
	fmt.Scan(&carsCount, &bikesCount)

	carsSum := 0.0
	bikesSum := 0.0

	fmt.Printf("Введите стоимость %d автомобилей (в долларах):\n", carsCount)
	for i := 0; i < carsCount; i++ {
		var price float64
		fmt.Scan(&price)
		carsSum += price
	}

	fmt.Printf("Введите стоимость %d мотоциклов (в долларах):\n", bikesCount)
	for i := 0; i < bikesCount; i++ {
		var price float64
		fmt.Scan(&price)
		bikesSum += price
	}

	carsAvg := carsSum / float64(carsCount)
	bikesAvg := bikesSum / float64(bikesCount)

	fmt.Printf("Средняя стоимость автомобилей: $%.2f\n", carsAvg)
	fmt.Printf("Средняя стоимость мотоциклов: $%.2f\n", bikesAvg)

	if carsAvg > bikesAvg*3 {
		fmt.Println("Верно: средняя стоимость автомобилей более чем в 3 раза выше")
	} else {
		fmt.Println("Неверно: средняя стоимость автомобилей не более чем в 3 раза выше")
	}
}

func ch7ex117() {
	fmt.Printf("Ex117. Сравнение среднего роста мальчиков и девочек\n")

	var n int
	fmt.Print("Введите количество учеников: ")
	fmt.Scan(&n)

	boysSum := 0.0
	boysCount := 0
	girlsSum := 0.0
	girlsCount := 0

	fmt.Printf("Введите рост %d учеников (рост мальчиков - отрицательные числа):\n", n)
	for i := 0; i < n; i++ {
		var height float64
		fmt.Scan(&height)

		if height < 0 {
			boysSum += math.Abs(height)
			boysCount++
		} else {
			girlsSum += height
			girlsCount++
		}
	}

	if boysCount > 0 && girlsCount > 0 {
		boysAvg := boysSum / float64(boysCount)
		girlsAvg := girlsSum / float64(girlsCount)

		fmt.Printf("Средний рост мальчиков: %.2f см\n", boysAvg)
		fmt.Printf("Средний рост девочек: %.2f см\n", girlsAvg)

		if boysAvg > girlsAvg+10 {
			fmt.Println("Верно: средний рост мальчиков более чем на 10 см выше")
		} else {
			fmt.Println("Неверно: средний рост мальчиков не более чем на 10 см выше")
		}
	} else {
		fmt.Println("Невозможно сравнить: нет одной из групп")
	}
}

func ch7ex118() {
	fmt.Printf("Ex118. Поиск чисел, равных 10\n")

	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	lastIndex := -1
	firstIndex := -1

	fmt.Printf("Введите %d целых чисел:\n", n)
	for i := 1; i <= n; i++ {
		var num int
		fmt.Scan(&num)

		if num == 10 {
			lastIndex = i
			if firstIndex == -1 {
				firstIndex = i
			}
		}
	}

	if lastIndex != -1 {
		fmt.Printf("a) Номер последнего числа 10: %d\n", lastIndex)
	} else {
		fmt.Println("a) Число 10 не найдено")
	}

	if firstIndex != -1 {
		fmt.Printf("б) Номер первого числа 10: %d\n", firstIndex)
	} else {
		fmt.Println("б) Число 10 не найдено")
	}

	fmt.Println("Замечание: В задаче (а) нельзя использовать оператор цикла с условием,")
	fmt.Println("так как нужно обработать все элементы. В задаче (б) можно было бы")
	fmt.Println("использовать цикл с условием для раннего выхода после нахождения первого.")
}

func ch7ex119() {
	fmt.Printf("Ex119. Номер последнего числа >100\n")

	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	lastIndex := -1

	fmt.Printf("Введите %d целых чисел:\n", n)
	for i := 1; i <= n; i++ {
		var num int
		fmt.Scan(&num)

		if num > 100 {
			lastIndex = i
		}
	}

	if lastIndex != -1 {
		fmt.Printf("Номер последнего числа >100: %d\n", lastIndex)
	} else {
		fmt.Println("Числа >100 не найдены")
	}
}

func ch7ex120() {
	fmt.Printf("Ex120. Номер последнего числа, равного 25\n")

	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	lastIndex := -1

	fmt.Printf("Введите %d целых чисел:\n", n)
	for i := 1; i <= n; i++ {
		var num int
		fmt.Scan(&num)

		if num == 25 {
			lastIndex = i
		}
	}

	if lastIndex != -1 {
		fmt.Printf("Номер последнего числа 25: %d\n", lastIndex)
	} else {
		fmt.Println("Число 25 не найдено")
	}
}

func ch7ex121() {
	fmt.Printf("Ex121. Номер последнего отрицательного числа\n")

	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	lastIndex := -1

	fmt.Printf("Введите %d целых чисел:\n", n)
	for i := 1; i <= n; i++ {
		var num int
		fmt.Scan(&num)

		if num < 0 {
			lastIndex = i
		}
	}

	if lastIndex != -1 {
		fmt.Printf("Номер последнего отрицательного числа: %d\n", lastIndex)
	} else {
		fmt.Println("Отрицательные числа не найдены")
	}
}

func ch7ex122() {
	fmt.Printf("Ex122. Номер последнего числа, оканчивающегося на 12\n")

	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	lastIndex := -1

	fmt.Printf("Введите %d целых чисел:\n", n)
	for i := 1; i <= n; i++ {
		var num int
		fmt.Scan(&num)

		if num%100 == 12 {
			lastIndex = i
		}
	}

	if lastIndex != -1 {
		fmt.Printf("Номер последнего числа, оканчивающегося на 12: %d\n", lastIndex)
	} else {
		fmt.Println("Числа, оканчивающиеся на 12, не найдены")
	}
}

func ch7ex123() {
	fmt.Printf("Ex123. Определение места нового ученика в упорядоченном списке\n")

	heights := make([]int, 15)

	fmt.Println("Введите рост 15 учеников (по убыванию):")
	for i := 0; i < 15; i++ {
		fmt.Scan(&heights[i])
	}

	var newHeight int
	fmt.Print("Введите рост нового ученика: ")
	fmt.Scan(&newHeight)

	position := 1
	for i := 0; i < 15; i++ {
		if newHeight > heights[i] {
			break
		}
		position++
	}

	fmt.Printf("Новый ученик займет %d-е место в списке\n", position)
}

func ch7ex124() {
	fmt.Printf("Ex124. Определение места команды по набранным очкам\n")

	scores := make([]int, 20)

	fmt.Println("Введите количество очков 20 команд (по убыванию):")
	for i := 0; i < 20; i++ {
		fmt.Scan(&scores[i])
	}

	var targetScore int
	fmt.Print("Введите количество очков целевой команды: ")
	fmt.Scan(&targetScore)

	position := 1
	for i := 0; i < 20; i++ {
		if targetScore >= scores[i] {
			position = i + 1
			break
		}
		if i == 19 {
			position = 21
		}
	}

	fmt.Printf("Команда займет %d-е место\n", position)
}

func ch7ex125() {
	fmt.Printf("Ex125. Обработка упорядоченной последовательности\n")

	var n int
	fmt.Print("Введите количество чисел в последовательности: ")
	fmt.Scan(&n)

	if n < 2 {
		fmt.Println("Последовательность должна содержать хотя бы 2 числа")
		return
	}

	y := make([]float64, n)
	fmt.Printf("Введите %d чисел (по возрастанию):\n", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&y[i])
	}

	var target float64
	fmt.Print("Введите число N (не равное элементам последовательности): ")
	fmt.Scan(&target)

	sumLess := 0.0
	for i := 0; i < n; i++ {
		if y[i] < target {
			sumLess += y[i]
		} else {
			break
		}
	}

	fmt.Printf("a) Сумма чисел, меньших %.2f: %.2f\n", target, sumLess)

	if target < y[0] || target > y[n-1] {
		fmt.Println("б) N находится вне границ последовательности")
		return
	}

	for i := 0; i < n-1; i++ {
		if y[i] < target && target < y[i+1] {
			fmt.Printf("б) N находится между элементами %d (%.2f) и %d (%.2f)\n",
				i+1, y[i], i+2, y[i+1])
			break
		}
	}
}

func ch7ex126() {
	fmt.Printf("Ex126. Поиск первого отрицательного числа\n")

	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	firstIndex := -1

	fmt.Printf("Введите %d вещественных чисел:\n", n)
	for i := 1; i <= n; i++ {
		var num float64
		fmt.Scan(&num)

		if num < 0 && firstIndex == -1 {
			firstIndex = i
		}
	}

	if firstIndex != -1 {
		fmt.Printf("Да, есть отрицательные числа. Первое на позиции %d\n", firstIndex)
	} else {
		fmt.Println("Нет отрицательных чисел")
	}
}

func ch7ex127() {
	fmt.Printf("Ex127. Поиск числа 666 в последовательности\n")

	fmt.Println("Введите последовательность целых чисел (оканчивается 100):")

	position := 0
	firstIndex := -1

	for {
		var num int
		fmt.Scan(&num)
		position++

		if num == 100 {
			break
		}

		if num == 666 && firstIndex == -1 {
			firstIndex = position
		}
	}

	if firstIndex != -1 {
		fmt.Printf("Число 666 найдено на позиции %d\n", firstIndex)
	} else {
		fmt.Println("Число 666 не найдено")
	}
}

func ch7ex128() {
	fmt.Printf("Ex128. Поиск числа, оканчивающегося на 7\n")

	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	firstIndex := -1

	fmt.Printf("Введите %d натуральных чисел:\n", n)
	for i := 1; i <= n; i++ {
		var num int
		fmt.Scan(&num)

		if num%10 == 7 && firstIndex == -1 {
			firstIndex = i
		}
	}

	if firstIndex != -1 {
		fmt.Printf("Да, есть число, оканчивающееся на 7. Первое на позиции %d\n", firstIndex)
	} else {
		fmt.Println("Нет чисел, оканчивающихся на 7")
	}
}

func ch7ex129() {
	fmt.Printf("Ex129. Поиск числа, кратного 7\n")

	fmt.Println("Введите последовательность целых чисел (оканчивается -1):")

	position := 0
	firstIndex := -1

	for {
		var num int
		fmt.Scan(&num)
		position++

		if num == -1 {
			break
		}

		if num%7 == 0 && firstIndex == -1 {
			firstIndex = position
		}
	}

	if firstIndex != -1 {
		fmt.Printf("Да, есть число, кратное 7. Первое на позиции %d\n", firstIndex)
	} else {
		fmt.Println("Нет чисел, кратных 7")
	}
}

func ch7ex130() {
	fmt.Printf("Ex130. Подсчет повторений числа в неубывающей последовательности\n")

	numbers := make([]int, 20)

	fmt.Println("Введите 20 чисел (неубывающая последовательность):")
	for i := 0; i < 20; i++ {
		fmt.Scan(&numbers[i])
	}

	repeats := 0
	for i := 1; i < 20; i++ {
		if numbers[i] == numbers[i-1] {
			repeats++
		}
	}

	fmt.Printf("Количество повторений чисел: %d\n", repeats)
}

func ch7ex131() {
	fmt.Printf("Ex131. Подсчет повторяющихся чисел в неубывающей последовательности\n")

	numbers := make([]int, 20)

	fmt.Println("Введите 20 чисел (неубывающая последовательность):")
	for i := 0; i < 20; i++ {
		fmt.Scan(&numbers[i])
	}

	repeatedCount := 0
	i := 0
	for i < 20 {
		count := 1

		for j := i + 1; j < 20 && numbers[j] == numbers[i]; j++ {
			count++
		}

		if count > 1 {
			repeatedCount += count
		}

		i += count
	}

	fmt.Printf("Количество повторяющихся чисел: %d\n", repeatedCount)
}

func ch7ex132() {
	fmt.Printf("Ex132. Подсчет различных чисел в неубывающей последовательности\n")

	numbers := make([]int, 30)

	fmt.Println("Введите 30 чисел (неубывающая последовательность):")
	for i := 0; i < 30; i++ {
		fmt.Scan(&numbers[i])
	}

	distinctCount := 1
	for i := 1; i < 30; i++ {
		if numbers[i] != numbers[i-1] {
			distinctCount++
		}
	}

	fmt.Printf("Количество различных чисел: %d\n", distinctCount)
}

func ch7ex133() {
	fmt.Printf("Ex133. Подсчет повторяющихся чисел (окончание 1000)\n")

	fmt.Println("Введите неубывающую последовательность (оканчивается 1000):")

	repeatedCount := 0
	var prev, current int

	fmt.Scan(&prev)
	if prev == 1000 {
		fmt.Println("Последовательность пуста")
		return
	}

	for {
		fmt.Scan(&current)

		if current == 1000 {
			break
		}

		if current == prev {
			repeatedCount++
		}

		prev = current
	}

	fmt.Printf("Количество повторяющихся чисел: %d\n", repeatedCount)
}

func ch7ex134() {
	fmt.Printf("Ex134. Подсчет различных чисел (окончание 0)\n")

	fmt.Println("Введите невозрастающую последовательность (оканчивается 0):")

	distinctCount := 0
	var prev, current float64

	fmt.Scan(&prev)
	if prev == 0 {
		fmt.Println("Последовательность пуста")
		return
	}

	distinctCount = 1

	for {
		fmt.Scan(&current)

		if current == 0 {
			break
		}

		if current != prev {
			distinctCount++
		}

		prev = current
	}

	fmt.Printf("Количество различных чисел: %d\n", distinctCount)
}

func ch7ex135() {
	fmt.Printf("Ex135. Поиск максимального и минимального элементов\n")

	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Некорректное количество чисел")
		return
	}

	var num float64
	fmt.Scan(&num)

	maxVal := num
	minVal := num

	for i := 1; i < n; i++ {
		fmt.Scan(&num)
		if num > maxVal {
			maxVal = num
		}
	}

	fmt.Printf("a) Максимальное значение: %.2f\n", maxVal)

	fmt.Print("Введите те же числа еще раз: ")
	fmt.Scan(&num)
	minVal = num

	for i := 1; i < n; i++ {
		fmt.Scan(&num)
		if num < minVal {
			minVal = num
		}
	}

	fmt.Printf("б) Минимальное значение: %.2f\n", minVal)

	fmt.Print("Введите те же числа еще раз: ")
	fmt.Scan(&num)
	maxVal = num
	minVal = num

	for i := 1; i < n; i++ {
		fmt.Scan(&num)
		if num > maxVal {
			maxVal = num
		}
		if num < minVal {
			minVal = num
		}
	}

	fmt.Printf("в) Максимальное: %.2f, Минимальное: %.2f\n", maxVal, minVal)
}

func ch7ex136() {
	fmt.Printf("Ex136. Лучший результат спортивных соревнований\n")

	var n int
	fmt.Print("Введите количество спортсменов: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Некорректное количество спортсменов")
		return
	}

	var bestTime float64
	fmt.Print("Введите время первого спортсмена: ")
	fmt.Scan(&bestTime)

	fmt.Println("Лучший результат после 1-го спортсмена:", bestTime)

	for i := 2; i <= n; i++ {
		var currentTime float64
		fmt.Printf("Введите время %d-го спортсмена: ", i)
		fmt.Scan(&currentTime)

		if currentTime < bestTime {
			bestTime = currentTime
		}

		fmt.Printf("Лучший результат после %d-го спортсмена: %.2f\n", i, bestTime)
	}
}

func ch7ex137() {
	fmt.Printf("Ex137. Самое большое расстояние от Москвы\n")

	var n int
	fmt.Print("Введите количество городов: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Некорректное количество городов")
		return
	}

	maxDistance := 0.0

	fmt.Println("Введите расстояния от Москвы до городов (в км):")
	for i := 0; i < n; i++ {
		var distance float64
		fmt.Scan(&distance)

		if distance > maxDistance {
			maxDistance = distance
		}
	}

	fmt.Printf("Максимальное расстояние от Москвы: %.2f км\n", maxDistance)
}

func ch7ex138() {
	fmt.Printf("Ex138. Максимальная температура месяца\n")

	var days int
	fmt.Print("Введите количество дней в месяце: ")
	fmt.Scan(&days)

	if days <= 0 {
		fmt.Println("Некорректное количество дней")
		return
	}

	maxTemp := -100.0

	fmt.Printf("Введите температуру за %d дней:\n", days)
	for i := 0; i < days; i++ {
		var temp float64
		fmt.Scan(&temp)

		if temp > maxTemp {
			maxTemp = temp
		}
	}

	fmt.Printf("Максимальная температура месяца: %.1f°C\n", maxTemp)
}

func ch7ex139() {
	fmt.Printf("Ex139. Самый быстрый автомобиль\n")

	maxSpeed := 0.0
	carNumber := 0

	fmt.Println("Введите максимальные скорости 20 марок автомобилей:")
	for i := 1; i <= 20; i++ {
		var speed float64
		fmt.Scan(&speed)

		if speed > maxSpeed {
			maxSpeed = speed
			carNumber = i
		}
	}

	fmt.Printf("Самый быстрый автомобиль №%d со скоростью %.1f км/ч\n", carNumber, maxSpeed)
}

func ch7ex140() {
	fmt.Printf("Ex140. Наименьший радиус круга\n")

	var n int
	fmt.Print("Введите количество кругов: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Некорректное количество кругов")
		return
	}

	minRadius := math.MaxFloat64

	fmt.Println("Введите радиусы кругов:")
	for i := 0; i < n; i++ {
		var radius float64
		fmt.Scan(&radius)

		if radius < minRadius {
			minRadius = radius
		}
	}

	fmt.Printf("Наименьший радиус: %.2f\n", minRadius)
}

func ch7ex141() {
	fmt.Printf("Ex141. Длина диагонали самого большого квадрата\n")

	var n int
	fmt.Print("Введите количество квадратов: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Некорректное количество квадратов")
		return
	}

	maxSide := 0.0

	fmt.Println("Введите длины сторон квадратов:")
	for i := 0; i < n; i++ {
		var side float64
		fmt.Scan(&side)

		if side > maxSide {
			maxSide = side
		}
	}

	diagonal := maxSide * math.Sqrt(2)
	fmt.Printf("Длина диагонали самого большого квадрата: %.2f\n", diagonal)
}

func ch7ex142() {
	fmt.Printf("Ex142. Максимальная плотность материала\n")

	maxDensity := 0.0

	fmt.Println("Введите массу и объем 20 тел:")
	for i := 1; i <= 20; i++ {
		var mass, volume float64
		fmt.Scan(&mass, &volume)

		if volume > 0 {
			density := mass / volume
			if density > maxDensity {
				maxDensity = density
			}
		}
	}

	fmt.Printf("Максимальная плотность материала: %.2f кг/см³\n", maxDensity)
}

func ch7ex143() {
	fmt.Printf("Ex143. Минимальная плотность населения\n")

	minDensity := math.MaxFloat64
	countryNumber := 0

	fmt.Println("Введите население (в млн) и площадь (в тыс. км²) 16 государств:")
	for i := 1; i <= 16; i++ {
		var population, area float64
		fmt.Scan(&population, &area)

		if area > 0 {
			density := population / area
			if density < minDensity {
				minDensity = density
				countryNumber = i
			}
		}
	}

	fmt.Printf("Минимальная плотность в государстве №%d: %.2f чел/км²\n", countryNumber, minDensity)
}

func ch7ex144() {
	fmt.Printf("Ex144. Максимальная сумма и минимальное произведение в парах\n")

	var n int
	fmt.Print("Введите количество пар: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Некорректное количество пар")
		return
	}

	fmt.Println("Введите пары чисел:")
	var a1, b1 float64
	fmt.Scan(&a1, &b1)
	maxSum := a1 + b1

	for i := 1; i < n; i++ {
		var a, b float64
		fmt.Scan(&a, &b)
		sum := a + b
		if sum > maxSum {
			maxSum = sum
		}
	}

	fmt.Printf("a) Максимальная сумма: %.2f\n", maxSum)

	fmt.Println("Введите те же пары еще раз:")
	fmt.Scan(&a1, &b1)
	minProduct := a1 * b1

	for i := 1; i < n; i++ {
		var a, b float64
		fmt.Scan(&a, &b)
		product := a * b
		if product < minProduct {
			minProduct = product
		}
	}

	fmt.Printf("б) Минимальное произведение: %.2f\n", minProduct)
}

func ch7ex145() {
	fmt.Printf("Ex145. Расчет итоговой оценки в спортивных соревнованиях\n")

	var judges int
	fmt.Print("Введите количество судей: ")
	fmt.Scan(&judges)

	if judges <= 2 {
		fmt.Println("Слишком мало судей")
		return
	}

	scores := make([]float64, judges)
	fmt.Printf("Введите %d оценок:\n", judges)
	for i := 0; i < judges; i++ {
		fmt.Scan(&scores[i])
	}

	maxScore := scores[0]
	minScore := scores[0]
	maxIndex := 0
	minIndex := 0

	for i := 1; i < judges; i++ {
		if scores[i] > maxScore {
			maxScore = scores[i]
			maxIndex = i
		}
		if scores[i] < minScore {
			minScore = scores[i]
			minIndex = i
		}
	}

	sum := 0.0
	count := 0
	for i := 0; i < judges; i++ {
		if i != maxIndex && i != minIndex {
			sum += scores[i]
			count++
		}
	}

	if count > 0 {
		finalScore := sum / float64(count)
		fmt.Printf("Итоговая оценка: %.2f\n", finalScore)
	} else {
		fmt.Println("Недостаточно оценок для расчета")
	}
}

func ch7ex146() {
	fmt.Printf("Ex146. Разница между самым высоким и самым низким ростом\n")

	var n int
	fmt.Print("Введите количество людей: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Некорректное количество людей")
		return
	}

	var height float64
	fmt.Scan(&height)

	maxHeight := height
	minHeight := height

	for i := 1; i < n; i++ {
		fmt.Scan(&height)

		if height > maxHeight {
			maxHeight = height
		}
		if height < minHeight {
			minHeight = height
		}
	}

	difference := maxHeight - minHeight
	fmt.Printf("Разница между самым высоким и самым низким: %.1f см\n", difference)
}

func ch7ex147() {
	fmt.Printf("Ex147. Разница между самым большим и самым маленьким классом\n")

	maxStudents := 0
	minStudents := math.MaxInt32

	fmt.Println("Введите количество учеников в 20 классах:")
	for i := 0; i < 20; i++ {
		var students int
		fmt.Scan(&students)

		if students > maxStudents {
			maxStudents = students
		}
		if students < minStudents {
			minStudents = students
		}
	}

	difference := maxStudents - minStudents
	fmt.Printf("Разница между самым большим и самым маленьким классом: %d учеников\n", difference)
}

func ch7ex148() {
	fmt.Printf("Ex148. Проверка разницы между максимальным и минимальным\n")

	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Некорректное количество чисел")
		return
	}

	var num int
	fmt.Scan(&num)

	maxVal := num
	minVal := num

	for i := 1; i < n; i++ {
		fmt.Scan(&num)

		if num > maxVal {
			maxVal = num
		}
		if num < minVal {
			minVal = num
		}
	}

	difference := maxVal - minVal
	fmt.Printf("Максимальное: %d, Минимальное: %d, Разница: %d\n", maxVal, minVal, difference)

	if difference <= 25 {
		fmt.Println("Верно: разница не превышает 25")
	} else {
		fmt.Println("Неверно: разница превышает 25")
	}
}

func ch7ex149() {
	fmt.Printf("Ex149. Проверка соотношения массы самого тяжелого и самого легкого\n")

	var n int
	fmt.Print("Введите количество людей: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Некорректное количество людей")
		return
	}

	var mass float64
	fmt.Scan(&mass)

	maxMass := mass
	minMass := mass

	for i := 1; i < n; i++ {
		fmt.Scan(&mass)

		if mass > maxMass {
			maxMass = mass
		}
		if mass < minMass {
			minMass = mass
		}
	}

	ratio := maxMass / minMass
	fmt.Printf("Самый тяжелый: %.1f кг, Самый легкий: %.1f кг, Соотношение: %.2f\n",
		maxMass, minMass, ratio)

	if ratio > 2 {
		fmt.Println("Верно: самый тяжелый более чем в 2 раза тяжелее самого легкого")
	} else {
		fmt.Println("Неверно: самый тяжелый не более чем в 2 раза тяжелее самого легкого")
	}
}

func ch7ex150() {
	fmt.Printf("Ex150. Наибольшая длина отрезка из четных чисел\n")

	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Некорректное количество чисел")
		return
	}

	maxLength := 0
	currentLength := 0

	fmt.Printf("Введите %d целых чисел:\n", n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)

		if num%2 == 0 {
			currentLength++
			if currentLength > maxLength {
				maxLength = currentLength
			}
		} else {
			currentLength = 0
		}
	}

	fmt.Printf("Наибольшая длина отрезка из четных чисел: %d\n", maxLength)
}

func ch7ex151() {
	fmt.Printf("Ex151. Наименьшая длина отрезка из нулей\n")

	var n int
	fmt.Print("Введите длину последовательности: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Некорректная длина")
		return
	}

	minLength := math.MaxInt32
	currentLength := 0

	fmt.Printf("Введите %d чисел (0 или 1):\n", n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)

		if num == 0 {
			currentLength++
		} else {
			if currentLength > 0 && currentLength < minLength {
				minLength = currentLength
			}
			currentLength = 0
		}
	}

	if currentLength > 0 && currentLength < minLength {
		minLength = currentLength
	}

	if minLength == math.MaxInt32 {
		fmt.Println("Нет отрезков из нулей")
	} else {
		fmt.Printf("Наименьшая длина отрезка из нулей: %d\n", minLength)
	}
}

func ch7ex152() {
	fmt.Printf("Ex152. Поиск ближайшего элемента к заданному числу\n")

	var n int
	fmt.Print("Введите количество чисел в упорядоченной последовательности: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Некорректное количество чисел")
		return
	}

	sequence := make([]float64, n)
	fmt.Printf("Введите %d чисел (по возрастанию):\n", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&sequence[i])
	}

	var target float64
	fmt.Print("Введите число N: ")
	fmt.Scan(&target)

	closestIndex := 0
	minDiff := math.Abs(sequence[0] - target)

	for i := 1; i < n; i++ {
		diff := math.Abs(sequence[i] - target)
		if diff < minDiff {
			minDiff = diff
			closestIndex = i
		}
	}

	fmt.Printf("Ближайший элемент: позиция %d, значение %.2f\n",
		closestIndex+1, sequence[closestIndex])
}

func ch7ex153() {
	fmt.Printf("Ex153. Максимальное четное число\n")

	maxEven := -1

	fmt.Println("Введите 14 целых чисел:")
	for i := 0; i < 14; i++ {
		var num int
		fmt.Scan(&num)

		if num%2 == 0 && num > maxEven {
			maxEven = num
		}
	}

	if maxEven != -1 {
		fmt.Printf("Максимальное четное число: %d\n", maxEven)
	} else {
		fmt.Println("Четных чисел нет")
	}
}

func ch7ex154() {
	fmt.Printf("Ex154. Наибольшее количество подряд идущих одинаковых элементов\n")

	var n int
	fmt.Print("Введите длину последовательности: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Некорректная длина")
		return
	}

	var prev, current int
	fmt.Scan(&prev)

	maxCount := 1
	currentCount := 1

	for i := 1; i < n; i++ {
		fmt.Scan(&current)

		if current == prev {
			currentCount++
			if currentCount > maxCount {
				maxCount = currentCount
			}
		} else {
			currentCount = 1
		}

		prev = current
	}

	fmt.Printf("Наибольшее количество подряд идущих одинаковых элементов: %d\n", maxCount)
}

func ch7ex155() {
	fmt.Printf("Ex155. Наибольшая длина монотонно возрастающего фрагмента\n")

	var n int
	fmt.Print("Введите длину последовательности: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Некорректная длина")
		return
	}

	var prev, current int
	fmt.Scan(&prev)

	maxLength := 1
	currentLength := 1

	for i := 1; i < n; i++ {
		fmt.Scan(&current)

		if current > prev {
			currentLength++
			if currentLength > maxLength {
				maxLength = currentLength
			}
		} else {
			currentLength = 1
		}

		prev = current
	}

	fmt.Printf("Наибольшая длина монотонно возрастающего фрагмента: %d\n", maxLength)
}
