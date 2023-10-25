package pattern

import "fmt"

type Visitor interface {
	VisitForSquare(*Square)
	VisitForCircle(*Circle)
	VisitForRectangle(*Rectangle)
}

type Shape interface {
	Accept(Visitor)
}

type Square struct {
	side int
}

func (s *Square) Accept(v Visitor) {
	v.VisitForSquare(s)
}

type Circle struct {
	radius int
}

func (c *Circle) Accept(v Visitor) {
	v.VisitForCircle(c)
}

type Rectangle struct {
	l int
	b int
}

func (r *Rectangle) Accept(v Visitor) {
	v.VisitForRectangle(r)
}

type AreaVisitor struct {
	area int
}

func (a *AreaVisitor) VisitForSquare(s *Square) {
	fmt.Println("Calculating area for square")
}

func (a *AreaVisitor) VisitForCircle(c *Circle) {
	fmt.Println("Calculating area for circle")
}

func (a *AreaVisitor) VisitForRectangle(r *Rectangle) {
	fmt.Println("Calculating area for rectangle")
}

/*
Применимость паттерна посетитель:
	Когда необходимо выполнить операцию над всеми объектами в сложной структуре объектов, например, дереве.
Когда структура объектов не поддерживает некоторые операции, которые необходимы в приложении, и требуется добавить эти операции, не изменяя классы объектов.

Преимущества паттерна посетитель:
	Упрощает добавление операций над объектами без изменения их структуры.
Позволяет собрать связанные операции в одном классе, улучшая читаемость и обслуживаемость кода.

Недостатки паттерна посетитель:
	Увеличивает количество классов в системе из-за введения новых классов посетителей.
Может затруднить добавление новых типов элементов, требующих обновления всех посетителей.

Примеры использования:
	Обработка элементов в структурах данных, таких как деревья или списки.
Обработка различных узлов в документах HTML или XML.
Переопределение операций для объектов различных типов в играх или графических приложениях.
*/
