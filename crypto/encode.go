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

func phi(n uint64) uint64 {
	result := uint64(1)
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
	p := uint64(primes[pindex1])
	q := uint64(primes[pindex2])
	n := p * q

	target := (p - 1) * (q - 1)
	bigRan := rand.Uint64()
	index := uint64(0)
	for {
		if bigRan >= 2 && bigRan < target-3 {
			break
		}
		bigRan = bigRan / 2
		index++
	}
	e := bigRan
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

	return n, e, d
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

	var bp big.Int
	var be big.Int
	var bn big.Int
	bp.SetUint64(p)
	be.SetUint64(e)
	bn.SetUint64(n)

	var limit big.Int
	c := limit.Exp(&bp, &be, &bn)
	return *c
}
