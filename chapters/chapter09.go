package chapters

import (
	"fmt"
	"sort"
)

func Run09() {
	tasks := []func(){
		ch9ex01, ch9ex02, ch9ex03, ch9ex04, ch9ex05, ch9ex06, ch9ex07, ch9ex08,
		ch9ex09, ch9ex10, ch9ex11, ch9ex12, ch9ex13, ch9ex14, ch9ex15, ch9ex16,
		ch9ex17, ch9ex18, ch9ex19, ch9ex20, ch9ex21, ch9ex22, ch9ex23, ch9ex24,
		ch9ex25, ch9ex26, ch9ex27, ch9ex28, ch9ex29, ch9ex30, ch9ex31, ch9ex32,
		ch9ex33, ch9ex34, ch9ex35, ch9ex36, ch9ex37, ch9ex38, ch9ex39, ch9ex40,
		ch9ex41, ch9ex42, ch9ex43, ch9ex44, ch9ex45, ch9ex46, ch9ex47, ch9ex48,
		ch9ex49, ch9ex50, ch9ex51, ch9ex52, ch9ex53, ch9ex54,
	}

	for {
		fmt.Print("\nГлава 9. Выберите задачу (1-54): ")
		var n int
		fmt.Scan(&n)

		if n == 0 {
			break
		}
		if n >= 1 && n <= 54 {
			tasks[n-1]()
		} else {
			fmt.Println("От 1 до 54! Для выхода введите 0.")
		}
	}
}

func ch9ex01() {
	fmt.Printf("Ex1. Заполнение массивов по заданным шаблонам\n\n")

	fmt.Println("а)")

	for i := 0; i < 5; i++ {
		fmt.Println("8 8 8")
	}
	fmt.Println()

	fmt.Println("б)")
	for i := 1; i <= 7; i++ {
		fmt.Printf("%d %d %d %d %d\n", i, i, i, i, i)
	}
	fmt.Println()

	fmt.Println("в)")
	for i := 10; i <= 80; i += 10 {
		fmt.Printf("%d %d %d %d\n", i, i, i, i)
	}
	fmt.Println()

	fmt.Println("г)")
	for i := 12; i <= 82; i += 10 {
		fmt.Printf("%d %d %d %d\n", i, i, i, i)
	}
	fmt.Println()

	fmt.Println("д)")
	for i := 0; i < 5; i++ {
		for j := 2; j <= 20; j++ {
			fmt.Print(j)
			if j < 20 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	fmt.Println("е)")
	for i := 0; i < 4; i++ {
		for j := 15; j >= 3; j-- {
			fmt.Print(j)
			if j <= 15 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	// я устал босс, это полный треш
}

func ch9ex02() {
	fmt.Println("Ex2. Таблица сложения")

	fmt.Println("\nа)")
	for j := 1; j <= 9; j++ {
		for i := 1; i <= 9; i++ {
			fmt.Printf("%d+%d=%2d  ", i, j, i+j)
			if i%9 == 0 {
				fmt.Println()
			}
		}
	}

	fmt.Println("\nб)")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			fmt.Printf("%d+%d=%2d  ", i, j, i+j)
			if j%9 == 0 {
				fmt.Println()
			}
		}
	}
}

func ch9ex03() {
	fmt.Println("Ex3. Таблица умножения")

	fmt.Println("\nа)")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			fmt.Printf("%d*%d=%2d  ", i, j, i*j)
			if j%9 == 0 {
				fmt.Println()
			}
		}
	}

	fmt.Println("\nб)")
	for j := 1; j <= 9; j++ {
		for i := 1; i <= 9; i++ {
			fmt.Printf("%d*%d=%2d  ", i, j, i*j)
			if i%9 == 0 {
				fmt.Println()
			}
		}
	}
}

func ch9ex04() {
	fmt.Println("Ex4. Оценки 12 учеников по 3 предметам")

	fmt.Println("Вариант 1 - ввод по строкам:")
	var marks1 [12][3]int
	sum1 := 0

	for i := 0; i < 12; i++ {
		fmt.Printf("Введите 3 оценки для ученика %d: ", i+1)
		for j := 0; j < 3; j++ {
			fmt.Scan(&marks1[i][j])
			sum1 += marks1[i][j]
		}
	}
	fmt.Printf("Сумма всех оценок: %d\n\n", sum1)

	fmt.Println("Вариант 2 - ввод по столбцам:")
	var marks2 [12][3]int
	sum2 := 0

	for j := 0; j < 3; j++ {
		fmt.Printf("Введите оценки для предмета %d (12 учеников): ", j+1)
		for i := 0; i < 12; i++ {
			fmt.Scan(&marks2[i][j])
			sum2 += marks2[i][j]
		}
	}
	fmt.Printf("Сумма всех оценок: %d\n", sum2)
}

func ch9ex05() {
	fmt.Println("Ex5. Зарплата работников за квартал")

	var salary [12][3]float64
	totalQuarter := 0.0
	var monthTotal [3]float64
	var workerTotal [12]float64

	fmt.Println("Введите зарплату для 12 работников за 3 месяца:")
	for i := 0; i < 12; i++ {
		fmt.Printf("Работник %d: ", i+1)
		for j := 0; j < 3; j++ {
			fmt.Scan(&salary[i][j])
			workerTotal[i] += salary[i][j]
			monthTotal[j] += salary[i][j]
			totalQuarter += salary[i][j]
		}
	}

	fmt.Printf("а) Общая сумма за квартал: %.2f\n", totalQuarter)

	fmt.Println("б) Зарплата каждого работника за квартал:")
	for i := 0; i < 12; i++ {
		fmt.Printf("  Работник %d: %.2f\n", i+1, workerTotal[i])
	}

	fmt.Println("в) Общая зарплата за каждый месяц:")
	for j := 0; j < 3; j++ {
		fmt.Printf("  Месяц %d: %.2f\n", j+1, monthTotal[j])
	}
}

