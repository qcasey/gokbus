package gokbus

import (
	"fmt"

	"github.com/qcasey/gokbus/pkg/ecu"
)

// Packet is a group of bytes read from the bus
type Packet struct {
	Source      byte
	length      byte
	Destination byte
	Data        []byte
	checksum    byte
}

func (p *Packet) getChecksum() byte {
	var derivedChecksum byte
	derivedChecksum = derivedChecksum ^ p.Source
	derivedChecksum = derivedChecksum ^ p.length
	derivedChecksum = derivedChecksum ^ p.Destination

	for _, b := range p.Data {
		derivedChecksum = derivedChecksum ^ b
	}
	return derivedChecksum
}

func (p *Packet) addChecksum() *Packet {
	p.checksum = p.getChecksum()
	return p
}

func (p *Packet) addLength() *Packet {
	p.length = byte(len(p.Data) + 2)
	return p
}

// Flatten returns the packet represented as a byte array
func (p *Packet) Flatten() []byte {
	var flatPacket []byte
	flatPacket = append(flatPacket, p.Source)
	flatPacket = append(flatPacket, p.length)
	flatPacket = append(flatPacket, p.Destination)
	for _, b := range p.Data {
		flatPacket = append(flatPacket, b)
	}
	flatPacket = append(flatPacket, p.checksum)
	return flatPacket
}

// IsValid checks a recieved Packet's checksum to ensure it was transmitted correctly
// This will fail on our own generated packets, since those haven't been given a Checksum or Length yet
func (p *Packet) IsValid() bool {
	return p.getChecksum() == p.checksum
}

// Pretty returns a well formatted string of the packet, with ECU locations resolved. Format: SOURCE -> DEST (LEN) DATA [CHECKSUM]
func (p *Packet) Pretty() string {
	if p.length == 0 {
		p.addLength().addChecksum()
	}
	return fmt.Sprintf("%s (0x%02X) -> %s (0x%02X) %02X (0x%02X bytes) [0x%02X]", ecu.ByAddress[p.Source], p.Source, ecu.ByAddress[p.Destination], p.Destination, p.Data, p.length, p.checksum)
}
