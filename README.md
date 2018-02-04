# layer2-wol
[![Go Report Card](https://goreportcard.com/badge/github.com/pthevenet/layer2-wol)](https://goreportcard.com/report/github.com/pthevenet/layer2-wol)

Library for creating and sending Wake-on-Lan layer-2 (ethernet for now) magic packets.  
This is used to start a (Wake-on-Lan activated) computer on the same Local Area Network.  
[Wake-on-Lan](https://en.wikipedia.org/wiki/Wake-on-LAN) is a Data Link Layer standard.

[go-wol](https://github.com/sabhiram/go-wol) uses UDP sockets (layer-4) for sending WOL packets, however we recommend using it over our library.  
