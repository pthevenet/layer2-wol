// A layer 2 Wake on Lan Magic Packet
package main

import (
	"encoding/hex"
	"fmt"
	"github.com/mdlayher/ethernet"
	"github.com/mdlayher/raw"
	"log"
	"net"
)

// MagicPacket is a broadcast frame containing FF FF FF FF FF FF followed by 16 repetitions of the target MAC address
type MagicPacket struct {
	target net.HardwareAddr // MAC address of target
}

// Bytes return the payload of the MagicPacket in a byte slice form
func (mpkt MagicPacket) Bytes() []byte {
	buf := make([]byte, 0)
	head, err := hex.DecodeString("ffffffffffff")
	if err != nil {
		log.Panic("cannot decode string into bytes: %v", err)
	}
	buf = append(buf, head...)

	for i := 0; i < 16; i++ {
		buf = append(buf, []byte(mpkt.target)...)
	}
	return buf
}

// SendMagicPacket sends a MagicPacket to the destination
func SendMagicPacket(mpkt MagicPacket, iface *net.Interface) error {
	conn, err := raw.ListenPacket(iface, 0x0842, nil)
	if err != nil {
		return fmt.Errorf("could not open ethernet socket: %v", err)
	}
	defer conn.Close()

	// frame := &ethernet.Frame{
	// 	Destination: ethernet.Broadcast,
	// 	Source:      net.HardwareAddr{0xde, 0xad, 0xbe, 0xef, 0xde, 0xad},
	// 	EtherType:   0x0842,
	// 	Payload:     mpkt.Bytes(),
	// }

	addr := &raw.Addr{HardwareAddr: ethernet.Broadcast}
	_, err = conn.WriteTo(mpkt.Bytes(), addr)

  if err != nil {
    return fmt.Errorf("could not write to ethernet socket: %v", err)
  }

  return nil
}
