// on-off sends wake-on-lan packets to a mac address on a LAN
// it also provides alive status of this computer
package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net"
)

func main() {
	var (
		macS  = flag.String("mac", "", "MAC address of the target computer")
		port = flag.Uint("port", 8080, "port for the webserver")
		ifacename = flag.String("interface", "lo", "interface name")
	)
	flag.Parse()

	mac, err := net.ParseMAC(*macS)
	if err != nil {
		log.Fatal("MAC address is invalid:", err)
	}

	iface, err :=  net.InterfaceByName(*ifacename)
	if err != nil {
		log.Fatal("Interface is invalid:", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/on", func(w http.ResponseWriter, r *http.Request) {
		onHandler(w, r, mac, iface)
	})
	r.HandleFunc("/off", offHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	http.Handle("/", r)

	http.ListenAndServe(":"+fmt.Sprint(*port), r)
}

func onHandler(w http.ResponseWriter, r *http.Request, mac net.HardwareAddr, iface *net.Interface) {
	log.Println("ON REQUEST target", mac.String(), "interface", iface.Name)

	// create magic packet
	mpkt := MagicPacket{mac}
	err := SendMagicPacket(mpkt, iface)
	if err != nil {
		log.Println("Error:", err)
	}
}

func offHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("OFF")
	// TODO
}
