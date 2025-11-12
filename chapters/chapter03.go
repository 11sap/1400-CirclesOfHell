package chapters

import (
	"fmt"
)

func Run03() {
	tasks := []func(){
		ch3ex01, ch3ex02, ch3ex03, ch3ex04, ch3ex05, ch3ex06, ch3ex07, ch3ex08,
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
