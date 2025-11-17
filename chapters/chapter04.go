package chapters

import (
	"fmt"
	"math"
)

func Run04() {
	tasks := []func(){
		ch4ex01, ch4ex02, ch4ex03, ch4ex04, ch4ex05, ch4ex06, ch4ex07, ch4ex08,
		ch4ex09, ch4ex10, ch4ex11, ch4ex12, ch4ex13, ch4ex14, ch4ex15, ch4ex16,
		ch4ex17, ch4ex18, ch4ex19, ch4ex20, ch4ex21, ch4ex22, ch4ex23, ch4ex24,
		ch4ex25, ch4ex26, ch4ex27, ch4ex28, ch4ex29, ch4ex30, ch4ex31, ch4ex32,
		ch4ex33, ch4ex34, ch4ex35, ch4ex36, ch4ex37, ch4ex38, ch4ex39, ch4ex40,
		ch4ex41, ch4ex42, ch4ex43, ch4ex44, ch4ex45, ch4ex46, ch4ex47, ch4ex48,
		ch4ex49, ch4ex50, ch4ex51, ch4ex52, ch4ex53, ch4ex54, ch4ex55, ch4ex56,
		ch4ex57, ch4ex58, ch4ex59, ch4ex60, ch4ex61,
	}

	for {
		fmt.Print("\nГлава 4. Выберите задачу (1-145): ")
		var n int
		fmt.Scan(&n)

		if n == 0 {
			break
		}
		if n >= 1 && n <= 145 {
			tasks[n-1]()
		} else {
			fmt.Println("От 1 до 145! Для выхода введите 0.")
		}
	}
}

func ch4ex01() {
	var a, b float64
	fmt.Printf("Ex01. Сравнение двух чисел\n")
	fmt.Printf("Ex01. Введите два числа: ")
	fmt.Scan(&a, &b)
	if a > b {
		a, b = b, a
	}

	fmt.Printf("Ex01. Ответ: %.2f > %.2f\n", a, b)
}

func ch4ex02() {
	var x float64
	fmt.Printf("Ex02. Вычисление функции\n")
	fmt.Printf("Ex02. Введите x: ")
	fmt.Scan(&x)

	var y float64
	if x > 0 {
		y = math.Sin(x) * math.Sin(x)
	} else {
		y = 1 - 2*math.Sin(x*x)
	}
	fmt.Printf("Ex02. Ответ: y = %.4f\n", y)
}

func ch4ex03() {
	var x float64
	fmt.Printf("Ex03. Вычисление функции\n")
	fmt.Printf("Ex03. Введите x: ")
	fmt.Scan(&x)

	var y float64
	if x > 0 {
		y = math.Sin(x * x)
	} else {
		y = 1 + 2*math.Sin(x)*math.Sin(x)
	}
	fmt.Printf("Ex03. Ответ: y = %.4f\n", y)
}

func ch4ex04() {
	var x, y float64
	fmt.Printf("Ex04. Область I или II (рис 4.1)\n")
	fmt.Printf("Ex04. Введите координаты x и y: ")
	fmt.Scan(&x, &y)

	if x < 4 {
		fmt.Printf("Ex04. Ответ: Область I\n")
	} else {
		fmt.Printf("Ex04. Ответ: Область II\n")
	}
}

func ch4ex05() {
	var x, y float64
	fmt.Printf("Ex05. Область I или II (рис 4.2)\n")
	fmt.Printf("Ex05. Введите координаты x и y: ")
	fmt.Scan(&x, &y)

	if x < 5 {
		fmt.Printf("Ex05. Ответ: Область I\n")
	} else {
		fmt.Printf("Ex05. Ответ: Область II\n")
	}
}

func ch4ex06() {
	var x float64
	fmt.Printf("Ex06. Значение функции по графику\n")
	fmt.Printf("Ex06. Введите x: ")
	fmt.Scan(&x)

	var y float64
	if x < 2 {
		y = 1
	} else if x > 2 {
		y = 2
	} else {
		y = 0
	}
	fmt.Printf("Ex06a. Ответ: y = %.1f\n", y)

	if x < 0 {
		y = -x
	} else {
		y = x
	}
	fmt.Printf("Ex06b. Ответ: y = %.1f\n", y)
}

func ch4ex07() {
	var x float64
	fmt.Printf("Ex07. Вычисление сложной функции\n")
	fmt.Printf("Ex07. Введите x: ")
	fmt.Scan(&x)

	var k float64
	if math.Sin(x) < 0 {
		k = x * x
	} else {
		k = math.Abs(x)
	}

	var f float64
	if k < x {
		f = k * x
	} else {
		f = k + x
	}
	fmt.Printf("Ex07. Ответ: f(%.2f) = %.4f\n", x, f)
}

func ch4ex08() {
	var x float64
	fmt.Printf("Ex08. Вычисление сложной функции\n")
	fmt.Printf("Ex08. Введите x: ")
	fmt.Scan(&x)

	var k float64
	if math.Sin(x) >= 0 {
		k = x * x
	} else {
		k = math.Abs(x)
	}

	var f float64
	if x < k {
		f = math.Abs(x)
	} else {
		f = k * x
	}
	fmt.Printf("Ex08. Ответ: f(%.2f) = %.4f\n", x, f)
}

