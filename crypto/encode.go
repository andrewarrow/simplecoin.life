package crypto

import "fmt"

func factors(a int, b int) int {
	for {
		m := a % b
		a = b
		b = m

		if b == 0 {
			break
		}

	}

	return a
}

func coprime(a int, b int) bool {
	return factors(a, b) == 1
}

func GenPubKey() string {
	p := 97
	q := 197
	n := p * q
	fmt.Println(n)

	e := 2
	target := (p - 1) * (q - 1)
	for {
		if e >= target {
			break
		}
		if coprime(e, target) {
		} else {
			fmt.Println(e)
		}

		e++
	}

	buffer := ""

	return buffer
}
func Encode(s string) string {

	buffer := ""

	return buffer
}
