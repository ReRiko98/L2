package main
 
import (
    "fmt"
    "os"
)
 
func Foo() error {
    var err *os.PathError = nil
    return err
}
 
func main() {
    err := Foo()
    fmt.Println(err)
    fmt.Println(err == nil)
}

Ответ: nill false

Это связано с тем, что функция Foo() возвращает переменную типа *os.PathError, которая по умолчанию будет равна nil. Однако, так как err является конкретным указателем, а не пустым интерфейсом, он не будет считаться nil при сравнении.