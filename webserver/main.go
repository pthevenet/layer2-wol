// on-off sends wake-on-lan packets to a mac address on a LAN
// it also provides alive status of this computer
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/pthevenet/on-off/wol"
)

func main() {
	var (
		macS       = flag.String("mac", "", "MAC address of the target computer")
		ipS        = flag.String("ip", "", "IP address of the target computer")
		ifacename  = flag.String("interface", "lo", "interface name")
		port       = flag.Uint("port", 8080, "port for the webserver")
		targetPort = flag.Uint("targetport", 8080, "port of the target 'off' service")
	)
	flag.Parse()

	mac, err := net.ParseMAC(*macS)
	if err != nil {
		log.Fatal("MAC address is invalid:", err)
	}

	ip := net.ParseIP(*ipS)
	if ip == nil {
		log.Fatal("IP address is invalid:", err)
	}

	iface, err := net.InterfaceByName(*ifacename)
	if err != nil {
		log.Fatal("Interface is invalid:", err)
	}

	http.HandleFunc("/on", func(w http.ResponseWriter, r *http.Request) {
		onHandler(w, r, mac, iface)
	})
	http.HandleFunc("/off", func(w http.ResponseWriter, r *http.Request) {
		offHandler(w, r, ip, *targetPort)
	})
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	http.ListenAndServe(":"+fmt.Sprint(*port), nil)
}

func onHandler(w http.ResponseWriter, r *http.Request, mac net.HardwareAddr, iface *net.Interface) {
	log.Println("ON REQUEST target", mac.String(), "interface", iface.Name)
	err := wol.WakeOnLan(mac, iface)
	if err != nil {
		log.Println("Error:", err)
		fmt.Fprint(w, "ON request executed with error: %v.\n", err)
		return
	}
	fmt.Fprint(w, "ON request executed with no error.\n")
}

func offHandler(w http.ResponseWriter, r *http.Request, ip net.IP, port uint) {
	log.Printf("OFF REQUEST target %v:%v", ip, port)
	url := fmt.Sprintf("http://%s:%d/off", ip, port)
	_, err := http.Get(url)
	if err != nil {
		log.Println("Error:", err)
		fmt.Fprint(w, "OFF request executed with error: %v.\n", err)
		return
	}
	fmt.Fprint(w, "OFF request executed with no error.\n")
}
