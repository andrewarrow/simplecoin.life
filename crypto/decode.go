package crypto

import "fmt"

//import "encoding/binary"
import "math/big"

func Decode(bc big.Int, n, d int64) string {

	//Plaintext = Cd mod n
	bd := big.NewInt(d)
	bn := big.NewInt(n)
	fmt.Println("-bc", bc, bd, bn)

	var limit big.Int
	p := limit.Exp(&bc, bd, bn)

	fmt.Println(p)

	return ""
}
