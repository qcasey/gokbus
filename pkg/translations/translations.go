package translations

// PacketMessageMeaning will define a generic meaning we can assign to Data packets
type PacketMessageMeaning uint

const (
	CarUnlocked PacketMessageMeaning = iota
	CarLocked
	PassengerDoorLocked
	DriverDoorLocked
	WindowDoorMessage
	TopClosed
	TopOpen

	KeyDetected
	KeyIn
	KeyOut
	KeyButtonReleased
	KeyNotDetected

	RecirculatingAirPressed
	RecirculatingAirReleased
	AuxHeatingOff

	RadioVolumeUp
	RadioVolumeDown
	RadioCDStatus
	RadioCDNext
	RadioCDPrev
	Radio1
	Radio2
	Radio3
	Radio4
	Radio5
	Radio6
	RadioModePressed
	RadioModeReleased
	RadioNextPlaylist
	RadioPrevPlaylist
	RadioSeekNextPressed
	RadioSeekNextReleased
	RadioSeekPrevPressed
	RadioSeekPrevReleased
	RadioFMPressed
	RadioFMReleased
	RadioAMPressed
	RadioAMReleased
	RadioMANPressed
	RadioMANReleased
	RadioSCRPPressed
	RadioSCRPReleased
	RadioReady
	RadioPowerToggled

	IkeStatus

	RainLightSensorStatus
	AllLightsTurnedOff
	AllLightsTurnedOn
	AllLightsOK

	SeatMemoryAny
	SeatMemory1
	SeatMemory2
	SeatMemory3

	SteeringWheelNextPressed
	SteeringWheelNextPressedLong
	SteeringWheelNextReleased

	SteeringWheelPreviousPressed
	SteeringWheelPreviousPresssedLong
	SteeringWheelPreviousReleased

	SteeringWheelSpeakPressed
	SteeringWheelSpeakPressedLong
	SteeringWheelSpeakReleased

	SteeringWheelRTPressed
	SteeringWheelDialReleased

	ClimateControl

	IgnitionOff
	Diagnostic
	OdometerRequest
	VehicleStatus
)