func ch9ex06() {
	fmt.Println("Ex6. Результаты спортсменов")

	var results [15][3]float64

	fmt.Println("Введите баллы для 15 спортсменов по 3 программам:")
	for i := 0; i < 15; i++ {
		fmt.Printf("Спортсмен %d: ", i+1)
		for j := 0; j < 3; j++ {
			fmt.Scan(&results[i][j])
		}
	}

	fmt.Println("а) Средний балл каждого спортсмена:")
	for i := 0; i < 15; i++ {
		sum := 0.0
		for j := 0; j < 3; j++ {
			sum += results[i][j]
		}
		average := sum / 3.0
		fmt.Printf("  Спортсмен %d: %.2f\n", i+1, average)
	}

	fmt.Println("б) Средний балл по каждой программе:")
	for j := 0; j < 3; j++ {
		sum := 0.0
		for i := 0; i < 15; i++ {
			sum += results[i][j]
		}
		average := sum / 15.0
		programName := ""
		switch j {
		case 0:
			programName = "Обязательная"
		case 1:
			programName = "Короткая"
		case 2:
			programName = "Произвольная"
		}
		fmt.Printf("  %s программа: %.2f\n", programName, average)
	}
}

func ch9ex07() {
	fmt.Println("Ex7. Оценки учеников")

	var marks [15][3]int

	fmt.Println("Введите оценки для 15 учеников по 3 предметам (2-5):")
	for i := 0; i < 15; i++ {
		fmt.Printf("Ученик %d: ", i+1)
		for j := 0; j < 3; j++ {
			fmt.Scan(&marks[i][j])
		}
	}

	fives := 0
	for i := 0; i < 15; i++ {
		for j := 0; j < 3; j++ {
			if marks[i][j] == 5 {
				fives++
			}
		}
	}
	fmt.Printf("а) Общее количество пятерок: %d\n", fives)

	fmt.Println("б) Количество двоек у каждого ученика:")
	for i := 0; i < 15; i++ {
		twos := 0
		for j := 0; j < 3; j++ {
			if marks[i][j] == 2 {
				twos++
			}
		}
		fmt.Printf("  Ученик %d: %d двоек\n", i+1, twos)
	}

	fmt.Println("в) Количество двоек по каждому предмету:")
	for j := 0; j < 3; j++ {
		twos := 0
		for i := 0; i < 15; i++ {
			if marks[i][j] == 2 {
				twos++
			}
		}
		fmt.Printf("  Предмет %d: %d двоек\n", j+1, twos)
	}
}

func ch9ex08() {
	fmt.Println("Ex8. Оценки студентов в сессию")

	var marks [14][3]int

	fmt.Println("Введите оценки для 14 студентов по 3 предметам (2-5):")
	for i := 0; i < 14; i++ {
		fmt.Printf("Студент %d: ", i+1)
		for j := 0; j < 3; j++ {
			fmt.Scan(&marks[i][j])
		}
	}

	noTwos := 0
	for i := 0; i < 14; i++ {
		hasTwo := false
		for j := 0; j < 3; j++ {
			if marks[i][j] == 2 {
				hasTwo = true
				break
			}
		}
		if !hasTwo {
			noTwos++
		}
	}
	fmt.Printf("а) Студентов без двоек: %d\n", noTwos)

	goodSubjects := 0
	for j := 0; j < 3; j++ {
		onlyGood := true
		for i := 0; i < 14; i++ {
			if marks[i][j] != 5 && marks[i][j] != 4 && marks[i][j] != 6 {
				onlyGood = false
				break
			}
		}
		if onlyGood {
			goodSubjects++
		}
	}
	fmt.Printf("б) Предметов только с 5+ и 4+: %d\n", goodSubjects)

	fmt.Println("в) Количество двоек по каждому предмету:")
	for j := 0; j < 3; j++ {
		twos := 0
		for i := 0; i < 14; i++ {
			if marks[i][j] == 2 {
				twos++
			}
		}
		fmt.Printf("  Предмет %d: %d двоек\n", j+1, twos)
	}
}

func ch9ex09() {
	fmt.Println("Ex9. Баллы спортсменов по пятиборью")

	var scores [8][5]int

	fmt.Println("Введите баллы для 8 спортсменов по 5 видам спорта:")
	for i := 0; i < 8; i++ {
		fmt.Printf("Спортсмен %d (5 видов): ", i+1)
		for j := 0; j < 5; j++ {
			fmt.Scan(&scores[i][j])
		}
	}

	maxScore := scores[0][0]
	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			if scores[i][j] > maxScore {
				maxScore = scores[i][j]
			}
		}
	}
	fmt.Printf("а) Максимальная оценка: %d\n", maxScore)

	winnerIndex := 0
	maxTotal := 0

	for i := 0; i < 8; i++ {
		total := 0
		for j := 0; j < 5; j++ {
			total += scores[i][j]
		}
		if total > maxTotal {
			maxTotal = total
			winnerIndex = i
		}
	}
	fmt.Printf("б) Победитель (спортсмен %d) набрал: %d баллов\n",
		winnerIndex+1, maxTotal)
}

func ch9ex10() {
	fmt.Println("Ex10. Зарплата работников за квартал")

	var salary [12][3]float64

	fmt.Println("Введите зарплату 12 работников за 3 месяца:")
	for i := 0; i < 12; i++ {
		fmt.Printf("Работник %d: ", i+1)
		for j := 0; j < 3; j++ {
			fmt.Scan(&salary[i][j])
		}
	}

	maxSalary := salary[0][0]
	for i := 0; i < 12; i++ {
		for j := 0; j < 3; j++ {
			if salary[i][j] > maxSalary {
				maxSalary = salary[i][j]
			}
		}
	}
	fmt.Printf("а) Максимальная зарплата: %.2f\n", maxSalary)

	maxTotal := 0.0
	workerWithMaxTotal := 0

	for i := 0; i < 12; i++ {
		total := 0.0
		for j := 0; j < 3; j++ {
			total += salary[i][j]
		}
		if total > maxTotal {
			maxTotal = total
			workerWithMaxTotal = i
		}
	}
	fmt.Printf("б) Работник %d получил наибольшую сумму: %.2f\n",
		workerWithMaxTotal+1, maxTotal)

	var monthTotal [3]float64
	for j := 0; j < 3; j++ {
		for i := 0; i < 12; i++ {
			monthTotal[j] += salary[i][j]
		}
	}

	maxMonthIndex := 0
	maxMonthTotal := monthTotal[0]
	for j := 1; j < 3; j++ {
		if monthTotal[j] > maxMonthTotal {
			maxMonthTotal = monthTotal[j]
			maxMonthIndex = j
		}
	}
	fmt.Printf("в) Максимальная общая зарплата была в месяце %d: %.2f\n",
		maxMonthIndex+1, maxMonthTotal)
}

