package pattern

type Fproduct interface {
	SetName(name string)
	GetName() string
}

type ConcreteProductA struct {
	name string
}

func (p *ConcreteProductA) SetName(name string) {
	p.name = name
}

func (p *ConcreteProductA) GetName() string {
	return p.name
}

type ConcreteProductB struct {
	name string
}

func (p *ConcreteProductB) SetName(name string) {
	p.name = name
}

func (p *ConcreteProductB) GetName() string {
	return p.name
}

type Creator interface {
	CreateProduct() Fproduct
}

type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) CreateProduct() Fproduct {
	return &ConcreteProductA{}
}

type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) CreateProduct() Fproduct {
	return &ConcreteProductB{}
}

/*
Применимость паттерна "Фабричный метод":
	Когда заранее неизвестно, объекты каких типов будут созданы.
Когда система должна быть независимой от процесса создания новых объектов и их типов.

Преимущества паттерна "Фабричный метод":
	Упрощает добавление новых продуктов в программу.
Избавляет клиентский код от зависимости от конкретных классов продуктов.

Недостатки паттерна "Фабричный метод":
	Увеличивает количество подклассов.

Примеры использования паттерна "Фабричный метод" на практике:
	Создание различных типов документов в текстовом редакторе (например, создание новых документов формата PDF, DOCX и т.д.).
Игровая разработка, где разные типы игровых персонажей могут быть созданы с помощью фабричного метода.
*/
