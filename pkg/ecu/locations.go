package ecu

// ByAddress is a mapping of hex codes seen in SRC/DST parts of packets.
var ByAddress = map[byte]string{
	0x00: "Broadcast",
	0x02: "Electronic Body Module",
	0x08: "Sunroof Control",
	0x12: "Digital Motor Electronics (DME)",
	0x18: "CDW - CDC CD-Player",
	0x24: "Boot Lid Control Unit",
	0x28: "Radio Controlled Clock",
	0x30: "Check Control Module",
	0x32: "Electronic gearbox control",
	0x3B: "NAV Navigation/Videomodule",
	0x3F: "Diagnostic",
	0x40: "Remote Control Central Locking (Key Fob)",
	0x43: "Menu Screen, Navigation - rear",
	0x44: "EWS Ignition, Immobiliser",
	0x46: "Central Information Display",
	0x47: "Rear Monitor Controls",
	0x50: "MFL Multi Functional Steering Wheel Buttons",
	0x51: "Mirror Memory, Left",
	0x52: "Convertible Folding Top Module",
	0x56: "Anti-Lock Braking System With ASC",
	0x57: "Steering Angle Sensor",
	0x5B: "Integrated Heating And Air Conditioning",
	0x60: "PDC Park Distance Control",
	0x65: "Electronic Fuel Pump",
	0x66: "Adaptive Headlight Unit",
	0x68: "RAD Radio",
	0x6A: "DSP Digital Sound Processor / Amplifier",
	0x6B: "Standing Heat",
	0x70: "Tire Pressure Control",
	0x72: "Seat Memory",
	0x73: "Sirius Radio",
	0x74: "Seat Occupancy Recognition Unit",
	0x76: "CD Changer DIN Size",
	0x7F: "Navigation Europe",
	0x80: "IKE Instrument Control Electronics",
	0x81: "Revolution Counter/Steering Column",
	0x9A: "Headlight Aim Control",
	0x9B: "Mirror Memory, Right",
	0x9C: "Convertible Soft Top",
	0xA0: "Rear Multi Info Display",
	0xA4: "Air Bag Module",
	0xA6: "Cruise Control",
	0xA7: "Auto Climate Control - Rear",
	0xA8: "Navigation China",
	0xAC: "Electronic height control",
	0xB0: "Speed Recognition System",
	0xB8: "DME (K2000 protocol)",
	0xBB: "Navigation Japan",
	0xBF: "Global Broadcast Address",
	0xC0: "MID Multi-Information Display Buttons",
	0xC8: "TEL Telephone",
	0xCA: "BMW Assist / Telematics Control Unit",
	0xD0: "Light Control Module",
	0xDA: "Seat Memory Second",
	0xE0: "Integrated Radio Information System",
	0xE7: "OBC TextBar, Front Display",
	0xE8: "Rain Light Sensor",
	0xEA: "Digital Sound Processor Controller",
	0xED: "Lights, Wipers, Seat Memory, Television",
	0xF0: "BMB Board Monitor Buttons, On Board Monitor Operating Part",
	0xF1: "Programmable Controller (Custom Unit)",
	0xF5: "Centre Switch Control Unit",
	0xFF: "Broadcast",
}

const BroadcastZero = 0x00
const ElectronicBody = 0x02
const SunroofControl = 0x08
const DigitalMotorElectronics = 0x12
const CDPlayer = 0x18
const BootLidControlUnit = 0x24
const RadioControlledClock = 0x28
const CheckControl = 0x30
const ElectronicGearboxControl = 0x32
const Navigation = 0x3B
const Diagnostic = 0x3F
const RemoteControlCentralLockingKeyFob = 0x40
const MenuScreenNavigationRear = 0x43
const IgnitionImmobiliser = 0x44
const CentralInformationDisplay = 0x46
const RearMonitorControls = 0x47
const MultiFunctionalSteeringWheelButtons = 0x50
const MirrorMemoryLeft = 0x51
const ConvertibleFoldingTop = 0x52
const AntiLockBrakingSystem = 0x56
const SteeringAngleSensor = 0x57
const HeatingAndAirConditioning = 0x5B
const ParkDistanceControl = 0x60
const ElectronicFuelPump = 0x65
const AdaptiveHeadlights = 0x66
const Radio = 0x68
const DigitalSoundProcessorAmplifier = 0x6A
const StandingHeat = 0x6B
const TirePressureControl = 0x70
const SeatMemory = 0x72
const SiriusRadio = 0x73
const SeatOccupancyRecognition = 0x74
const CDChanger = 0x76
const NavigationEurope = 0x7F
const InstrumentControlElectronics = 0x80
const RevolutionCounter = 0x81
const HeadlightAimControl = 0x9A
const MirrorMemoryRight = 0x9B
const ConvertibleSoftTop = 0x9C
const RearMultiInfoDisplay = 0xA0
const AirBag = 0xA4
const CruiseControl = 0xA6
const AutoClimateControl = 0xA7
const NavigationChina = 0xA8
const ElectronicHeightControl = 0xAC
const SpeedRecognitionSystem = 0xB0
const DigitalMotorElectronicsK2000 = 0xB8
const NavigationJapan = 0xBB
const GlobalBroadcastAddress = 0xBF
const MultiInformationDisplayButtons = 0xC0
const Telephone = 0xC8
const BMWAssist = 0xCA
const LightControl = 0xD0
const SeatMemorySecond = 0xDA
const IntegratedRadioInformation = 0xE0
const FrontDisplay = 0xE7
const RainLightSensor = 0xE8
const DigitalSoundProcessorController = 0xEA
const LightsWipersSeatMemoryTelevisionController = 0xED
const BoardMonitorButtons = 0xF0
const ProgrammableController = 0xF1
const CentreSwitchControl = 0xF5
const Broadcast = 0xFF

// Aliases
const DME = DigitalMotorElectronics
const DSP = DigitalSoundProcessorAmplifier
const AC = HeatingAndAirConditioning
const PDC = ParkDistanceControl
const ABS = AntiLockBrakingSystem
const EWS = IgnitionImmobiliser
const DIAG = Diagnostic
const EGC = ElectronicGearboxControl
const SAS = SteeringAngleSensor
const TRUNK = BootLidControlUnit
const ICE = InstrumentControlElectronics
const VERT = ConvertibleSoftTop
