package main

import "fmt"

type Context struct {
    state State
}

func NewContext() *Context {
    return &Context{state: &ConcreteStateA{}}
}

func (c *Context) Request() {
    c.state.Handle(c)
}

func (c *Context) SetState(state State) {
    c.state = state
}

type State interface {
    Handle(context *Context)
}

type ConcreteStateA struct{}

func (csa *ConcreteStateA) Handle(context *Context) {
    fmt.Println("Handling request in State A")
    context.SetState(&ConcreteStateB{})
}

type ConcreteStateB struct{}

func (csb *ConcreteStateB) Handle(context *Context) {
    fmt.Println("Handling request in State B")
    context.SetState(&ConcreteStateA{})
}

// func main() {
//     context := NewContext()

//     context.Request()
//     context.Request()
// }
