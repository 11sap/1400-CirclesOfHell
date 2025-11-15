package chapters

import (
	"fmt"
)

func Run03() {
	tasks := []func(){
		ch3ex01, ch3ex02, ch3ex03, ch3ex04, ch3ex05, ch3ex06, ch3ex07, ch3ex08,
		ch3ex09, ch3ex10, ch3ex11, ch3ex12, ch3ex13, ch3ex14, ch3ex15, ch3ex16,
		ch3ex17, ch3ex18, ch3ex19, ch3ex20, ch3ex21, ch3ex22, ch3ex23, ch3ex24,
		ch3ex25, ch3ex26, ch3ex27, ch3ex28, ch3ex29, ch3ex30, ch3ex31, ch3ex32,
		ch3ex33,
	}

	for {
		fmt.Print("\nГлава 3. Выберите задачу (1-51): ")
		var n int
		fmt.Scan(&n)

		if n == 0 {
			break
		}
		if n >= 1 && n <= 51 {
			tasks[n-1]()
		} else {
			fmt.Println("От 1 до 51! Для выхода введите 0.")
		}
	}
}

func ch3ex01() {
	var x int
	fmt.Printf("Ex01. Вычисление метров из сантиметров\n")
	fmt.Printf("Ex01. Введите расстояние в сантиметрах: ")
	fmt.Scan(&x)
	fmt.Printf("Ex01. Расстояние: %d (м)\n", x/100)
}

func ch3ex02() {
	var x int
	fmt.Printf("Ex02. Вычисление центнеров из килограммов\n")
	fmt.Printf("Ex02. Введите массу в кг: ")
	fmt.Scan(&x)
	fmt.Printf("Ex02. Масса: %d (ц)\n", x/100)
}

func ch3ex03() {
	fmt.Printf("Ex03. Вычисление кол-ва полных недель в 234 днях\n")
	fmt.Printf("Ex03. Ответ: %d нед.\n", 234/7)
}

func ch3ex04() {
	var x, y int
	fmt.Printf("Ex04. Вычисление задачи с яблоками\n")
	fmt.Printf("Ex04. Введите кол-во школьников и яблок: ")
	fmt.Scan(&x, &y)
	fmt.Printf("Ex04. Ответ: Каждому досталось по %d яблок(-а)\nВ корзинке - %d \n", y/x, y%x)
}

func ch3ex05() {
	fmt.Printf("Ex05. Вычисление задачи с прямоугольником\n")
	fmt.Printf("Ex05. Ответ: %d квадратов со стороной 130мм\n", 543/130)
}

func ch3ex06() {
	var z int
	x, y := 9, 4
	fmt.Printf("Ex06. Вычисление номера купе по месту\n")
	fmt.Printf("Ex06. Введите номер места:\n")
	fmt.Scan(&z)

	if z <= x*y {
		fmt.Printf("Ex06. Ответ: Ваше купе №%d\n", z/y+1)
	} else {
		fmt.Printf("Ex06. Ответ: Ваше купе в другом вагоне\n")
	}
}

func ch3ex07() {
	var z int
	x, y := 5, 3
	fmt.Printf("Ex07. Вычисление этажа по номеру кввартиры'\n")
	fmt.Printf("Ex07. Введите номер квартиры:")
	fmt.Scan(&z)

	if z <= x*y {
		fmt.Printf("\nEx07. Ответ: Ваш этаж №%d\n", z/y+1)
	} else {
		fmt.Printf("Ex07. Ответ: Ваша квартира в другом подъезде\n")
	}
}

func ch3ex08() {
	var z int
	const (
		firstNum    = 1643
		rows        = 20
		seatsPerRow = 15
	)
	fmt.Printf("Ex08. Вычисление ряда нахождения места'\n")
	fmt.Printf("Ex08. Введите номер билета:")
	fmt.Scan(&z)

	i := z - firstNum

	if z >= firstNum {
		fmt.Printf("\nEx08. Ответ: Ваше место №%d в %d ряду\n", i%seatsPerRow+1, i/seatsPerRow+1)
	} else {
		fmt.Printf("Ex08. Ответ: Билет не существует\n")
	}
}

func ch3ex09() {
	var seconds int
	fmt.Printf("Ex09. Вычисление времени\n")
	fmt.Printf("Ex09. Введите количество секунд: ")
	fmt.Scan(&seconds)
	fmt.Printf("Ex09. Ответ:\nа) %d\nб) %d\nв) %d\n", seconds/3600, (seconds%3600)/60, seconds%60)
}

