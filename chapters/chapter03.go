package chapters

import (
	"fmt"
)

func Run03() {
	tasks := []func(){
		ch3ex01,
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
