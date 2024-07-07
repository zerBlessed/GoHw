package main

import (
	"fmt"
	"sync"
	"time"
)

// Необходимо реализовать паттерн Producer-Consumer (производитель-потребитель)
// Многопоточная программа, в которой несколько продюсеров создают задачи, а несколько консюмеров получают задачи обрабатывают их
// Для имитации работы можно использовать defaultTaskCallback

type Task struct {
	id           int
	taskCallback func(id int) error
}

func defaultTaskCallback(id int) error {
	fmt.Printf("worker %d doing useful work...\n", id)
	time.Sleep(1 * time.Second)
	return nil
}

func main() {
	const taskCount = 100   // // Количество задач для производства
	const producerCount = 3 // Количество продюсеров
	const consumerCount = 7 // Кол-во консюмеров

	taskChannel := make(chan Task, taskCount)

	wgProducer := &sync.WaitGroup{}
	wgProducer.Add(producerCount)

	wgConsumer := &sync.WaitGroup{}
	wgConsumer.Add(consumerCount)

	for i := 0; i < producerCount; i++ {
		go func(producerID int) {
			defer wgProducer.Done()
			for j := 0; j < taskCount/producerCount; j++ {
				task := Task{
					id:           producerID*taskCount/producerCount + j,
					taskCallback: defaultTaskCallback,
				}
				taskChannel <- task
			}
		}(i)
	}

	wgProducer.Wait()
	close(taskChannel)

	for i := 0; i < consumerCount; i++ {
		go func() {
			defer wgConsumer.Done()
			for task := range taskChannel {
				err := task.taskCallback(task.id)
				if err != nil {
					fmt.Printf("error\n")
				}
			}
		}()
	}

	wgConsumer.Wait()

	fmt.Println("Все задачи обработаны!")
}
