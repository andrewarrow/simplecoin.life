package crypto

import "fmt"

func greatest_common_divisor(a int, b int) int {
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
	return greatest_common_divisor(a, b) == 1
}

func phi(n int) int {
	result := 1
	i := 2
	for {
		if greatest_common_divisor(i, n) == 1 {
			result++
		}
		i++
		if i >= n {
			break
		}
	}
	return result
}

func GenKeys() string {
	p := 97
	q := 197
	p = 7
	q = 13
	n := p * q

	target := (p - 1) * (q - 1)
	e := target - 2
	for {
		if e == 2 {
			break
		}
		if coprime(e, target) {
			break
		}

		e--
	}
	e = 5
	p1 := phi(n)
	fmt.Println(n, e)

	//7*d mod 40 = 1
	//		e*d % p1 == 1
	// d such that e*d = 1 mod phi(n)
	target = (p - 1) * (q - 1)
	d := target - 2
	for {
		if d == 2 {
			break
		}
		if (e*d)%p1 == 1 {
			break
		}

		d--
	}

	fmt.Println(n, d)
	buffer := ""

	return buffer
}

func Encode(s string) string {

	buffer := ""

	return buffer
}