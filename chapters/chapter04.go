package chapters

import (
	"fmt"
	"math"
)

func Run04() {
	tasks := []func(){
		ch4ex01, ch4ex02, ch4ex03, ch4ex04, ch4ex05, ch4ex06, ch4ex07, ch4ex08,
		ch4ex09, ch4ex10, ch4ex11, ch4ex12, ch4ex13, ch4ex14, ch4ex15, ch4ex16,
		ch4ex17, ch4ex18, ch4ex19, ch4ex20, ch4ex21, ch4ex22, ch4ex23,
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
