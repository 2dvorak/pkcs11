// Package pkcs11 provides an interface to the PKCS#11 API.
//
//
package pkcs11

/*
#cgo LDFLAGS: -lltdl
#include <string.h>
#include "pkcs11c"

CK_SLOT_INFO_PTR SlotIndex(CK_SLOT_INFO_PTR *l, int i) { return l[i]; } 
CK_TOKEN_INFO_PTR TokenIndex(CK_TOKEN_INFO_PTR *l, int i) { return l[i]; } 
CK_SLOT_INFO_PTR SlotNew() { CK_SLOT_INFO_PTR s = NULL; return s; }
CK_TOKEN_INFO_PTR TokenNew() { CK_TOKEN_INFO_PTR t = NULL; return t; }
CK_ULONG UlongNew() { CK_ULONG u = 0; return u; }
*/
import "C"

import (
	"strconv"
	"unsafe"
)

// Pkcs11Error represents an error from the PKCS#11 library.
type Pkcs11Error struct {
	err string // error text
	rv  int    // return value from pkcs11 api
}

func newPkcs11Error(s string, rv C.CK_RV) *Pkcs11Error {
	return &Pkcs11Error{err: s, rv: int(rv)}
}

func (e *Pkcs11Error) Error() string {
	s := "pkcs11: " + e.err + "(rv: " + strconv.Itoa(e.rv) + ")"
	return s
}

type Pkcs11 struct {
	ctx *C.struct_ctx
}

// New returns a new instance of a pkcs11. The dynamic PKCS#11 library is
// loaded. New returns nil on error.
func New(module string) *Pkcs11 {
	p := new(Pkcs11)
	mod := C.CString(module)
	defer C.free(unsafe.Pointer(mod))
	p.ctx = C.New(mod)
	if p.ctx == nil {
		return nil
	}
	return p
}

// Wraps PKCS#11's C_Initialize.
func (p *Pkcs11) C_Initialize() error {
	e := C.Go_C_Initialize(p.ctx)
	if e != C.CKR_OK {
		return newPkcs11Error("", e)
	}
	return nil
}

// Wraps PKCS#11's C_Finalize.
func (p *Pkcs11) C_Finalize() error {
	e := C.Go_C_Finalize(p.ctx)
	if e != C.CKR_OK {
		return newPkcs11Error("", e)
	}
	return nil
}

// Destroy unload the module and frees any remaining memory.
func (p *Pkcs11) Destroy() {
	if p == nil {
		return
	}
	C.Destroy(p.ctx)
}

func (p *Pkcs11) C_GetInfo() (*Info, error) {
	var pInfo C.CK_INFO_PTR
	defer C.free(unsafe.Pointer(pInfo))
	e := C.Go_C_GetInfo(p.ctx, &pInfo)
	if e != C.CKR_OK {
		return nil, newPkcs11Error("", e)
	}
	return infoFromC(pInfo), nil
}

func versionFromC(v C.CK_VERSION) *Version {
	v1 := new(Version)
	v1.Major = byte(v.major)
	v1.Minor = byte(v.minor)
	return v1
}

func stringFromC(p unsafe.Pointer, i int) string {
	return string(C.GoBytes(p, C.int(i)))
}

func infoFromC(pInfo C.CK_INFO_PTR) *Info {
	i := new(Info)
	i.ManufacturerID = stringFromC(unsafe.Pointer(&(pInfo.manufacturerID)), 32)
	i.CryptokiVersion = versionFromC((*pInfo).cryptokiVersion)
	i.Flags = uint(pInfo.flags)
	i.LibraryDescription = stringFromC(unsafe.Pointer(&(pInfo.libraryDescription)), 32)
	i.LibraryVersion = versionFromC(pInfo.libraryVersion)
	return i
}

/*
// Slots returns all available slots in the system.
func (p *Pkcs11) Slots() (s []*Slot, e error) {
	slots := C.SlotNew()
	tokens := C.TokenNew()
	nslots := C.UlongNew()
	if rv := C.Slots(p.ctx, &slots, &tokens, &nslots); rv != 0 {
		return nil, nil // TODO(mg): error
	}
	for i := 0; i < int(nslots); i++ {
		o := new(Slot)
		o.slotId = i
		o.Description = string(C.GoBytes(unsafe.Pointer(&C.SlotIndex(&slots, C.int(i)).slotDescription), 64))
		o.Manufacturer = string(C.GoBytes(unsafe.Pointer(&C.SlotIndex(&slots, C.int(i)).manufacturerID), 32))
		o.Removable = int(C.SlotIndex(&slots, C.int(i)).flags)&C.CKF_REMOVABLE_DEVICE == C.CKF_REMOVABLE_DEVICE
		if C.TokenIndex(&tokens, C.int(i)) != nil {
			t := new(Token)
			t.parent = p
			t.slotId = o.slotId
			t.Label = string(C.GoBytes(unsafe.Pointer(&C.TokenIndex(&tokens, C.int(i)).label), 32))
			t.Manufacturer = string(C.GoBytes(unsafe.Pointer(&C.TokenIndex(&tokens, C.int(i)).manufacturerID), 32))
			t.Model = string(C.GoBytes(unsafe.Pointer(&C.TokenIndex(&tokens, C.int(i)).manufacturerID), 16))
			t.Serial = string(C.GoBytes(unsafe.Pointer(&C.TokenIndex(&tokens, C.int(i)).serialNumber), 16))

			t.HasRandGenerator = int(C.TokenIndex(&tokens, C.int(i)).flags)&C.CKF_RNG == C.CKF_RNG
			t.ReadOnly = int(C.TokenIndex(&tokens, C.int(i)).flags)&C.CKF_WRITE_PROTECTED == C.CKF_WRITE_PROTECTED
			t.LoginRequired = int(C.TokenIndex(&tokens, C.int(i)).flags)&C.CKF_LOGIN_REQUIRED == C.CKF_LOGIN_REQUIRED
			t.UserPinSet = int(C.TokenIndex(&tokens, C.int(i)).flags)&C.CKF_USER_PIN_INITIALIZED == C.CKF_USER_PIN_INITIALIZED
			t.Initialized = int(C.TokenIndex(&tokens, C.int(i)).flags)&C.CKF_TOKEN_INITIALIZED == C.CKF_TOKEN_INITIALIZED

			o.Token = t
		}
		s = append(s, o)
	}
	return
}

// Init initializes a token.
func (t *Token) Init(sopin, label string) error {
	cpin := C.CString(sopin)
	clab := C.CString(label) // 32 bytes, padded with spaces
	t1 := C.TokenNew()
	defer C.free(unsafe.Pointer(cpin))
	defer C.free(unsafe.Pointer(clab))
	defer C.free(unsafe.Pointer(t1))

	rv := C.InitToken(t.parent.ctx, &t1, C.uint(t.slotId), cpin, C.uint(len(sopin)), clab)
	if rv != 0 {
		return nil // TODO(mg): error
	}
	// TODO: more
	t.Label = string(C.GoBytes(unsafe.Pointer(&t1.label), 32))
	t.Initialized = int(t1.flags)&C.CKF_TOKEN_INITIALIZED == C.CKF_TOKEN_INITIALIZED
	return nil
}
*/
