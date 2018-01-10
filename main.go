// on-off sends wake-on-lan packets to a mac address on a LAN
// it also provides alive status of this computer
package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	var (
		mac  = flag.String("mac", "", "MAC address of the target computer")
		port = flag.Uint("port", 8080, "port for the webserver")
	)
	flag.Parse()

	if *mac == "" {
		log.Fatal("MAC address must be given")
	}

	r := mux.NewRouter()
	r.HandleFunc("/on", onHandler)
	r.HandleFunc("/off", offHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))
	http.Handle("/", r)

	http.ListenAndServe(":"+fmt.Sprint(*port), r)
}

func onHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ON")
	// TODO
}

func offHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("OFF")
	// TODO
}