func ch9ex11() {
	fmt.Println("Ex11. Зарплата работников за квартал (доп.)")

	var salary [12][3]float64

	fmt.Println("Введите зарплату 12 работников за 3 месяца:")
	for i := 0; i < 12; i++ {
		fmt.Printf("Работник %d: ", i+1)
		for j := 0; j < 3; j++ {
			fmt.Scan(&salary[i][j])
		}
	}

	fmt.Println("а) Месяц с наибольшей зарплатой для каждого работника:")
	for i := 0; i < 12; i++ {
		maxMonth := 0
		maxSalary := salary[i][0]

		for j := 1; j < 3; j++ {
			if salary[i][j] > maxSalary {
				maxSalary = salary[i][j]
				maxMonth = j
			}
		}
		fmt.Printf("  Работник %d: месяц %d (%.2f)\n",
			i+1, maxMonth+1, maxSalary)
	}

	fmt.Println("б) Работник с наибольшей зарплатой в каждом месяце:")
	for j := 0; j < 3; j++ {
		maxWorker := 0
		maxSalary := salary[0][j]

		for i := 1; i < 12; i++ {
			if salary[i][j] > maxSalary {
				maxSalary = salary[i][j]
				maxWorker = i
			}
		}
		fmt.Printf("  Месяц %d: работник %d (%.2f)\n",
			j+1, maxWorker+1, maxSalary)
	}
}

func ch9ex12() {
	fmt.Println("Ex12. Количество учеников в классах")

	var students [11][4]int

	data := [11][4]int{
		{23, 25, 27, 22},
		{24, 26, 25, 23},
		{22, 24, 26, 21},
		{25, 27, 24, 22},
		{23, 25, 26, 24},
		{24, 26, 25, 23},
		{22, 25, 27, 21},
		{23, 24, 26, 22},
		{25, 26, 24, 23},
		{24, 25, 26, 22},
		{20, 25, 21, 26},
	}
	students = data

	minStudents := students[0][0]
	for i := 0; i < 11; i++ {
		for j := 0; j < 4; j++ {
			if students[i][j] < minStudents {
				minStudents = students[i][j]
			}
		}
	}
	fmt.Printf("а) В самом малочисленном классе: %d учеников\n", minStudents)

	fmt.Println("б) Минимальное количество учеников в параллели:")
	minParallelTotal := 0
	for i := 0; i < 11; i++ {
		parallelTotal := 0
		for j := 0; j < 4; j++ {
			parallelTotal += students[i][j]
		}
		if i == 0 || parallelTotal < minParallelTotal {
			minParallelTotal = parallelTotal
		}
	}
	fmt.Printf("  Минимальное: %d учеников\n", minParallelTotal)

	fmt.Println("в) Минимальное количество учеников по классам (А,Б,В,Г):")
	var classTotal [4]int
	for j := 0; j < 4; j++ {
		for i := 0; i < 11; i++ {
			classTotal[j] += students[i][j]
		}
	}

	minClassTotal := classTotal[0]
	for j := 1; j < 4; j++ {
		if classTotal[j] < minClassTotal {
			minClassTotal = classTotal[j]
		}
	}
	fmt.Printf("  Минимальное: %d учеников\n", minClassTotal)
}

func ch9ex13() {
	fmt.Println("Ex13. Самый малочисленный класс")

	var students [11][4]int

	data := [11][4]int{
		{23, 25, 27, 22},
		{24, 26, 25, 23},
		{22, 24, 26, 21},
		{25, 27, 24, 22},
		{23, 25, 26, 24},
		{24, 26, 25, 23},
		{22, 25, 27, 21},
		{23, 24, 26, 22},
		{25, 26, 24, 23},
		{24, 25, 26, 22},
		{20, 25, 21, 26},
	}
	students = data

	minStudents := students[0][0]
	minParallel := 0

	for i := 0; i < 11; i++ {
		for j := 0; j < 4; j++ {
			if students[i][j] < minStudents {
				minStudents = students[i][j]
				minParallel = i
			}
		}
	}

	classLetters := []string{"А", "Б", "В", "Г"}
	fmt.Printf("а) Самый малочисленный класс находится в параллели: %d\n",
		minParallel+1)

	fmt.Println("б) Самый малочисленный класс среди классов с буквой:")
	for j := 0; j < 4; j++ {
		minForLetter := students[0][j]
		for i := 1; i < 11; i++ {
			if students[i][j] < minForLetter {
				minForLetter = students[i][j]
			}
		}
		fmt.Printf("  Класс %s: %d учеников\n", classLetters[j], minForLetter)
	}
}

func ch9ex14() {
	fmt.Println("Ex14. Доход магазинов")

	var income [3][10]float64

	fmt.Println("Введите доход 3 магазинов за 10 дней:")
	for i := 0; i < 3; i++ {
		fmt.Printf("Магазин %d: ", i+1)
		for j := 0; j < 10; j++ {
			fmt.Scan(&income[i][j])
		}
	}

	var shopTotal [3]float64
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			shopTotal[i] += income[i][j]
		}
	}

	maxShopIndex := 0
	maxShopTotal := shopTotal[0]
	for i := 1; i < 3; i++ {
		if shopTotal[i] > maxShopTotal {
			maxShopTotal = shopTotal[i]
			maxShopIndex = i
		}
	}
	fmt.Printf("а) Магазин %d получил максимальный общий доход: %.2f\n",
		maxShopIndex+1, maxShopTotal)

	var dayTotal [10]float64
	for j := 0; j < 10; j++ {
		for i := 0; i < 3; i++ {
			dayTotal[j] += income[i][j]
		}
	}

	maxDayIndex := 0
	maxDayTotal := dayTotal[0]
	for j := 1; j < 10; j++ {
		if dayTotal[j] > maxDayTotal {
			maxDayTotal = dayTotal[j]
			maxDayIndex = j
		}
	}
	fmt.Printf("б) Максимальный общий доход был в день %d: %.2f\n",
		maxDayIndex+1, maxDayTotal)

	maxIncome := income[0][0]
	maxShopDay := 0
	maxDayShop := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			if income[i][j] > maxIncome {
				maxIncome = income[i][j]
				maxShopDay = i
				maxDayShop = j
			}
		}
	}
	fmt.Printf("в) Магазин %d в день %d получил максимальный доход за день: %.2f\n",
		maxShopDay+1, maxDayShop+1, maxIncome)
}

