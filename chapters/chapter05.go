package chapters

import (
	"fmt"
)

func Run05() {
	tasks := []func(){
		ch5ex01,
	}

	for {
		fmt.Print("\nГлава 5. Выберите задачу (1-94): ")
		var n int
		fmt.Scan(&n)

		if n == 0 {
			break
		}
		if n >= 1 && n <= 94 {
			tasks[n-1]()
		} else {
			fmt.Println("От 1 до 94! Для выхода введите 0.")
		}
	}
}

func ch5ex01() {
	fmt.Printf("Ex01. Ряд чисел 20\n")
	for i := 0; i < 8; i++ {
		fmt.Printf("20 ")
	}
	fmt.Println()
}