func ch4ex09() {
	var a, b float64
	fmt.Printf("Ex09. Максимум и минимум\n")
	fmt.Printf("Ex09. Введите два числа: ")
	fmt.Scan(&a, &b)

	var max, min float64
	if a > b {
		max, min = a, b
	} else {
		max, min = b, a
	}
	fmt.Printf("Ex09. Ответ: Максимум = %.2f, Минимум = %.2f\n", max, min)
}

func ch4ex10() {
	var km, feet float64
	fmt.Printf("Ex10. Сравнение расстояний\n")
	fmt.Printf("Ex10. Введите расстояние в км и в футах: ")
	fmt.Scan(&km, &feet)

	metersFromKm := km * 1000
	metersFromFeet := feet * 0.3048

	if metersFromKm < metersFromFeet {
		fmt.Printf("Ex10. Ответ: Меньше расстояние в километрах (%.2f м)\n", metersFromKm)
	} else {
		fmt.Printf("Ex10. Ответ: Меньше расстояние в футах (%.2f м)\n", metersFromFeet)
	}
}

func ch4ex11() {
	var kmh, ms float64
	fmt.Printf("Ex11. Сравнение скоростей\n")
	fmt.Printf("Ex11. Введите скорость в км/ч и м/с: ")
	fmt.Scan(&kmh, &ms)

	msFromKmh := kmh * 1000 / 3600

	if msFromKmh > ms {
		fmt.Printf("Ex11. Ответ: Больше скорость в км/ч (%.2f м/с)\n", msFromKmh)
	} else {
		fmt.Printf("Ex11. Ответ: Больше скорость в м/с (%.2f м/с)\n", ms)
	}
}

func ch4ex12() {
	var radius, side float64
	fmt.Printf("Ex12. Сравнение площадей\n")
	fmt.Printf("Ex12. Введите радиус круга и сторону квадрата: ")
	fmt.Scan(&radius, &side)

	circleArea := math.Pi * radius * radius
	squareArea := side * side

	if circleArea > squareArea {
		fmt.Printf("Ex12. Ответ: Больше площадь круга (%.2f)\n", circleArea)
	} else {
		fmt.Printf("Ex12. Ответ: Больше площадь квадрата (%.2f)\n", squareArea)
	}
}

func ch4ex13() {
	var v1, m1, v2, m2 float64
	fmt.Printf("Ex13. Сравнение плотностей\n")
	fmt.Printf("Ex13. Введите объем и массу первого тела: ")
	fmt.Scan(&v1, &m1)
	fmt.Printf("Ex13. Введите объем и массу второго тела: ")
	fmt.Scan(&v2, &m2)

	density1 := m1 / v1
	density2 := m2 / v2

	if density1 > density2 {
		fmt.Printf("Ex13. Ответ: Больше плотность первого материала (%.2f)\n", density1)
	} else {
		fmt.Printf("Ex13. Ответ: Больше плотность второго материала (%.2f)\n", density2)
	}
}

func ch4ex14() {
	var r1, u1, r2, u2 float64
	fmt.Printf("Ex14. Сравнение токов\n")
	fmt.Printf("Ex14. Введите сопротивление и напряжение первого участка: ")
	fmt.Scan(&r1, &u1)
	fmt.Printf("Ex14. Введите сопротивление и напряжение второго участка: ")
	fmt.Scan(&r2, &u2)

	i1 := u1 / r1
	i2 := u2 / r2

	if i1 < i2 {
		fmt.Printf("Ex14. Ответ: Меньший ток на первом участке (%.2f А)\n", i1)
	} else {
		fmt.Printf("Ex14. Ответ: Меньший ток на втором участке (%.2f А)\n", i2)
	}
}

func ch4ex15() {
	var a, b, c float64
	fmt.Printf("Ex15. Проверка корней квадратного уравнения\n")
	fmt.Printf("Ex15. Введите коэффициенты a, b, c: ")
	fmt.Scan(&a, &b, &c)

	discriminant := b*b - 4*a*c

	if discriminant >= 0 {
		fmt.Printf("Ex15. Ответ: Уравнение имеет вещественные корни\n")
	} else {
		fmt.Printf("Ex15. Ответ: Уравнение не имеет вещественных корней\n")
	}
}

func ch4ex16() {
	var a, b, c float64
	fmt.Printf("Ex16. Решение квадратного уравнения\n")
	fmt.Printf("Ex16. Введите коэффициенты a, b, c: ")
	fmt.Scan(&a, &b, &c)

	discriminant := b*b - 4*a*c

	if discriminant >= 0 {
		x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		x2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		fmt.Printf("Ex16. Ответ: Корни: %.2f и %.2f\n", x1, x2)
	} else {
		fmt.Printf("Ex16. Ответ: Уравнение не имеет вещественных корней\n")
	}
}

func ch4ex17() {
	var birthYear, birthMonth, currentYear, currentMonth int
	fmt.Printf("Ex17. Определение возраста\n")
	fmt.Printf("Ex17. Введите год и месяц рождения: ")
	fmt.Scan(&birthYear, &birthMonth)
	fmt.Printf("Ex17. Введите текущий год и месяц: ")
	fmt.Scan(&currentYear, &currentMonth)

	age := currentYear - birthYear
	if currentMonth < birthMonth {
		age--
	}
	fmt.Printf("Ex17. Ответ: Возраст: %d лет\n", age)
}

