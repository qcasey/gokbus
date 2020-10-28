// Package prepackets are a collection of KBus packets prepared for a 2005 BMW Vert, they may or may not work for your model car
package prepackets

import "github.com/qcasey/gokbus"

var PressMode = []gokbus.Packet{
	// PUSH
	{
		Source:      0xF0,
		Destination: 0x68,
		Data:        []byte{0x48, 0x23},
	},
	// RELEASE
	{
		Source:      0xF0,
		Destination: 0x68,
		Data:        []byte{0x48, 0xA3},
	},
}

var PressNum1 = []gokbus.Packet{
	// PUSH
	{
		Source:      0xF0,
		Destination: 0x68,
		Data:        []byte{0x48, 0x11},
	},
	// RELEASE
	{
		Source:      0xF0,
		Destination: 0x68,
		Data:        []byte{0x48, 0x91},
	},
}

var PressNum2 = []gokbus.Packet{
	// PUSH
	{
		Source:      0xF0,
		Destination: 0x68,
		Data:        []byte{0x48, 0x01},
	},
	// RELEASE
	{
		Source:      0xF0,
		Destination: 0x68,
		Data:        []byte{0x48, 0x81},
	},
}

var PressNum3 = []gokbus.Packet{
	// PUSH
	{
		Source:      0xF0,
		Destination: 0x68,
		Data:        []byte{0x48, 0x12},
	},
	// RELEASE
	{
		Source:      0xF0,
		Destination: 0x68,
		Data:        []byte{0x48, 0x92},
	},
}

var PressNum4 = []gokbus.Packet{
	// PUSH
	{
		Source:      0xF0,
		Destination: 0x68,
		Data:        []byte{0x48, 0x02},
	},
	// RELEASE
	{
		Source:      0xF0,
		Destination: 0x68,
		Data:        []byte{0x48, 0x82},
	},
}

var PressNum5 = []gokbus.Packet{
	// PUSH
	{
		Source:      0xF0,
		Destination: 0x68,
		Data:        []byte{0x48, 0x13},
	},
	// RELEASE
	{
		Source:      0xF0,
		Destination: 0x68,
		Data:        []byte{0x48, 0x93},
	},
}

var PressNum6 = []gokbus.Packet{
	// PUSH
	{
		Source:      0xF0,
		Destination: 0x68,
		Data:        []byte{0x48, 0x03},
	},
	// RELEASE
	{
		Source:      0xF0,
		Destination: 0x68,
		Data:        []byte{0x48, 0x83},
	},
}

var PressAM = []gokbus.Packet{
	// PUSH
	{
		Source:      0xF0,
		Destination: 0x68,
		Data:        []byte{0x48, 0x31},
	},
	// RELEASE
	{
		Source:      0xF0,
		Destination: 0x68,
		Data:        []byte{0x48, 0xA1},
	},
}

var PressFM = []gokbus.Packet{
	// PUSH
	{
		Source:      0xF0,
		Destination: 0x68,
		Data:        []byte{0x48, 0x31},
	},
	// RELEASE
	{
		Source:      0xF0,
		Destination: 0x68,
		Data:        []byte{0x48, 0xB1},
	},
}

var PressNext = []gokbus.Packet{
	// PUSH
	{
		Source:      0xF0,
		Destination: 0x68,
		Data:        []byte{0x48, 0x00},
	},
	// RELEASE
	{
		Source:      0xF0,
		Destination: 0x68,
		Data:        []byte{0x48, 0x80},
	},
}

var PressPrev = []gokbus.Packet{
	// PUSH
	{
		Source:      0xF0,
		Destination: 0x68,
		Data:        []byte{0x48, 0x10},
	},
	// RELEASE
	{
		Source:      0xF0,
		Destination: 0x68,
		Data:        []byte{0x48, 0x90},
	},
}

var PressStereoPower = []gokbus.Packet{
	// PUSH
	{
		Source:      0xF0,
		Destination: 0x68,
		Data:        []byte{0x48, 0x06},
	},
	// RELEASE
	{
		Source:      0xF0,
		Destination: 0x68,
		Data:        []byte{0x48, 0x86},
	},
}

var PressRecirculatingAir = []gokbus.Packet{
	// PUSH
	{
		Source:      0xF0,
		Destination: 0x68,
		Data:        []byte{0x48, 0x01},
	},
	// RELEASE
	{
		Source:      0xF0,
		Destination: 0x68,
		Data:        []byte{0x48, 0x00},
	},
}

var OpenTrunk = gokbus.Packet{
	Source:      0x3F,
	Destination: 0x00,
	Data:        []byte{0x0C, 0x02, 0x01},
}

