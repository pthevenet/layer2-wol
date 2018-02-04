package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/pthevenet/layer2-wol/wol"
)

func main() {

	// user input flags parsing

	targetMAC := flag.String("mac", "", "target MAC address (Required)")
	ifaceName := flag.String("iface", "*", "interface")
	flag.Parse()

	if *targetMAC == "" {
		// target is required
		flag.PrintDefaults()
		os.Exit(1)
	}

	mac, err := net.ParseMAC(*targetMAC)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ifaces := []net.Interface{}

	if *ifaceName != "*" {
		iface, err := net.InterfaceByName(*ifaceName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else if iface == nil {
			fmt.Println("no such network interface")
			os.Exit(1)
		}
		ifaces = append(ifaces, *iface)
	} else {
		ifaces, err = net.Interfaces()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// for each interface given, send a wol packet to the target

	for _, iface := range ifaces {
		log.Printf("Sending Wake-on-Lan packet to %v on interface %v.\n", mac, iface.Name)
		wol.WakeOnLan(mac, &iface)
	}

}
