package chapters

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func Run02() {
	tasks := []func(){
		ch2ex01, ch2ex02, ch2ex03, ch2ex04, ch2ex05, ch2ex06, ch2ex07,
		ch2ex08, ch2ex09, ch2ex10, ch2ex11, ch2ex12, ch2ex13,
		ch2ex14, ch2ex15, ch2ex16, ch2ex17, ch2ex18, ch2ex19,
		ch2ex20, ch2ex21, ch2ex22, ch2ex23, ch2ex24, ch2ex25,
		ch2ex26, ch2ex27, ch2ex28, ch2ex29, ch2ex30, ch2ex31,
		ch2ex32, ch2ex33, ch2ex34, ch2ex35, ch2ex36, ch2ex37, ch2ex38,
	}

	for {
		fmt.Print("\nГлава 2. Выберите задачу (1-38): ")
		var n int
		fmt.Scan(&n)

		if n == 0 {
			break
		}
		if n >= 1 && n <= 38 {
			tasks[n-1]()
		} else {
			fmt.Println("\nОт 1 до 38! Для выхода введите 0.")
		}
	}
}

func ch2ex01() {
	var x int
	fmt.Printf("Ex01a. y = 17x² - 6x + 13\n")
	fmt.Printf("Ex01a. Введите x: ")
	fmt.Scan(&x)
	fmt.Printf("Ex01a. y = %d\n", 17*x*x-6*x+13)
	var a int
	fmt.Printf("Ex01b. y = 3a² + 5a - 21\n")
	fmt.Printf("Ex01b. Введите a: ")
	fmt.Scan(&a)
	fmt.Printf("Ex01b. y = %d\n", 3*a*a+5*a-21)
}

func ch2ex02() {
	var a float64
	fmt.Printf("Ex02. y = (a²+10)/sqrt(a²+1)\n")
	fmt.Printf("Ex02. Введите a: ")
	fmt.Scan(&a)
	fmt.Printf("Ex02. y = %.3f\n", (a*a+10)/math.Sqrt(a*a+1))
}

func ch2ex03() {
	var a float64
	fmt.Printf("Ex03a. y = sqrt((2a+sin|3a|)/3,56)\n")
	fmt.Printf("Ex03a. Введите a: ")
	fmt.Scan(&a)
	fmt.Printf("Ex03a. y = %.3f\n", math.Sqrt((2*a+math.Sin(math.Abs(3*a)))/3.56))

	var x float64
	fmt.Printf("Ex03b. y = sin((3.2+sqrt(1+x))/|5x|)\n")
	fmt.Printf("Ex03b. Введите x: ")
	fmt.Scan(&x)
	fmt.Printf("Ex03b. y = %.3f\n", math.Sin((3.2+math.Sqrt(1+x))/math.Abs(5*x)))
}

func ch2ex04() {
	var x float64
	fmt.Printf("Ex04. Нахождение периметра квадрата:\n")
	fmt.Printf("Ex04. Введите сторону: ")
	fmt.Scan(&x)
	fmt.Printf("Ex04. P = %.3f\n", x*x)
}

func ch2ex05() {
	var x float64
	fmt.Printf("Ex05. Нахождение диаметра окружности:\n")
	fmt.Printf("Ex05. Введите радиус: ")
	fmt.Scan(&x)
	fmt.Printf("Ex05. D = %.3f\n", x+x)
}

func ch2ex06() {
	var h float64
	R := 6350000.0
	fmt.Printf("Ex06. Расстояние до линии горизонта от точки над землей:\n")
	fmt.Printf("Ex06. Введите высоту точки наблюдения (в метрах): ")
	fmt.Scan(&h)
	fmt.Printf("Ex06. X = %.3f км\n", math.Sqrt(2*R*h+h*h)/1000)
}

func ch2ex07() {
	var x float64
	fmt.Printf("Ex07. Нахождение объема и площади боковой поверхности куба:\n")
	fmt.Printf("Ex07. Введите длину ребра: ")
	fmt.Scan(&x)
	fmt.Printf("Ex07. Объём куба: %.3f\n", x*x*x)
	fmt.Printf("Ex07. Площадь боковой поверхности = %.3f\n", x*x*4)
}