func ch9ex15() {
	fmt.Println("Ex15. Доход магазинов (доп.)")

	var income [3][10]float64

	fmt.Println("Введите доход 3 магазинов за 10 дней:")
	for i := 0; i < 3; i++ {
		fmt.Printf("Магазин %d: ", i+1)
		for j := 0; j < 10; j++ {
			fmt.Scan(&income[i][j])
		}
	}

	fmt.Println("а) День с максимальным доходом для каждого магазина:")
	for i := 0; i < 3; i++ {
		maxDay := 0
		maxDayIncome := income[i][0]

		for j := 1; j < 10; j++ {
			if income[i][j] > maxDayIncome {
				maxDayIncome = income[i][j]
				maxDay = j
			}
		}
		fmt.Printf("  Магазин %d: день %d (%.2f)\n",
			i+1, maxDay+1, maxDayIncome)
	}

	fmt.Println("б) Магазин с максимальным доходом в каждый день:")
	for j := 0; j < 10; j++ {
		maxShop := 0
		maxIncome := income[0][j]

		for i := 1; i < 3; i++ {
			if income[i][j] > maxIncome {
				maxIncome = income[i][j]
				maxShop = i
			}
		}
		fmt.Printf("  День %d: магазин %d (%.2f)\n",
			j+1, maxShop+1, maxIncome)
	}
}

func ch9ex16() {
	fmt.Println("Ex16. Количество студентов в группах")

	var students [5][6]int

	fmt.Println("Введите количество студентов (5 курсов по 6 групп):")
	for i := 0; i < 5; i++ {
		fmt.Printf("Курс %d: ", i+1)
		for j := 0; j < 6; j++ {
			fmt.Scan(&students[i][j])
		}
	}

	var courseTotal [5]int
	for i := 0; i < 5; i++ {
		for j := 0; j < 6; j++ {
			courseTotal[i] += students[i][j]
		}
	}

	maxCourseIndex := 0
	maxCourseTotal := courseTotal[0]
	for i := 1; i < 5; i++ {
		if courseTotal[i] > maxCourseTotal {
			maxCourseTotal = courseTotal[i]
			maxCourseIndex = i
		}
	}
	fmt.Printf("а) Больше всего студентов на курсе %d: %d студентов\n",
		maxCourseIndex+1, maxCourseTotal)

	minStudents := students[0][0]
	minCourse := 0
	minGroup := 0

	for i := 0; i < 5; i++ {
		for j := 0; j < 6; j++ {
			if students[i][j] < minStudents {
				minStudents = students[i][j]
				minCourse = i
				minGroup = j
			}
		}
	}
	fmt.Printf("б) Самая малочисленная группа: курс %d, группа %d (%d студентов)\n",
		minCourse+1, minGroup+1, minStudents)

	fmt.Println("в) Самая малочисленная группа на каждом курсе:")
	for i := 0; i < 5; i++ {
		minGroupOnCourse := 0
		minOnCourse := students[i][0]

		for j := 1; j < 6; j++ {
			if students[i][j] < minOnCourse {
				minOnCourse = students[i][j]
				minGroupOnCourse = j
			}
		}
		fmt.Printf("  Курс %d: группа %d (%d студентов)\n",
			i+1, minGroupOnCourse+1, minOnCourse)
	}
}

func ch9ex17() {
	fmt.Println("Ex17. Продажи товаров")

	var prices [5]float64
	var quantities [5][6]int

	fmt.Println("Введите цены 5 видов товаров:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Товар %d: ", i+1)
		fmt.Scan(&prices[i])
	}

	fmt.Println("Введите количество проданных товаров за 6 дней:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Товар %d (6 дней): ", i+1)
		for j := 0; j < 6; j++ {
			fmt.Scan(&quantities[i][j])
		}
	}

	fmt.Println("а) Общий доход от каждого вида товара:")
	var goodsIncome [5]float64
	for i := 0; i < 5; i++ {
		totalQty := 0
		for j := 0; j < 6; j++ {
			totalQty += quantities[i][j]
		}
		goodsIncome[i] = float64(totalQty) * prices[i]
		fmt.Printf("  Товар %d: %.2f\n", i+1, goodsIncome[i])
	}

	fmt.Println("б) Общий доход за каждый день:")
	var daysIncome [6]float64
	for j := 0; j < 6; j++ {
		for i := 0; i < 5; i++ {
			daysIncome[j] += float64(quantities[i][j]) * prices[i]
		}
		fmt.Printf("  День %d: %.2f\n", j+1, daysIncome[j])
	}

	totalIncome := 0.0
	for j := 0; j < 6; j++ {
		totalIncome += daysIncome[j]
	}
	fmt.Printf("в) Общий доход за 6 дней: %.2f\n", totalIncome)

	maxGoodIndex := 0
	maxGoodIncome := goodsIncome[0]
	for i := 1; i < 5; i++ {
		if goodsIncome[i] > maxGoodIncome {
			maxGoodIncome = goodsIncome[i]
			maxGoodIndex = i
		}
	}
	fmt.Printf("г) Максимальный доход от товара %d: %.2f\n",
		maxGoodIndex+1, maxGoodIncome)

	maxDayIndex := 0
	maxDayIncome := daysIncome[0]
	for j := 1; j < 6; j++ {
		if daysIncome[j] > maxDayIncome {
			maxDayIncome = daysIncome[j]
			maxDayIndex = j
		}
	}
	fmt.Printf("д) Максимальный общий доход был в день %d: %.2f\n",
		maxDayIndex+1, maxDayIncome)

	threshold := 1000.0
	daysAboveThreshold := 0
	for j := 0; j < 6; j++ {
		if daysIncome[j] > threshold {
			daysAboveThreshold++
		}
	}
	fmt.Printf("е) Дней с доходом выше %.2f: %d\n", threshold, daysAboveThreshold)
}

func ch9ex18() {
	fmt.Println("Ex18. Лучшая группа по среднему баллу")

	var groups [3][20][3]float64

	studentsInGroup := [3]int{18, 20, 17}

	fmt.Println("Введите оценки студентов (3 группы, 3 экзамена):")
	for g := 0; g < 3; g++ {
		fmt.Printf("Группа %d (%d студентов):\n", g+1, studentsInGroup[g])
		for s := 0; s < studentsInGroup[g]; s++ {
			fmt.Printf("  Студент %d: ", s+1)
			for e := 0; e < 3; e++ {
				fmt.Scan(&groups[g][s][e])
			}
		}
	}

	var groupAverages [3]float64

	for g := 0; g < 3; g++ {
		groupTotal := 0.0
		totalGrades := 0

		for s := 0; s < studentsInGroup[g]; s++ {
			studentTotal := 0.0
			for e := 0; e < 3; e++ {
				studentTotal += groups[g][s][e]
			}
			groupTotal += studentTotal / 3.0
			totalGrades++
		}

		groupAverages[g] = groupTotal / float64(totalGrades)
	}

	bestGroup := 0
	bestAverage := groupAverages[0]

	for g := 1; g < 3; g++ {
		if groupAverages[g] > bestAverage {
			bestAverage = groupAverages[g]
			bestGroup = g
		}
	}

	fmt.Printf("Лучшая группа: №%d со средним баллом: %.2f\n",
		bestGroup+1, bestAverage)
}

