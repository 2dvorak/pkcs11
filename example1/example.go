package main

import (
	"fmt"
	"github.com/miekg/pkcs11"
)

func yesno(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

func main() {
	p := pkcs11.New("/home/miekg/libsofthsm.so")
	if p == nil {
		return
	}
	if e := p.C_Initialize(); e != nil {
		fmt.Printf("init error %s\n", e.Error())
		return
	}

	defer p.Destroy()
	defer p.C_Finalize()
	if info, err := p.C_GetInfo(); err == nil {
		fmt.Printf("%s\n", info.ManufacturerID)
	} else {
		fmt.Printf("error %s\n", err.Error())
	}
	slots, e := p.C_GetSlotList(true)
	fmt.Printf("slots %v\n", slots)
	if e != nil {
		fmt.Printf("%s\n", e.Error())
	}
	// Only works on initialized tokens

	session, e := p.C_OpenSession(slots[0], pkcs11.CKF_SERIAL_SESSION|pkcs11.CKF_RW_SESSION)
	if e != nil {
		fmt.Printf("%s\n", e.Error())
	}

	pub, priv, e := p.C_GenerateKeyPair(session, &pkcs11.CKM_RSA_PKCS_KEY_PAIR_GEN{},
		[]pkcs11.Attribute{&pkcs11.CKA_MODULUS_BITS{1024}}, []pkcs11.Attribute{&pkcs11.CKA_TOKEN{true}, &pkcs11.CKA_PRIVATE{false}})
	if e != nil {
		fmt.Printf("%s\n", e.Error())
	}
	println(pub)
	println(priv)
}