func ch3ex10() {
	var k, d int
	fmt.Printf("Ex10. Определение дня недели\n")
	fmt.Printf("Ex10. Введите номер дня года (1-365) и день недели 1 января (1-7): ")
	fmt.Scan(&k, &d)

	n := (k+d-2)%7 + 1
	if n == 7 {
		n = 0
	}

	fmt.Printf("Ex10. Ответ: %d\n", n)
}

func ch3ex11() {
	var n int
	fmt.Printf("Ex11. Определение месяца\n")
	fmt.Printf("Ex11. Введите количество прошедших месяцев: ")
	fmt.Scan(&n)
	fmt.Printf("Ex11. Ответ: %d\n", n%12+1)
}

func ch3ex12() {
	var flat int
	fmt.Printf("Ex12. Определение этажа и позиции квартиры\n")
	fmt.Printf("Ex12. Введите номер квартиры (1-20): ")
	fmt.Scan(&flat)
	fmt.Printf("Ex12. Ответ: Этаж %d, Квартира по счету %d\n", (flat-1)/4+1, (flat-1)%4+1)
}

func ch3ex13() {
	var n int
	fmt.Printf("Ex13. Определение позиции в таблице\n")
	fmt.Printf("Ex13. Введите порядковый номер: ")
	fmt.Scan(&n)
	fmt.Printf("Ex13. Ответ: Строка %d, Столбец %d\n", (n-1)/5+1, (n-1)%5+1)
}

func ch3ex14() {
	var flat int
	fmt.Printf("Ex14. Определение подъезда, этажа и позиции\n")
	fmt.Printf("Ex14. Введите номер квартиры: ")
	fmt.Scan(&flat)
	flatsPerEntrance := 9 * 6
	flatsPerFloor := 6
	entrance := (flat-1)/flatsPerEntrance + 1
	floorInEntrance := ((flat-1)%flatsPerEntrance)/flatsPerFloor + 1
	positionOnFloor := (flat-1)%flatsPerFloor + 1
	fmt.Printf("Ex14. Ответ: Подъезд %d, Этаж %d, Позиция %d\n", entrance, floorInEntrance, positionOnFloor)
}

func ch3ex15() {
	var place int
	fmt.Printf("Ex15. Определение секции и яруса\n")
	fmt.Printf("Ex15. Введите номер места: ")
	fmt.Scan(&place)

	placesPerTier := 8 * 15
	tier := (place-1)/placesPerTier + 1
	section := ((place-1)%placesPerTier)/15 + 1

	fmt.Printf("Ex15. Ответ: Ярус %d, Секция %d\n", tier, section)
}

func ch3ex16() {
	var n int
	fmt.Printf("Ex16. Анализ двузначного числа\n")
	fmt.Printf("Ex16. Введите двузначное число: ")
	fmt.Scan(&n)
	fmt.Printf("Ex16. Ответ: Десятков: %d, Единиц: %d\n", n/10, n%10)
}

func ch3ex17() {
	var n int
	fmt.Printf("Ex17. Сумма цифр двузначного числа\n")
	fmt.Printf("Ex17. Введите двузначное число: ")
	fmt.Scan(&n)
	fmt.Printf("Ex17. Ответ: %d\n", n/10+n%10)
}

func ch3ex18() {
	var n int
	fmt.Printf("Ex18. Перестановка цифр\n")
	fmt.Printf("Ex18. Введите двузначное число: ")
	fmt.Scan(&n)
	fmt.Printf("Ex18. Ответ: %d\n", n%10*10+n/10)
}

func ch3ex19() {
	var n int
	fmt.Printf("Ex19. Цифры трехзначного числа\n")
	fmt.Printf("Ex19. Введите трехзначное число: ")
	fmt.Scan(&n)
	fmt.Printf("Ex19. Ответ: %d, %d, %d\n", n/100, n/10%10, n%10)
}

func ch3ex20() {
	var n int
	fmt.Printf("Ex20. Анализ трехзначного числа\n")
	fmt.Printf("Ex20. Введите трехзначное число: ")
	fmt.Scan(&n)
	h, t, u := n/100, n/10%10, n%10
	fmt.Printf("Ex20. Ответ: Единиц: %d, Десятков: %d, Сумма: %d, Произведение: %d\n", u, t, h+t+u, h*t*u)
}

func ch3ex21() {
	var n int
	fmt.Printf("Ex21. Число справа налево\n")
	fmt.Printf("Ex21. Введите трехзначное число: ")
	fmt.Scan(&n)
	fmt.Printf("Ex21. Ответ: %d\n", n%10*100+n/10%10*10+n/100)
}

