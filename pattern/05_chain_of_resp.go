package pattern

import "fmt"

type Handler interface {
	Execute(int)
	SetNext(Handler)
}

type ConcreteHandler struct {
	next Handler
}

func (c *ConcreteHandler) SetNext(next Handler) {
	c.next = next
}

type ConcreteHandlerOne struct {
	ConcreteHandler
}

func (c *ConcreteHandlerOne) Execute(i int) {
	if i < 0 {
		fmt.Printf("Число %d обработано ConcreteHandlerOne\n", i)
	} else if c.next != nil {
		c.next.Execute(i)
	}
}

type ConcreteHandlerTwo struct {
	ConcreteHandler
}

func (c *ConcreteHandlerTwo) Execute(i int) {
	if i == 0 {
		fmt.Printf("Число %d обработано ConcreteHandlerTwo\n", i)
	} else if c.next != nil {
		c.next.Execute(i)
	}
}

type ConcreteHandlerThree struct {
	ConcreteHandler
}

func (c *ConcreteHandlerThree) Execute(i int) {
	if i > 0 {
		fmt.Printf("Число %d обработано ConcreteHandlerThree\n", i)
	} else if c.next != nil {
		c.next.Execute(i)
	}
}

/*
Применимость паттерна "Цепочка обязанностей":
	Когда необходимо передавать запросы по цепочке объектов для их обработки.
Когда набор объектов, способных обработать запрос, должен быть определен динамически.

Преимущества паттерна "Цепочка обязанностей":
	Уменьшает зависимость между отправителем запроса и получателем.
Позволяет настраивать поведение системы динамически.

Недостатки паттерна "Цепочка обязанностей":
	Запрос может остаться необработанным.

Примеры использования паттерна "Цепочка обязанностей" на практике:
	Сетевые протоколы, такие как HTTP, используют цепочку обработчиков для маршрутизации запросов.
Обработка событий в пользовательских интерфейсах.
*/
