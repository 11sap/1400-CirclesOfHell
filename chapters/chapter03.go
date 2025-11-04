package chapters

import (
	"fmt"
)

func Run03() {
	tasks := []func(){}

	for {
		fmt.Print("\nГлава 3. Выберите задачу (1-51): ")
		var n int
		fmt.Scan(&n)

		if n == 0 {
			break
		}
		if n >= 1 && n <= 17 {
			tasks[n-1]()
		} else {
			fmt.Println("От 1 до 51! Для выхода введите 0.")
		}
	}
}
