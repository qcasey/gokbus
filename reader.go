package gokbus

func (kbus *KBUS) startReader() {
	for {
		newPacket, err := kbus.ReadPacket()
		if err != nil {
			kbus.addError(err)
		}

		select {
		case kbus.ReadChannel <- newPacket:
			continue
		default:
			continue
		}
	}
}
