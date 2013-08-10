// Package pkcs11 is a thin wrapper around the PKCS#11 crypto library.
package pkcs11

/*
#cgo LDFLAGS: -lltdl
#define CK_PTR *
#ifndef NULL_PTR
#define NULL_PTR 0
#endif
#define CK_DEFINE_FUNCTION(returnType, name) returnType name
#define CK_DECLARE_FUNCTION(returnType, name) returnType name
#define CK_DECLARE_FUNCTION_POINTER(returnType, name) returnType (* name)
#define CK_CALLBACK_FUNCTION(returnType, name) returnType (* name)

#include <stdlib.h>
#include <stdio.h>
#include <ltdl.h>
#include <unistd.h>
#include "pkcs11.h"

struct ctx {
        lt_dlhandle handle;
        CK_FUNCTION_LIST_PTR sym;
};

// New initializes a ctx and fills the symbol table.
struct ctx * New(const char *module) {
        if (lt_dlinit() != 0) {
                return NULL;
        }
        CK_C_GetFunctionList list;
        struct ctx *c=  calloc(1, sizeof(struct ctx));
        c->handle = lt_dlopen(module);
        if (c->handle == NULL) {
                free(c);
                return NULL;
        }
        list = (CK_C_GetFunctionList) lt_dlsym(c->handle, "C_GetFunctionList");
        if (list == NULL) {
                free(c);
                return NULL;
        }
        list(&c->sym);
        return c;
}

// Destroy cleans up a ctx.
void Destroy(struct ctx *c) {
        if (!c) {
                return;
        }
        if (c->handle == NULL) {
                return;
        }
        if (lt_dlclose(c->handle) < 0) {
                return;
        }
        lt_dlexit();
        free(c);
}

CK_RV Initialize(struct ctx* c, CK_VOID_PTR initArgs) {
	return c->sym->C_Initialize(initArgs);
}

CK_RV Finalize(struct ctx* c) {
	return c->sym->C_Finalize(NULL_PTR);
}

CK_RV GetSlotList(struct ctx* c, CK_BBOOL tokenPresent, CK_ULONG_PTR *slotList, CK_ULONG_PTR ulCount) {
	CK_RV e = c->sym->C_GetSlotList(tokenPresent, NULL_PTR, ulCount);
	if (e != CKR_OK) {
		return e;
	}
	*slotList = calloc(1, sizeof(CK_SLOT_ID) * *ulCount);
	e = c->sym->C_GetSlotList(tokenPresent, *slotList, ulCount);
	return e;
}

CK_RV OpenSession(struct ctx* c, CK_ULONG slotID, CK_ULONG flags, CK_SESSION_HANDLE_PTR session) {
	CK_RV e = c->sym->C_OpenSession((CK_SLOT_ID)slotID, (CK_FLAGS)flags, NULL_PTR, NULL_PTR, session);
	return e;
}

CK_RV CloseSession(struct ctx *c, CK_SESSION_HANDLE session) {
	CK_RV e = c->sym->C_CloseSession(session);
	return e;
}

CK_RV CloseAllSessions(struct ctx *c, CK_ULONG slotID) {
	CK_RV e = c->sym->C_CloseAllSessions(slotID);
	return e;
}

CK_RV Login(struct ctx* c, CK_SESSION_HANDLE session, CK_USER_TYPE userType, char* pin, CK_ULONG pinLen) {
	CK_RV e = c->sym->C_Login(session, userType, (CK_UTF8CHAR_PTR)pin, pinLen);
	return e;
}

CK_RV Logout(struct ctx *c, CK_SESSION_HANDLE session) {
	CK_RV e = c->sym->C_Logout(session);
	return e;
}

CK_RV DigestInit(struct ctx *c, CK_SESSION_HANDLE session, CK_MECHANISM_PTR mechanism) {
	CK_RV e = c->sym->C_DigestInit(session, mechanism);
	return e;
}

CK_RV Digest(struct ctx *c, CK_SESSION_HANDLE session, CK_BYTE_PTR message, CK_ULONG mlen, CK_BYTE_PTR *hash, CK_ULONG_PTR hashlen) {
        CK_RV rv = c->sym->C_Digest(session, message, mlen, NULL, hashlen);
        if (rv != CKR_OK) {
                return rv;
        }
        *hash = malloc(*hashlen * sizeof(CK_BYTE));
	if (*hash == NULL) {
		return CKR_HOST_MEMORY;
	}
        rv = c->sym->C_Digest(session, message, mlen, *hash, hashlen);
	return rv;
}

CK_RV GenerateKey(struct ctx* c, CK_SESSION_HANDLE session, CK_MECHANISM_PTR mechanism,
	CK_ATTRIBUTE_PTR temp, CK_ULONG tempCount, CK_OBJECT_HANDLE_PTR key) {
	CK_RV e = c->sym->C_GenerateKey(session, mechanism, temp, tempCount, key);
	return e;
}

CK_RV GenerateKeyPair(struct ctx* c, CK_SESSION_HANDLE session, CK_MECHANISM_PTR mechanism,
	CK_ATTRIBUTE_PTR pub, CK_ULONG pubCount, CK_ATTRIBUTE_PTR priv, CK_ULONG privCount,
	CK_OBJECT_HANDLE_PTR pubkey, CK_OBJECT_HANDLE_PTR privkey) {
	CK_RV e = c->sym->C_GenerateKeyPair(session, mechanism, pub, pubCount, priv, privCount,
					pubkey, privkey);
	return e;
}

CK_RV SignInit(struct ctx* c, CK_SESSION_HANDLE session, CK_MECHANISM_PTR mechanism, CK_OBJECT_HANDLE key) {
	CK_RV e = c->sym->C_SignInit(session, mechanism, key);
	return e;
}

CK_RV Sign(struct ctx *c, CK_SESSION_HANDLE session, CK_BYTE_PTR message, CK_ULONG mlen, CK_BYTE_PTR *sig, CK_ULONG_PTR siglen) {
        CK_RV rv = c->sym->C_Sign(session, message, mlen, NULL, siglen);
        if (rv != CKR_OK) {
                return rv;
        }
        *sig = malloc(*siglen * sizeof(CK_BYTE));
	if (*sig == NULL) {
		return CKR_HOST_MEMORY;
	}
        rv = c->sym->C_Sign(session, message, mlen, *sig, siglen);
	return rv;
}

CK_RV EncryptInit(struct ctx* c, CK_SESSION_HANDLE session, CK_MECHANISM_PTR mechanism, CK_OBJECT_HANDLE key) {
	CK_RV e = c->sym->C_EncryptInit(session, mechanism, key);
	return e;
}

CK_RV Encrypt(struct ctx *c, CK_SESSION_HANDLE session, CK_BYTE_PTR message, CK_ULONG mlen, CK_BYTE_PTR *enc, CK_ULONG_PTR enclen) {
        CK_RV rv = c->sym->C_Encrypt(session, message, mlen, NULL, enclen);
        if (rv != CKR_OK) {
                return rv;
        }
        *enc = malloc(*enclen * sizeof(CK_BYTE));
	if (*enc == NULL) {
		return CKR_HOST_MEMORY;
	}
        rv = c->sym->C_Encrypt(session, message, mlen, *enc, enclen);
	return rv;
}

CK_RV CreateObject(struct ctx* c, CK_SESSION_HANDLE session, CK_ATTRIBUTE_PTR temp, 
		CK_ULONG tempCount, CK_OBJECT_HANDLE_PTR obj) {
	CK_RV e = c->sym->C_CreateObject(session, temp, tempCount, obj);
	return e;
}

// TODO(miek): CopyObject

CK_RV DestroyObject(struct ctx* c, CK_SESSION_HANDLE session, CK_OBJECT_HANDLE object) {
	CK_RV e = c->sym->C_DestroyObject(session, object);
	return e;
}

CK_RV GetObjectSize(struct ctx* c, CK_SESSION_HANDLE session, CK_OBJECT_HANDLE object, CK_ULONG_PTR size) {
	CK_RV e = c->sym->C_GetObjectSize(session, object, size);
	return e;
}

// TODO(miek): GetAttributeValue
// TODO(miek): SetAttributeValue

CK_RV FindObjectsInit(struct ctx* c, CK_SESSION_HANDLE session, CK_ATTRIBUTE_PTR temp, CK_ULONG tempCount) {
	CK_RV e = c->sym->C_FindObjectsInit(session, temp, tempCount);
	return e;
}

//CK_RV FindObjects(struct ctx* c, CK_SESSION_HANDLE session
*/
import "C"

