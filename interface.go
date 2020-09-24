package gokbus

import (
	"time"

	"github.com/tarm/serial"
)

// KBUS is the main interface to communicate with the BMW I/K-Bus
type KBUS struct {
	Path string
	Baud int

	ReadChannel  chan Packet
	WriteChannel chan Packet
	Port         *serial.Port
}

// New returns a new KBUS device interface
func New(devicePath string, baudrate int) (*KBUS, error) {
	serialConfig := serial.Config{
		Name:        devicePath,
		Baud:        baudrate,
		Size:        serial.DefaultSize,
		StopBits:    serial.Stop1,
		Parity:      serial.ParityEven,
		ReadTimeout: time.Millisecond * 500,
	}

	s, err := serial.OpenPort(&serialConfig)
	if err != nil {
		return nil, err
	}

	newKbus := &KBUS{
		Path:         devicePath,
		Baud:         baudrate,
		Port:         s,
		ReadChannel:  make(chan Packet, 1),
		WriteChannel: make(chan Packet, 1),
	}

	return newKbus, nil
}

// Start will block, reading and writing packets to the KBus with the provided channels
func (kbus *KBUS) Start() {
	for {
		kbus.ReadChannel <- kbus.ReadPacket()

		// We know the bus is open if we just read a packet, so write is now viable
		select {
		case p := <-kbus.WriteChannel:
			p.addLength().addChecksum() // ensure metadata is valid
			kbus.WritePacket(p)
		default:
			continue
		}
	}
}

// Close the KBUS port
func (kbus *KBUS) Close() {
	kbus.Port.Close()
}

// ReadPacket will block until it returns the next Packet read from the KBus
func (kbus *KBUS) ReadPacket() Packet {
	p := Packet{}
	p.Source = kbus.getSourceByte()
	p.length = kbus.readBytes()[0]
	p.Destination = kbus.readBytes()[0]
	p.Data = kbus.getDataBytes(int(p.length) - 2)
	p.checksum = kbus.readBytes()[0]

	return p
}

// WritePacket will attempt to write a packet to the kbus
func (kbus *KBUS) WritePacket(p Packet) (int, error) {
	flatPacket := p.Flatten()

	n, err := kbus.Port.Write(flatPacket)
	if err != nil {
		return 0, err
	}

	err = kbus.Port.Flush()
	if err != nil {
		return 0, err
	}

	return n, nil
}
