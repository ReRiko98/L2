package main

import "fmt"

type PaymentStrategy interface {
    Pay(amount float64)
}

type CardPayment struct{}

func (cp *CardPayment) Pay(amount float64) {
    fmt.Printf("Paying %.2f using card\n", amount)
}

type CashPayment struct{}

func (cp *CashPayment) Pay(amount float64) {
    fmt.Printf("Paying %.2f using cash\n", amount)
}

type ShoppingCart struct {
    paymentMethod PaymentStrategy
}

func NewShoppingCart(paymentMethod PaymentStrategy) *ShoppingCart {
    return &ShoppingCart{paymentMethod: paymentMethod}
}

func (sc *ShoppingCart) Pay(amount float64) {
    sc.paymentMethod.Pay(amount)
}

// func main() {
//     cardPayment := &CardPayment{}
//     cashPayment := &CashPayment{}

//     cart := NewShoppingCart(cardPayment)
//     cart.Pay(100)

//     cart.paymentMethod = cashPayment
//     cart.Pay(50)
// }