func ch9ex19() {
	fmt.Println("Ex19. Количество делителей чисел 120-140")

	fmt.Println("Число | Кол-во делителей")
	fmt.Println("-----------------------")

	for num := 120; num <= 140; num++ {
		divisors := 0

		for d := 1; d <= num; d++ {
			if num%d == 0 {
				divisors++
			}
		}

		fmt.Printf("%5d | %d\n", num, divisors)
	}
}

func ch9ex20() {
	fmt.Println("Ex20. График количества делителей")

	var n int
	fmt.Print("Введите n: ")
	fmt.Scan(&n)

	fmt.Printf("Количество делителей для чисел от 1 до %d:\n", n)

	for num := 1; num <= n; num++ {

		divisors := 0
		for d := 1; d <= num; d++ {
			if num%d == 0 {
				divisors++
			}
		}

		fmt.Printf("%d", num)

		for i := 0; i < divisors; i++ {
			fmt.Print("+")
		}

		fmt.Println()
	}
}

func ch9ex21() {
	fmt.Println("Ex21. Числа от 1 до 500 с 5 делителями:")

	count := 0
	for num := 1; num <= 500; num++ {
		divisors := 0
		for d := 1; d <= num; d++ {
			if num%d == 0 {
				divisors++
			}
		}
		if divisors == 5 {
			fmt.Printf("%d ", num)
			count++
		}
	}
	fmt.Printf("\nНайдено чисел: %d\n", count)
}

func ch9ex22() {
	fmt.Println("Ex22. Числа от 200 до 500 с 6 делителями:")

	count := 0
	for num := 200; num <= 500; num++ {
		divisors := 0
		for d := 1; d <= num; d++ {
			if num%d == 0 {
				divisors++
			}
		}
		if divisors == 6 {
			fmt.Printf("%d ", num)
			count++
		}
	}
	fmt.Printf("\nНайдено чисел: %d\n", count)
}

func ch9ex23() {
	fmt.Println("Ex23. Числа с заданным количеством делителей")

	var a, b, k int
	fmt.Print("Введите a, b и k: ")
	fmt.Scan(&a, &b, &k)

	fmt.Printf("Числа от %d до %d с %d делителями:\n", a, b, k)

	count := 0
	for num := a; num <= b; num++ {
		divisors := 0
		for d := 1; d <= num; d++ {
			if num%d == 0 {
				divisors++
			}
		}
		if divisors == k {
			fmt.Printf("%d ", num)
			count++
		}
	}
	if count == 0 {
		fmt.Println("Таких чисел нет")
	} else {
		fmt.Printf("\nНайдено чисел: %d\n", count)
	}
}

func ch9ex24() {
	fmt.Println("Ex24. Число с максимальным количеством делителей")

	var a, b int
	fmt.Print("Введите a и b: ")
	fmt.Scan(&a, &b)

	maxDivisors := 0
	var numbersWithMax []int

	for num := a; num <= b; num++ {
		divisors := 0
		for d := 1; d <= num; d++ {
			if num%d == 0 {
				divisors++
			}
		}

		if divisors > maxDivisors {
			maxDivisors = divisors
			numbersWithMax = []int{num}
		} else if divisors == maxDivisors {
			numbersWithMax = append(numbersWithMax, num)
		}
	}

	fmt.Printf("Максимальное количество делителей: %d\n", maxDivisors)
	fmt.Printf("Числа с таким количеством делителей: %v\n", numbersWithMax)

	if len(numbersWithMax) > 0 {
		maxNum := numbersWithMax[0]
		for _, num := range numbersWithMax {
			if num > maxNum {
				maxNum = num
			}
		}
		fmt.Printf("а) Максимальное число: %d\n", maxNum)

		minNum := numbersWithMax[0]
		for _, num := range numbersWithMax {
			if num < minNum {
				minNum = num
			}
		}
		fmt.Printf("б) Минимальное число: %d\n", minNum)
	}
}

func ch9ex25() {
	fmt.Println("Ex25. Трехзначные простые числа:")

	count := 0
	for num := 100; num <= 999; num++ {
		isPrime := true

		for d := 2; d*d <= num; d++ {
			if num%d == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d ", num)
			count++
			if count%10 == 0 {
				fmt.Println()
			}
		}
	}
	fmt.Printf("\nВсего: %d простых чисел\n", count)
}

func ch9ex26() {
	fmt.Println("Ex26. Первые 100 простых чисел:")

	count := 0
	num := 2

	for count < 100 {
		isPrime := true
		for d := 2; d*d <= num; d++ {
			if num%d == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%4d ", num)
			count++
			if count%10 == 0 {
				fmt.Println()
			}
		}
		num++
	}
	fmt.Println()
}

func ch9ex27() {
	fmt.Println("Ex27. Сумма делителей чисел от 50 до 70:")

	fmt.Println("Число | Сумма делителей")
	fmt.Println("----------------------")
	for num := 50; num <= 70; num++ {
		sum := 0
		for d := 1; d <= num; d++ {
			if num%d == 0 {
				sum += d
			}
		}
		fmt.Printf("%5d | %d\n", num, sum)
	}
}

func ch9ex28() {
	fmt.Println("Ex28. Числа от 100 до 300 с суммой делителей = 50:")

	found := false
	for num := 100; num <= 300; num++ {
		sum := 0
		for d := 1; d <= num; d++ {
			if num%d == 0 {
				sum += d
			}
		}
		if sum == 50 {
			fmt.Printf("%d ", num)
			found = true
		}
	}
	if !found {
		fmt.Println("Таких чисел нет")
	} else {
		fmt.Println()
	}
}

func ch9ex29() {
	fmt.Println("Ex29. Числа от 300 до 600 с суммой делителей кратной 10:")

	count := 0
	for num := 300; num <= 600; num++ {
		sum := 0
		for d := 1; d <= num; d++ {
			if num%d == 0 {
				sum += d
			}
		}
		if sum%10 == 0 {
			fmt.Printf("%d (сумма: %d) ", num, sum)
			count++
			if count%5 == 0 {
				fmt.Println()
			}
		}
	}
	fmt.Printf("\nВсего: %d чисел\n", count)
}

