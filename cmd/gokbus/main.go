package main

import (
	"log"
	"os"

	"github.com/qcasey/gokbus"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Provide the device path to the kbus interface")
	}
	kbus, err := gokbus.New(os.Args[1], 9600)
	if err != nil {
		panic(err)
	}
	defer kbus.Port.Close()

	// Start the read and write channels
	go kbus.Start()

	// Read packets
	for {
		select {
		case newPacket := <-kbus.ReadChannel:
			log.Print(newPacket.Pretty())
		}
	}
}