func ch4ex18() {
	var circleArea, squareArea float64
	fmt.Printf("Ex18. Вписывание круга и квадрата\n")
	fmt.Printf("Ex18. Введите площадь круга и квадрата: ")
	fmt.Scan(&circleArea, &squareArea)

	circleRadius := math.Sqrt(circleArea / math.Pi)
	squareSide := math.Sqrt(squareArea)
	circleDiameter := 2 * circleRadius
	squareDiagonal := squareSide * math.Sqrt(2)

	if circleDiameter <= squareSide {
		fmt.Printf("Ex18a. Ответ: Круг уместится в квадрате\n")
	} else {
		fmt.Printf("Ex18a. Ответ: Круг не уместится в квадрате\n")
	}

	if squareDiagonal <= circleDiameter {
		fmt.Printf("Ex18b. Ответ: Квадрат уместится в круге\n")
	} else {
		fmt.Printf("Ex18b. Ответ: Квадрат не уместится в круге\n")
	}
}

func ch4ex19() {
	var circleArea, triangleArea float64
	fmt.Printf("Ex19. Вписывание круга и треугольника\n")
	fmt.Printf("Ex19. Введите площадь круга и треугольника: ")
	fmt.Scan(&circleArea, &triangleArea)

	circleRadius := math.Sqrt(circleArea / math.Pi)
	triangleSide := math.Sqrt(4 * triangleArea / math.Sqrt(3))

	inscribedRadius := triangleSide * math.Sqrt(3) / 6
	circumscribedRadius := triangleSide * math.Sqrt(3) / 3

	if circleRadius <= inscribedRadius {
		fmt.Printf("Ex19a. Ответ: Круг уместится в треугольнике\n")
	} else {
		fmt.Printf("Ex19a. Ответ: Круг не уместится в треугольнике\n")
	}

	if circumscribedRadius <= circleRadius {
		fmt.Printf("Ex19b. Ответ: Треугольник уместится в круге\n")
	} else {
		fmt.Printf("Ex19b. Ответ: Треугольник не уместится в круге\n")
	}
}

func ch4ex20() {
	var x1l, y1l, x1r, y1r, x2l, y2l, x2r, y2r float64

	fmt.Printf("Ex20. Минимальный охватывающий прямоугольник\n")
	fmt.Printf("Ex20. Введите координаты левого нижнего и правого верхнего углов первого прямоугольника: ")
	fmt.Scan(&x1l, &y1l, &x1r, &y1r)
	fmt.Printf("Ex20. Введите координаты левого нижнего и правого верхнего углов второго прямоугольника: ")
	fmt.Scan(&x2l, &y2l, &x2r, &y2r)

	minX, minY := math.Min(x1l, x2l), math.Min(y1l, y2l)
	maxX, maxY := math.Max(x1r, x2r), math.Max(y1r, y2r)

	fmt.Printf("Ex20. Ответ: Левый нижний: (%.2f, %.2f), Правый верхний: (%.2f, %.2f)\n",
		minX, minY, maxX, maxY)
}

func ch4ex21() {
	var x1, y1, w1, h1, x2, y2, w2, h2 float64

	fmt.Printf("Ex21. Минимальный охватывающий прямоугольник\n")
	fmt.Printf("Ex21. Введите координаты левого нижнего угла и размеры первого прямоугольника: ")
	fmt.Scan(&x1, &y1, &w1, &h1)
	fmt.Printf("Ex21. Введите координаты левого нижнего угла и размеры второго прямоугольника: ")
	fmt.Scan(&x2, &y2, &w2, &h2)

	minX, minY := math.Min(x1, x2), math.Min(y1, y2)
	maxX, maxY := math.Max(x1+w1, x2+w2), math.Max(y1+h1, y2+h2)

	fmt.Printf("Ex21. Ответ: Левый нижний: (%.2f, %.2f), Правый верхний: (%.2f, %.2f)\n",
		minX, minY, maxX, maxY)
}

func ch4ex22() {
	var n, n1 int
	fmt.Printf("Ex22. Проверка делимости\n")
	fmt.Printf("Ex22. Введите два числа: ")
	fmt.Scan(&n, &n1)

	if n%n1 == 0 {
		fmt.Printf("Ex22. Ответ: Частное: %d\n", n/n1)
	} else {
		fmt.Printf("Ex22. Ответ: %d на %d нацело не делится\n", n, n1)
	}
}

func ch4ex23() {
	var a, n int
	fmt.Printf("Ex23. Проверка делителя\n")
	fmt.Printf("Ex23. Введите число и возможный делитель: ")
	fmt.Scan(&n, &a)

	if n%a == 0 {
		fmt.Printf("Ex23. Ответ: %d является делителем %d\n", a, n)
	} else {
		fmt.Printf("Ex23. Ответ: %d не является делителем %d\n", a, n)
	}
}

func ch4ex24() {
	var n int
	fmt.Printf("Ex24. Анализ числа\n")
	fmt.Printf("Ex24. Введите натуральное число: ")
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Printf("Ex24a. Ответ: Число четное\n")
	} else {
		fmt.Printf("Ex24a. Ответ: Число нечетное\n")
	}

	if n%10 == 7 {
		fmt.Printf("Ex24b. Ответ: Число оканчивается на 7\n")
	} else {
		fmt.Printf("Ex24b. Ответ: Число не оканчивается на 7\n")
	}
}