func ch9ex30() {
	fmt.Println("Ex30. Трехзначное совершенное число:")

	for num := 100; num <= 999; num++ {
		sum := 0
		for d := 1; d <= num/2; d++ {
			if num%d == 0 {
				sum += d
			}
		}
		if sum == num {
			fmt.Printf("Найдено совершенное число: %d\n", num)
			fmt.Printf("Его делители (кроме самого числа): ")
			for d := 1; d <= num/2; d++ {
				if num%d == 0 {
					fmt.Printf("%d ", d)
				}
			}
			fmt.Println()
			return
		}
	}
	fmt.Println("Трехзначных совершенных чисел нет")
}

func ch9ex31() {
	fmt.Println("Ex31. Совершенные числа меньше 100000:")

	num := 1
	count := 0

	for {
		sum := 0
		for d := 1; d <= num/2; d++ {
			if num%d == 0 {
				sum += d
			}
		}

		if sum == num && num < 100000 {
			fmt.Printf("%d ", num)
			count++
		}

		num++
		if num >= 100000 {
			break
		}
	}

	if count == 0 {
		fmt.Println("Совершенных чисел не найдено")
	} else {
		fmt.Printf("\nВсего: %d совершенных чисел\n", count)
	}
}

func ch9ex32() {
	fmt.Println("Ex32. Число с максимальной суммой делителей")

	var a, b int
	fmt.Print("Введите a и b: ")
	fmt.Scan(&a, &b)

	maxSum := 0
	numberWithMaxSum := a

	for num := a; num <= b; num++ {
		sum := 0
		for d := 1; d <= num; d++ {
			if num%d == 0 {
				sum += d
			}
		}

		if sum > maxSum {
			maxSum = sum
			numberWithMaxSum = num
		}
	}

	fmt.Printf("Число с максимальной суммой делителей: %d\n", numberWithMaxSum)
	fmt.Printf("Сумма его делителей: %d\n", maxSum)
}

func ch9ex33() {
	fmt.Println("Ex33. Дружественные числа меньше 50000:")

	divisorSums := make([]int, 50001)

	for num := 1; num <= 50000; num++ {
		sum := 0
		for d := 1; d <= num/2; d++ {
			if num%d == 0 {
				sum += d
			}
		}
		divisorSums[num] = sum
	}

	count := 0
	for a := 1; a <= 50000; a++ {
		b := divisorSums[a]
		if b <= 50000 && b > a && divisorSums[b] == a {
			fmt.Printf("(%d, %d) ", a, b)
			count++
		}
	}

	if count == 0 {
		fmt.Println("Дружественных чисел не найдено")
	} else {
		fmt.Printf("\nВсего пар: %d\n", count)
	}
}

func ch9ex34() {
	fmt.Println("Ex34. Способы выплаты суммы")

	var n int
	fmt.Print("Введите сумму n (n < 100): ")
	fmt.Scan(&n)

	count := 0
	fmt.Println("\na) Подсчет количества способов:")

	for tens := 0; tens <= n/10; tens++ {
		remaining := n - tens*10
		for fives := 0; fives <= remaining/5; fives++ {
			remaining2 := remaining - fives*5
			for twos := 0; twos <= remaining2/2; twos++ {
				count++
			}
		}
	}

	fmt.Printf("Количество способов выплаты: %d\n", count)

	fmt.Println("\nb) Все способы выплаты:")
	fmt.Println("10р 5р 2р 1р")
	fmt.Println("------------")

	wayNum := 1
	for tens := 0; tens <= n/10; tens++ {
		remaining := n - tens*10
		for fives := 0; fives <= remaining/5; fives++ {
			remaining2 := remaining - fives*5
			for twos := 0; twos <= remaining2/2; twos++ {
				ones := remaining2 - twos*2
				fmt.Printf("%2d. %2d  %2d  %2d  %2d\n",
					wayNum, tens, fives, twos, ones)
				wayNum++
			}
		}
	}
}

func ch9ex35() {
	fmt.Println("Ex35. Минимальное количество купюр для выплаты")

	var n int
	fmt.Print("Введите n: ")
	fmt.Scan(&n)

	denominations := []int{64, 32, 16, 8, 4, 2, 1}

	for amount := n; amount <= n+10; amount++ {
		fmt.Printf("\nСумма %d рублей:\n", amount)
		remaining := amount

		for _, denom := range denominations {
			if remaining == 0 {
				break
			}
			count := remaining / denom
			if count > 0 {
				fmt.Printf("  Купюры по %d руб: %d шт.\n", denom, count)
				remaining -= count * denom
			}
		}
	}
}

func ch9ex36() {
	fmt.Println("Ex36. Прямоугольники с заданной площадью")

	var s int
	fmt.Print("Введите площадь s: ")
	fmt.Scan(&s)

	fmt.Println("\na) Решения разными (с перестановкой):")
	countA := 0
	for a := 1; a <= s; a++ {
		if s%a == 0 {
			b := s / a
			fmt.Printf("  %d x %d\n", a, b)
			countA++
		}
	}
	fmt.Printf("Всего решений: %d\n", countA)

	fmt.Println("\nb) Решения совпадающими (без перестановки):")
	countB := 0
	for a := 1; a*a <= s; a++ {
		if s%a == 0 {
			b := s / a
			fmt.Printf("  %d x %d\n", a, b)
			countB++
		}
	}
	fmt.Printf("Всего решений: %d\n", countB)
}

func ch9ex37() {
	fmt.Println("Ex37. Параллелепипеды с заданным объемом")

	var v int
	fmt.Print("Введите объем v: ")
	fmt.Scan(&v)

	fmt.Println("\na) Решения разными (с перестановкой):")
	countA := 0
	for a := 1; a <= v; a++ {
		if v%a == 0 {
			for b := 1; b <= v/a; b++ {
				if (v/a)%b == 0 {
					c := v / (a * b)
					if a*b*c == v {
						fmt.Printf("  %d x %d x %d\n", a, b, c)
						countA++
					}
				}
			}
		}
	}
	fmt.Printf("Всего решений: %d\n", countA)

	fmt.Println("\nb) Решения совпадающими (без перестановки):")
	countB := 0
	for a := 1; a*a*a <= v; a++ {
		if v%a == 0 {
			for b := a; b*b <= v/a; b++ {
				if (v/a)%b == 0 {
					c := v / (a * b)
					if a*b*c == v && c >= b {
						fmt.Printf("  %d x %d x %d\n", a, b, c)
						countB++
					}
				}
			}
		}
	}
	fmt.Printf("Всего решений: %d\n", countB)
}

