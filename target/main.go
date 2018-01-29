// Target server, listening for 'off' http requests and shutdowns for trusted ip source
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

// main opens an http server on port, listening for /off http request and shutdowns if the request is legitimate
// a request is legitimate if it has a source IP inside given subnet
// anyone in the given network will have power to shutdown the machine
// IP spoofing will work, therefore do not use this outside the LAN and use only local IP CIDRs
func main() {
	port := flag.Int("port", 8080, "port for the off service")
	trustedNetwork := flag.String("net", "192.168.0.0/24", "subnet CIDR for accepting requests")
	flag.Parse()

	_, trustedNet, err := net.ParseCIDR(*trustedNetwork)
	if err != nil {
		log.Fatal("error net must be a cidr network:", err)
	}

	startHttpServer(*port, *trustedNet)
}

func startHttpServer(port int, trustedNet net.IPNet) {
	srv := &http.Server{Addr: ":" + strconv.Itoa(port)}

	http.HandleFunc("/off", func(w http.ResponseWriter, r *http.Request) {
		ipString := strings.Split(r.RemoteAddr, ":")[0]

		ip := net.ParseIP(ipString)
		if ip == nil {
			http.Error(w, "request error", http.StatusBadRequest)
			return
		}

		log.Println("OFF Request by", ip)

		if !trustedNet.Contains(ip) {
			log.Printf("ILLIGITIMATE OFF REQUEST from: %v.", ip.String())
			http.Error(w, "not authorized : "+ip.String(), http.StatusUnauthorized)
			return
		}

		log.Println("LEGITIMATE")
		fmt.Fprint(w, "LEGITIMATE OFF request received from %v.\n", ip.String())
		fmt.Fprint(w, "system will shutdown")

		// trusted request

		err := off()

		if err != nil {
			log.Printf("shutdown command finished with error: %v", err)
		} else {
			log.Printf("shutdown command finished with no error")
			// shutdown
			if err := srv.Shutdown(nil); err != nil {
				panic(err)
			}
		}
	})

	err := srv.ListenAndServe()
	if err != nil {
		log.Printf("http server error: %v", err)
	}
}

// off executes a "shutdown -h +1" command
func off() error {
	cmd := exec.Command("shutdown", "-h", "+1")
	return cmd.Run()
}
