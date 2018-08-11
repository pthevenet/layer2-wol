# layer2-wol
[![Go Report Card](https://goreportcard.com/badge/github.com/pthevenet/layer2-wol)](https://goreportcard.com/report/github.com/pthevenet/layer2-wol)


layer2-wol is a CLI tool and a library to send and create layer-2 (Ethernet) Wake-on-Lan "magic" packets.  
This is used to start a (Wake-on-Lan activated) computer on the same Local Area Network.  
[Wake-on-Lan](https://en.wikipedia.org/wiki/Wake-on-LAN) is a Data Link Layer standard.

## Usage



### CLI Tool

to build the tool binary:

```shell
go build
```

```shell
./layer2-wol -h
Usage of ./layer2-wol:
  -iface string
    	interface (default "*")
  -mac string
    	target MAC address (Required)
```
#### Example

> send a wol packet to aa:aa:aa:aa:aa:aa on interface eth0

```shell

./layer2-wol -iface eth0 -mac "aa:aa:aa:aa:aa:aa"
```

### GO Library

```go
package main

import (
	"fmt"
	"log"
	"net"

	"github.com/pthevenet/layer2-wol/wol"
)

func main() {
	target_mac := "aa:aa:aa:aa:aa:aa"
	target_addr, err := net.ParseMAC(target_mac)
	if err != nil {
		log.Fatalf("could not parse mac %s: %v", target_mac, err)
	}

	magic_pkt := wol.NewMagicPacket(target_addr)

	fmt.Println(magic_pkt)
}
```

## Notes

[go-wol](https://github.com/sabhiram/go-wol) uses UDP sockets (layer-4) for sending WOL packets, however we recommend using it over this library.  
