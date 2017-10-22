package peer

import "fmt"
import "net"

func FindPeers() {
	ips, err := net.LookupIP("dnsseed.simplecoin.life")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, ip := range ips {
		fmt.Println(ip.String())
	}
}
