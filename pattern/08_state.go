package pattern

import "fmt"

type State interface {
	doAction(c *StateContext)
}

type StartState struct{}

func (s *StartState) doAction(c *StateContext) {
	fmt.Println("Player is in start state")
	c.setState(s)
}

type StopState struct{}

func (s *StopState) doAction(c *StateContext) {
	fmt.Println("Player is in stop state")
	c.setState(s)
}

type StateContext struct {
	state State
}

func (c *StateContext) setState(state State) {
	c.state = state
}

func (c *StateContext) getState() State {
	return c.state
}

/*
Применимость паттерна "Состояние":
	Когда поведение объекта должно зависеть от его состояния и должно меняться во время выполнения.
Когда в коде есть множество больших условных операторов, которые контролируют поведение объекта в зависимости от его состояния.

Преимущества паттерна "Состояние":
	Позволяет управлять переходами между состояниями объекта.
Упрощает код, так как устраняет большие блоки условных операторов.

Недостатки паттерна "Состояние":
	Может привести к увеличению числа классов и объектов из-за дополнительных состояний.

Примеры использования паттерна "Состояние" на практике:
	Реализация системы торгового автомата с различными состояниями (например, "ожидание монеты", "выбор товара", "выдача товара" и т.д.).
Разработка игрового персонажа с различными состояниями (например, "бег", "прыжок", "атака" и т.д.).
*/
