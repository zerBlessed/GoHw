package main

// Нужно создать имитацию чата
// Есть 3 пользователя, которые в случайные моменты времени пишут сообщения
// Есть 1 админ, который в случайный момент времени (от 10 до 30 секунд) может закрыть чат
// Когда любой пользователь отправляет сообщение, нужно вывести в консоль имя пользователя и само сообщение
// Сигнатуры функций можно изменять

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	Author  string
	Content string
}

// Имитация пользователя. Пользователь в случайный момент времени отправляет сообщение
func simulateUser(name string, chat chan<- Message) {
	for {
		time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)

		message := Message{
			Author:  name,
			Content: fmt.Sprintf("%s: привет!", name),
		}

		chat <- message
	}
}

// Имитация админа. Админ сообщения не отправляет, но может закрыть общий чат
func simulateAdmin(chat chan Message) {
	for {
		time.Sleep(time.Duration(rand.Intn(20)+10) * time.Second)

		close(chat)
		fmt.Println("Админ закрыл чат.")
		break
	}
}

// Логика чата. Получаем сообщение -- печатаем в консоль. Помнить, что админ может закрыть чат в любой момент
func chatManager(chat chan Message) {
	for {
		message, ok := <-chat
		if !ok {
			fmt.Println("Чат закрыт.")
			return
		}
		fmt.Println(message.Content)
	}
}

func main() {
	chat := make(chan Message)

	go simulateUser("1", chat)
	go simulateUser("2", chat)
	go simulateUser("3", chat)

	go simulateAdmin(chat)

	chatManager(chat)
}
