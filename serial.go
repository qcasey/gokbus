package gokbus

import (
	"time"
)

// Waits for the bus to clear, then returns the Source of the next packet
func (kbus *KBUS) getSourceByte() ([]byte, error) {
	var lastByteRecievedTime = time.Now()
	for {
		srcPacket, err := kbus.readBytes()
		if err != nil {
			return nil, err
		}
		if srcPacket == nil {
			return nil, nil
		}
		if time.Since(lastByteRecievedTime) > time.Millisecond*10 {
			return srcPacket, nil
		}

		lastByteRecievedTime = time.Now()
	}
}

// Gets the remaining packets in this data stream
func (kbus *KBUS) getDataBytes(length int) ([]byte, error) {
	var returnBytes []byte
	for i := 0; i < length; i++ {
		newByte, err := kbus.readBytes()
		if err != nil {
			return nil, err
		}
		if newByte == nil {
			return nil, nil
		}
		returnBytes = append(returnBytes, newByte[0])
	}
	return returnBytes, nil
}

func (kbus *KBUS) readBytes() ([]byte, error) {
	buf := make([]byte, 32)

	// Read until there is activity on the bus, or 200ms have passed
	n, err := kbus.Port.Read(buf)
	if n == 0 {
		return nil, nil
	}
	return buf[:n], err
}