func ch4ex25() {
	var n int
	fmt.Printf("Ex25. Следующее четное число\n")
	fmt.Printf("Ex25. Введите целое число: ")
	fmt.Scan(&n)

	nextEven := n + 2 - n%2
	fmt.Printf("Ex25. Ответ: %d\n", nextEven)
}

func ch4ex26() {
	var n int
	fmt.Printf("Ex26. Анализ двузначного числа\n")
	fmt.Printf("Ex26. Введите двузначное число: ")
	fmt.Scan(&n)

	a, b := n/10, n%10

	if a > b {
		fmt.Printf("Ex26a. Ответ: Первая цифра больше\n")
	} else if a < b {
		fmt.Printf("Ex26a. Ответ: Вторая цифра больше\n")
	} else {
		fmt.Printf("Ex26a. Ответ: Цифры равны\n")
	}

	if a == b {
		fmt.Printf("Ex26b. Ответ: Цифры одинаковы\n")
	} else {
		fmt.Printf("Ex26b. Ответ: Цифры различны\n")
	}
}

func ch4ex27() {
	var n int
	fmt.Printf("Ex27. Проверка равенства квадрата и суммы кубов\n")
	fmt.Printf("Ex27. Введите двузначное число: ")
	fmt.Scan(&n)

	a, b := n/10, n%10
	square := n * n
	sumCubes := a*a*a + b*b*b

	if square == 4*sumCubes {
		fmt.Printf("Ex27. Ответ: Равенство выполняется\n")
	} else {
		fmt.Printf("Ex27. Ответ: Равенство не выполняется\n")
	}
}

func ch4ex28() {
	var n, a int
	fmt.Printf("Ex28. Анализ суммы цифр\n")
	fmt.Printf("Ex28. Введите двузначное число и число a: ")
	fmt.Scan(&n, &a)

	sum := n/10 + n%10

	if sum >= 10 {
		fmt.Printf("Ex28a. Ответ: Сумма цифр двузначная\n")
	} else {
		fmt.Printf("Ex28a. Ответ: Сумма цифр не двузначная\n")
	}

	if sum > a {
		fmt.Printf("Ex28b. Ответ: Сумма цифр больше a\n")
	} else {
		fmt.Printf("Ex28b. Ответ: Сумма цифр не больше a\n")
	}
}

func ch4ex29() {
	var n int
	fmt.Printf("Ex29. Кратность суммы цифр\n")
	fmt.Printf("Ex29. Введите двузначное число: ")
	fmt.Scan(&n)

	sum := n/10 + n%10

	if sum%3 == 0 {
		fmt.Printf("Ex29a. Ответ: Сумма цифр кратна трем\n")
	} else {
		fmt.Printf("Ex29a. Ответ: Сумма цифр не кратна трем\n")
	}

	if sum%6 == 0 {
		fmt.Printf("Ex29b. Ответ: Сумма цифр кратна шести\n")
	} else {
		fmt.Printf("Ex29b. Ответ: Сумма цифр не кратна шести\n")
	}
}

func ch4ex30() {
	var n int
	fmt.Printf("Ex30. Палиндром\n")
	fmt.Printf("Ex30. Введите трехзначное число: ")
	fmt.Scan(&n)

	a, c := n/100, n%10

	if a == c {
		fmt.Printf("Ex30. Ответ: Число палиндром\n")
	} else {
		fmt.Printf("Ex30. Ответ: Число не палиндром\n")
	}
}

func ch4ex31() {
	var n int
	fmt.Printf("Ex31. Сравнение цифр трехзначного числа\n")
	fmt.Printf("Ex31. Введите трехзначное число: ")
	fmt.Scan(&n)

	a, b, c := n/100, n/10%10, n%10

	if a > c {
		fmt.Printf("Ex31a. Ответ: Первая цифра больше\n")
	} else if a < c {
		fmt.Printf("Ex31a. Ответ: Последняя цифра больше\n")
	} else {
		fmt.Printf("Ex31a. Ответ: Первая и последняя цифры равны\n")
	}

	if a > b {
		fmt.Printf("Ex31b. Ответ: Первая цифра больше\n")
	} else if a < b {
		fmt.Printf("Ex31b. Ответ: Вторая цифра больше\n")
	} else {
		fmt.Printf("Ex31b. Ответ: Первая и вторая цифры равны\n")
	}

	if b > c {
		fmt.Printf("Ex31c. Ответ: Вторая цифра больше\n")
	} else if b < c {
		fmt.Printf("Ex31c. Ответ: Последняя цифра больше\n")
	} else {
		fmt.Printf("Ex31c. Ответ: Вторая и последняя цифры равны\n")
	}
}

func ch4ex32() {
	var n int
	fmt.Printf("Ex32. Квадрат и сумма кубов цифр\n")
	fmt.Printf("Ex32. Введите трехзначное число: ")
	fmt.Scan(&n)

	a, b, c := n/100, n/10%10, n%10
	square := n * n
	sumCubes := a*a*a + b*b*b + c*c*c

	if square == sumCubes {
		fmt.Printf("Ex32. Ответ: Равенство выполняется\n")
	} else {
		fmt.Printf("Ex32. Ответ: Равенство не выполняется\n")
	}
}

