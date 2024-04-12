package main
 
import (
    "fmt"
    "math/rand"
    "time"
)
 
func asChan(vs ...int) <-chan int {
   c := make(chan int)
 
   go func() {
       for _, v := range vs {
           c <- v
           time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
      }
 
      close(c)
  }()
  return c
}
 
func merge(a, b <-chan int) <-chan int {
   c := make(chan int)
   go func() {
       for {
           select {
               case v := <-a:
                   c <- v
              case v := <-b:
                   c <- v
           }
      }
   }()
 return c
}
 
func main() {
 
   a := asChan(1, 3, 5, 7)
   b := asChan(2, 4 ,6, 8)
   c := merge(a, b )
   for v := range c {
       fmt.Println(v)
   }
}

Ответ: от 1 до 8

Эта программа создает два канала, a и b, которые отправляют значения в каждом канале по порядку. Затем она использует функцию merge, чтобы объединить значения из этих двух каналов в один канал c, который она затем читает и выводит на экран в основной функции main().

Таким образом, программа выведет числа от 1 до 8 в произвольном порядке, так как каналы a и b отправляют значения асинхронно.