func ch2ex08() {
	var x float64
	fmt.Printf("Ex08. Нахождение длины и площади круга:\n")
	fmt.Printf("Ex08. Введите радиус: ")
	fmt.Scan(&x)
	fmt.Printf("Ex08. Длина окружности: %.3f\n", 2*math.Pi*x)
	fmt.Printf("Ex08. Площадь круга: %.3f\n", math.Pi*x*x)
}

func ch2ex09() {
	var x, y float64
	fmt.Printf("Ex09a. Значения функции z = 2x³ - 3.44xy + 2.3x² - 7.1y + 2\n")
	fmt.Printf("Ex09a. Введите x и y: ")
	fmt.Scan(&x, &y)
	fmt.Printf("Ex09a. z = %.3f\n", 2*x*x*x-3.44*x*y+2.3*x*x-7.1*y+2)

	fmt.Printf("Ex09a. Значения функции x = 3.14(a + b)³ + 2.75b² - 12.7a - 4.1\n")
	fmt.Printf("Ex09a. Введите a и b: ")
	fmt.Scan(&x, &y)
	fmt.Printf("Ex09a. x = %.3f\n", 3.14*(x+y)*(x+y)*(x+y)+2.75*y*y-12.7*x-4.1)
}

func ch2ex10() {
	var a, b float64
	fmt.Printf("Ex10. Нахождение среднего арифметическое и геометрического числа:\n")
	fmt.Printf("Ex10. Введите два целых числа: ")
	fmt.Scan(&a, &b)
	fmt.Printf("Ex10. Среднее арифметическое: %.3f\n", (a+b)/2)
	fmt.Printf("Ex10. Среднее геометрическое: %.3f\n", math.Sqrt(a*b))
}

func ch2ex11() {
	var a, b float64
	fmt.Printf("Ex11. Определение плотности материала:\n")
	fmt.Printf("Ex11. Введите объём и массу тела: ")
	fmt.Scan(&a, &b)
	fmt.Printf("Ex11. Плотность материала: %.3f\n", a/b)
}

func ch2ex12() {
	var a, b float64
	fmt.Printf("Ex12. Определение плотности населения:\n")
	fmt.Printf("Ex12. Введите кол-во жителей и площадь территории: ")
	fmt.Scan(&a, &b)
	fmt.Printf("Ex12. Плотность: %.3f\n", a/b+0.5)
}

func ch2ex13() {
	var a, b float64
	fmt.Printf("Ex13. Решение ax+b=0 (a!=0):\n")
	fmt.Printf("Ex13. Введите a и b: ")
	fmt.Scan(&a, &b)
	fmt.Printf("Ex13. x = %.3f\n", -b/a)
}

func ch2ex14() {
	var a, b float64
	fmt.Printf("Ex14. Нахождение гипотенузы треугольника:\n")
	fmt.Printf("Ex14. Введите длину катетов a и b: ")
	fmt.Scan(&a, &b)
	fmt.Printf("Ex14. Гипотенуза: %.3f\n", math.Sqrt(a*a+b*b))
}

func ch2ex15() {
	var a, b float64
	fmt.Printf("Ex15. Нахождение площади кольца:\n")
	fmt.Printf("Ex15. Введите внешний и внутренний радиус: ")
	fmt.Scan(&a, &b)
	fmt.Printf("Ex15. Площадь кольца: %.3f\n", math.Pi*(a*a-b*b))
}

func ch2ex16() {
	var a, b float64
	fmt.Printf("Ex16. Нахождение периметра прямоугольного треугольника:\n")
	fmt.Printf("Ex16. Введите катеты a и b: ")
	fmt.Scan(&a, &b)
	fmt.Printf("Ex16. Периметр: %.3f\n", math.Sqrt(a*a+b*b)+a+b)
}

