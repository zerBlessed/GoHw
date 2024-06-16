// Запроси данные из источников в параллельном режиме

// go func() -- запуск горутины
// sync.WaitGroup -- для координации горутин

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ulrs := []string{
		"google.com",
		"yandex.ru",
		"linkedin.com",
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(ulrs))

	for _, url := range ulrs {
		go func(url string) {
			defer wg.Done() // Уменьшаем счетчик при завершении горутины
			checkUrl(url)
		}(url)
	}

	wg.Wait()
	fmt.Println("Горутины завершили выполнение")
}

func checkUrl(url string) bool {
	fmt.Println("checked url:", url)
	time.Sleep(1 * time.Second)
	return true
}
