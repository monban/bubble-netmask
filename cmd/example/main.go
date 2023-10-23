package main

import (
	"net"

	"github.com/charmbracelet/log"
	netmask "github.com/monban/bubble-netmask"
)

func main() {
	mask, err := netmask.New("10.1.1.0").Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("netmask selected", "mask", net.IP(mask))
}
