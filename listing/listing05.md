package main
 
type customError struct {
     msg string
}
 
func (e *customError) Error() string {
    return e.msg
}
 
func test() *customError {
     {
         // do something
     }
     return nil
}
 
func main() {
    var err error
    err = test()
    if err != nil {
        println("error")
        return
    }
    println("ok")
}

Ответ: ok

Это происходит потому, что функция test() возвращает nil, что означает отсутствие ошибки. В условии if err != nil, переменная err сравнивается с nil, и поскольку test() вернула nil, условие не выполняется, и программа печатает "ok".