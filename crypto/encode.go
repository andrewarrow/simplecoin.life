package crypto

//import "fmt"
import "math/rand"
import "time"

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

func GenKeys() (int, int, int) {
	primes := generatePrimes()
	rand.Seed(time.Now().UnixNano())
	pindex1 := rand.Intn(len(primes))
	pindex2 := rand.Intn(len(primes))
	p := primes[pindex1]
	q := primes[pindex2]
	n := p * q

	target := (p - 1) * (q - 1)
	e := rand.Intn(target-3) + 2
	for {
		if e == 2 {
			break
		}
		if coprime(e, target) {
			break
		}

		e--
	}
	p1 := phi(n)
	//fmt.Println(n, e)

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

	//fmt.Println(n, d)

	return n, e, d
}

func Encode(s string) string {

	buffer := ""

	return buffer
}