func ch4ex33() {
	var n, a int
	fmt.Printf("Ex33. Анализ трехзначного числа\n")
	fmt.Printf("Ex33. Введите трехзначное число и число a: ")
	fmt.Scan(&n, &a)

	x, y, z := n/100, n/10%10, n%10
	sum := x + y + z
	product := x * y * z

	if sum >= 10 {
		fmt.Printf("Ex33a. Ответ: Сумма цифр двузначная\n")
	} else {
		fmt.Printf("Ex33a. Ответ: Сумма цифр не двузначная\n")
	}

	if product >= 100 {
		fmt.Printf("Ex33b. Ответ: Произведение цифр трехзначное\n")
	} else {
		fmt.Printf("Ex33b. Ответ: Произведение цифр не трехзначное\n")
	}

	if product > a {
		fmt.Printf("Ex33c. Ответ: Произведение цифр больше a\n")
	} else {
		fmt.Printf("Ex33c. Ответ: Произведение цифр не больше a\n")
	}

	if sum%5 == 0 {
		fmt.Printf("Ex33d. Ответ: Сумма цифр кратна пяти\n")
	} else {
		fmt.Printf("Ex33d. Ответ: Сумма цифр не кратна пяти\n")
	}

	if sum%a == 0 {
		fmt.Printf("Ex33e. Ответ: Сумма цифр кратна a\n")
	} else {
		fmt.Printf("Ex33e. Ответ: Сумма цифр не кратна a\n")
	}
}

func ch4ex34() {
	var n int
	fmt.Printf("Ex34. Анализ одинаковых цифр\n")
	fmt.Printf("Ex34. Введите трехзначное число: ")
	fmt.Scan(&n)

	a, b, c := n/100, n/10%10, n%10

	if a == b && b == c {
		fmt.Printf("Ex34a. Ответ: Все цифры одинаковые\n")
	} else {
		fmt.Printf("Ex34a. Ответ: Не все цифры одинаковые\n")
	}

	if a == b || b == c || a == c {
		fmt.Printf("Ex34b. Ответ: Есть одинаковые цифры\n")
	} else {
		fmt.Printf("Ex34b. Ответ: Нет одинаковых цифр\n")
	}
}

func ch4ex35() {
	var n, a int
	fmt.Printf("Ex35. Анализ четырехзначного числа\n")
	fmt.Printf("Ex35. Введите четырехзначное число и число a: ")
	fmt.Scan(&n, &a)

	w, x, y, z := n/1000, n/100%10, n/10%10, n%10

	if w+x == y+z {
		fmt.Printf("Ex35a. Ответ: Суммы пар цифр равны\n")
	} else {
		fmt.Printf("Ex35a. Ответ: Суммы пар цифр не равны\n")
	}

	sum := w + x + y + z
	if sum%3 == 0 {
		fmt.Printf("Ex35b. Ответ: Сумма цифр кратна трем\n")
	} else {
		fmt.Printf("Ex35b. Ответ: Сумма цифр не кратна трем\n")
	}

	product := w * x * y * z
	if product%4 == 0 {
		fmt.Printf("Ex35c. Ответ: Произведение цифр кратно четырем\n")
	} else {
		fmt.Printf("Ex35c. Ответ: Произведение цифр не кратно четырем\n")
	}

	if product%a == 0 {
		fmt.Printf("Ex35d. Ответ: Произведение цифр кратно a\n")
	} else {
		fmt.Printf("Ex35d. Ответ: Произведение цифр не кратно a\n")
	}
}

func ch4ex36() {
	var n int
	fmt.Printf("Ex36. Анализ последней цифры\n")
	fmt.Printf("Ex36. Введите натуральное число: ")
	fmt.Scan(&n)

	lastDigit := n % 10

	switch lastDigit % 2 {
	case 0:
		fmt.Printf("Ex36a. Ответ: Число оканчивается четной цифрой\n")
		fmt.Printf("Ex36b. Ответ: Число не оканчивается нечетной цифрой\n")
	case 1:
		fmt.Printf("Ex36a. Ответ: Число не оканчивается четной цифрой\n")
		fmt.Printf("Ex36b. Ответ: Число оканчивается нечетной цифрой\n")
	}
}

func ch4ex37() {
	var a, b int
	fmt.Printf("Ex37. Взаимная делимость\n")
	fmt.Printf("Ex37. Введите два числа: ")
	fmt.Scan(&a, &b)

	if b%a == 0 {
		fmt.Printf("Ex37. Ответ: a является делителем b\n")
	} else {
		fmt.Printf("Ex37. Ответ: a не является делителем b\n")
	}

	if a%b == 0 {
		fmt.Printf("Ex37. Ответ: b является делителем a\n")
	} else {
		fmt.Printf("Ex37. Ответ: b не является делителем a\n")
	}
}

