package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	// Максимальное количество горутин
	const maxGoroutinesCount = 10
	// Создай канал для обмена числами.
	ch := make(chan int)
	// Создай WaitGroup, чтобы контролировать запущенные горутины.
	var wg sync.WaitGroup
	wg.Add(maxGoroutinesCount + 1) // Увеличиваем счетчик на количество горутин

	// Запуск горутин, генерирующих числа
	for i := 0; i < maxGoroutinesCount; i++ {
		go func() {
			defer wg.Done() // Уменьшаем счетчик при завершении горутины
			j := 0
			for { // Каждая горутина генерирует 10 чисел
				randomNumber := rand.Intn(1000)
				if randomNumber%2 != 0 { // Проверка на четность
					ch <- randomNumber
					j++
				}

				if j == 10 {
					break
				}
			}
		}()
	}

	// Создание горутины для чтения из канала
	go func() {
		defer wg.Done() // Завершаем горутину
		numbers := make([]int, 0, 100)
		for i := 0; i < 100; i++ { // Ожидаем получения 100 чисел
			numbers = append(numbers, <-ch)
		}
		fmt.Println("Результат:", numbers)
	}()

	wg.Wait() // Ждем завершения всех горутин
	fmt.Println("Все горутины завершены!")
}
