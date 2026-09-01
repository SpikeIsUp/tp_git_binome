package main
import "fmt"
func Add(a, b int) int { return a + b }
func main() {
fmt.Println("Add(2, 3) =", Add(2, 3))
Divide(10, 0)
}
//go run main.go # Sortie: Add(2, 3) = 5
//git add main.go && git commit -m "feat: fonction Add de base"
//git push -u origin main



Divide(a, b int) int { return a / b }