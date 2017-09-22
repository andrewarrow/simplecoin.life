package crypto

import "fmt"
import "encoding/binary"
import "math/rand"
import "math"
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

	debug := false

	if debug {
		fmt.Println("n", n)
	}

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
	if debug {
		fmt.Println("e", e)
	}
	p1 := phi(n)
	if debug {
		fmt.Println("p1", p1)
	}
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
	if debug {
		fmt.Println("d", d)
	}

	//fmt.Println(n, d)

	return n, e, d
}

func Encode(s string, n, e int) string {

	fmt.Println(s)
	byteArray := []byte(s)
	byteArray = append(byteArray, 0)
	byteArray = append(byteArray, 0)
	byteArray = append(byteArray, 0)
	fmt.Println(byteArray)
	p := int64(binary.LittleEndian.Uint64(byteArray))
	fmt.Println(p)

	pow := uint64(math.Pow(float64(p), float64(e)))
	fmt.Println(pow)
	c := pow % uint64(n)
	fmt.Println(c)

	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, c)
	fmt.Println(b)

	return string(b)
}
