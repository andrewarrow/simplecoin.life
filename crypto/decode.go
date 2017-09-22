package crypto

import "fmt"
import "encoding/binary"
import "math"

func Decode(s string, n, d int) string {

	//Plaintext = Cd mod n

	//byteArray := []byte{153, 239, 5, 0}
	pow := uint64(math.Pow(float64(389017), float64(d)))
	plain := pow % uint64(n)

	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, plain)
	fmt.Println(b)

	return string(b)
}
