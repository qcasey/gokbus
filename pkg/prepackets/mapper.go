package prepackets

import "github.com/qcasey/gokbus"

// RequestToPacket maps a string to a slice of relevant packets
func RequestToPacket(request string) []gokbus.Packet {
	switch request {
	case "PressMode":
		return PressMode
	case "PressNum1":
		return PressNum1
	case "PressNum2":
		return PressNum2
	case "PressNum3":
		return PressNum3
	case "PressNum4":
		return PressNum4
	case "PressNum5":
		return PressNum5
	case "PressNum6":
		return PressNum6
	case "PressAM":
		return PressAM
	case "PressFM":
		return PressFM
	case "PressNext":
		return PressNext
	case "PressPrev":
		return PressPrev
	case "PressStereoPower":
		return PressStereoPower
	case "PressRecirculatingAir":
		return PressRecirculatingAir
	case "OpenTrunk":
		return []gokbus.Packet{OpenTrunk}
	case "TurnOnClownNose":
		return []gokbus.Packet{TurnOnClownNose}
	case "FlashLowBeams":
		return []gokbus.Packet{FlashLowBeams}
	case "FlashLowBeamsAndHazards":
		return []gokbus.Packet{FlashLowBeamsAndHazards}
	case "FlashHighBeams":
		return []gokbus.Packet{FlashHighBeams}
	case "FlashHighBeamsAndHazards":
		return []gokbus.Packet{FlashHighBeamsAndHazards}
	case "FlashHighBeamsAndLowBeams":
		return []gokbus.Packet{FlashHighBeamsAndLowBeams}
	case "FlashHazards":
		return []gokbus.Packet{FlashHazards}
	case "TurnOffAllExteriorLights":
		return []gokbus.Packet{TurnOffAllExteriorLights}
	case "ToggleInteriorLights":
		return []gokbus.Packet{ToggleInteriorLights}
	case "ToggleDoorLocks":
		return []gokbus.Packet{ToggleDoorLocks}
	case "LockDriverDoor":
		return []gokbus.Packet{LockDriverDoor}
	case "LockPassengerDoor":
		return []gokbus.Packet{LockPassengerDoor}
	case "RequestDoorStatus":
		return []gokbus.Packet{RequestDoorStatus}
	case "RequestIgnitionStatus":
		return []gokbus.Packet{RequestIgnitionStatus}
	case "RequestLightStatus":
		return []gokbus.Packet{RequestLightStatus}
	case "RequestTimeStatus":
		return []gokbus.Packet{RequestTimeStatus}
	case "RequestOdometer":
		return []gokbus.Packet{RequestOdometer}
	case "RequestPDCStatus":
		return []gokbus.Packet{RequestPDCStatus}
	case "RequestVehicleStatus":
		return []gokbus.Packet{RequestVehicleStatus}
	case "RequestTemperatureStatus":
		return []gokbus.Packet{RequestTemperatureStatus}
	case "PopWindowsUp":
		return PopWindowsUp
	case "PopWindowsDown":
		return PopWindowsDown
	}
	return nil
}
