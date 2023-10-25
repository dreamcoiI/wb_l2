package pattern

import "fmt"

type Product struct {
	Name  string
	Price float64
}

type Order struct {
	products []Product
}

func (o *Order) addProduct(p Product) {
	o.products = append(o.products, p)
}

type Shipping struct{}

func (s *Shipping) calculateShippingCost(order Order) float64 {
	return float64(len(order.products)) * 5.0
}

type OrderFacade struct {
	ProductSubsystem  *Product
	OrderSubsystem    *Order
	ShippingSubsystem *Shipping
}

func (of *OrderFacade) addProductToOrder(productName string, price float64) {
	product := Product{Name: productName, Price: price}
	of.OrderSubsystem.addProduct(product)
}

func (of *OrderFacade) placeOrder() {
	fmt.Println("Order placed. Total cost:", of.ShippingSubsystem.calculateShippingCost(*of.OrderSubsystem))
}

/*
Паттерн используется когда нужно представить простой интерфейс для работы со сложной подсистемой
Плюсы:
- Упрощение работы с сложной системой путем предоставления унифицированного интерфейса.
- Сокрытие деталей реализации подсистемы от клиентского кода, что делает код более управляемым.
- Сокращение зависимостей между компонентами системы, что позволяет улучшить модульность и гибкость кода.

Минусы:
- Возможное усложнение отладки и тестирования всей системы из-за упрощения интерфейса.
- Возможное увеличение накладных расходов из-за дополнительного слоя абстракции между клиентским кодом и подсистемой.
- Риск создания недостаточно гибкого интерфейса, который не удовлетворит все потребности клиентов.

Реальный пример-графические библиотеки(например OpenGL) представляют фасад для упрощения работы
*/