func ch2ex17() {
	var a, b, c float64
	fmt.Printf("Ex17. Нахождение периметра равнобедренной трапеции:\n")
	fmt.Printf("Ex17. Введите основание1, основание2 и высоту: ")
	fmt.Scan(&a, &b, &c)
	fmt.Printf("Ex17. Периметр: %.3f\n", a+b+2*math.Sqrt(c*c+(a-b)*(a-b)/4))
}

func ch2ex18() {
	var x, y float64
	fmt.Printf("Ex18. Вычисление значения больших функций:\n")
	fmt.Printf("Ex18. Введите x, y: ")
	fmt.Scan(&x, &y)
	fmt.Printf("Ex18. z = %.3f\n", (x+(2+y)/(x*x))/(y+1/(math.Sqrt(x*x+10))))
	fmt.Printf("Ex18. q = %.3f\n", 7.25*math.Sin(x*math.Pi/180)-math.Abs(y))
}

func ch2ex19() {
	var a, b float64
	fmt.Printf("Ex19. Вычисление значения больших функций:\n")
	fmt.Printf("Ex19. Введите a, b: ")
	fmt.Scan(&a, &b)
	fmt.Printf("Ex19. x = %.3f\n", (2/(a*a+25)+b)/(math.Sqrt(b)+(a+b)/2))
	fmt.Printf("Ex19. y = %.3f\n", (math.Abs(a)+2*math.Sin(b*math.Pi/180))/(5.5*a))
}

func ch2ex20() {
	var e, f, g, h float64
	fmt.Printf("Ex20. Вычисление значения больших функций:\n")
	fmt.Printf("Ex20. Введите e, f, g и h: ")
	fmt.Scan(&e, &f, &g, &h)
	fmt.Printf("Ex20. a = %.3f\n", math.Sqrt(math.Pow(math.Abs(e-3/f), 3)+g))
	fmt.Printf("Ex20. b = %.3f\n", math.Sin(e*math.Pi/180)+math.Pow(math.Cos(h*math.Pi/180), 2))
	fmt.Printf("Ex20. c = %.3f\n", (33*g)/(e*f-3))
}

func ch2ex21() {
	var e, f, g, h float64
	fmt.Printf("Ex21. Вычисление значения больших функций:\n")
	fmt.Printf("Ex21. Введите e, f, g и h: ")
	fmt.Scan(&e, &f, &g, &h)
	fmt.Printf("Ex21. a = %.3f\n", (e+f/2)/3)
	fmt.Printf("Ex21. b = %.3f\n", math.Abs(h*h-g))
	fmt.Printf("Ex21. c = %.3f\n", math.Sqrt(math.Pow(g-h, 2)-3*math.Sin(e*math.Pi/180)))
}

func ch2ex22() {
	var a, b float64
	fmt.Printf("Ex22. Вычисление среднего арифметического и геометрического модулей чисел:\n")
	fmt.Printf("Ex22. Введите a и b: ")
	fmt.Scan(&a, &b)
	fmt.Printf("Ex22. Среднеe арифметическое модулей: %.3f\n", (math.Abs(a)+math.Abs(b))/2)
	fmt.Printf("Ex22. Среднеe геометрическое модулей: %.3f\n", math.Pow(math.Abs(a*b), 1.0/2))
}

func ch2ex23() {
	var a, b float64
	fmt.Printf("Ex23. Вычисление переметра и длину диагонали прямоугольника:\n")
	fmt.Printf("Ex23. Введите a и b: ")
	fmt.Scan(&a, &b)
	fmt.Printf("Ex23. Периметр: %.3f\n", (a+b)*2)
	fmt.Printf("Ex23. Длина диагонали: %.3f\n", math.Sqrt(a*a+b*b))
}

func ch2ex24() {
	var a, b float64
	fmt.Printf("Ex24. Вычисление суммы, разности, произведения и частного:\n")
	fmt.Printf("Ex24. Введите a и b: ")
	fmt.Scan(&a, &b)
	fmt.Printf("Ex24. Сумма: %.3f\n", a+b)
	fmt.Printf("Ex24. Разность: %.3f\n", a-b)
	fmt.Printf("Ex24. Произведение: %.3f\n", a*b)
	fmt.Printf("Ex24. Частное: %.3f\n", a/b)
}

