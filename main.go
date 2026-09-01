package main
import "fmt"
func Add(a, b int) int { return a + b }
func main() {

    fmt.Println("Add(2, 3) =", Add(2, 3))
    Multiply(a, b int) int { return a * b }
    Divide (10, 0)
}

func Divide(a, b int) (int, error) {
    if b == 0 { return 0, errors.New("division by zero") }
    return a / b, nil
}