package main

import "fmt"

type Product interface {
    Name() string
}

type ConcreteProduct1 struct{}

func (cp1 *ConcreteProduct1) Name() string {
    return "ConcreteProduct1"
}

type ConcreteProduct2 struct{}

func (cp2 *ConcreteProduct2) Name() string {
    return "ConcreteProduct2"
}

type Creator interface {
    FactoryMethod() Product
}

type ConcreteCreator1 struct{}

func (cc1 *ConcreteCreator1) FactoryMethod() Product {
    return &ConcreteProduct1{}
}

type ConcreteCreator2 struct{}

func (cc2 *ConcreteCreator2) FactoryMethod() Product {
    return &ConcreteProduct2{}
}

// func main() {
//     creator1 := &ConcreteCreator1{}
//     product1 := creator1.FactoryMethod()
//     fmt.Println("Product1:", product1.Name())

//     creator2 := &ConcreteCreator2{}
//     product2 := creator2.FactoryMethod()
//     fmt.Println("Product2:", product2.Name())
// }
