package peer

import "fmt"
import "net"

func FindPeers() []string {
	result := []string{}
	ips, err := net.LookupIP("dnsseed.simplecoin.life")
	if err != nil {
		fmt.Println(err)
		return result
	}
	for _, ip := range ips {
		result = append(result, ip.String())
		fmt.Println(ip.String())
	}
	result = append(result, "127.0.0.1")
	return result
}
