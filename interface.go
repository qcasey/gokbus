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
	ErrorChannel chan error
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
		ErrorChannel: make(chan error, 1),
	}

	return newKbus, nil
}

// Start will begin reading and writing packets to the KBus with the provided channels
func (kbus *KBUS) Start() {
	go kbus.startReader()
	go kbus.startWriter()
}

// Close the KBUS port
func (kbus *KBUS) Close() {
	kbus.Port.Close()
}

// ReadPacket will block until it either recieves the next Packet or times out because of an empty bus
func (kbus *KBUS) ReadPacket() (Packet, error) {
	var err error
	p := Packet{}

	p.Source, err = kbus.getSourceByte()
	if err != nil {
		return p, err
	}

	length, err := kbus.readBytes()
	if err != nil {
		return p, err
	}
	p.length = length[0]

	destination, err := kbus.readBytes()
	if err != nil {
		return p, err
	}
	p.Destination = destination[0]

	p.Data = kbus.getDataBytes(int(p.length) - 2)

	checksum, err := kbus.readBytes()
	if err != nil {
		return p, err
	}
	p.checksum = checksum[0]

	return p, nil
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
