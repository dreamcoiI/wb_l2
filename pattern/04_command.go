package pattern

import "fmt"

type Command interface {
	Execute()
}

type LightOnCommand struct {
	Light *Light
}

func (c *LightOnCommand) Execute() {
	c.Light.On()
}

type LightOffCommand struct {
	Light *Light
}

func (c *LightOffCommand) Execute() {
	c.Light.Off()
}

type Light struct{}

func (l *Light) On() {
	fmt.Println("Свет включен")
}

func (l *Light) Off() {
	fmt.Println("Свет выключен")
}

type RemoteControl struct {
	Command Command
}

func (r *RemoteControl) PressButton() {
	r.Command.Execute()
}

/*
Применимость паттерна "Команда":
	Когда требуется параметризовать объекты выполняемыми действиями.
Когда необходима поддержка отмены операций.
Когда необходима поддержка выполнения операций в разное время или в разных потоках.

Преимущества паттерна "Команда":
	Уменьшает связанность между объектами отправителя и получателя.
Позволяет легко добавлять новые команды в систему.

Недостатки паттерна "Команда":
	Увеличивает количество классов и объектов в системе.

Примеры использования паттерна "Команда" на практике:
	Интерфейс команд в текстовом редакторе для выполнения операций "отменить" и "повторить".
Управление действиями в видеоиграх с помощью кнопок на геймпаде.
Контроль выполнения операций в очереди задач в системах управления задачами.
*/