func ch4ex38() {
	var a, b, c, d int
	fmt.Printf("Ex38. Размещение прямоугольников на столе\n")
	fmt.Printf("Ex38. Введите размеры стола (a b): ")
	fmt.Scan(&a, &b)
	fmt.Printf("Ex38. Введите размеры картонных прямоугольников (c d): ")
	fmt.Scan(&c, &d)

	count1 := (a / c) * (b / d)
	count2 := (a / d) * (b / c)

	if count1 > count2 {
		fmt.Printf("Ex38. Ответ: Выгоднее размещать длинной стороной вдоль длинной (шт: %d)\n", count1)
	} else if count2 > count1 {
		fmt.Printf("Ex38. Ответ: Выгоднее размещать длинной стороной вдоль короткой (шт: %d)\n", count2)
	} else {
		fmt.Printf("Ex38. Ответ: Оба способа дают одинаковый результат (шт: %d)\n", count1)
	}
}

func ch4ex39() {
	var t float64
	fmt.Printf("Ex39. Светофор для пешеходов\n")
	fmt.Printf("Ex39. Введите время в минутах с начала часа: ")
	fmt.Scan(&t)

	timeInCycle := math.Mod(t, 5)

	if timeInCycle < 3 {
		fmt.Printf("Ex39. Ответ: Горит зеленый сигнал\n")
	} else {
		fmt.Printf("Ex39. Ответ: Горит красный сигнал\n")
	}
}

func ch4ex40() {
	var x float64
	fmt.Printf("Ex40. Принадлежность интервалу\n")
	fmt.Printf("Ex40. Введите число: ")
	fmt.Scan(&x)

	if x > -5 && x < 3 {
		fmt.Printf("Ex40. Ответ: Число принадлежит интервалу (-5, 3)\n")
	} else {
		fmt.Printf("Ex40. Ответ: Число не принадлежит интервалу (-5, 3)\n")
	}
}

func ch4ex41() {
	var n int
	fmt.Printf("Ex41. Проверка на двузначность\n")
	fmt.Printf("Ex41. Введите натуральное число: ")
	fmt.Scan(&n)

	if n >= 10 && n <= 99 {
		fmt.Printf("Ex41. Ответ: Число двузначное\n")
	} else {
		fmt.Printf("Ex41. Ответ: Число не двузначное\n")
	}
}

func ch4ex42() {
	var x, y float64
	fmt.Printf("Ex42. Принадлежность области\n")
	fmt.Printf("Ex42. Введите координаты x и y: ")
	fmt.Scan(&x, &y)

	if (x-1)*(x-1)+(y-1)*(y-1) < 1 {
		fmt.Printf("Ex42a. Ответ: Точка принадлежит области\n")
	} else {
		fmt.Printf("Ex42a. Ответ: Точка не принадлежит области\n")
	}

	if x > -2 && x < 2 && y > -2 && y < 2 && (x*x+y*y > 1) {
		fmt.Printf("Ex42b. Ответ: Точка принадлежит области\n")
	} else {
		fmt.Printf("Ex42b. Ответ: Точка не принадлежит области\n")
	}

	if x > 0 && x*x+y*y < 4 {
		fmt.Printf("Ex42c. Ответ: Точка принадлежит области\n")
	} else {
		fmt.Printf("Ex42c. Ответ: Точка не принадлежит области\n")
	}

	if x*x+y*y > 4 && x*x+y*y < 16 {
		fmt.Printf("Ex42d. Ответ: Точка принадлежит области\n")
	} else {
		fmt.Printf("Ex42d. Ответ: Точка не принадлежит области\n")
	}
}

func ch4ex43() {
	var x, y float64
	fmt.Printf("Ex43. Принадлежность области I\n")
	fmt.Printf("Ex43. Введите координаты x и y: ")
	fmt.Scan(&x, &y)

	if x > 0 && x < 2 && y > 0 && y < 2 {
		fmt.Printf("Ex43. Ответ: Точка принадлежит области I\n")
	} else {
		fmt.Printf("Ex43. Ответ: Точка не принадлежит области I\n")
	}
}

func ch4ex44() {
	var x, y float64
	fmt.Printf("Ex44. Принадлежность области I или III\n")
	fmt.Printf("Ex44. Введите координаты x и y: ")
	fmt.Scan(&x, &y)

	inRegionI := (x > -1 && x < 0 && y > 0)
	inRegionIII := (x > 5 && x < 7 && y < 0)

	if inRegionI {
		fmt.Printf("Ex44. Ответ: Точка принадлежит области I\n")
	} else if inRegionIII {
		fmt.Printf("Ex44. Ответ: Точка принадлежит области III\n")
	} else {
		fmt.Printf("Ex44. Ответ: Точка не принадлежит ни области I, ни области III\n")
	}
}

func ch4ex45() {
	var x float64
	fmt.Printf("Ex45. Вычисление функции\n")
	fmt.Printf("Ex45. Введите x: ")
	fmt.Scan(&x)

	var f float64
	if x >= -2.4 && x <= 5.7 {
		f = x * x
	} else {
		f = 4
	}
	fmt.Printf("Ex45. Ответ: f(%.2f) = %.4f\n", x, f)
}

func ch4ex46() {
	var x float64
	fmt.Printf("Ex46. Вычисление функции\n")
	fmt.Printf("Ex46. Введите x: ")
	fmt.Scan(&x)

	var f float64
	if x >= 0.2 && x <= 0.9 {
		f = math.Sin(x)
	} else {
		f = 1
	}
	fmt.Printf("Ex46. Ответ: f(%.2f) = %.4f\n", x, f)
}