import "unsafe"

// Ctx contains the current pkcs11 context.
type Ctx struct {
	ctx         *C.struct_ctx
	initialized bool
}

// New creates a new context.
func New(module string) *Ctx {
	c := new(Ctx)
	mod := C.CString(module)
	defer C.free(unsafe.Pointer(mod))
	c.ctx = C.New(mod)
	if c.ctx == nil {
		return nil
	}
	return c
}

// Destroy unload the module and frees any remaining memory.
func (c *Ctx) Destroy() {
	if c == nil {
		return
	}
	C.Destroy(c.ctx)
}

/* Initialize initializes the Cryptoki library. */
func (c *Ctx) Initialize() error {
	args := &C.CK_C_INITIALIZE_ARGS{nil, nil, nil, nil, C.CKF_OS_LOCKING_OK, nil}
	e := C.Initialize(c.ctx, C.CK_VOID_PTR(args))
	if e == C.CKR_OK {
		c.initialized = true // TODO(miek): keep?
	}
	return toError(e)
}

/* Finalize indicates that an application is done with the Cryptoki library. */
func (c *Ctx) Finalize() error {
	if !c.initialized {
		panic("pkcs11: context not initialized")
	}
	e := C.Finalize(c.ctx)
	return toError(e)
}

/* GetSlotList obtains a list of slots in the system. */
func (c *Ctx) GetSlotList(tokenPresent bool) ([]uint, error) {
	var (
		slotList C.CK_ULONG_PTR
		ulCount  C.CK_ULONG
	)
	e := C.GetSlotList(c.ctx, cBBool(tokenPresent), &slotList, &ulCount)
	if toError(e) == nil {
		l := toList(slotList, ulCount)
		return l, nil
	}
	return nil, toError(e)
}

