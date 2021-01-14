package gokbus

func (kbus *KBUS) startWriter() {
	for {
		select {
		case p := <-kbus.WriteChannel:
			p.addLength().addChecksum() // ensure metadata is valid
			kbus.WritePacket(p)
		default:
			continue
		}
	}
}
