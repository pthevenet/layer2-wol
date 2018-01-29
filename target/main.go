// Target server, listening for 'off' http requests and shutdowns for trusted ip source
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"
	"strings"
)

// main opens an http server on port, listening for /off http request and shutdowns if the request is legitimate
// a request is legitimate if it has a source IP inside given subnet
// anyone in the given network will have power to shutdown the machine
// IP spoofing will work, therefore do not use this outside the LAN and use only local IP CIDRs
func main() {
	port := flag.Uint("port", 8080, "port for the off service")
	trustedNetwork := flag.String("net", "192.168.0.0/24", "network CIDR for accepting requests")
	flag.Parse()

	_, trustedNet, err := net.ParseCIDR(*trustedNetwork)
	if err != nil {
		log.Fatal("error net must be a cidr network:", err)
	}

	http.HandleFunc("/off", func(w http.ResponseWriter, r *http.Request) {
		ipString := strings.Split(r.RemoteAddr, ":")[0]

		ip := net.ParseIP(ipString)
		if ip == nil {
			return
		}
		log.Println("OFF Request by", ip)
		if !trustedNet.Contains(ip) {
			log.Println("ILLIGITIMATE")
			fmt.Fprint(w, "ILLIGITIMATE OFF request received.")
			return
		}
		log.Println("LEGITIMATE")
		fmt.Fprint(w, "LEGITIMATE OFF request received.\n")
		fmt.Fprint(w, "system will shutdown")

		// trusted request

		err := off()

		if err != nil {
			log.Printf("shutdown command finished with error: %v", err)
		} else {
			log.Printf("shutdown command finished with no error")
		}
	})
	err = http.ListenAndServe(fmt.Sprintf(":%v", *port), nil)

	if err != nil {
		fmt.Println("error listening to port:", *port, ":", err)
	}
}

// off executes a "shutdown -h +1" command
func off() error {
	cmd := exec.Command("shutdown", "-h", "+1")
	return cmd.Run()
}