func ch3ex22() {
	var n int
	fmt.Printf("Ex22. Перемещение первой цифры в конец\n")
	fmt.Printf("Ex22. Введите трехзначное число: ")
	fmt.Scan(&n)
	fmt.Printf("Ex22. Ответ: %d\n", n%100*10+n/100)
}

func ch3ex23() {
	var n int
	fmt.Printf("Ex23. Последнюю цифру в начало\n")
	fmt.Printf("Ex23. Введите трехзначное число: ")
	fmt.Scan(&n)
	fmt.Printf("Ex23. Ответ: %d\n", n%10*100+n/10)
}

func ch3ex24() {
	var n int
	fmt.Printf("Ex24. Перестановка первой и второй цифр\n")
	fmt.Printf("Ex24. Введите трехзначное число: ")
	fmt.Scan(&n)
	fmt.Printf("Ex24. Ответ: %d\n", n/100%10*100+n/100*10+n%10)
}

func ch3ex25() {
	var n int
	fmt.Printf("Ex25. Перестановка второй и третьей цифр\n")
	fmt.Printf("Ex25. Введите трехзначное число: ")
	fmt.Scan(&n)
	fmt.Printf("Ex25. Ответ: %d\n", n/100*100+n%10*10+n/10%10)
}

func ch3ex26() {
	var n int
	fmt.Printf("Ex26. Все перестановки цифр\n")
	fmt.Printf("Ex26. Введите трехзначное число: ")
	fmt.Scan(&n)
	a, b, c := n/100, n/10%10, n%10
	fmt.Printf("Ex26. Ответ: %d %d %d %d %d %d\n", a*100+b*10+c, a*100+c*10+b, b*100+a*10+c, b*100+c*10+a, c*100+a*10+b, c*100+b*10+a)
}

func ch3ex27() {
	var n int
	fmt.Printf("Ex27. Сумма цифр\n")
	fmt.Printf("Ex27. Введите 4-значное или 5-значное число: ")
	fmt.Scan(&n)
	sum := 0
	for t := n; t > 0; t /= 10 {
		sum += t % 10
	}
	fmt.Printf("Ex27. Ответ: %d\n", sum)
}

func ch3ex28() {
	var n int
	fmt.Printf("Ex28. Операции с четырехзначным числом\n")
	fmt.Printf("Ex28. Введите четырехзначное число: ")
	fmt.Scan(&n)
	a, b, c, d := n/1000, n/100%10, n/10%10, n%10

	fmt.Printf("а) %d\n", d*1000+c*100+b*10+a)
	fmt.Printf("б) %d\n", b*1000+a*100+d*10+c)
	fmt.Printf("в) %d\n", a*1000+c*100+b*10+d)
	fmt.Printf("г1) %d\n", c*1000+d*1000+a*10+b)
	fmt.Printf("г2) %d\n", n%100*100+n/100)
}

func ch3ex29() {
	var n int
	fmt.Printf("Ex29. Единицы и десятки\n")
	fmt.Printf("Ex29. Введите натуральное число (>9): ")
	fmt.Scan(&n)
	fmt.Printf("Ex29. Ответ: Единиц: %d, Десятков: %d\n", n%10, n/10%10)
}

func ch3ex30() {
	var n int
	fmt.Printf("Ex30. Десятки и сотни\n")
	fmt.Printf("Ex30. Введите натуральное число (>99): ")
	fmt.Scan(&n)
	fmt.Printf("Ex30. Ответ: Десятков: %d, Сотен: %d\n", n/10%10, n/100%10)
}

func ch3ex31() {
	var n int
	fmt.Printf("Ex31. Сотни и тысячи\n")
	fmt.Printf("Ex31. Введите натуральное число (>999): ")
	fmt.Scan(&n)
	fmt.Printf("Ex31. Ответ: Сотен: %d, Тысяч: %d\n", n/100%10, n/1000)
}

func ch3ex32() {
	fmt.Printf("Ex32. Восстановление числа\n")
	result := 237
	lastDigit := result / 100
	firstTwo := result % 100
	original := firstTwo*10 + lastDigit
	fmt.Printf("Ex32. Ответ: %d\n", original)
}

func ch3ex33() {
	fmt.Printf("Ex33. Восстановление трехзначного числа\n")
	var result int
	fmt.Printf("Ex33. Введите полученное число: ")
	fmt.Scan(&result)
	lastDigit := result / 100
	firstTwo := result % 100
	original := firstTwo*10 + lastDigit
	fmt.Printf("Ex33. Ответ: %d\n", original)
}