func ch9ex38() {
	fmt.Println("Ex38. Решения уравнения x² + y² = k² (1-30):")

	solutions := make(map[string]bool)
	count := 0

	for x := 1; x <= 30; x++ {
		for y := x; y <= 30; y++ {
			for k := 1; k <= 30; k++ {
				if x*x+y*y == k*k {
					key := fmt.Sprintf("%d² + %d² = %d²", x, y, k)
					if !solutions[key] {
						fmt.Printf("  %s\n", key)
						solutions[key] = true
						count++
					}
				}
			}
		}
	}

	fmt.Printf("Всего решений: %d\n", count)
}

func ch9ex39() {
	fmt.Println("Ex39. Сумма степеней 1ⁿ + 2ⁿ + ... + mⁿ")

	var m, n int
	fmt.Print("Введите m и n: ")
	fmt.Scan(&m, &n)

	sum := 0
	for i := 1; i <= m; i++ {

		power := 1
		for j := 0; j < n; j++ {
			power *= i
		}
		sum += power
	}

	fmt.Printf("1^%d + 2^%d + ... + %d^%d = %d\n", n, n, m, n, sum)
}

func ch9ex40() {
	fmt.Println("Ex40. Сумма 1¹ + 2² + ... + nⁿ")

	var n int
	fmt.Print("Введите n: ")
	fmt.Scan(&n)

	sum := 0
	for i := 1; i <= n; i++ {

		power := 1
		for j := 0; j < i; j++ {
			power *= i
		}
		sum += power
	}

	fmt.Printf("1¹ + 2² + ... + %d^%d = %d\n", n, n, sum)
}

func ch9ex41() {
	fmt.Println("Ex41. Трехзначные числа с заданной суммой цифр")

	var n int
	fmt.Print("Введите n (1-27): ")
	fmt.Scan(&n)

	if n < 1 || n > 27 {
		fmt.Println("n должно быть от 1 до 27")
		return
	}

	fmt.Printf("Трехзначные числа с суммой цифр %d:\n", n)
	count := 0

	for num := 100; num <= 999; num++ {

		str := fmt.Sprintf("%d", num)

		sumDigits := 0
		for _, ch := range str {
			sumDigits += int(ch - '0')
		}

		if sumDigits == n {
			fmt.Printf("%d ", num)
			count++
			if count%10 == 0 {
				fmt.Println()
			}
		}
	}

	if count == 0 {
		fmt.Println("Таких чисел нет")
	} else {
		fmt.Printf("\nВсего: %d чисел\n", count)
	}
}

func ch9ex42() {
	fmt.Println("Ex42. Трехзначные числа без повторяющихся цифр:")

	count := 0

	for num := 100; num <= 999; num++ {
		str := fmt.Sprintf("%d", num)

		a, b, c := str[0], str[1], str[2]
		if a != b && a != c && b != c {
			fmt.Printf("%d ", num)
			count++
			if count%10 == 0 {
				fmt.Println()
			}
		}
	}

	fmt.Printf("\nВсего: %d чисел\n", count)
}

