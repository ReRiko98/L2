package main

import "fmt"

type Handler interface {
    SetNext(handler Handler)
    Handle(request int)
}

type ConcreteHandler1 struct {
    next Handler
}

func (ch *ConcreteHandler1) SetNext(handler Handler) {
    ch.next = handler
}

func (ch *ConcreteHandler1) Handle(request int) {
    if request >= 0 && request < 10 {
        fmt.Println("Handled by ConcreteHandler1")
    } else if ch.next != nil {
        ch.next.Handle(request)
    }
}

type ConcreteHandler2 struct {
    next Handler
}

func (ch *ConcreteHandler2) SetNext(handler Handler) {
    ch.next = handler
}

func (ch *ConcreteHandler2) Handle(request int) {
    if request >= 10 && request < 20 {
        fmt.Println("Handled by ConcreteHandler2")
    } else if ch.next != nil {
        ch.next.Handle(request)
    }
}

// func main() {
//     handler1 := &ConcreteHandler1{}
//     handler2 := &ConcreteHandler2{}

//     handler1.SetNext(handler2)

//     requests := []int{5, 15, 25}

//     for _, req := range requests {
//         handler1.Handle(req)
//     }
// }