func ch2ex25() {
	var a, b, c float64
	fmt.Printf("Ex25. Вычисление объема и площади боковой поверхности прямоугольного параллелепипеда:\n")
	fmt.Printf("Ex25. Введите длины сторон a, b, c: ")
	fmt.Scan(&a, &b, &c)
	fmt.Printf("Ex25. Объем: %.3f\n", a*b*c)
	fmt.Printf("Ex25. Площадь боковой поверхности: %.3f\n", (a+b)*c*2)
}

func ch2ex26() {
	var x1, y1, x2, y2 float64
	fmt.Printf("Ex26. Вычисление расстояния между точками на плоскости:\n")
	fmt.Printf("Ex26. Введите координаты точек x1, y1, x2, y2: ")
	fmt.Scan(&x1, &y1, &x2, &y2)
	fmt.Printf("Ex26. Расстояние: %.3f\n", math.Sqrt(math.Pow(x1-x2, 2)+math.Pow(y1-y2, 2)))
}

func ch2ex27() {
	var a, b, c float64
	fmt.Printf("Ex27(17). Нахождение периметра равнобедренной трапеции:\n")
	fmt.Printf("Ex27(17). Введите основание1, основание2 и высоту: ")
	fmt.Scan(&a, &b, &c)
	fmt.Printf("Ex27(17). Периметр: %.3f\n", a+b+2*math.Sqrt(c*c+(a-b)*(a-b)/4))
}

func ch2ex28() {
	var a, b, c float64
	fmt.Printf("Ex28. Нахождение площади равнобедренной трапеции:\n")
	fmt.Printf("Ex28. Введите основание1, основание2 и угол при большем основании: ")
	fmt.Scan(&a, &b, &c)
	fmt.Printf("Ex28. Площадь: %.3f\n", (b*b-a*a)/4*math.Tan(c*math.Pi/180))
}

func ch2ex29() {
	var x1, y1, x2, y2, x3, y3 float64

	fmt.Printf("Ex29. Найти периметр и площадь треугольника:\n")
	fmt.Printf("Ex29. Введите кординаты вершин x1, y1, x2, y2, x3, y3: ")
	fmt.Scan(&x1, &y1, &x2, &y2, &x3, &y3)

	a := distance(x1, y1, x2, y2)
	b := distance(x2, y2, x3, y3)
	c := distance(x3, y3, x1, y1)
	fmt.Printf("Ex29. Периметр: %.3f\n", a+b+c)

	p := (a + b + c) / 2
	s := math.Sqrt(p * (p - a) * (p - b) * (p - c))
	fmt.Printf("Ex29. Площадь: %.3f\n", s)
}

func ch2ex30() {
	var x1, y1, x2, y2, x3, y3, x4, y4 float64

	fmt.Printf("Ex30. Найти площадь четырехугольников:\n")
	fmt.Printf("Ex30. Введите кординаты вершин x1, y1, x2, y2, x3, y3, x4, y4: ")
	fmt.Scan(&x1, &y1, &x2, &y2, &x3, &y3, &x4, &y4)

	a := distance(x1, y1, x2, y2)
	b := distance(x2, y2, x3, y3)
	c := distance(x3, y3, x4, y4)
	d := distance(x4, y4, x1, y1)
	diag := distance(x1, y1, x3, y3)

	p1 := (a + b + diag) / 2
	p2 := (c + d + diag) / 2

	s1 := math.Sqrt(p1 * (p1 - a) * (p1 - b) * (p1 - diag))
	s2 := math.Sqrt(p2 * (p2 - c) * (p2 - d) * (p2 - diag))

	fmt.Printf("Ex30. Стороны: AB=%.3f, BC=%.3f, CD=%.3f, DA=%.3f\n", a, b, c, d)
	fmt.Printf("Ex30. Площадь треугольника ABD: %.3f\n", s1)
	fmt.Printf("Ex30. Площадь треугольника BCD: %.3f\n", s2)
	fmt.Printf("Ex30. Площадь четырёхугольника ABCD: %.3f\n", s1+s2)
}