func ch4ex47() {
	var a, b, c float64
	fmt.Printf("Ex47. Проверка неравенств\n")
	fmt.Printf("Ex47. Введите три числа: ")
	fmt.Scan(&a, &b, &c)

	if a < b && b < c {
		fmt.Printf("Ex47a. Ответ: Неравенство a < b < c выполняется\n")
	} else {
		fmt.Printf("Ex47a. Ответ: Неравенство a < b < c не выполняется\n")
	}

	if b > a && a > c {
		fmt.Printf("Ex47b. Ответ: Неравенство b > a > c выполняется\n")
	} else {
		fmt.Printf("Ex47b. Ответ: Неравенство b > a > c не выполняется\n")
	}
}

func ch4ex48() {
	var a, b int
	fmt.Printf("Ex48. Взаимная делимость\n")
	fmt.Printf("Ex48. Введите два целых числа: ")
	fmt.Scan(&a, &b)

	if a%b == 0 || b%a == 0 {
		fmt.Printf("Ex48. Ответ: Хотя бы одно число является делителем другого\n")
	} else {
		fmt.Printf("Ex48. Ответ: Ни одно число не является делителем другого\n")
	}
}

func ch4ex49() {
	var n int
	fmt.Printf("Ex49. Поиск цифр в двузначном числе\n")
	fmt.Printf("Ex49. Введите двузначное число: ")
	fmt.Scan(&n)

	a, b := n/10, n%10

	if a == 3 || b == 3 {
		fmt.Printf("Ex49a. Ответ: Цифра 3 входит в число\n")
	} else {
		fmt.Printf("Ex49a. Ответ: Цифра 3 не входит в число\n")
	}

	if a == 4 || b == 4 {
		fmt.Printf("Ex49b. Ответ: Цифра 4 входит в число\n")
	} else {
		fmt.Printf("Ex49b. Ответ: Цифра 4 не входит в число\n")
	}
}

func ch4ex50() {
	var n int
	fmt.Printf("Ex50. Поиск цифр в двузначном числе\n")
	fmt.Printf("Ex50. Введите двузначное число: ")
	fmt.Scan(&n)

	a, b := n/10, n%10

	if a == 4 || b == 4 || a == 7 || b == 7 {
		fmt.Printf("Ex50a. Ответ: Цифры 4 или 7 входят в число\n")
	} else {
		fmt.Printf("Ex50a. Ответ: Цифры 4 или 7 не входят в число\n")
	}

	if a == 3 || b == 3 || a == 6 || b == 6 || a == 9 || b == 9 {
		fmt.Printf("Ex50b. Ответ: Цифры 3, 6 или 9 входят в число\n")
	} else {
		fmt.Printf("Ex50b. Ответ: Цифры 3, 6 или 9 не входят в число\n")
	}
}

func ch4ex51() {
	var n int
	fmt.Printf("Ex51. Поиск цифр в трехзначном числе\n")
	fmt.Printf("Ex51. Введите трехзначное число: ")
	fmt.Scan(&n)

	a, b, c := n/100, n/10%10, n%10

	if a == 6 || b == 6 || c == 6 {
		fmt.Printf("Ex51a. Ответ: Цифра 6 входит в число\n")
	} else {
		fmt.Printf("Ex51a. Ответ: Цифра 6 не входит в число\n")
	}

	if a == 5 || b == 5 || c == 5 {
		fmt.Printf("Ex51b. Ответ: Цифра 5 входит в число\n")
	} else {
		fmt.Printf("Ex51b. Ответ: Цифра 5 не входит в число\n")
	}
}

func ch4ex52() {
	var n int
	fmt.Printf("Ex52. Поиск цифры 6 в трехзначном числе\n")
	fmt.Printf("Ex52. Введите трехзначное число: ")
	fmt.Scan(&n)

	a, b, c := n/100, n/10%10, n%10

	if a == 6 || b == 6 || c == 6 {
		fmt.Printf("Ex52. Ответ: Цифра 6 входит в число\n")
	} else {
		fmt.Printf("Ex52. Ответ: Цифра 6 не входит в число\n")
	}
}

func ch4ex53() {
	var n int
	fmt.Printf("Ex53. Поиск цифр в трехзначном числе\n")
	fmt.Printf("Ex53. Введите трехзначное число: ")
	fmt.Scan(&n)

	a, b, c := n/100, n/10%10, n%10

	if a == 4 || b == 4 || c == 4 || a == 7 || b == 7 || c == 7 {
		fmt.Printf("Ex53a. Ответ: Цифры 4 или 7 входят в число\n")
	} else {
		fmt.Printf("Ex53a. Ответ: Цифры 4 или 7 не входят в число\n")
	}

	if a == 3 || b == 3 || c == 3 || a == 6 || b == 6 || c == 6 || a == 9 || b == 9 || c == 9 {
		fmt.Printf("Ex53b. Ответ: Цифры 3, 6 или 9 входят в число\n")
	} else {
		fmt.Printf("Ex53b. Ответ: Цифры 3, 6 или 9 не входят в число\n")
	}
}

