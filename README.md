# Go KBUS

Go KBus is a golang module designed to interface with the BMW I/K Bus. This protocol is similar to CAN, and controls the non-essential components in the car.

This module can be used to read, write, and interpret BMW serial commands. It's loosely based on ezeakeal's excellent [pyBus](https://github.com/ezeakeal/pyBus) \(and my subsequent [fork](https://github.com/qcasey/pyBus)\).

## Overview
* USB serial interface can be acquired from [Reslers.de](http://www.reslers.de/IBUS/). Other USB interfaces should work just as well, although are untested. 
* Communicates on the I-BUS to send and receive diagnostics, running status, button presses, vehicle info, etc. 

## Install

```go get github.com/qcasey/gokbus```

## Example module usage

```golang
package main

import (
	"log"

	"github.com/qcasey/gokbus"
)

func main() {
	kbus, err := gokbus.New("/dev/ttyUSB0", 9600)
	if err != nil {
		panic(err)
	}
	defer kbus.Port.Close()

	// Start the read and write channels
	// This isn't required - we can ReadPacket() / WritePacket() separately without channels,
	// but timing will be less precise and many packets will be dropped.
	go kbus.Start()

	// Read 5 packets
	for i := 0; i < 5; i++ {
		select {
		case newPacket := <-kbus.ReadChannel:
			log.Print(newPacket.Pretty())
		}
	}

	// Prepare a write packet, this one opens the trunk
	openTrunkPacket := gokbus.Packet{
		Source:      0x3F,
		Destination: 0x00,
		Data:        []byte{0x0C, 0x02, 0x01},
	}
	
	// alternatively, a list of known packets are provided in /pkg/prepackets
	// flashLowBeamsPacket := prepackets.FlashLowBeams

	// Block until packet is written to the bus
	kbus.WriteChannel <- openTrunkPacket

	// Read another 5 packets
	for i := 0; i < 5; i++ {
		select {
		case newPacket := <-kbus.ReadChannel:
			log.Print(newPacket.Pretty())
		}
	}

}
```

Missing write confirmation that pyBus offers

### Useful links
http://web.archive.org/web/20041204074622/www.openbmw.org/bus/  
https://web.archive.org/web/20190210133920/http://web.comhem.se/bengt-olof.swing/ibusdevicesandoperations.htm   
https://web.archive.org/web/20180511152050/http://www.online-rubin.de/BMW/I-Bus/I-BUS%20Codes%20e46.htm