func ch2ex31() {
	var candies, cookies, apples, a, b, c float64
	fmt.Printf("Ex31. Вычисление стоимости покупки:\n")
	fmt.Printf("Ex31. Введите цену за 1 кг конфет, печенья, яблок: ")
	fmt.Scan(&candies, &cookies, &apples)
	fmt.Printf("Ex31. Введите кол-во купленных конфет, печенья и яблок (кг): ")
	fmt.Scan(&a, &b, &c)
	fmt.Printf("Ex31. Стоимость покупки: %.2f\n", a*candies+b*cookies+c*apples)
}

func ch2ex32() {
	var a, b, c, d, e float64
	fmt.Printf("Ex32. Вычисление стоимости компьютеров:\n")
	fmt.Printf("Ex32. Введите стоимость монитора, системного блока, клавиатуры и мыши: ")
	fmt.Scan(&a, &b, &c, &d)
	cost := a + b + c + d
	fmt.Printf("Ex32. Три компьютера будет стоить: %.2f\n", cost*3)
	fmt.Printf("Ex32. Введите кол-во компьютеров: ")
	fmt.Scan(&e)
	fmt.Printf("Ex32. Стоимость %d компьютеров: %.2f\n", cost*e)
}

func ch2ex33() {
	var a, b float64
	fmt.Printf("Ex33. Вычисление среднего возраста:\n")
	fmt.Printf("Ex33. Введите возраст Тани и Мити: ")
	fmt.Scan(&a, &b)
	avg := (a + b) / 2
	fmt.Printf("Ex33. Средний возраст: %.2f\n", avg)
	fmt.Printf("Ex33. Возраст Тани отличается от среднего на %d лет", avg-a)
	fmt.Printf("Ex33. Возраст Мити отличается от среднего на %d лет", avg-b)
}

func ch2ex34() {
	var a, b, s float64
	fmt.Printf("Ex34. Вычисление время встречи автомобилей:\n")
	fmt.Printf("Ex34. Введите скорости 1 и 2 автомобиля (км/ч) и расстояние между ними (км): ")
	fmt.Scan(&a, &b, &s)
	fmt.Printf("Ex34. Автомобили встретятся через: %.2f\n часа(-ов)", s/(a+b))
}

func ch2ex35() {
	var a, b, s float64
	fmt.Printf("Ex35. Определить, какое расстояние будет через 30 мин. после опережения.\n")
	fmt.Printf("Ex35. Введите скорости (км/ч) 1 и 2 (V₁ > V₂) автомобилей и расстояние (км), на которое первый опережает второй:\n")
	fmt.Scan(&a, &b, &s)
	fmt.Printf("Ex35. Расстояние между автомобилями через 30 минут: %.3f км\n", s+0.5*(a-b))
}

func ch2ex36() {
	var a float64
	fmt.Printf("Ex36. Перевод цельсия в Кельвина и Фаренгейта.\n")
	fmt.Printf("Ex36. Введите температуру (°C):\n")
	fmt.Scan(&a)
	fmt.Printf("Ex36. Это: %.2f°F\n", a*1.8+32)
	fmt.Printf("Ex36. Это: %.2f°K\n", a-273.15)
}

func ch2ex37() {
	fmt.Printf("Ex37. 450 °F\n")
	fmt.Printf("Ex37. Это: %.2f°C\n", (450-32)/1.8)
}

func ch2ex38() {
	var a, b int
	fmt.Printf("Ex38. Вычисление суммы, разности, произведения, частного и среднего ариф.:\n")
	fmt.Printf("Ex38. Введите 2 целых числа: ")
	fmt.Scan(&a, &b)

	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d = %.2f\n", a, b, float64(a)/float64(b))
	fmt.Printf("(%d + %d) / 2 = %.2f\n", a, b, float64(a+b)/2.0)
}
