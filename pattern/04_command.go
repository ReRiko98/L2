package main

import "fmt"

type Command interface {
    Execute()
}

type Receiver struct{}

func (r *Receiver) Action() {
    fmt.Println("Receiver executing action")
}

type ConcreteCommand struct {
    receiver *Receiver
}

func NewConcreteCommand(receiver *Receiver) *ConcreteCommand {
    return &ConcreteCommand{receiver: receiver}
}

func (cc *ConcreteCommand) Execute() {
    cc.receiver.Action()
}

type Invoker struct {
    command Command
}

func (i *Invoker) SetCommand(command Command) {
    i.command = command
}

func (i *Invoker) ExecuteCommand() {
    i.command.Execute()
}

// func main() {
//     receiver := &Receiver{}
//     command := NewConcreteCommand(receiver)
//     invoker := &Invoker{}

//     invoker.SetCommand(command)
//     invoker.ExecuteCommand()
// }
