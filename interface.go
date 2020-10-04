package gokbus

import (
	"fmt"
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
	fmt.Println("starting")
	for {
		newPacket, readyForWrite, err := kbus.ReadPacket()
		if err == nil {
			kbus.ReadChannel <- newPacket
		}

		if readyForWrite {
			select {
			case p := <-kbus.WriteChannel:
				p.addLength().addChecksum() // ensure metadata is valid
				kbus.WritePacket(p)
			default:
				continue
			}
		}
	}
}

// Close the KBUS port
func (kbus *KBUS) Close() {
	kbus.Port.Close()
}

// ReadPacket will block until it either recieves the next Packet or times out because of an empty bus
func (kbus *KBUS) ReadPacket() (Packet, bool, error) {
	var err error
	p := Packet{}

	p.Source, err = kbus.getSourceByte()
	if err != nil {
		return p, true, err
	}

	p.length = kbus.readBytes()[0]
	p.Destination = kbus.readBytes()[0]
	p.Data = kbus.getDataBytes(int(p.length) - 2)
	p.checksum = kbus.readBytes()[0]

	return p, true, nil
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
