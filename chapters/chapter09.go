package chapters

import "fmt"

func Run09() {
	tasks := []func(){
		ch9ex01,
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
	fmt.Printf("Ex1. Сумма 10 чисел\n")
	sum := 0
	for i := 1; i <= 10; i++ {
		var num int
		fmt.Scan(&num)
		sum += num
	}
	fmt.Printf("Сумма: %d\n", sum)
}
