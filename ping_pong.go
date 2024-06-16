// Напиши программу, в которой две горутины передают int друг другу 10 раз
// Каждый раз, когда горутина получает int, его нужно выводить
// Увеличивай int каждый раз при передаче
// Когда int == 10, правильно заверши программу

package main

import (
	"fmt"
	"sync"
)

// ch := make(chan int) -- создание небуферизированного канала
func main() {

	ch := make(chan int, 1)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func(ch chan int) {
		defer wg.Done()
		val := 0
		ok := false
		for {
			ch <- val

			val, ok = <-ch

			if ok == false {
				break
			}

			fmt.Println("Горутина 1 получила:", val)

			if val == 10 {
				close(ch)
				break
			}

			val++
		}
	}(ch)

	go func(ch chan int) {
		defer wg.Done()
		for {
			val, ok := <-ch

			if ok == false {
				break
			}

			fmt.Println("Горутина 2 получила:", val)

			if val == 10 {
				close(ch)
				break
			}

			ch <- (val + 1)
		}
	}(ch)

	wg.Wait()
	fmt.Println("Горутины завершили выполнение")
}