func ch9ex43() {
	fmt.Println("Ex43. НОД нескольких чисел (алгоритм Евклида)")

	var n int
	fmt.Print("Сколько чисел? ")
	fmt.Scan(&n)

	if n < 2 {
		fmt.Println("Нужно как минимум 2 числа")
		return
	}

	numbers := make([]int, n)
	fmt.Printf("Введите %d чисел: ", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	result := numbers[0]
	for i := 1; i < n; i++ {
		result = gcd(result, numbers[i])
	}

	fmt.Printf("НОД всех чисел: %d\n", result)
}

func ch9ex44() {
	fmt.Println("Ex44. Способы составления веса из гирь")

	weights := []int{100, 200, 500, 500, 1000, 1200, 1400, 1500, 2000, 3000}
	n := len(weights)

	var v int
	fmt.Print("Введите вес v (в граммах): ")
	fmt.Scan(&v)

	count := 0

	var findCombinations func(int, int)
	findCombinations = func(index int, currentSum int) {
		if currentSum == v {
			count++
			return
		}
		if currentSum > v || index >= n {
			return
		}

		findCombinations(index+1, currentSum)

		findCombinations(index+1, currentSum+weights[index])
	}

	findCombinations(0, 0)

	fmt.Printf("Число способов составить вес %d г: %d\n", v, count)
}

func ch9ex45() {
	fmt.Println("Ex45. Числа с квадратом суммы цифр = m")

	var m, n int
	fmt.Print("Введите m и n: ")
	fmt.Scan(&m, &n)

	fmt.Printf("Числа < %d, у которых (сумма цифр)² = %d:\n", n, m)

	count := 0
	for num := 1; num < n; num++ {

		sumDigits := 0
		temp := num
		for temp > 0 {
			sumDigits += temp % 10
			temp /= 10
		}

		if sumDigits*sumDigits == m {
			fmt.Printf("%d (сумма цифр: %d) ", num, sumDigits)
			count++
			if count%5 == 0 {
				fmt.Println()
			}
		}
	}

	if count == 0 {
		fmt.Println("Таких чисел нет")
	} else {
		fmt.Printf("\nВсего: %d чисел\n", count)
	}
}

func ch9ex46() {
	fmt.Println("Ex46. Цифровой корень числа")

	var num int
	fmt.Print("Введите натуральное число: ")
	fmt.Scan(&num)

	original := num

	for num >= 10 {
		sum := 0
		temp := num
		for temp > 0 {
			sum += temp % 10
			temp /= 10
		}
		num = sum
	}

	fmt.Printf("Цифровой корень числа %d: %d\n", original, num)
}

func ch9ex47() {
	fmt.Println("Ex47. Покупка скота (100 рублей, 100 голов)")
	fmt.Println("Бык - 10 руб, Корова - 5 руб, Теленок - 0.5 руб")
	fmt.Println("Варианты покупки 100 голов на 100 рублей:")

	count := 0

	for bulls := 0; bulls <= 10; bulls++ {
		for cows := 0; cows <= 20; cows++ {
			calves := 100 - bulls - cows
			if calves >= 0 {
				totalCost := float64(bulls)*10 + float64(cows)*5 + float64(calves)*0.5
				if totalCost == 100 {
					fmt.Printf("Быков: %d, Коров: %d, Телят: %d\n",
						bulls, cows, calves)
					count++
				}
			}
		}
	}

	if count == 0 {
		fmt.Println("Нет возможных вариантов")
	} else {
		fmt.Printf("Всего вариантов: %d\n", count)
	}
}

func ch9ex48() {
	fmt.Println("Ex48. Разложение на простые множители")

	var n int
	fmt.Print("Введите натуральное число: ")
	fmt.Scan(&n)

	fmt.Printf("Разложение числа %d:\n", n)

	fmt.Println("1) Каждый множитель один раз:")
	factorsSet := make(map[int]bool)
	temp := n
	divisor := 2

	for temp > 1 {
		if temp%divisor == 0 {
			factorsSet[divisor] = true
			temp /= divisor
		} else {
			divisor++
		}
	}

	sortedFactors := make([]int, 0, len(factorsSet))
	for factor := range factorsSet {
		sortedFactors = append(sortedFactors, factor)
	}
	sort.Ints(sortedFactors)

	for i, factor := range sortedFactors {
		if i > 0 {
			fmt.Print(" * ")
		}
		fmt.Print(factor)
	}
	if len(sortedFactors) == 0 {
		fmt.Print("1")
	}
	fmt.Println()

	fmt.Println("2) Множители с кратностями:")
	temp = n
	divisor = 2
	first := true

	for temp > 1 {
		count := 0
		for temp%divisor == 0 {
			count++
			temp /= divisor
		}

		if count > 0 {
			if !first {
				fmt.Print(" * ")
			}
			if count == 1 {
				fmt.Print(divisor)
			} else {
				fmt.Printf("%d^%d", divisor, count)
			}
			first = false
		}
		divisor++
	}

	if first {
		fmt.Print("1")
	}
	fmt.Println()
}

func ch9ex49() {
	fmt.Println("Ex49. Все простые делители числа")

	var n int
	fmt.Print("Введите натуральное число: ")
	fmt.Scan(&n)

	fmt.Printf("Простые делители числа %d: ", n)

	temp := n
	divisor := 2
	first := true

	for temp > 1 {
		if temp%divisor == 0 {
			if !first {
				fmt.Print(", ")
			}
			fmt.Print(divisor)
			first = false

			for temp%divisor == 0 {
				temp /= divisor
			}
		}
		divisor++
	}

	if first {
		fmt.Print("нет (или 1)")
	}
	fmt.Println()
}

func ch9ex50() {
	fmt.Println("Ex50. Числа, взаимно простые с n")

	var n int
	fmt.Print("Введите натуральное число n: ")
	fmt.Scan(&n)

	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	fmt.Printf("Числа < %d, взаимно простые с %d:\n", n, n)

	count := 0
	for i := 1; i < n; i++ {
		if gcd(i, n) == 1 {
			fmt.Printf("%d ", i)
			count++
			if count%10 == 0 {
				fmt.Println()
			}
		}
	}

	if count == 0 {
		fmt.Println("Таких чисел нет")
	} else {
		fmt.Printf("\nВсего: %d чисел\n", count)
	}
}

func ch9ex51() {
	fmt.Println("Ex51. Числа, взаимно простые с p, меньшие n")

	var n, p int
	fmt.Print("Введите n и p: ")
	fmt.Scan(&n, &p)

	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	fmt.Printf("Числа < %d, взаимно простые с %d:\n", n, p)

	count := 0
	for i := 1; i < n; i++ {
		if gcd(i, p) == 1 {
			fmt.Printf("%d ", i)
			count++
			if count%10 == 0 {
				fmt.Println()
			}
		}
	}

	if count == 0 {
		fmt.Println("Таких чисел нет")
	} else {
		fmt.Printf("\nВсего: %d чисел\n", count)
	}
}

func ch9ex52() {
	fmt.Println("Ex52. Делители q, взаимно простые с p")

	var p, q int
	fmt.Print("Введите p и q: ")
	fmt.Scan(&p, &q)

	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	fmt.Printf("Делители %d, взаимно простые с %d:\n", q, p)

	count := 0
	for d := 1; d <= q; d++ {
		if q%d == 0 && gcd(d, p) == 1 {
			fmt.Printf("%d ", d)
			count++
			if count%10 == 0 {
				fmt.Println()
			}
		}
	}

	if count == 0 {
		fmt.Println("Таких делителей нет")
	} else {
		fmt.Printf("\nВсего: %d делителей\n", count)
	}
}

func ch9ex53() {
	fmt.Println("Ex53. Число, представимое двумя способами как сумма кубов")

	limit := 10000
	sums := make(map[int][]string)

	for a := 1; a*a*a < limit; a++ {
		for b := a; b*b*b < limit; b++ {
			sum := a*a*a + b*b*b
			if sum < limit {
				representation := fmt.Sprintf("%d³ + %d³", a, b)
				sums[sum] = append(sums[sum], representation)
			}
		}
	}

	smallest := -1
	for sum, representations := range sums {
		if len(representations) >= 2 {
			if smallest == -1 || sum < smallest {
				smallest = sum
			}
		}
	}

	if smallest != -1 {
		fmt.Printf("Наименьшее такое число: %d\n", smallest)
		fmt.Println("Представления:")
		for _, rep := range sums[smallest] {
			fmt.Printf("  %s = %d\n", rep, smallest)
		}
	} else {
		fmt.Println("Такое число не найдено в пределах", limit)
	}
}

func ch9ex54() {
	fmt.Println("Ex54. Простые несократимые дроби 0 < дробь < 1, знаменатель ≤ 7")

	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	fractions := make([]struct {
		num, den int
		str      string
	}, 0)

	for denominator := 2; denominator <= 7; denominator++ {
		for numerator := 1; numerator < denominator; numerator++ {
			if gcd(numerator, denominator) == 1 {
				fractions = append(fractions, struct {
					num, den int
					str      string
				}{
					num: numerator,
					den: denominator,
					str: fmt.Sprintf("%d/%d", numerator, denominator),
				})
			}
		}
	}

	sort.Slice(fractions, func(i, j int) bool {
		return fractions[i].num*fractions[j].den < fractions[j].num*fractions[i].den
	})

	fmt.Println("Дроби в порядке возрастания:")
	for i, f := range fractions {
		fmt.Printf("%s ", f.str)
		if (i+1)%6 == 0 {
			fmt.Println()
		}
	}
	fmt.Printf("\nВсего: %d дробей\n", len(fractions))
}
