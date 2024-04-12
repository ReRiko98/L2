package main

import "fmt"

type Visitor interface {
    VisitElementA(element *ElementA)
    VisitElementB(element *ElementB)
}

type ConcreteVisitor struct{}

func (cv *ConcreteVisitor) VisitElementA(element *ElementA) {
    fmt.Println("Visited Element A")
}

func (cv *ConcreteVisitor) VisitElementB(element *ElementB) {
    fmt.Println("Visited Element B")
}

type Element interface {
    Accept(visitor Visitor)
}

type ElementA struct{}

func (ea *ElementA) Accept(visitor Visitor) {
    visitor.VisitElementA(ea)
}

type ElementB struct{}

func (eb *ElementB) Accept(visitor Visitor) {
    visitor.VisitElementB(eb)
}

type ObjectStructure struct {
    elements []Element
}

func (os *ObjectStructure) Attach(element Element) {
    os.elements = append(os.elements, element)
}

func (os *ObjectStructure) Accept(visitor Visitor) {
    for _, e := range os.elements {
        e.Accept(visitor)
    }
}

// func main() {
//     visitor := &ConcreteVisitor{}

//     objectStructure := &ObjectStructure{}
//     objectStructure.Attach(&ElementA{})
//     objectStructure.Attach(&ElementB{})

//     objectStructure.Accept(visitor)
// }