/* OpenSession opens a session between an application and a token. */
func (c *Ctx) OpenSession(slotID uint, flags uint) (SessionHandle, error) {
	var s C.CK_SESSION_HANDLE
	e := C.OpenSession(c.ctx, C.CK_ULONG(slotID), C.CK_ULONG(flags), C.CK_SESSION_HANDLE_PTR(&s))
	return SessionHandle(s), toError(e)
}

/* CloseSession closes a session between an application and a token. */
func (c *Ctx) CloseSession(sh SessionHandle) error {
	e := C.CloseSession(c.ctx, C.CK_SESSION_HANDLE(sh))
	return toError(e)
}

/* CloseAllSessions closes all sessions with a token. */
func (c *Ctx) CloseAllSessions(slotID uint) error {
	e := C.CloseAllSessions(c.ctx, C.CK_ULONG(slotID))
	return toError(e)
}

/* Login logs a user into a token. */
func (c *Ctx) Login(sh SessionHandle, userType uint, pin string) error {
	p := C.CString(pin)
	defer C.free(unsafe.Pointer(p))
	e := C.Login(c.ctx, C.CK_SESSION_HANDLE(sh), C.CK_USER_TYPE(userType), p, C.CK_ULONG(len(pin)))
	return toError(e)
}

/* Logout logs a user out from a token. */
func (c *Ctx) Logout(sh SessionHandle) error {
	e := C.Logout(c.ctx, C.CK_SESSION_HANDLE(sh))
	return toError(e)
}

/* GenerateKey generates a secret key, creating a new key object. */
func (c *Ctx) GenerateKey(sh SessionHandle, m []*Mechanism, temp []*Attribute) (ObjectHandle, error) {
	var key C.CK_OBJECT_HANDLE
	t, tcount := cAttributeList(temp)
	mech, _ := cMechanismList(m)
	e := C.GenerateKey(c.ctx, C.CK_SESSION_HANDLE(sh), mech, t, tcount, C.CK_OBJECT_HANDLE_PTR(&key))
	e1 := toError(e)
	if e1 == nil {
		return ObjectHandle(key), nil
	}
	return 0, e1
}

/* GenerateKeyPair generates a public-key/private-key pair creating new key objects. */
func (c *Ctx) GenerateKeyPair(sh SessionHandle, m []*Mechanism, public, private []*Attribute) (ObjectHandle, ObjectHandle, error) {
	var (
		pubkey  C.CK_OBJECT_HANDLE
		privkey C.CK_OBJECT_HANDLE
	)
	pub, pubcount := cAttributeList(public)
	priv, privcount := cAttributeList(private)
	mech, _ := cMechanismList(m)
	e := C.GenerateKeyPair(c.ctx, C.CK_SESSION_HANDLE(sh), mech, pub, pubcount, priv, privcount, C.CK_OBJECT_HANDLE_PTR(&pubkey), C.CK_OBJECT_HANDLE_PTR(&privkey))
	e1 := toError(e)
	if e1 == nil {
		return ObjectHandle(pubkey), ObjectHandle(privkey), nil
	}
	return 0, 0, e1
}

func (c *Ctx) DigestInit(sh SessionHandle, m []*Mechanism) error {
	mech, _ := cMechanismList(m)
	e := C.DigestInit(c.ctx, C.CK_SESSION_HANDLE(sh), mech)
	return toError(e)
}

