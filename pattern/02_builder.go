package main

import "fmt"

type Product struct {
    part1 string
    part2 string
}

type Builder interface {
    BuildPart1()
    BuildPart2()
    GetProduct() *Product
}

type ConcreteBuilder struct {
    product *Product
}

func NewConcreteBuilder() *ConcreteBuilder {
    return &ConcreteBuilder{product: &Product{}}
}

func (b *ConcreteBuilder) BuildPart1() {
    b.product.part1 = "Part 1"
}

func (b *ConcreteBuilder) BuildPart2() {
    b.product.part2 = "Part 2"
}

func (b *ConcreteBuilder) GetProduct() *Product {
    return b.product
}

type Director struct {
    builder Builder
}

func NewDirector(builder Builder) *Director {
    return &Director{builder: builder}
}

func (d *Director) Construct() *Product {
    d.builder.BuildPart1()
    d.builder.BuildPart2()
    return d.builder.GetProduct()
}

func main() {
    builder := NewConcreteBuilder()
    director := NewDirector(builder)

    product := director.Construct()

    fmt.Println("Product part 1:", product.part1)
    fmt.Println("Product part 2:", product.part2)
}
