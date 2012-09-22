package pkcs11

import (
	"unsafe"
)

// Translation table
//CK_UTF8CHAR -> string
//U_LONG	-> uint

// Wraps CK_VERSION
type Version struct {
	Major byte
	Minor byte
}

type Info struct {
	CryptokiVersion    *Version
	ManufacturerID     string
	Flags              uint
	LibraryDescription string
	LibraryVersion     *Version
}

type SlotInfo struct {
	SlotDescription string
	ManufacturerID  string
	Flags           uint
	HardwareVersion *Version
	FirmwareVersion *Version
}

type TokenInfo struct {
	Label              string
	ManufacturerID     string
	Model              string
	SerialNumber       string
	Flags              uint
	MaxSessionCount    uint
	SessionCount       uint
	MaxRwSessionCount  uint
	RwSessionCount     uint
	MaxPinLen          uint
	MinPinLen          uint
	TotalPublicMemory  uint
	FreePublicMemory   uint
	TotalPrivateMemory uint
	FreePrivateMemory  uint
	HardwareVersion    *Version
	FirmwareVersion    *Version
	UTCTime            string
}

type SessionInfo struct {
	SlotID		uint
	State		uint
	Flags		uint
	DeviceError	uint
}

type Attribute struct {
	AttributeType	uint
	Value		unsafe.Pointer
	ValueLen	uint
}

// type Date struct {} ??
// type Mechanism struct {} ??
// type MechanismInfo struct {} ??
// callback functions