func (c *Ctx) Digest(sh SessionHandle, message []byte) ([]byte, error) {
	var (
		hash    C.CK_BYTE_PTR
		hashlen C.CK_ULONG
	)
	e := C.Digest(c.ctx, C.CK_SESSION_HANDLE(sh), C.CK_BYTE_PTR(unsafe.Pointer(&message[0])), C.CK_ULONG(len(message)), &hash, &hashlen)
	if toError(e) != nil {
		return nil, toError(e)
	}
	h := C.GoBytes(unsafe.Pointer(hash), C.int(hashlen))
	return h, nil
}

func (c *Ctx) SignInit(sh SessionHandle, m []*Mechanism, o ObjectHandle) error {
	mech, _ := cMechanismList(m) // Only the first is used, but still use a list.
	e := C.SignInit(c.ctx, C.CK_SESSION_HANDLE(sh), mech, C.CK_OBJECT_HANDLE(o))
	return toError(e)
}

// Sign signs (encrypts with private key) data in a single part, where the signature
// is (will be) an appendix to the data, and plaintext cannot be recovered from the signature.
func (c *Ctx) Sign(sh SessionHandle, message []byte) ([]byte, error) {
	var (
		sig    C.CK_BYTE_PTR
		siglen C.CK_ULONG
	)
	e := C.Sign(c.ctx, C.CK_SESSION_HANDLE(sh), C.CK_BYTE_PTR(unsafe.Pointer(&message[0])), C.CK_ULONG(len(message)), &sig, &siglen)
	if toError(e) != nil {
		return nil, toError(e)
	}
	s := C.GoBytes(unsafe.Pointer(sig), C.int(siglen))
	return s, nil
}

/* EncryptInit initializes an encryption operation. */
func (c *Ctx) EncryptInit(sh SessionHandle, m []*Mechanism, o ObjectHandle) error {
	mech, _ := cMechanismList(m)
	e := C.EncryptInit(c.ctx, C.CK_SESSION_HANDLE(sh), mech, C.CK_OBJECT_HANDLE(o))
	return toError(e)
}

/* Encrypt encrypts single-part data. */
func (c *Ctx) Encrypt(sh SessionHandle, message []byte) ([]byte, error) {
	var (
		enc    C.CK_BYTE_PTR
		enclen C.CK_ULONG
	)
	e := C.Encrypt(c.ctx, C.CK_SESSION_HANDLE(sh), C.CK_BYTE_PTR(unsafe.Pointer(&message[0])), C.CK_ULONG(len(message)), &enc, &enclen)
	if toError(e) != nil {
		return nil, toError(e)
	}
	s := C.GoBytes(unsafe.Pointer(enc), C.int(enclen))
	return s, nil
}

/* C_CreateObject creates a new object. */
func (c *Ctx) CreateObject(sh SessionHandle, temp []*Attribute) (ObjectHandle, error) {
	var obj C.CK_OBJECT_HANDLE
	t, tcount := cAttributeList(temp)
	e := C.CreateObject(c.ctx, C.CK_SESSION_HANDLE(sh), t, tcount, C.CK_OBJECT_HANDLE_PTR(&obj))
	e1 := toError(e)
	if e1 == nil {
		return ObjectHandle(obj), nil
	}
	return 0, e1
}

// TODO(miek): CopyObject here

/* C_DestroyObject destroys an object. */
func (c *Ctx) DestroyObject(sh SessionHandle, oh ObjectHandle) error {
	e := C.DestroyObject(c.ctx, C.CK_SESSION_HANDLE(sh), C.CK_OBJECT_HANDLE(oh))
	return toError(e)
}

/* C_GetObjectSize gets the size of an object in bytes. */
func (c *Ctx) GetObjectSize(sh SessionHandle, oh ObjectHandle) (uint, error) {
	var size C.CK_ULONG
	e := C.GetObjectSize(c.ctx, C.CK_SESSION_HANDLE(sh), C.CK_OBJECT_HANDLE(oh), &size)
	return uint(size), toError(e)
}

// TODO(miek): GetAttributeValue
// TODO(miek): SetAttributeValue

func (c *Ctx) FindObjectsInit(sh SessionHandle, temp []*Attribute) error {
	t, tcount := cAttributeList(temp)
	e := C.FindObjectsInit(c.ctx, C.CK_SESSION_HANDLE(sh), t, tcount)
	return toError(e)
}

// C_FindObjects continues a search for token and session
// objects that match a template, obtaining additional object
// handles. The returned boolean indicates if the list would
// have been larger than max.
// func (c *Ctx) FindOBjects(sh SessionHandle, max int)
