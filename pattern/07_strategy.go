package pattern

type Strategy interface {
	Execute(a, b int) int
}

type ConcreteStrategyAdd struct{}

func (s *ConcreteStrategyAdd) Execute(a, b int) int {
	return a + b
}

type ConcreteStrategySubtract struct{}

func (s *ConcreteStrategySubtract) Execute(a, b int) int {
	return a - b
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.Execute(a, b)
}

/*
Применимость паттерна "Стратегия":
	Когда есть несколько похожих классов, которые отличаются только поведением.
Когда необходимо менять поведение объектов во время выполнения.

Преимущества паттерна "Стратегия":
	Позволяет заменять алгоритмы независимо от контекста и клиентского кода.
Увеличивает переиспользование кода.

Недостатки паттерна "Стратегия":
	Может привести к увеличению числа классов и объектов из-за дополнительных интерфейсов.

Примеры использования паттерна "Стратегия" на практике:
	Различные стратегии оплаты в интернет-магазине (например, кредитная карта, PayPal, наличные и т.д.).
Различные алгоритмы сортировки в приложении (например, быстрая сортировка, сортировка пузырьком и т.д.).
*/
