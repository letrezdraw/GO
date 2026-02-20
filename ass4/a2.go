//p-rint multiplication of two numberws using method
package main
import "fmt"
type Multiplier struct{
	a,b int
	
}

func (m Multiplier) Multiply() int {
	return m.a * m.b
}
func main() {
	m := Multiplier{a: 5, b: 10}
	result := m.Multiply()
	fmt.Printf("The multiplication of %d and %d is: %d\n", m.a, m.b, result)
}
