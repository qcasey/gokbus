package gokbus

func (kbus *KBUS) addError(err error) {
	select {
	case kbus.ErrorChannel <- err:
		return
	default:
		return
	}
}