func ch4ex54() {
	var n int
	fmt.Printf("Ex54. Поиск цифр в четырехзначном числе\n")
	fmt.Printf("Ex54. Введите четырехзначное число: ")
	fmt.Scan(&n)

	a, b, c, d := n/1000, n/100%10, n/10%10, n%10

	if a == 4 || b == 4 || c == 4 || d == 4 {
		fmt.Printf("Ex54a. Ответ: Цифра 4 входит в число\n")
	} else {
		fmt.Printf("Ex54a. Ответ: Цифра 4 не входит в число\n")
	}

	if a == 5 || b == 5 || c == 5 || d == 5 {
		fmt.Printf("Ex54b. Ответ: Цифра 5 входит в число\n")
	} else {
		fmt.Printf("Ex54b. Ответ: Цифра 5 не входит в число\n")
	}
}

func ch4ex55() {
	var n int
	fmt.Printf("Ex55. Поиск цифр в четырехзначном числе\n")
	fmt.Printf("Ex55. Введите четырехзначное число: ")
	fmt.Scan(&n)

	a, b, c, d := n/1000, n/100%10, n/10%10, n%10

	// a) цифры 2 или 7
	if a == 2 || b == 2 || c == 2 || d == 2 || a == 7 || b == 7 || c == 7 || d == 7 {
		fmt.Printf("Ex55a. Ответ: Цифры 2 или 7 входят в число\n")
	} else {
		fmt.Printf("Ex55a. Ответ: Цифры 2 или 7 не входят в число\n")
	}

	// b) цифры 3, 6 или 9
	if a == 3 || b == 3 || c == 3 || d == 3 || a == 6 || b == 6 || c == 6 || d == 6 || a == 9 || b == 9 || c == 9 || d == 9 {
		fmt.Printf("Ex55b. Ответ: Цифры 3, 6 или 9 входят в число\n")
	} else {
		fmt.Printf("Ex55b. Ответ: Цифры 3, 6 или 9 не входят в число\n")
	}
}

func ch4ex56() {
	var n int
	fmt.Printf("Ex56. Симметричность четырехзначного числа\n")
	fmt.Printf("Ex56. Введите четырехзначное число: ")
	fmt.Scan(&n)

	a, b, c, d := n/1000, n/100%10, n/10%10, n%10

	if a == d && b == c {
		fmt.Printf("Ex56. Ответ: Число симметричное\n")
	} else {
		fmt.Printf("Ex56. Ответ: Число не симметричное\n")
	}
}

func ch4ex57() {
	var a, b, c, d int
	fmt.Printf("Ex57. Проверка остатка от деления\n")
	fmt.Printf("Ex57. Введите a, b, c, d: ")
	fmt.Scan(&a, &b, &c, &d)

	remainder := a % b
	if remainder == c || remainder == d {
		fmt.Printf("Ex57. Ответ: Остаток равен c или d\n")
	} else {
		fmt.Printf("Ex57. Ответ: Остаток не равен ни c, ни d\n")
	}
}

func ch4ex58() {
	var a, b, c float64
	fmt.Printf("Ex58. Поиск равных чисел\n")
	fmt.Printf("Ex58. Введите три числа: ")
	fmt.Scan(&a, &b, &c)

	if a == b || b == c || a == c {
		fmt.Printf("Ex58. Ответ: Есть хотя бы одна пара равных чисел\n")
	} else {
		fmt.Printf("Ex58. Ответ: Нет равных чисел\n")
	}
}

func ch4ex59() {
	var a, b, c float64
	fmt.Printf("Ex59. Тип треугольника\n")
	fmt.Printf("Ex59. Введите три стороны треугольника: ")
	fmt.Scan(&a, &b, &c)

	if a == b && b == c {
		fmt.Printf("Ex59a. Ответ: Треугольник равносторонний\n")
	} else {
		fmt.Printf("Ex59a. Ответ: Треугольник не равносторонний\n")
	}

	if a == b || b == c || a == c {
		fmt.Printf("Ex59b. Ответ: Треугольник равнобедренный\n")
	} else {
		fmt.Printf("Ex59b. Ответ: Треугольник не равнобедренный\n")
	}
}

func ch4ex60() {
	var h1, h2, h3 float64
	fmt.Printf("Ex60. Сравнение роста\n")
	fmt.Printf("Ex60. Введите рост трех человек: ")
	fmt.Scan(&h1, &h2, &h3)

	if h1 == h2 && h2 == h3 {
		fmt.Printf("Ex60. Ответ: Рост всех трех человек одинаковый\n")
	} else {
		fmt.Printf("Ex60. Ответ: Рост не одинаковый\n")
	}
}

func ch4ex61() {
	var a, b, c float64
	fmt.Printf("Ex61. Решение квадратного уравнения\n")
	fmt.Printf("Ex61. Введите коэффициенты a, b, c: ")
	fmt.Scan(&a, &b, &c)

	discriminant := b*b - 4*a*c

	if discriminant > 0 {
		x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		x2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		fmt.Printf("Ex61. Ответ: Два корня: %.2f и %.2f\n", x1, x2)
	} else if discriminant == 0 {
		x := -b / (2 * a)
		fmt.Printf("Ex61. Ответ: Один корень: %.2f\n", x)
	} else {
		fmt.Printf("Ex61. Ответ: Действительных корней нет\n")
	}
}
