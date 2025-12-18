package chapters

import "fmt"

func Run11() {
	tasks := []func(){
		ch11ex01,
	}

	for {
		fmt.Print("\nГлава 11. Выберите задачу (1-261): ")
		var n int
		fmt.Scan(&n)

		if n == 0 {
			break
		}
		if n >= 1 && n <= 261 {
			tasks[n-1]()
		} else {
			fmt.Println("От 1 до 261! Для выхода введите 0.")
		}
	}
}

func ch11ex01() {
	fmt.Println("11.1. Заполнение массива из 8 элементов")

	arr := [8]int{37, 0, 50, 46, 34, 46, 0, 13}

	fmt.Print("Массив: [")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d", arr[i])
		if i < len(arr)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println("]")
}
