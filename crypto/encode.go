package crypto

import "fmt"

//import "encoding/binary"
import "math/rand"
import "time"
import "math/big"

func greatest_common_divisor(a uint64, b uint64) uint64 {
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

func coprime(a uint64, b uint64) bool {
	return greatest_common_divisor(a, b) == 1
}

func phi(n uint64) int {
	result := 1
	i := uint64(2)
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

func GenKeys() (uint64, uint64, uint64) {
	primes := generatePrimes()
	rand.Seed(time.Now().UnixNano())
	pindex1 := rand.Intn(len(primes))
	pindex2 := rand.Intn(len(primes))
	p := primes[pindex1]
	q := primes[pindex2]
	n := uint64(p) * uint64(q)

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
		if coprime(uint64(e), uint64(target)) {
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

	return n, uint64(e), uint64(d)
}

func EncodeString(s string, n, e uint64) []big.Int {
	list := []byte(s)
	buffer := []big.Int{}
	for _, p := range list {
		buffer = append(buffer, Encode(uint64(p), n, e))
	}
	return buffer
}

func Encode(p, n, e uint64) big.Int {

	if p >= n {
		fmt.Println("ERR p >= n")
		return *big.NewInt(0)
	}

	// C = Pe mod n
	//fmt.Println("encode", s, n, e)
	//byteArray := []byte(s)
	//byteArray = append(byteArray, 0)
	//byteArray = append(byteArray, 0)
	//byteArray = append(byteArray, 0)
	//fmt.Println(byteArray)
	//p := binary.LittleEndian.Uint64(byteArray)
	//p = 55010
	//fmt.Println("p", p)

	var bp big.Int
	var be big.Int
	var bn big.Int
	bp.SetUint64(p)
	be.SetUint64(e)
	bn.SetUint64(n)

	//fmt.Println("bn", bn)
	var limit big.Int
	c := limit.Exp(&bp, &be, &bn)
	//fmt.Println("ec", c)
	return *c
}