// PacketTranslateMap maps a packet's SRC / DEST / DATA (flattened into string) into a generic meaning for easier parsing
var PacketTranslateMap = map[string]map[string]map[string]PacketMessageMeaning{
	"00": {
		"BF": {
			"0200B9": CarUnlocked,       // Unlocked via key
			"1100":   IgnitionOff,       // Engine off
			"7202":   KeyButtonReleased, // Key, no button pressed (released)
			"7212":   CarLocked,         // Locked via key
			"7222":   CarUnlocked,       // Unlocked via key
			"7A":     WindowDoorMessage, // Response to window / door status request
			"7D0000": TopClosed,         // Technically "sunroof locked", but responds when top closed
			"7D00":   TopOpen,           // Last packet is the state / progress of open
		},
	},

	"3F": {
		"00": {
			"0C3401": CarUnlocked, // All doors unlocked
			"0C4601": PassengerDoorLocked,
			"0C4701": DriverDoorLocked,
			"OTHER":  Diagnostic,
		},
	},
	"44": { // EWS Ignition / Immobilizer
		"80": {
			"16": OdometerRequest, // Odometer request
		},
		"BF": { // Global
			"7404":   KeyDetected,    // Last byte is key //
			"740500": KeyIn,          // Key in, to 2nd position
			"7401FF": KeyOut,         // No_key_detected - Triggred by pulling the key out
			"7400FF": KeyNotDetected, // No_key_detected - Triggered in response to a request
			//"7400" : "d_keyIn",
			//"7A" : "d_windowDoorMessage"
		},
	},
	"50": { // MF Steering Wheel Buttons
		"5B": {
			"3A01": RecirculatingAirPressed,  // AUC - recirculating air pressed
			"3A00": RecirculatingAirReleased, // AUC - recirculating air released
		},
		"68": { // RADIO
			"3210": RadioVolumeDown, // Volume Down
			"3211": RadioVolumeUp,   // Volume Up
			"3B01": SteeringWheelNextPressed,
			"3B11": SteeringWheelNextPressedLong, // Next, long press
			"3B21": SteeringWheelNextReleased,    // Next Released
			"3B08": SteeringWheelPreviousPressed,
			"3B18": SteeringWheelPreviousPresssedLong, // Prev, long press
			"3B28": SteeringWheelPreviousReleased,     // Prev Released
		},
		"C8": {
			"01":   SteeringWheelRTPressed, // This can happen via RT button or ignition
			"019A": SteeringWheelRTPressed, // RT button?
			// "3B40" :  // reset
			"3B80": SteeringWheelSpeakPressed,     // Dial button
			"3B90": SteeringWheelSpeakPressedLong, // Dial button, long press
			"3BA0": SteeringWheelDialReleased,     // Dial button, released
		},
		"FF": {
			"3B40": SteeringWheelRTPressed, // also RT Button?
			"3B00": SteeringWheelRTPressed,
		},
	},
	"5B": {
		"80": {
			"*": ClimateControl,
		},
	},
	"68": {
		"18": {
			"380000": RadioCDStatus, // d_cdSendStatus
			//"380100" : None,
			//"380300" : None,
			"380A00": RadioCDNext, // Next button on Radio
			"380A01": RadioCDPrev, // Prev button on Radio
			//"380700" : None,
			//"380701" : None,
			"380601": Radio1,            // 1 pressed
			"380602": Radio2,            // 2 pressed
			"380603": Radio3,            // 3 pressed
			"380604": Radio4,            // 4 pressed
			"380605": Radio5,            // 5 pressed
			"380606": Radio6,            // 6 pressed
			"380400": RadioPrevPlaylist, // prev Playlist function?
			"380401": RadioNextPlaylist, // next Playlist function?
			//"380800" : None,
			//"380801" : None
		},
	},
	"72": {
		"BF": {
			"780000": SeatMemoryAny, // One of the seat memory buttons were pushed
			"780100": SeatMemory1,   // Button 1 released
			"780200": SeatMemory2,   // Button 2 released
			"780400": SeatMemory3,   // Button 3 released
		},
	},
	"80": {
		"68": { // Radio buttons
			"31000007": RadioSeekNextPressed,  // Seek > pressed
			"31000047": RadioSeekNextReleased, // Seek > released
			"31000006": RadioSeekPrevPressed,  // Seek < pressed
			"31000046": RadioSeekPrevReleased, // Seek < released
			"31000009": RadioFMPressed,        // FM pressed
			"31000049": RadioFMReleased,       // FM released
			"3100000A": RadioAMPressed,        // AM pressed
			"3100004A": RadioAMReleased,       // AM released
			"3100000C": RadioMANPressed,       // MAN pressed
			"3100004C": RadioMANReleased,      // MAN released
			"3100000E": RadioSCRPPressed,      // SC/RP pressed
			"3100004E": RadioSCRPReleased,     // SC/RP released
		},
		"BF": {
			"0201":  RadioReady, // Device status ready after Reset
			"OTHER": IkeStatus,  // Use ALL to send all data to a particular function
		},
		"E7": {
			"2A0000": AuxHeatingOff, // NAVCODER - Not totally sure what this refers to ("Aux_Heating_LED = Off")
		},
	},
	"9C": { // Technically Sunroof module, operates w/ convertible tops
		"BF": {
			"7C0174": TopClosed,
		},
	},
	"C0": { // telephone module
		"80": { // IKE
			"234220": SteeringWheelRTPressed, // Telephone invoked (from steering wheel?)
		},
		"68": {
			"3100000B": RadioModePressed,  // Mode button pressed
			"3100134B": RadioModeReleased, // Mode button released
		},
	},
	"D0": {
		"80": { // Vehicle Data, in response to request
			"*": VehicleStatus,
		},
		"BF": {
			"5B6100040001": AllLightsTurnedOn,  // lights turned on (including interior)
			"5B0100000001": AllLightsTurnedOff, // lights turned off (including interior)
			"5B6000040000": AllLightsOK,        // Indicator_Left Indicator_Right Indicator_sync  All_OK
		},
	},
	"E8": {
		"D0": {
			"*": RainLightSensorStatus,
		},
	},
	"F0": { // Board Monitor Buttons
		"68": { // Radio
			"4806": RadioPowerToggled, // Radio power toggled
		},
	},
}
