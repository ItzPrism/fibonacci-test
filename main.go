package main
import "fmt"
import "math/big"
func main() {
	a := big.NewInt(0)
	b := big.NewInt(1)
	for (
		a.Add(a, b)
		fmt.Println(a)
		b.Add(b, a)
		fmt.Println(b)
	)
}
