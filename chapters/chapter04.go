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
		ch4ex57, ch4ex58, ch4ex59, ch4ex60, ch4ex61, ch4ex62, ch4ex63, ch4ex64,
		ch4ex65, ch4ex66, ch4ex67, ch4ex68, ch4ex69, ch4ex70, ch4ex71, ch4ex72,
		ch4ex73, ch4ex74, ch4ex75, ch4ex76, ch4ex77, ch4ex78, ch4ex79, ch4ex80,
		ch4ex81, ch4ex82, ch4ex83, ch4ex84, ch4ex85, ch4ex86, ch4ex87, ch4ex88,
		ch4ex89, ch4ex90, ch4ex91, ch4ex92, ch4ex93, ch4ex94, ch4ex95, ch4ex96,
		ch4ex97, ch4ex98, ch4ex99, ch4ex100, ch4ex101, ch4ex102, ch4ex103, ch4ex104,
		ch4ex105, ch4ex106, ch4ex107, ch4ex108, ch4ex109, ch4ex110, ch4ex111, ch4ex112,
		ch4ex113, ch4ex114, ch4ex115, ch4ex116, ch4ex117, ch4ex118, ch4ex119, ch4ex120,
		ch4ex121, ch4ex122, ch4ex123, ch4ex124, ch4ex125, ch4ex126, ch4ex127, ch4ex128,
		ch4ex129, ch4ex130, ch4ex131, ch4ex132, ch4ex133, ch4ex134, ch4ex135,
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

	include27, include369 := false, false

	for temp := n; temp > 0; temp /= 10 {
		digit := temp % 10
		if digit == 2 || digit == 7 {
			include27 = true
		}
		if digit == 3 || digit == 6 || digit == 9 {
			include369 = true
		}
	}

	if include27 {
		fmt.Printf("Ex55a. Ответ: Цифры 2 или 7 входят в число\n")
	} else {
		fmt.Printf("Ex55a. Ответ: Цифры 2 или 7 не входят в число\n")
	}

	if include369 {
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

func ch4ex62() {
	var a, b, c, d float64
	fmt.Printf("Ex62. Размещение прямоугольников\n")
	fmt.Printf("Ex62. Введите стороны первого прямоугольника (a b): ")
	fmt.Scan(&a, &b)
	fmt.Printf("Ex62. Введите стороны второго прямоугольника (c d): ")
	fmt.Scan(&c, &d)

	fits1 := (a <= c && b <= d) || (a <= d && b <= c)
	fits2 := (b <= c && a <= d) || (b <= d && a <= c)

	if fits1 || fits2 {
		fmt.Printf("Ex62. Ответ: Прямоугольник можно разместить\n")
	} else {
		fmt.Printf("Ex62. Ответ: Прямоугольник нельзя разместить\n")
	}
}

func ch4ex63() {
	var a, b, c, d float64
	fmt.Printf("Ex63. Открытка в конверте\n")
	fmt.Printf("Ex63. Введите размеры конверта (a b мм): ")
	fmt.Scan(&a, &b)
	fmt.Printf("Ex63. Введите размеры открытки (c d мм): ")
	fmt.Scan(&c, &d)

	availableWidth := a - 2
	availableHeight := b - 2

	fits := (c <= availableWidth && d <= availableHeight) ||
		(d <= availableWidth && c <= availableHeight)

	if fits {
		fmt.Printf("Ex63. Ответ: Открытка помещается в конверт\n")
	} else {
		fmt.Printf("Ex63. Ответ: Открытка не помещается в конверт\n")
	}
}

func ch4ex64() {
	var a, b, d float64
	fmt.Printf("Ex64. Голова в форточке\n")
	fmt.Printf("Ex64. Введите размеры форточки (a b см): ")
	fmt.Scan(&a, &b)
	fmt.Printf("Ex64. Введите диаметр головы (d см): ")
	fmt.Scan(&d)

	availableWidth := a - 2
	availableHeight := b - 2

	passes := d <= availableWidth && d <= availableHeight

	if passes {
		fmt.Printf("Ex64. Ответ: Вася сможет высунуть голову\n")
	} else {
		fmt.Printf("Ex64. Ответ: Вася не сможет высунуть голову\n")
	}
}

func ch4ex65() {
	var a, b, c, x, y float64
	fmt.Printf("Ex65. Кирпич в отверстии\n")
	fmt.Printf("Ex65. Введите размеры кирпича (a b c): ")
	fmt.Scan(&a, &b, &c)
	fmt.Printf("Ex65. Введите размеры отверстия (x y): ")
	fmt.Scan(&x, &y)

	passes := (a <= x && b <= y) || (a <= y && b <= x) ||
		(a <= x && c <= y) || (a <= y && c <= x) ||
		(b <= x && c <= y) || (b <= y && c <= x)

	if passes {
		fmt.Printf("Ex65. Ответ: Кирпич проходит в отверстие\n")
	} else {
		fmt.Printf("Ex65. Ответ: Кирпич не проходит в отверстие\n")
	}
}

func ch4ex66() {
	var a1, a2, a3, b1, b2, b3 float64
	fmt.Printf("Ex66. Коробка в чемодане\n")
	fmt.Printf("Ex66. Введите размеры чемодана (a1 a2 a3 см): ")
	fmt.Scan(&a1, &a2, &a3)
	fmt.Printf("Ex66. Введите размеры коробки (b1 b2 b3 см): ")
	fmt.Scan(&b1, b2, b3)

	fits := (b1 <= a1 && b2 <= a2 && b3 <= a3) ||
		(b1 <= a1 && b3 <= a2 && b2 <= a3) ||
		(b2 <= a1 && b1 <= a2 && b3 <= a3) ||
		(b2 <= a1 && b3 <= a2 && b1 <= a3) ||
		(b3 <= a1 && b1 <= a2 && b2 <= a3) ||
		(b3 <= a1 && b2 <= a2 && b1 <= a3)

	if fits {
		fmt.Printf("Ex66. Ответ: Можно сэкономить, поместив коробку в чемодан\n")
	} else {
		fmt.Printf("Ex66. Ответ: Нельзя сэкономить\n")
	}
}

func ch4ex67() {
	var n int
	fmt.Printf("Ex67. Счастливый билет\n")
	fmt.Printf("Ex67. Введите шестизначное число: ")
	fmt.Scan(&n)

	d1 := n / 100000
	d2 := n / 10000 % 10
	d3 := n / 1000 % 10
	d4 := n / 100 % 10
	d5 := n / 10 % 10
	d6 := n % 10

	sumFirst := d1 + d2 + d3
	sumLast := d4 + d5 + d6

	if sumFirst == sumLast {
		fmt.Printf("Ex67. Ответ: Число счастливое\n")
	} else {
		fmt.Printf("Ex67. Ответ: Число не счастливое\n")
	}
}

func ch4ex68() {
	var year int
	fmt.Printf("Ex68. Високосный год\n")
	fmt.Printf("Ex68. Введите год: ")
	fmt.Scan(&year)

	isLeap := (year%4 == 0 && year%100 != 0) || (year%400 == 0)

	if isLeap {
		fmt.Printf("Ex68. Ответ: Год високосный\n")
	} else {
		fmt.Printf("Ex68. Ответ: Год не високосный\n")
	}
}

func ch4ex69() {
	var a, b, c, d, e int
	fmt.Printf("Ex69. Кости домино на столе\n")
	fmt.Printf("Ex69. Введите размеры стола (a b): ")
	fmt.Scan(&a, &b)
	fmt.Printf("Ex69. Введите размеры кости домино (c d e): ")
	fmt.Scan(&c, &d, &e)

	countFace1 := (a / c) * (b / d)
	countFace2 := (a / d) * (b / c)
	countFace3 := (a / c) * (b / e)
	countFace4 := (a / e) * (b / c)
	countFace5 := (a / d) * (b / e)
	countFace6 := (a / e) * (b / d)

	maxCount := countFace1
	if countFace2 > maxCount {
		maxCount = countFace2
	}
	if countFace3 > maxCount {
		maxCount = countFace3
	}
	if countFace4 > maxCount {
		maxCount = countFace4
	}
	if countFace5 > maxCount {
		maxCount = countFace5
	}
	if countFace6 > maxCount {
		maxCount = countFace6
	}

	fmt.Printf("Ex69. Ответ: Максимальное количество костей: %d\n", maxCount)
}

func ch4ex70() {
	var k int
	fmt.Printf("Ex70. Определение выходного дня\n")
	fmt.Printf("Ex70. Введите номер дня года (1-365): ")
	fmt.Scan(&k)

	dayOfWeek := (k + 0) % 7

	if dayOfWeek == 0 || dayOfWeek == 6 {
		fmt.Printf("Ex70. Ответ: Выходной день\n")
	} else {
		fmt.Printf("Ex70. Ответ: Рабочий день\n")
	}
}

func ch4ex71() {
	var alpha, v, R, h float64
	fmt.Printf("Ex71. Попадание снаряда в цель\n")
	fmt.Printf("Ex71. Введите угол α (град), начальную скорость v, расстояние R и высоту H: ")
	fmt.Scan(&alpha, &v, &R, &h)

	alphaRad := alpha * math.Pi / 180
	t := R / (v * math.Cos(alphaRad))
	y := v*t*math.Sin(alphaRad) - (9.8*t*t)/2

	if math.Abs(y-h) < 0.1 {
		fmt.Printf("Ex71. Ответ: Снаряд поразит цель\n")
	} else {
		fmt.Printf("Ex71. Ответ: Снаряд не поразит цель\n")
	}
}

func ch4ex72() {
	var x1, y1, w1, h1 float64
	var x2, y2, w2, h2 float64

	fmt.Printf("Ex72. Операции с прямоугольниками\n")
	fmt.Printf("Ex72. Введите координаты и размеры первого прямоугольника: ")
	fmt.Scan(&x1, &y1, &w1, &h1)
	fmt.Printf("Ex72. Введите координаты и размеры второго прямоугольника: ")
	fmt.Scan(&x2, &y2, &w2, &h2)

	x1r, y1r := x1+w1, y1+h1
	x2r, y2r := x2+w2, y2+h2

	firstInSecond := x1 >= x2 && y1 >= y2 && x1r <= x2r && y1r <= y2r
	secondInFirst := x2 >= x1 && y2 >= y1 && x2r <= x1r && y2r <= y1r
	oneInOther := firstInSecond || secondInFirst
	intersect := x1 < x2r && x1r > x2 && y1 < y2r && y1r > y2

	fmt.Printf("Ex72a. Первый во втором: %v\n", firstInSecond)
	fmt.Printf("Ex72b. Один в другом: %v\n", oneInOther)
	fmt.Printf("Ex72c. Пересекаются: %v\n", intersect)
}

func ch4ex73() {
	var a1, a2, b int
	fmt.Printf("Ex73. Разность двузначного и однозначного чисел по цифрам\n")
	fmt.Printf("Ex73. Введите цифры (единицы2х, десятки2х, однозначное): ")
	fmt.Scan(&a1, &a2, &b)

	units := a1 - b
	if units < 0 {
		units += 10
		a2--
	}

	tens := a2

	fmt.Printf("Ex73. Ответ: Десятки: %d, Единицы: %d\n", tens, units)
}

func ch4ex74() {
	var a1, a2, b1, b2 int
	fmt.Printf("Ex74. Разность двузначных чисел по цифрам\n")
	fmt.Printf("Ex74. Введите цифры (единицы1, десятки1, единицы2, десятки2): ")
	fmt.Scan(&a1, &a2, &b1, &b2)

	units := a1 - b1
	if units < 0 {
		units += 10
		a2--
	}

	tens := a2 - b2

	fmt.Printf("Ex74. Ответ: Десятки: %d, Единицы: %d\n", tens, units)
}

func ch4ex75() {
	var a1, a2, a3, b1, b2 int
	fmt.Printf("Ex75. Сумма трехзначного и двузначного чисел по цифрам\n")
	fmt.Printf("Ex75. Введите цифры (единицы3х, десятки3х, сотни3х, единицы2х, десятки2х): ")
	fmt.Scan(&a1, &a2, &a3, &b1, &b2)

	units := (a1 + b1) % 10
	carry1 := (a1 + b1) / 10

	tens := (a2 + b2 + carry1) % 10
	carry2 := (a2 + b2 + carry1) / 10

	hundreds := a3 + carry2

	fmt.Printf("Ex75. Ответ: Сотни: %d, Десятки: %d, Единицы: %d\n", hundreds, tens, units)
}

func ch4ex76() {
	var a, b, c, d int
	fmt.Printf("Ex76. Шахматные фигуры\n")
	fmt.Printf("Ex76. Введите координаты (a b c d): ")
	fmt.Scan(&a, &b, &c, &d)

	rookThreat := a == c || b == d
	fmt.Printf("Ex76a. Ладья: %v\n", rookThreat)

	bishopThreat := math.Abs(float64(a-c)) == math.Abs(float64(b-d))
	fmt.Printf("Ex76b. Слон: %v\n", bishopThreat)

	kingMove := math.Abs(float64(a-c)) <= 1 && math.Abs(float64(b-d)) <= 1
	fmt.Printf("Ex76c. Король: %v\n", kingMove)

	queenThreat := rookThreat || bishopThreat
	fmt.Printf("Ex76d. Ферзь: %v\n", queenThreat)

	whitePawnNormal := a == c && d-b == 1
	whitePawnCapture := math.Abs(float64(a-c)) == 1 && d-b == 1
	fmt.Printf("Ex76e. Белая пешка (норм/взятие): %v / %v\n", whitePawnNormal, whitePawnCapture)

	blackPawnNormal := a == c && b-d == 1
	blackPawnCapture := math.Abs(float64(a-c)) == 1 && b-d == 1
	fmt.Printf("Ex76f. Черная пешка (норм/взятие): %v / %v\n", blackPawnNormal, blackPawnCapture)

	knightThreat := (math.Abs(float64(a-c)) == 2 && math.Abs(float64(b-d)) == 1) ||
		(math.Abs(float64(a-c)) == 1 && math.Abs(float64(b-d)) == 2)
	fmt.Printf("Ex76g. Конь: %v\n", knightThreat)
}

func ch4ex77() {
	var a, b, c, d, e, f int
	fmt.Printf("Ex77. Ход белой фигуры под ударом черной\n")
	fmt.Printf("Ex77. Введите координаты (белая a b, черная c d, цель e f): ")
	fmt.Scan(&a, &b, &c, &d, &e, &f)

	isRookMove := func(x1, y1, x2, y2 int) bool {
		return x1 == x2 || y1 == y2
	}

	isBishopMove := func(x1, y1, x2, y2 int) bool {
		return math.Abs(float64(x1-x2)) == math.Abs(float64(y1-y2))
	}

	isQueenMove := func(x1, y1, x2, y2 int) bool {
		return isRookMove(x1, y1, x2, y2) || isBishopMove(x1, y1, x2, y2)
	}

	isKnightMove := func(x1, y1, x2, y2 int) bool {
		dx := math.Abs(float64(x1 - x2))
		dy := math.Abs(float64(y1 - y2))
		return (dx == 2 && dy == 1) || (dx == 1 && dy == 2)
	}

	isKingMove := func(x1, y1, x2, y2 int) bool {
		return math.Abs(float64(x1-x2)) <= 1 && math.Abs(float64(y1-y2)) <= 1
	}

	var variant string
	fmt.Printf("Выберите вариант (a-ф): ")
	fmt.Scan(&variant)

	var whiteCanMove, blackCanAttack bool

	switch variant {
	case "a": // ладья и ладья
		whiteCanMove = isRookMove(a, b, e, f)
		blackCanAttack = isRookMove(c, d, e, f)
	case "б": // ладья и ферзь
		whiteCanMove = isRookMove(a, b, e, f)
		blackCanAttack = isQueenMove(c, d, e, f)
	case "в": // ладья и конь
		whiteCanMove = isRookMove(a, b, e, f)
		blackCanAttack = isKnightMove(c, d, e, f)
	case "г": // ладья и слон
		whiteCanMove = isRookMove(a, b, e, f)
		blackCanAttack = isBishopMove(c, d, e, f)
	case "д": // ферзь и ферзь
		whiteCanMove = isQueenMove(a, b, e, f)
		blackCanAttack = isQueenMove(c, d, e, f)
	case "е": // ферзь и ладья
		whiteCanMove = isQueenMove(a, b, e, f)
		blackCanAttack = isRookMove(c, d, e, f)
	case "ж": // ферзь и конь
		whiteCanMove = isQueenMove(a, b, e, f)
		blackCanAttack = isKnightMove(c, d, e, f)
	case "з": // ферзь и слон
		whiteCanMove = isQueenMove(a, b, e, f)
		blackCanAttack = isBishopMove(c, d, e, f)
	case "и": // конь и конь
		whiteCanMove = isKnightMove(a, b, e, f)
		blackCanAttack = isKnightMove(c, d, e, f)
	case "к": // конь и ладья
		whiteCanMove = isKnightMove(a, b, e, f)
		blackCanAttack = isRookMove(c, d, e, f)
	case "л": // конь и ферзь
		whiteCanMove = isKnightMove(a, b, e, f)
		blackCanAttack = isQueenMove(c, d, e, f)
	case "м": // конь и слон
		whiteCanMove = isKnightMove(a, b, e, f)
		blackCanAttack = isBishopMove(c, d, e, f)
	case "н": // слон и слон
		whiteCanMove = isBishopMove(a, b, e, f)
		blackCanAttack = isBishopMove(c, d, e, f)
	case "о": // слон и ферзь
		whiteCanMove = isBishopMove(a, b, e, f)
		blackCanAttack = isQueenMove(c, d, e, f)
	case "п": // слон и конь
		whiteCanMove = isBishopMove(a, b, e, f)
		blackCanAttack = isKnightMove(c, d, e, f)
	case "р": // слон и ладья
		whiteCanMove = isBishopMove(a, b, e, f)
		blackCanAttack = isRookMove(c, d, e, f)
	case "с": // король и слон
		whiteCanMove = isKingMove(a, b, e, f)
		blackCanAttack = isBishopMove(c, d, e, f)
	case "т": // король и ферзь
		whiteCanMove = isKingMove(a, b, e, f)
		blackCanAttack = isQueenMove(c, d, e, f)
	case "у": // король и конь
		whiteCanMove = isKingMove(a, b, e, f)
		blackCanAttack = isKnightMove(c, d, e, f)
	case "ф": // король и ладья
		whiteCanMove = isKingMove(a, b, e, f)
		blackCanAttack = isRookMove(c, d, e, f)
	}

	fmt.Printf("Ex77. Ответ: Белая фигура %v пойти на поле (%d,%d) безопасно\n",
		map[bool]string{true: "может", false: "не может"}[whiteCanMove && !blackCanAttack], e, f)
}

func ch4ex78() {
	var a, b, c, d int
	fmt.Printf("Ex78. Цвет полей шахматной доски\n")
	fmt.Printf("Ex78. Введите координаты двух полей (a b c d): ")
	fmt.Scan(&a, &b, &c, &d)

	color1 := (a + b) % 2
	color2 := (c + d) % 2

	fmt.Printf("Ex78. Ответ: Поля %v цвета\n",
		map[bool]string{true: "одного", false: "разного"}[color1 == color2])
}

func ch4ex79() {
	var a, b, c float64
	fmt.Printf("Ex79. Существование треугольника\n")
	fmt.Printf("Ex79. Введите три стороны: ")
	fmt.Scan(&a, &b, &c)

	exists := a+b > c && a+c > b && b+c > a
	fmt.Printf("Ex79. Ответ: Треугольник %s\n",
		map[bool]string{true: "существует", false: "не существует"}[exists])
}

func ch4ex80() {
	var a, b, c float64
	fmt.Printf("Ex80. Прямоугольный треугольник\n")
	fmt.Printf("Ex80. Введите три стороны: ")
	fmt.Scan(&a, &b, &c)

	exists := a+b > c && a+c > b && b+c > a
	isRight := exists && (a*a+b*b == c*c || a*a+c*c == b*b || b*b+c*c == a*a)
	fmt.Printf("Ex80. Ответ: Треугольник %s\n",
		map[bool]string{true: "прямоугольный", false: "не прямоугольный"}[isRight])
}

func ch4ex81() {
	var a, b, c float64
	fmt.Printf("Ex81. Вид треугольника\n")
	fmt.Printf("Ex81. Введите три стороны: ")
	fmt.Scan(&a, &b, &c)

	exists := a+b > c && a+c > b && b+c > a
	if !exists {
		fmt.Printf("Ex81. Ответ: Треугольник не существует\n")
		return
	}

	var angleType string
	if a*a+b*b == c*c || a*a+c*c == b*b || b*b+c*c == a*a {
		angleType = "прямоугольный"
	} else if a*a+b*b > c*c && a*a+c*c > b*b && b*b+c*c > a*a {
		angleType = "остроугольный"
	} else {
		angleType = "тупоугольный"
	}

	var sideType string
	if a == b && b == c {
		sideType = "равносторонний"
	} else if a == b || a == c || b == c {
		sideType = "равнобедренный"
	} else {
		sideType = "разносторонний"
	}

	fmt.Printf("Ex81a. Ответ: %s\n", angleType)
	fmt.Printf("Ex81b. Ответ: %s, %s\n", angleType, sideType)
}

func ch4ex82() {
	var n int
	fmt.Printf("Ex82. Склонение лет\n")
	fmt.Printf("Ex82. Введите возраст: ")
	fmt.Scan(&n)

	word := "лет"
	if n%10 == 1 && n%100 != 11 {
		word = "год"
	} else if n%10 >= 2 && n%10 <= 4 && (n%100 < 10 || n%100 >= 20) {
		word = "года"
	}

	fmt.Printf("Ex82. Ответ: мне %d %s\n", n, word)
}

func ch4ex83() {
	var k int
	fmt.Printf("Ex83. Склонение грибов\n")
	fmt.Printf("Ex83. Введите количество грибов: ")
	fmt.Scan(&k)

	word := "грибов"
	if k%10 == 1 && k%100 != 11 {
		word = "гриб"
	} else if k%10 >= 2 && k%10 <= 4 && (k%100 < 10 || k%100 >= 20) {
		word = "гриба"
	}

	fmt.Printf("Ex83. Ответ: мы нашли %d %s в лесу\n", k, word)
}

func ch4ex84() {
	var n int
	fmt.Printf("Ex84. Стоимость в рублях и копейках\n")
	fmt.Printf("Ex84. Введите стоимость в копейках: ")
	fmt.Scan(&n)

	rubles := n / 100
	kopeyki := n % 100

	rubleWord := map[int]string{1: "рубль", 2: "рубля", 5: "рублей"}[func() int {
		if rubles%10 == 1 && rubles%100 != 11 {
			return 1
		}
		if rubles%10 >= 2 && rubles%10 <= 4 && (rubles%100 < 10 || rubles%100 >= 20) {
			return 2
		}
		return 5
	}()]

	kopeykiWord := map[int]string{1: "копейка", 2: "копейки", 5: "копеек"}[func() int {
		if kopeyki%10 == 1 && kopeyki%100 != 11 {
			return 1
		}
		if kopeyki%10 >= 2 && kopeyki%10 <= 4 && (kopeyki%100 < 10 || kopeyki%100 >= 20) {
			return 2
		}
		return 5
	}()]

	if kopeyki == 0 {
		fmt.Printf("Ex84. Ответ: %d %s ровно\n", rubles, rubleWord)
	} else {
		fmt.Printf("Ex84. Ответ: %d %s %d %s\n", rubles, rubleWord, kopeyki, kopeykiWord)
	}
}

func ch4ex85() {
	var n int
	fmt.Printf("Ex85. Возраст в годах и месяцах\n")
	fmt.Printf("Ex85. Введите возраст в месяцах: ")
	fmt.Scan(&n)

	years := n / 12
	months := n % 12

	yearWord := map[int]string{1: "год", 2: "года", 5: "лет"}[func() int {
		if years%10 == 1 && years%100 != 11 {
			return 1
		}
		if years%10 >= 2 && years%10 <= 4 && (years%100 < 10 || years%100 >= 20) {
			return 2
		}
		return 5
	}()]

	monthWord := map[int]string{1: "месяц", 2: "месяца", 5: "месяцев"}[func() int {
		if months%10 == 1 && months%100 != 11 {
			return 1
		}
		if months%10 >= 2 && months%10 <= 4 && (months%100 < 10 || months%100 >= 20) {
			return 2
		}
		return 5
	}()]

	if months == 0 {
		fmt.Printf("Ex85. Ответ: %d %s ровно\n", years, yearWord)
	} else {
		fmt.Printf("Ex85. Ответ: %d %s %d %s\n", years, yearWord, months, monthWord)
	}
}

func ch4ex86() {
	var birthYear, birthMonth, birthDay int
	var currentYear, currentMonth, currentDay int

	fmt.Printf("Ex86. Возраст в полных годах\n")
	fmt.Printf("Ex86. Введите дату рождения (год месяц день): ")
	fmt.Scan(&birthYear, &birthMonth, &birthDay)
	fmt.Printf("Ex86. Введите текущую дату (год месяц день): ")
	fmt.Scan(&currentYear, &currentMonth, &currentDay)

	age := currentYear - birthYear
	if currentMonth < birthMonth || (currentMonth == birthMonth && currentDay < birthDay) {
		age--
	}

	fmt.Printf("Ex86. Ответ: %d %s\n", age,
		map[int]string{1: "год", 2: "года", 5: "лет"}[func() int {
			if age%10 == 1 && age%100 != 11 {
				return 1
			}
			if age%10 >= 2 && age%10 <= 4 && (age%100 < 10 || age%100 >= 20) {
				return 2
			}
			return 5
		}()])
}

func ch4ex87() {
	var year1, month1, day1 int
	var year2, month2, day2 int

	fmt.Printf("Ex87. Сравнение возраста\n")
	fmt.Printf("Ex87. Введите дату рождения первого (год месяц день): ")
	fmt.Scan(&year1, &month1, &day1)
	fmt.Printf("Ex87. Введите дату рождения второго (год месяц день): ")
	fmt.Scan(&year2, &month2, &day2)

	older := map[bool]string{true: "первый", false: "второй"}[year1 < year2 || (year1 == year2 && month1 < month2) ||
		(year1 == year2 && month1 == month2 && day1 < day2)]

	fmt.Printf("Ex87. Ответ: Старше %s человек\n", older)
}

func ch4ex88() {
	var birthYear, birthMonth int
	var currentYear, currentMonth int

	fmt.Printf("Ex88. Возраст в годах и месяцах\n")
	fmt.Printf("Ex88. Введите год и месяц рождения: ")
	fmt.Scan(&birthYear, &birthMonth)
	fmt.Printf("Ex88. Введите текущий год и месяц: ")
	fmt.Scan(&currentYear, &currentMonth)

	years := currentYear - birthYear
	months := currentMonth - birthMonth

	if months < 0 {
		years--
		months += 12
	}

	yearWord := map[int]string{1: "год", 2: "года", 5: "лет"}[func() int {
		if years%10 == 1 && years%100 != 11 {
			return 1
		}
		if years%10 >= 2 && years%10 <= 4 && (years%100 < 10 || years%100 >= 20) {
			return 2
		}
		return 5
	}()]

	monthWord := map[int]string{1: "месяц", 2: "месяца", 5: "месяцев"}[func() int {
		if months%10 == 1 && months%100 != 11 {
			return 1
		}
		if months%10 >= 2 && months%10 <= 4 && (months%100 < 10 || months%100 >= 20) {
			return 2
		}
		return 5
	}()]

	fmt.Printf("Ex88. Ответ: %d %s %d %s\n", years, yearWord, months, monthWord)
}

func ch4ex89() {
	var a, b, c, d, n, m int
	fmt.Printf("Ex89. Поезд на платформе\n")
	fmt.Printf("Ex89. Введите прибытие (a b), отправление (c d), время пассажира (n m): ")
	fmt.Scan(&a, &b, &c, &d, &n, &m)

	arrival := a*60 + b
	departure := c*60 + d
	passenger := n*60 + m

	onPlatform := passenger >= arrival && passenger < departure
	fmt.Printf("Ex89. Ответ: Поезд %s на платформе\n",
		map[bool]string{true: "стоит", false: "не стоит"}[onPlatform])
}

func ch4ex90() {
	var n, m int
	fmt.Printf("Ex90. Соседние даты\n")
	fmt.Printf("Ex90. Введите число и месяц: ")
	fmt.Scan(&n, &m)

	daysInMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	prevN, prevM := n-1, m
	if prevN == 0 {
		prevM = m - 1
		prevN = daysInMonth[prevM-1]
	}

	nextN, nextM := n+1, m
	if nextN > daysInMonth[m-1] {
		nextM = m + 1
		nextN = 1
	}

	fmt.Printf("Ex90a. Ответ: %d.%d\n", prevN, prevM)
	fmt.Printf("Ex90b. Ответ: %d.%d\n", nextN, nextM)
}

func ch4ex91() {
	var g, n, m int
	fmt.Printf("Ex91. Соседние даты с годом\n")
	fmt.Printf("Ex91. Введите год, число и месяц: ")
	fmt.Scan(&g, &n, &m)

	isLeap := (g%4 == 0 && g%100 != 0) || (g%400 == 0)
	daysInMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if isLeap {
		daysInMonth[1] = 29
	}

	prevG, prevN, prevM := g, n-1, m
	if prevN == 0 {
		prevM = m - 1
		if prevM == 0 {
			prevM = 12
			prevG = g - 1
		}
		prevN = daysInMonth[prevM-1]
	}

	nextG, nextN, nextM := g, n+1, m
	if nextN > daysInMonth[m-1] {
		nextM = m + 1
		nextN = 1
		if nextM > 12 {
			nextM = 1
			nextG = g + 1
		}
	}

	fmt.Printf("Ex91a. Ответ: %d.%d.%d\n", prevN, prevM, prevG)
	fmt.Printf("Ex91b. Ответ: %d.%d.%d\n", nextN, nextM, nextG)
}

func ch4ex92() {
	var t float64
	fmt.Printf("Ex92. Светофор для водителей\n")
	fmt.Printf("Ex92. Введите время в минутах с начала часа: ")
	fmt.Scan(&t)

	timeInCycle := math.Mod(t, 6)

	var color string
	switch {
	case timeInCycle < 3:
		color = "зеленый"
	case timeInCycle < 4:
		color = "желтый"
	default:
		color = "красный"
	}

	fmt.Printf("Ex92. Ответ: Горит %s сигнал\n", color)
}

func ch4ex93() {
	var k, d int
	fmt.Printf("Ex93. День недели года\n")
	fmt.Printf("Ex93. Введите номер дня года и день недели 1 января: ")
	fmt.Scan(&k, &d)

	dayOfWeek := (k + d - 2) % 7

	var dayType string
	switch dayOfWeek {
	case 0:
		dayType = "воскресенье"
	case 6:
		dayType = "суббота"
	default:
		dayType = "рабочий день"
	}

	fmt.Printf("Ex93. Ответ: %s\n", dayType)
}

func ch4ex94() {
	var k int
	fmt.Printf("Ex94. Введите k (1-180): ")
	fmt.Scan(&k)

	number := 10 + (k-1)/2
	if k%2 == 1 {
		fmt.Printf("Ex94. Ответ: %d\n", number/10)
	} else {
		fmt.Printf("Ex94. Ответ: %d\n", number%10)
	}
}

func ch4ex95() {
	var n int
	fmt.Printf("Ex95. Введите n (1-32): ")
	fmt.Scan(&n)

	if n <= 9 {
		fmt.Printf("Ex95. Ответ: %d\n", n)
	} else if n <= 28 {
		number := 10 + (n-10)/2
		if n%2 == 0 {
			fmt.Printf("Ex95. Ответ: %d\n", number/10)
		} else {
			fmt.Printf("Ex95. Ответ: %d\n", number%10)
		}
	} else {
		fmt.Printf("Ex95. Ответ: %d\n", []int{2, 0}[n-29])
	}
}

func ch4ex96() {
	var k int
	fmt.Printf("Ex96. k-я цифра в последовательности 50-250\n")
	fmt.Printf("Ex96. Введите k (1-252): ")
	fmt.Scan(&k)

	if k <= 102 {
		number := 50 + (k-1)/2
		if k%2 == 1 {
			fmt.Printf("Ex96. Ответ: %d\n", number/10)
		} else {
			fmt.Printf("Ex96. Ответ: %d\n", number%10)
		}
	} else {
		number := 100 + (k-103)/3
		pos := (k - 103) % 3
		fmt.Printf("Ex96. Ответ: %d\n", []int{number / 100, number / 10 % 10, number % 10}[pos])
	}
}

func ch4ex97() {
	var k int
	fmt.Printf("Ex97. k-я цифра в последовательности 1-110\n")
	fmt.Printf("Ex97. Введите k (1-222): ")
	fmt.Scan(&k)

	if k <= 9 {
		fmt.Printf("Ex97. Ответ: %d\n", k)
	} else if k <= 189 {
		number := 10 + (k-10)/2
		if k%2 == 0 {
			fmt.Printf("Ex97. Ответ: %d\n", number/10)
		} else {
			fmt.Printf("Ex97. Ответ: %d\n", number%10)
		}
	} else {
		number := 100 + (k-190)/3
		pos := (k - 190) % 3
		fmt.Printf("Ex97. Ответ: %d\n", []int{number / 100, number / 10 % 10, number % 10}[pos])
	}
}

func ch4ex98() {
	var n, a int
	fmt.Printf("Ex98. Четность суммы номеров квартир\n")
	fmt.Printf("Ex98. Введите количество квартир и номер первой квартиры: ")
	fmt.Scan(&n, &a)

	isEven := (n*(2*a+n-1))%2 == 0
	fmt.Printf("Ex98. Ответ: Сумма номеров %s числом\n",
		map[bool]string{true: "четное", false: "нечетное"}[isEven])
}

func ch4ex99() {
	var a, b float64
	fmt.Printf("Ex99. Наибольшее из двух чисел\n")
	fmt.Printf("Ex99. Введите два числа: ")
	fmt.Scan(&a, &b)

	max := a
	if b > max {
		max = b
	}
	fmt.Printf("Ex99a. Ответ: Наибольшее: %.2f\n", max)

	if a > b {
		fmt.Printf("Ex99b. Ответ: Наибольшее: %.2f\n", a)
	} else {
		fmt.Printf("Ex99b. Ответ: Наибольшее: %.2f\n", b)
	}
}

func ch4ex100() {
	var a, b float64
	fmt.Printf("Ex100. Наибольшее и наименьшее из двух чисел\n")
	fmt.Printf("Ex100. Введите два числа: ")
	fmt.Scan(&a, &b)

	max, min := a, a
	if b > max {
		max = b
	}
	if b < min {
		min = b
	}
	fmt.Printf("Ex100a. Ответ: Наибольшее: %.2f, Наименьшее: %.2f\n", max, min)

	if a > b {
		fmt.Printf("Ex100b. Ответ: Наибольшее: %.2f, Наименьшее: %.2f\n", a, b)
	} else {
		fmt.Printf("Ex100b. Ответ: Наибольшее: %.2f, Наименьшее: %.2f\n", b, a)
	}
}

func ch4ex101() {
	var a, b, c float64
	fmt.Printf("Ex101. Наибольшее и наименьшее из трех чисел\n")
	fmt.Printf("Ex101. Введите три числа: ")
	fmt.Scan(&a, &b, &c)

	maximum := math.Min(math.Min(a, b), c)
	fmt.Printf("Ex101a. Ответ: Наибольшее: %.2f\n", maximum)

	minimum := math.Min(math.Min(a, b), c)
	fmt.Printf("Ex101b. Ответ: Наименьшее: %.2f\n", minimum)
}

func ch4ex102() {
	var a, b, c, d float64
	fmt.Printf("Ex102. Наибольшее и наименьшее из четырех чисел\n")
	fmt.Printf("Ex102. Введите четыре числа: ")
	fmt.Scan(&a, &b, &c, &d)

	maximum := math.Min(math.Min(a, b), math.Min(c, d))
	fmt.Printf("Ex102a. Ответ: Наибольшее: %.2f\n", maximum)

	minimum := math.Min(math.Min(a, b), math.Min(c, d))
	fmt.Printf("Ex102b. Ответ: Наименьшее: %.2f\n", minimum)
}

func ch4ex103() {
	var x float64
	fmt.Printf("Ex103. Абсолютная величина\n")
	fmt.Printf("Ex103. Введите число: ")
	fmt.Scan(&x)

	fmt.Printf("Ex103. Ответ: |%.2f| = %.2f\n", x, math.Abs(x))
}

func ch4ex104() {
	var a, b float64
	fmt.Printf("Ex104. Введите два числа: ")
	fmt.Scan(&a, &b)

	absA, absB := math.Abs(a), math.Abs(b)

	fmt.Printf("Ex104a. Полусумма: %.2f\n", (absA+absB)/2)
	fmt.Printf("Ex104b. Корень из произведения: %.2f\n", math.Sqrt(absA*absB))
}

func ch4ex105() {
	var a, b float64
	fmt.Printf("Ex105. Введите два числа: ")
	fmt.Scan(&a, &b)

	if math.Abs(a) > math.Abs(b) {
		a /= 2
	}

	fmt.Printf("Ex105. Ответ: %.2f\n", a)
}

func ch4ex106() {
	var a, b float64
	fmt.Printf("Ex106. Введите два числа: ")
	fmt.Scan(&a, &b)

	if math.Sqrt(b) < a {
		b *= 5
	}

	fmt.Printf("Ex106. Ответ: %.2f\n", b)
}

func ch4ex107() {
	var a, b, c int
	fmt.Printf("Ex107. Введите три целых числа: ")
	fmt.Scan(&a, &b, &c)

	fmt.Printf("Ex107. Четные числа: ")
	if a%2 == 0 {
		fmt.Printf("%d ", a)
	}
	if b%2 == 0 {
		fmt.Printf("%d ", b)
	}
	if c%2 == 0 {
		fmt.Printf("%d ", c)
	}
}

func ch4ex108() {
	var a, b, c float64
	fmt.Printf("Ex108. Введите три числа: ")
	fmt.Scan(&a, &b, &c)

	numbers := []float64{a, b, c}
	for i, x := range numbers {
		if x >= 0 {
			numbers[i] = x * x
		}
	}
	fmt.Printf("Ex108. Ответ: %.2f %.2f %.2f\n", numbers[0], numbers[1], numbers[2])
}

func ch4ex109() {
	var a, b, c float64
	fmt.Printf("Ex109. Введите три числа: ")
	fmt.Scan(&a, &b, &c)

	result := []float64{}
	for _, x := range []float64{a, b, c} {
		if x >= 1.6 && x <= 5.8 {
			result = append(result, x)
		}
	}

	fmt.Printf("Ex109. Числа в интервале [1.6, 5.8]: %v\n", result)
}

func ch4ex110() {
	var a, b, c, d float64
	fmt.Printf("Ex110. Введите четыре числа: ")
	fmt.Scan(&a, &b, &c, &d)

	count := 0
	for _, x := range []float64{a, b, c, d} {
		if x < 0 {
			count++
		}
	}

	fmt.Printf("Ex110. Отрицательных чисел: %d\n", count)
}

func ch4ex111() {
	var a, b, c, d int
	fmt.Printf("Ex111. Введите четыре целых числа: ")
	fmt.Scan(&a, &b, &c, &d)

	count := 0
	for _, x := range []int{a, b, c, d} {
		if x%2 == 0 {
			count++
		}
	}

	fmt.Printf("Ex111. Четных чисел: %d\n", count)
}

func ch4ex112() {
	var a, b int
	fmt.Printf("Ex112. Введите два целых числа: ")
	fmt.Scan(&a, &b)

	sum := 0
	for _, x := range []int{a, b} {
		if x > 0 {
			sum += x
		}
	}

	fmt.Printf("Ex112. Сумма положительных: %d\n", sum)
}

func ch4ex113() {
	var a, b, c, d float64
	fmt.Printf("Ex113. Введите четыре числа: ")
	fmt.Scan(&a, &b, &c, &d)

	sum := 0.0
	for _, x := range []float64{a, b, c, d} {
		if x > 5 {
			sum += x
		}
	}

	fmt.Printf("Ex113. Сумма чисел больше пяти: %.2f\n", sum)
}

func ch4ex114() {
	var a, b, c, d int
	fmt.Printf("Ex114. Введите четыре целых числа: ")
	fmt.Scan(&a, &b, &c, &d)

	sum := 0
	for _, x := range []int{a, b, c, d} {
		if x%3 == 0 {
			sum += x
		}
	}

	fmt.Printf("Ex114. Сумма кратных трем: %d\n", sum)
}

func ch4ex115() {
	var a, b, c, d, e, f int
	fmt.Printf("Ex115. Введите шесть целых чисел: ")
	fmt.Scan(&a, &b, &c, &d, &e, &f)

	sum := 0
	for _, x := range []int{a, b, c, d, e, f} {
		if x > 0 {
			sum += x
		}
	}

	fmt.Printf("Ex115. Сумма положительных: %d\n", sum)
}

func ch4ex116() {
	var x float64
	fmt.Printf("Ex116. Вычисление функции y(x)\n")
	fmt.Printf("Ex116. Введите x: ")
	fmt.Scan(&x)

	y := 1.0
	if x < -1 {
		y = -1
	} else if x > -1 {
		y = x
	}

	fmt.Printf("Ex116. Ответ: y(%.2f) = %.2f\n", x, y)
}

func ch4ex117() {
	var a float64
	fmt.Printf("Ex117. Знак числа\n")
	fmt.Printf("Ex117. Введите число a: ")
	fmt.Scan(&a)

	var z int
	switch {
	case a > 0:
		z = 1
	case a < 0:
		z = -1
	default:
		z = 0
	}

	fmt.Printf("Ex117. Ответ: z = %d\n", z)
}

func ch4ex118() {
	var x float64
	fmt.Printf("Ex118. Вычисление функции f(x)\n")
	fmt.Printf("Ex118. Введите x: ")
	fmt.Scan(&x)

	var f float64
	switch {
	case x <= 0:
		f = 0
	case x <= 1:
		f = x
	default:
		f = x * x
	}

	fmt.Printf("Ex118. Ответ: f(%.2f) = %.2f\n", x, f)
}

func ch4ex119() {
	var y float64
	fmt.Printf("Ex119. Вычисление функции f(y)\n")
	fmt.Printf("Ex119. Введите y: ")
	fmt.Scan(&y)

	var f float64
	switch {
	case y > 2:
		f = 2
	case y > 0:
		f = 0
	default:
		f = -3 * y
	}

	fmt.Printf("Ex119. Ответ: f(%.2f) = %.2f\n", y, f)
}

func ch4ex120() {
	var x float64
	fmt.Printf("Ex120. Значение функции по графику\n")
	fmt.Printf("Ex120. Введите x: ")
	fmt.Scan(&x)

	var ya float64
	switch {
	case x < -2:
		ya = -1
	case x < 2:
		ya = 0
	default:
		ya = 1
	}

	yb := math.Abs(x)

	var yc float64
	if x >= -1 && x <= 1 {
		yc = x * x
	} else {
		yc = 1
	}

	var yd float64
	switch {
	case x < -1:
		yd = -1
	case x < 1:
		yd = x
	default:
		yd = 1
	}

	fmt.Printf("Ex120a. Ответ: y = %.2f\n", ya)
	fmt.Printf("Ex120b. Ответ: y = %.2f\n", yb)
	fmt.Printf("Ex120c. Ответ: y = %.2f\n", yc)
	fmt.Printf("Ex120d. Ответ: y = %.2f\n", yd)
}

func ch4ex121() {
	var x, y float64
	fmt.Printf("Ex121. Область точки (рис 4.9)\n")
	fmt.Printf("Ex121. Введите координаты x и y: ")
	fmt.Scan(&x, &y)

	var region string
	switch {
	case x < 1:
		region = "I"
	case x < 5:
		region = "II"
	default:
		region = "III"
	}

	fmt.Printf("Ex121. Ответ: Точка в области %s\n", region)
}

func ch4ex122() {
	var x, y float64
	fmt.Printf("Ex122. Область точки (рис 4.10)\n")
	fmt.Printf("Ex122. Введите координаты x и y: ")
	fmt.Scan(&x, &y)

	var region string
	switch {
	case y > 5.0:
		region = "I"
	case y > 2.5:
		region = "II"
	default:
		region = "III"
	}

	fmt.Printf("Ex122. Ответ: Точка в области %s\n", region)
}

func ch4ex123() {
	var points int
	fmt.Printf("Ex123. Результат футбольной игры\n")
	fmt.Printf("Ex123. Введите количество очков (0, 1 или 3): ")
	fmt.Scan(&points)

	result := map[int]string{
		0: "проигрыш",
		1: "ничья",
		3: "выигрыш",
	}[points]

	fmt.Printf("Ex123. Ответ: %s\n", result)
}

func ch4ex124() {
	var mitya, vasya int
	fmt.Printf("Ex124. Сравнение возраста\n")
	fmt.Printf("Ex124. Введите возраст Мити и Васи: ")
	fmt.Scan(&mitya, &vasya)

	result := map[bool]string{
		true:  "Митя старше",
		false: "Вася старше",
	}[mitya > vasya]

	if mitya == vasya {
		result = "Одного возраста"
	}

	fmt.Printf("Ex124. Ответ: %s\n", result)
}

func ch4ex125() {
	var weight float64
	fmt.Printf("Ex125. Весовая категория боксера\n")
	fmt.Printf("Ex125. Введите вес боксера: ")
	fmt.Scan(&weight)

	category := map[bool]string{
		true: "легкий вес",
		false: map[bool]string{
			true:  "первый полусредний вес",
			false: "полусредний вес",
		}[weight < 64],
	}[weight < 60]

	fmt.Printf("Ex125. Ответ: %s\n", category)
}

func ch4ex126() {
	var a, b int
	fmt.Printf("Ex126. Взаимная делимость\n")
	fmt.Printf("Ex126. Введите два различных числа: ")
	fmt.Scan(&a, &b)

	result := map[bool]string{
		true: "a является делителем b",
		false: map[bool]string{
			true:  "b является делителем a",
			false: "ни одно не является делителем другого",
		}[a%b == 0],
	}[b%a == 0]

	fmt.Printf("Ex126. Ответ: %s\n", result)
}

func ch4ex127() {
	var a, b, c int
	fmt.Printf("Ex127. Анализ трех чисел\n")
	fmt.Printf("Ex127. Введите три различных числа: ")
	fmt.Scan(&a, &b, &c)

	maxPos := map[bool]string{
		true: "первое",
		false: map[bool]string{
			true:  "второе",
			false: "третье",
		}[b > c],
	}[a > b && a > c]

	minPos := map[bool]string{
		true: "первое",
		false: map[bool]string{
			true:  "второе",
			false: "третье",
		}[b < c],
	}[a < b && a < c]

	midPos := map[bool]string{
		true: "первое",
		false: map[bool]string{
			true:  "второе",
			false: "третье",
		}[(b > a && b < c) || (b < a && b > c)],
	}[(a > b && a < c) || (a < b && a > c)]

	fmt.Printf("Ex127a. Ответ: самое большое - %s\n", maxPos)
	fmt.Printf("Ex127b. Ответ: самое маленькое - %s\n", minPos)
	fmt.Printf("Ex127c. Ответ: среднее - %s\n", midPos)
}

func ch4ex128() {
	var a, b, c float64
	fmt.Printf("Ex128. Максимум и минимум из трех чисел\n")
	fmt.Printf("Ex128. Введите три различных числа: ")
	fmt.Scan(&a, &b, &c)

	maximum := map[bool]float64{
		true: a,
		false: map[bool]float64{
			true:  b,
			false: c,
		}[b > c],
	}[a > b && a > c]

	minimum := map[bool]float64{
		true: a,
		false: map[bool]float64{
			true:  b,
			false: c,
		}[b < c],
	}[a < b && a < c]

	fmt.Printf("Ex128. Ответ: максимум = %.2f, минимум = %.2f\n", maximum, minimum)
}

func ch4ex129() {
	var a, b, c float64
	fmt.Printf("Ex129. Сумма двух наибольших\n")
	fmt.Printf("Ex129. Введите три различных числа: ")
	fmt.Scan(&a, &b, &c)

	sum := map[bool]float64{
		true: a + b,
		false: map[bool]float64{
			true:  a + c,
			false: b + c,
		}[a > c],
	}[a > c && b > c]

	fmt.Printf("Ex129. Ответ: сумма двух наибольших = %.2f\n", sum)
}

func ch4ex130() {
	var a, b, c float64
	fmt.Printf("Ex130. Произведение двух наименьших\n")
	fmt.Printf("Ex130. Введите три различных числа: ")
	fmt.Scan(&a, &b, &c)

	product := map[bool]float64{
		true: a * b,
		false: map[bool]float64{
			true:  a * c,
			false: b * c,
		}[a < c],
	}[a < c && b < c]

	fmt.Printf("Ex130. Ответ: произведение двух наименьших = %.2f\n", product)
}

func ch4ex131() {
	var a1, b1, c1, a2, b2, c2 float64
	fmt.Printf("Ex131. Среднее арифметическое средних чисел\n")
	fmt.Printf("Ex131. Введите первую тройку чисел: ")
	fmt.Scan(&a1, &b1, &c1)
	fmt.Printf("Ex131. Введите вторую тройку чисел: ")
	fmt.Scan(&a2, &b2, &c2)

	mid1 := map[bool]float64{
		true: a1,
		false: map[bool]float64{
			true:  b1,
			false: c1,
		}[(b1 > a1 && b1 < c1) || (b1 < a1 && b1 > c1)],
	}[(a1 > b1 && a1 < c1) || (a1 < b1 && a1 > c1)]

	mid2 := map[bool]float64{
		true: a2,
		false: map[bool]float64{
			true:  b2,
			false: c2,
		}[(b2 > a2 && b2 < c2) || (b2 < a2 && b2 > c2)],
	}[(a2 > b2 && a2 < c2) || (a2 < b2 && a2 > c2)]

	average := (mid1 + mid2) / 2
	fmt.Printf("Ex131. Ответ: среднее арифметическое = %.2f\n", average)
}

func ch4ex132() {
	var x, y float64
	fmt.Printf("Ex132. Четверть координатной плоскости\n")
	fmt.Printf("Ex132. Введите координаты x и y: ")
	fmt.Scan(&x, &y)

	quarter := map[bool]string{
		true: map[bool]string{
			true:  "первая",
			false: "четвертая",
		}[y > 0],
		false: map[bool]string{
			true:  "вторая",
			false: "третья",
		}[y > 0],
	}[x > 0]

	fmt.Printf("Ex132. Ответ: %s четверть\n", quarter)
}

func ch4ex133() {
	var day int
	fmt.Printf("Ex133. Название дня недели\n")
	fmt.Printf("Ex133. Введите номер дня недели (1-7): ")
	fmt.Scan(&day)

	dayName := map[int]string{
		1: "понедельник",
		2: "вторник",
		3: "среда",
		4: "четверг",
		5: "пятница",
		6: "суббота",
		7: "воскресенье",
	}[day]

	fmt.Printf("Ex133. Ответ: %s\n", dayName)
}

func ch4ex134() {
	var month int
	fmt.Printf("Ex134. Название месяца\n")
	fmt.Printf("Ex134. Введите номер месяца (1-12): ")
	fmt.Scan(&month)

	monthName := map[int]string{
		1: "январь", 2: "февраль", 3: "март", 4: "апрель",
		5: "май", 6: "июнь", 7: "июль", 8: "август",
		9: "сентябрь", 10: "октябрь", 11: "ноябрь", 12: "декабрь",
	}[month]

	fmt.Printf("Ex134. Ответ: %s\n", monthName)
}

func ch4ex135() {
	var month int
	fmt.Printf("Ex135. Время года\n")
	fmt.Printf("Ex135. Введите номер месяца (1-12): ")
	fmt.Scan(&month)

	season := map[bool]string{
		true: "зима",
		false: map[bool]string{
			true: "весна",
			false: map[bool]string{
				true:  "лето",
				false: "осень",
			}[month >= 6 && month <= 8],
		}[month >= 3 && month <= 5],
	}[month == 12 || month <= 2]

	fmt.Printf("Ex135. Ответ: %s\n", season)
}

func ch4ex136() {
	var month int
	var isLeap bool
	fmt.Printf("Ex136. Количество дней в месяце\n")
	fmt.Printf("Ex136. Введите номер месяца (1-12): ")
	fmt.Scan(&month)
	fmt.Printf("Ex136. Год високосный? (true/false): ")
	fmt.Scan(&isLeap)

	days := map[int]int{
		1: 31, 2: 28, 3: 31, 4: 30, 5: 31, 6: 30,
		7: 31, 8: 31, 9: 30, 10: 31, 11: 30, 12: 31,
	}[month]

	if month == 2 && isLeap {
		days = 29
	}

	fmt.Printf("Ex136. Ответ: %d дней\n", days)
}

func ch4ex137() {
	var month int
	fmt.Printf("Ex137. Количество дней в месяце (не високосный)\n")
	fmt.Printf("Ex137. Введите номер месяца (1-12): ")
	fmt.Scan(&month)

	days := map[bool]int{
		true: 31,
		false: map[bool]int{
			true: 28,
			false: map[bool]int{
				true:  30,
				false: 31,
			}[month == 4 || month == 6 || month == 9 || month == 11],
		}[month == 2],
	}[month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12]

	fmt.Printf("Ex137. Ответ: %d дней\n", days)
}

func ch4ex138() {
	var m int
	fmt.Printf("Ex138. Название масти карт\n")
	fmt.Printf("Ex138. Введите номер масти (1-4): ")
	fmt.Scan(&m)

	suit := map[int]string{
		1: "пики",
		2: "трефы",
		3: "бубны",
		4: "червы",
	}[m]

	fmt.Printf("Ex138. Ответ: %s\n", suit)
}

func ch4ex139() {
	var k int
	fmt.Printf("Ex139. Достоинство карты\n")
	fmt.Printf("Ex139. Введите номер карты (6-14): ")
	fmt.Scan(&k)

	rank := map[int]string{
		6: "шестерка", 7: "семерка", 8: "восьмерка", 9: "девятка", 10: "десятка",
		11: "валет", 12: "дама", 13: "король", 14: "туз",
	}[k]

	fmt.Printf("Ex139. Ответ: %s\n", rank)
}

func ch4ex140() {
	var m, k int
	fmt.Printf("Ex140. Полное название карты\n")
	fmt.Printf("Ex140. Введите номер масти (1-4) и номер карты (6-14): ")
	fmt.Scan(&m, &k)

	suit := map[int]string{
		1: "пик", 2: "треф", 3: "бубен", 4: "червей",
	}[m]

	rank := map[int]string{
		6: "Шестерка", 7: "Семерка", 8: "Восьмерка", 9: "Девятка", 10: "Десятка",
		11: "Валет", 12: "Дама", 13: "Король", 14: "Туз",
	}[k]

	fmt.Printf("Ex140. Ответ: %s %s\n", rank, suit)
}

func ch4ex141() {
	var k, d int
	fmt.Printf("Ex141. День недели по номеру дня года\n")
	fmt.Printf("Ex141. Введите номер дня года и день недели 1 января: ")
	fmt.Scan(&k, &d)

	dayOfWeek := (k + d - 2) % 7

	dayName := map[int]string{
		0: "воскресенье", 1: "понедельник", 2: "вторник", 3: "среда",
		4: "четверг", 5: "пятница", 6: "суббота",
	}[dayOfWeek]

	fmt.Printf("Ex141. Ответ: %s\n", dayName)
}

func ch4ex142() {
	var n int
	fmt.Printf("Ex142. Название месяца с 2010 года\n")
	fmt.Printf("Ex142. Введите количество прошедших месяцев: ")
	fmt.Scan(&n)

	monthName := map[int]string{
		1: "январь", 2: "февраль", 3: "март", 4: "апрель",
		5: "май", 6: "июнь", 7: "июль", 8: "август",
		9: "сентябрь", 10: "октябрь", 11: "ноябрь", 12: "декабрь",
	}[n%12+1]

	fmt.Printf("Ex142. Ответ: %s\n", monthName)
}

func ch4ex143() {
	var n, m int
	fmt.Printf("Ex143. Соседние даты (без года)\n")
	fmt.Printf("Ex143. Введите число и месяц: ")
	fmt.Scan(&n, &m)

	daysInMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	prevN, prevM := n-1, m
	if prevN == 0 {
		prevM = m - 1
		prevN = daysInMonth[prevM-1]
	}

	nextN, nextM := n+1, m
	if nextN > daysInMonth[m-1] {
		nextM = m + 1
		nextN = 1
	}

	fmt.Printf("Ex143a. Ответ: %d.%d\n", prevN, prevM)
	fmt.Printf("Ex143b. Ответ: %d.%d\n", nextN, nextM)
}

func ch4ex144() {
	var g, n, m int
	fmt.Printf("Ex144. Соседние даты (с годом)\n")
	fmt.Printf("Ex144. Введите год, число и месяц: ")
	fmt.Scan(&g, &n, &m)

	isLeap := (g%4 == 0 && g%100 != 0) || (g%400 == 0)
	daysInMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if isLeap {
		daysInMonth[1] = 29
	}

	prevG, prevN, prevM := g, n-1, m
	if prevN == 0 {
		prevM = m - 1
		if prevM == 0 {
			prevM = 12
			prevG = g - 1
		}
		prevN = daysInMonth[prevM-1]
	}

	nextG, nextN, nextM := g, n+1, m
	if nextN > daysInMonth[m-1] {
		nextM = m + 1
		nextN = 1
		if nextM > 12 {
			nextM = 1
			nextG = g + 1
		}
	}

	fmt.Printf("Ex144a. Ответ: %d.%d.%d\n", prevN, prevM, prevG)
	fmt.Printf("Ex144b. Ответ: %d.%d.%d\n", nextN, nextM, nextG)
}

func ch4ex145() {
	var year int
	fmt.Printf("Ex145. Восточный календарь\n")
	fmt.Printf("Ex145. Введите год: ")
	fmt.Scan(&year)

	animals := []string{"Крыса", "Корова", "Тигр", "Заяц", "Дракон", "Змея",
		"Лошадь", "Овца", "Обезьяна", "Петух", "Собака", "Свинья"}

	colors := []string{"Зеленый", "Красный", "Желтый", "Белый", "Черный"}

	cycleYear := (year - 1984) % 60
	if cycleYear < 0 {
		cycleYear += 60
	}

	animalIndex := cycleYear % 12
	colorIndex := (cycleYear / 2) % 5

	animal := animals[animalIndex]
	color := colors[colorIndex]

	fmt.Printf("Ex145. Ответ: %s, %s\n", animal, color)
}
