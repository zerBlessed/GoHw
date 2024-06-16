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
			defer wg.Done()
			j := 0
			for {
				randomNumber := rand.Intn(1000)
				if randomNumber%2 != 0 {
					ch <- randomNumber
					j++
				}

				if j == 10 {
					break
				}
			}
		}()
	}

	go func() {
		defer wg.Done()
		numbers := make([]int, 0, 100)
		for i := 0; i < 100; i++ {
			numbers = append(numbers, <-ch)
		}
		fmt.Println("Результат:", numbers)
	}()

	wg.Wait()
	fmt.Println("Все горутины завершены!")
}