// TurnOnClownNose for 3 seconds
var TurnOnClownNose = gokbus.Packet{
	Source:      0x3F,
	Destination: 0x00,
	Data:        []byte{0x0C, 0x4E, 0x01},
}

var FlashLowBeams = gokbus.Packet{
	Source:      0x00,
	Destination: 0xBF,
	Data:        []byte{0x76, 0x04},
}

var FlashLowBeamsAndHazards = gokbus.Packet{
	Source:      0x00,
	Destination: 0xBF,
	Data:        []byte{0x76, 0x06},
}

var FlashHighBeams = gokbus.Packet{
	Source:      0x00,
	Destination: 0xBF,
	Data:        []byte{0x76, 0x08},
}

var FlashHighBeamsAndHazards = gokbus.Packet{
	Source:      0x00,
	Destination: 0xBF,
	Data:        []byte{0x76, 0x0A},
}

var FlashHighBeamsAndLowBeams = gokbus.Packet{
	Source:      0x00,
	Destination: 0xBF,
	Data:        []byte{0x76, 0x0C},
}

var FlashHazards = gokbus.Packet{
	Source:      0x00,
	Destination: 0xBF,
	Data:        []byte{0x76, 0x02},
}

var TurnOffAllExteriorLights = gokbus.Packet{
	Source:      0x00,
	Destination: 0xBF,
	Data:        []byte{0x76, 0x00},
}

var ToggleInteriorLights = gokbus.Packet{
	Source:      0x3F,
	Destination: 0x00,
	Data:        []byte{0x0C, 0x01, 0x01},
}

var ToggleDoorLocks = gokbus.Packet{
	Source:      0x3F,
	Destination: 0x00,
	Data:        []byte{0x0C, 0x34, 0x01},
}

var LockDriverDoor = gokbus.Packet{
	Source:      0x3F,
	Destination: 0x00,
	Data:        []byte{0x0C, 0x47, 0x01},
}

var LockPassengerDoor = gokbus.Packet{
	Source:      0x3F,
	Destination: 0x00,
	Data:        []byte{0x0C, 0x46, 0x01},
}

var RequestDoorStatus = gokbus.Packet{
	Source:      0x9C,
	Destination: 0x00,
	Data:        []byte{0x79},
}

var RequestIgnitionStatus = gokbus.Packet{
	Source:      0x00,
	Destination: 0x80,
	Data:        []byte{0x10},
}

var RequestLightStatus = gokbus.Packet{
	Source:      0x00,
	Destination: 0xD0,
	Data:        []byte{0x5A},
}

var RequestTimeStatus = gokbus.Packet{
	Source:      0xA4,
	Destination: 0x80,
	Data:        []byte{0x41, 0x01, 0x01},
}

var RequestOdometer = gokbus.Packet{
	Source:      0x44,
	Destination: 0x03,
	Data:        []byte{0x80, 0x16, 0xD1},
}

var RequestPDCStatus = gokbus.Packet{
	Source:      0x3F,
	Destination: 0x60,
	Data:        []byte{0x1B},
}

var RequestVehicleStatus = gokbus.Packet{
	Source:      0x80,
	Destination: 0xD0,
	Data:        []byte{0x53},
}

var RequestTemperatureStatus = gokbus.Packet{
	Source:      0xD0,
	Destination: 0x80,
	Data:        []byte{0x1D},
}

// PopWindowsUp about 40%
var PopWindowsUp = []gokbus.Packet{
	// Window 1
	{
		Source:      0x3F,
		Destination: 0x00,
		Data:        []byte{0x0C, 0x53, 0x01},
	},
	// Window 2
	{
		Source:      0x3F,
		Destination: 0x00,
		Data:        []byte{0x0C, 0x55, 0x01},
	},
	// Window 3
	{
		Source:      0x3F,
		Destination: 0x00,
		Data:        []byte{0x0C, 0x42, 0x01},
	},
	// Window 4
	{
		Source:      0x3F,
		Destination: 0x00,
		Data:        []byte{0x0C, 0x43, 0x01},
	},
}

// PopWindowsDown about 40%
var PopWindowsDown = []gokbus.Packet{
	// Window 1
	{
		Source:      0x3F,
		Destination: 0x00,
		Data:        []byte{0x0C, 0x52, 0x01},
	},
	// Window 2
	{
		Source:      0x3F,
		Destination: 0x00,
		Data:        []byte{0x0C, 0x54, 0x01},
	},
	// Window 3
	{
		Source:      0x3F,
		Destination: 0x00,
		Data:        []byte{0x0C, 0x41, 0x01},
	},
	// Window 4
	{
		Source:      0x3F,
		Destination: 0x00,
		Data:        []byte{0x0C, 0x44, 0x01},
	},
}
