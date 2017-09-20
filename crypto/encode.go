package crypto

import "fmt"

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
		fmt.Println(e)

		e++
	}

	buffer := ""

	return buffer
}
func Encode(s string) string {

	buffer := ""

	return buffer
}
