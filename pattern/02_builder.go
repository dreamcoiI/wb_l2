package pattern

type Pizza struct {
	dough   string
	sauce   string
	topping string
}

type PizzaBuilder interface {
	SetDough() PizzaBuilder
	SetSauce() PizzaBuilder
	SetTopping() PizzaBuilder
	GetPizza() Pizza
}

type HawaiianPizzaBuilder struct {
	pizza Pizza
}

func (b *HawaiianPizzaBuilder) SetDough() PizzaBuilder {
	b.pizza.dough = "cross"
	return b
}

func (b *HawaiianPizzaBuilder) SetSauce() PizzaBuilder {
	b.pizza.sauce = "mild"
	return b
}

func (b *HawaiianPizzaBuilder) SetTopping() PizzaBuilder {
	b.pizza.topping = "ham+pineapple"
	return b
}

func (b *HawaiianPizzaBuilder) GetPizza() Pizza {
	return b.pizza
}

type Waiter struct {
	builder PizzaBuilder
}

func (w *Waiter) SetBuilder(b PizzaBuilder) {
	w.builder = b
}

func (w *Waiter) Construct() {
	w.builder.SetDough().SetSauce().SetTopping()

}

/*
Паттерн "Строитель" (Builder) применяется в случаях, когда требуется создавать сложные объекты, имеющие различные
вариации или конфигурации.

Плюсы:
	- Упрощает процесс создания сложных объектов, разделяя его на отдельные шаги.
	- Позволяет создавать различные вариации объектов, используя один и тот же процесс конструирования.
	- Улучшает управляемость и читаемость кода, делая его более структурированным.

Минусы:
	- Может привести к созданию большого количества классов, если вариаций объектов слишком много.
	- Увеличивает сложность кода, особенно если объект имеет множество частей или шагов конструирования.

Пример применения:
	- Построение SQL-запросов: Паттерн "Строитель" может использоваться для поэтапного создания сложных SQL-запросов
	  с различными параметрами.
*/
