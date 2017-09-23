package crypto

import "math/big"

func Decode(bc big.Int, n, d uint64) uint64 {

	//Plaintext = Cd mod n
	var bd big.Int
	var bn big.Int
	bd.SetUint64(d)
	bn.SetUint64(n)

	var limit big.Int
	p := limit.Exp(&bc, &bd, &bn)
	return p.Uint64()
}
