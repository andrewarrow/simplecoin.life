package crypto

//import "fmt"
//import "encoding/binary"
import "math/big"

func Decode(bc big.Int, n, d uint64) big.Int {

	//Plaintext = Cd mod n
	var bd big.Int
	var bn big.Int
	bd.SetUint64(d)
	bn.SetUint64(n)
	//fmt.Println("-bc", bc, bd, bn)

	var limit big.Int
	p := limit.Exp(&bc, &bd, &bn)
	return *p
}
