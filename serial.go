package gokbus

import (
	"fmt"
	"time"
)

// Waits for the bus to clear, then returns the Source of the next packet
func (kbus *KBUS) getSourceByte() (byte, error) {
	var lastByteRecievedTime = time.Now()
	for {
		srcPacket := kbus.readBytes()
		if len(srcPacket) == 0 {
			return 0, fmt.Errorf("Bus is empty")
		}
		if time.Since(lastByteRecievedTime) > time.Millisecond*10 {
			return srcPacket[0], nil
		}

		lastByteRecievedTime = time.Now()
	}
}

// Gets the remaining packets in this data stream
func (kbus *KBUS) getDataBytes(len int) []byte {
	var returnBytes []byte
	for i := 0; i < len; i++ {
		returnBytes = append(returnBytes, kbus.readBytes()[0])
	}
	return returnBytes
}

func (kbus *KBUS) readBytes() []byte {
	buf := make([]byte, 32)
	var (
		n   int
		err error
	)

	timeout := time.Now()
	// Read until there is activity on the bus, or 500ms have passed
	for {
		n, err = kbus.Port.Read(buf)
		if time.Since(timeout) > time.Millisecond*500 {
			return []byte{}
		}
		if err == nil {
			break
		}
	}

	return buf[:n]
}
