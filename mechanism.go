package pkcs11

import (
	"unsafe"
)

// Generated with:
// grep CKM *t.h|grep '#define' | sed 's/^#define //' | awk ' { print $1 "=" $2 } '

type CKM_RSA_PKCS_KEY_PAIR_GEN struct{}

func (m *CKM_RSA_PKCS_KEY_PAIR_GEN) Type() uint                { return cKM_RSA_PKCS_KEY_PAIR_GEN }
func (m *CKM_RSA_PKCS_KEY_PAIR_GEN) Parameter() unsafe.Pointer { return nil }
func (m *CKM_RSA_PKCS_KEY_PAIR_GEN) Len() uint                 { return 0 }

// Signing and MACing

type CKM_RSA_PKCS struct{}

func (m *CKM_RSA_PKCS) Type() uint                { return cKM_RSA_PKCS }
func (m *CKM_RSA_PKCS) Parameter() unsafe.Pointer { return nil }
func (m *CKM_RSA_PKCS) Len() uint                 { return 0 }

type CKM_SHA1_RSA_PKCS struct{}

func (m *CKM_SHA1_RSA_PKCS) Type() uint                { return cKM_SHA1_RSA_PKCS }
func (m *CKM_SHA1_RSA_PKCS) Parameter() unsafe.Pointer { return nil }
func (m *CKM_SHA1_RSA_PKCS) Len() uint                 { return 0 }

type CKM_SHA256_RSA_PKCS struct{}

func (m *CKM_SHA256_RSA_PKCS) Type() uint                { return cKM_SHA256_RSA_PKCS }
func (m *CKM_SHA256_RSA_PKCS) Parameter() unsafe.Pointer { return nil }
func (m *CKM_SHA256_RSA_PKCS) Len() uint                 { return 0 }

type CKM_SHA384_RSA_PKCS struct{}

func (m *CKM_SHA384_RSA_PKCS) Type() uint                { return cKM_SHA384_RSA_PKCS }
func (m *CKM_SHA384_RSA_PKCS) Parameter() unsafe.Pointer { return nil }
func (m *CKM_SHA384_RSA_PKCS) Len() uint                 { return 0 }

type CKM_SHA512_RSA_PKCS struct{}

func (m *CKM_SHA512_RSA_PKCS) Type() uint                { return cKM_SHA512_RSA_PKCS }
func (m *CKM_SHA512_RSA_PKCS) Parameter() unsafe.Pointer { return nil }
func (m *CKM_SHA512_RSA_PKCS) Len() uint                 { return 0 }

const (
	cKM_RSA_PKCS_KEY_PAIR_GEN          = 0x00000000
	cKM_RSA_PKCS                       = 0x00000001
	cKM_RSA_9796                       = 0x00000002
	cKM_RSA_X_509                      = 0x00000003
	cKM_MD2_RSA_PKCS                   = 0x00000004
	cKM_MD5_RSA_PKCS                   = 0x00000005
	cKM_SHA1_RSA_PKCS                  = 0x00000006
	cKM_RIPEMD128_RSA_PKCS             = 0x00000007
	cKM_RIPEMD160_RSA_PKCS             = 0x00000008
	cKM_RSA_PKCS_OAEP                  = 0x00000009
	cKM_RSA_X9_31_KEY_PAIR_GEN         = 0x0000000A
	cKM_RSA_X9_31                      = 0x0000000B
	cKM_SHA1_RSA_X9_31                 = 0x0000000C
	cKM_RSA_PKCS_PSS                   = 0x0000000D
	cKM_SHA1_RSA_PKCS_PSS              = 0x0000000E
	cKM_DSA_KEY_PAIR_GEN               = 0x00000010
	cKM_DSA                            = 0x00000011
	cKM_DSA_SHA1                       = 0x00000012
	cKM_DH_PKCS_KEY_PAIR_GEN           = 0x00000020
	cKM_DH_PKCS_DERIVE                 = 0x00000021
	cKM_X9_42_DH_KEY_PAIR_GEN          = 0x00000030
	cKM_X9_42_DH_DERIVE                = 0x00000031
	cKM_X9_42_DH_HYBRID_DERIVE         = 0x00000032
	cKM_X9_42_MQV_DERIVE               = 0x00000033
	cKM_SHA256_RSA_PKCS                = 0x00000040
	cKM_SHA384_RSA_PKCS                = 0x00000041
	cKM_SHA512_RSA_PKCS                = 0x00000042
	cKM_SHA256_RSA_PKCS_PSS            = 0x00000043
	cKM_SHA384_RSA_PKCS_PSS            = 0x00000044
	cKM_SHA512_RSA_PKCS_PSS            = 0x00000045
	cKM_SHA224_RSA_PKCS                = 0x00000046
	cKM_SHA224_RSA_PKCS_PSS            = 0x00000047
	cKM_RC2_KEY_GEN                    = 0x00000100
	cKM_RC2_ECB                        = 0x00000101
	cKM_RC2_CBC                        = 0x00000102
	cKM_RC2_MAC                        = 0x00000103
	cKM_RC2_MAC_GENERAL                = 0x00000104
	cKM_RC2_CBC_PAD                    = 0x00000105
	cKM_RC4_KEY_GEN                    = 0x00000110
	cKM_RC4                            = 0x00000111
	cKM_DES_KEY_GEN                    = 0x00000120
	cKM_DES_ECB                        = 0x00000121
	cKM_DES_CBC                        = 0x00000122
	cKM_DES_MAC                        = 0x00000123
	cKM_DES_MAC_GENERAL                = 0x00000124
	cKM_DES_CBC_PAD                    = 0x00000125
	cKM_DES2_KEY_GEN                   = 0x00000130
	cKM_DES3_KEY_GEN                   = 0x00000131
	cKM_DES3_ECB                       = 0x00000132
	cKM_DES3_CBC                       = 0x00000133
	cKM_DES3_MAC                       = 0x00000134
	cKM_DES3_MAC_GENERAL               = 0x00000135
	cKM_DES3_CBC_PAD                   = 0x00000136
	cKM_CDMF_KEY_GEN                   = 0x00000140
	cKM_CDMF_ECB                       = 0x00000141
	cKM_CDMF_CBC                       = 0x00000142
	cKM_CDMF_MAC                       = 0x00000143
	cKM_CDMF_MAC_GENERAL               = 0x00000144
	cKM_CDMF_CBC_PAD                   = 0x00000145
	cKM_DES_OFB64                      = 0x00000150
	cKM_DES_OFB8                       = 0x00000151
	cKM_DES_CFB64                      = 0x00000152
	cKM_DES_CFB8                       = 0x00000153
	cKM_MD2                            = 0x00000200
	cKM_MD2_HMAC                       = 0x00000201
	cKM_MD2_HMAC_GENERAL               = 0x00000202
	cKM_MD5                            = 0x00000210
	cKM_MD5_HMAC                       = 0x00000211
	cKM_MD5_HMAC_GENERAL               = 0x00000212
	cKM_SHA_1                          = 0x00000220
	cKM_SHA_1_HMAC                     = 0x00000221
	cKM_SHA_1_HMAC_GENERAL             = 0x00000222
	cKM_RIPEMD128                      = 0x00000230
	cKM_RIPEMD128_HMAC                 = 0x00000231
	cKM_RIPEMD128_HMAC_GENERAL         = 0x00000232
	cKM_RIPEMD160                      = 0x00000240
	cKM_RIPEMD160_HMAC                 = 0x00000241
	cKM_RIPEMD160_HMAC_GENERAL         = 0x00000242
	cKM_SHA256                         = 0x00000250
	cKM_SHA256_HMAC                    = 0x00000251
	cKM_SHA256_HMAC_GENERAL            = 0x00000252
	cKM_SHA224                         = 0x00000255
	cKM_SHA224_HMAC                    = 0x00000256
	cKM_SHA224_HMAC_GENERAL            = 0x00000257
	cKM_SHA384                         = 0x00000260
	cKM_SHA384_HMAC                    = 0x00000261
	cKM_SHA384_HMAC_GENERAL            = 0x00000262
	cKM_SHA512                         = 0x00000270
	cKM_SHA512_HMAC                    = 0x00000271
	cKM_SHA512_HMAC_GENERAL            = 0x00000272
	cKM_SECURID_KEY_GEN                = 0x00000280
	cKM_SECURID                        = 0x00000282
	cKM_HOTP_KEY_GEN                   = 0x00000290
	cKM_HOTP                           = 0x00000291
	cKM_ACTI                           = 0x000002A0
	cKM_ACTI_KEY_GEN                   = 0x000002A1
	cKM_CAST_KEY_GEN                   = 0x00000300
	cKM_CAST_ECB                       = 0x00000301
	cKM_CAST_CBC                       = 0x00000302
	cKM_CAST_MAC                       = 0x00000303
	cKM_CAST_MAC_GENERAL               = 0x00000304
	cKM_CAST_CBC_PAD                   = 0x00000305
	cKM_CAST3_KEY_GEN                  = 0x00000310
	cKM_CAST3_ECB                      = 0x00000311
	cKM_CAST3_CBC                      = 0x00000312
	cKM_CAST3_MAC                      = 0x00000313
	cKM_CAST3_MAC_GENERAL              = 0x00000314
	cKM_CAST3_CBC_PAD                  = 0x00000315
	cKM_CAST5_KEY_GEN                  = 0x00000320
	cKM_CAST128_KEY_GEN                = 0x00000320
	cKM_CAST5_ECB                      = 0x00000321
	cKM_CAST128_ECB                    = 0x00000321
	cKM_CAST5_CBC                      = 0x00000322
	cKM_CAST128_CBC                    = 0x00000322
	cKM_CAST5_MAC                      = 0x00000323
	cKM_CAST128_MAC                    = 0x00000323
	cKM_CAST5_MAC_GENERAL              = 0x00000324
	cKM_CAST128_MAC_GENERAL            = 0x00000324
	cKM_CAST5_CBC_PAD                  = 0x00000325
	cKM_CAST128_CBC_PAD                = 0x00000325
	cKM_RC5_KEY_GEN                    = 0x00000330
	cKM_RC5_ECB                        = 0x00000331
	cKM_RC5_CBC                        = 0x00000332
	cKM_RC5_MAC                        = 0x00000333
	cKM_RC5_MAC_GENERAL                = 0x00000334
	cKM_RC5_CBC_PAD                    = 0x00000335
	cKM_IDEA_KEY_GEN                   = 0x00000340
	cKM_IDEA_ECB                       = 0x00000341
	cKM_IDEA_CBC                       = 0x00000342
	cKM_IDEA_MAC                       = 0x00000343
	cKM_IDEA_MAC_GENERAL               = 0x00000344
	cKM_IDEA_CBC_PAD                   = 0x00000345
	cKM_GENERIC_SECRET_KEY_GEN         = 0x00000350
	cKM_CONCATENATE_BASE_AND_KEY       = 0x00000360
	cKM_CONCATENATE_BASE_AND_DATA      = 0x00000362
	cKM_CONCATENATE_DATA_AND_BASE      = 0x00000363
	cKM_XOR_BASE_AND_DATA              = 0x00000364
	cKM_EXTRACT_KEY_FROM_KEY           = 0x00000365
	cKM_SSL3_PRE_MASTER_KEY_GEN        = 0x00000370
	cKM_SSL3_MASTER_KEY_DERIVE         = 0x00000371
	cKM_SSL3_KEY_AND_MAC_DERIVE        = 0x00000372
	cKM_SSL3_MASTER_KEY_DERIVE_DH      = 0x00000373
	cKM_TLS_PRE_MASTER_KEY_GEN         = 0x00000374
	cKM_TLS_MASTER_KEY_DERIVE          = 0x00000375
	cKM_TLS_KEY_AND_MAC_DERIVE         = 0x00000376
	cKM_TLS_MASTER_KEY_DERIVE_DH       = 0x00000377
	cKM_TLS_PRF                        = 0x00000378
	cKM_SSL3_MD5_MAC                   = 0x00000380
	cKM_SSL3_SHA1_MAC                  = 0x00000381
	cKM_MD5_KEY_DERIVATION             = 0x00000390
	cKM_MD2_KEY_DERIVATION             = 0x00000391
	cKM_SHA1_KEY_DERIVATION            = 0x00000392
	cKM_SHA256_KEY_DERIVATION          = 0x00000393
	cKM_SHA384_KEY_DERIVATION          = 0x00000394
	cKM_SHA512_KEY_DERIVATION          = 0x00000395
	cKM_SHA224_KEY_DERIVATION          = 0x00000396
	cKM_PBE_MD2_DES_CBC                = 0x000003A0
	cKM_PBE_MD5_DES_CBC                = 0x000003A1
	cKM_PBE_MD5_CAST_CBC               = 0x000003A2
	cKM_PBE_MD5_CAST3_CBC              = 0x000003A3
	cKM_PBE_MD5_CAST5_CBC              = 0x000003A4
	cKM_PBE_MD5_CAST128_CBC            = 0x000003A4
	cKM_PBE_SHA1_CAST5_CBC             = 0x000003A5
	cKM_PBE_SHA1_CAST128_CBC           = 0x000003A5
	cKM_PBE_SHA1_RC4_128               = 0x000003A6
	cKM_PBE_SHA1_RC4_40                = 0x000003A7
	cKM_PBE_SHA1_DES3_EDE_CBC          = 0x000003A8
	cKM_PBE_SHA1_DES2_EDE_CBC          = 0x000003A9
	cKM_PBE_SHA1_RC2_128_CBC           = 0x000003AA
	cKM_PBE_SHA1_RC2_40_CBC            = 0x000003AB
	cKM_PKCS5_PBKD2                    = 0x000003B0
	cKM_PBA_SHA1_WITH_SHA1_HMAC        = 0x000003C0
	cKM_WTLS_PRE_MASTER_KEY_GEN        = 0x000003D0
	cKM_WTLS_MASTER_KEY_DERIVE         = 0x000003D1
	cKM_WTLS_MASTER_KEY_DERIVE_DH_ECC  = 0x000003D2
	cKM_WTLS_PRF                       = 0x000003D3
	cKM_WTLS_SERVER_KEY_AND_MAC_DERIVE = 0x000003D4
	cKM_WTLS_CLIENT_KEY_AND_MAC_DERIVE = 0x000003D5
	cKM_KEY_WRAP_LYNKS                 = 0x00000400
	cKM_KEY_WRAP_SET_OAEP              = 0x00000401
	cKM_CMS_SIG                        = 0x00000500
	cKM_KIP_DERIVE                     = 0x00000510
	cKM_KIP_WRAP                       = 0x00000511
	cKM_KIP_MAC                        = 0x00000512
	cKM_CAMELLIA_KEY_GEN               = 0x00000550
	cKM_CAMELLIA_ECB                   = 0x00000551
	cKM_CAMELLIA_CBC                   = 0x00000552
	cKM_CAMELLIA_MAC                   = 0x00000553
	cKM_CAMELLIA_MAC_GENERAL           = 0x00000554
	cKM_CAMELLIA_CBC_PAD               = 0x00000555
	cKM_CAMELLIA_ECB_ENCRYPT_DATA      = 0x00000556
	cKM_CAMELLIA_CBC_ENCRYPT_DATA      = 0x00000557
	cKM_CAMELLIA_CTR                   = 0x00000558
	cKM_ARIA_KEY_GEN                   = 0x00000560
	cKM_ARIA_ECB                       = 0x00000561
	cKM_ARIA_CBC                       = 0x00000562
	cKM_ARIA_MAC                       = 0x00000563
	cKM_ARIA_MAC_GENERAL               = 0x00000564
	cKM_ARIA_CBC_PAD                   = 0x00000565
	cKM_ARIA_ECB_ENCRYPT_DATA          = 0x00000566
	cKM_ARIA_CBC_ENCRYPT_DATA          = 0x00000567
	cKM_SKIPJACK_KEY_GEN               = 0x00001000
	cKM_SKIPJACK_ECB64                 = 0x00001001
	cKM_SKIPJACK_CBC64                 = 0x00001002
	cKM_SKIPJACK_OFB64                 = 0x00001003
	cKM_SKIPJACK_CFB64                 = 0x00001004
	cKM_SKIPJACK_CFB32                 = 0x00001005
	cKM_SKIPJACK_CFB16                 = 0x00001006
	cKM_SKIPJACK_CFB8                  = 0x00001007
	cKM_SKIPJACK_WRAP                  = 0x00001008
	cKM_SKIPJACK_PRIVATE_WRAP          = 0x00001009
	cKM_SKIPJACK_RELAYX                = 0x0000100a
	cKM_KEA_KEY_PAIR_GEN               = 0x00001010
	cKM_KEA_KEY_DERIVE                 = 0x00001011
	cKM_FORTEZZA_TIMESTAMP             = 0x00001020
	cKM_BATON_KEY_GEN                  = 0x00001030
	cKM_BATON_ECB128                   = 0x00001031
	cKM_BATON_ECB96                    = 0x00001032
	cKM_BATON_CBC128                   = 0x00001033
	cKM_BATON_COUNTER                  = 0x00001034
	cKM_BATON_SHUFFLE                  = 0x00001035
	cKM_BATON_WRAP                     = 0x00001036
	cKM_ECDSA_KEY_PAIR_GEN             = 0x00001040
	cKM_EC_KEY_PAIR_GEN                = 0x00001040
	cKM_ECDSA                          = 0x00001041
	cKM_ECDSA_SHA1                     = 0x00001042
	cKM_ECDH1_DERIVE                   = 0x00001050
	cKM_ECDH1_COFACTOR_DERIVE          = 0x00001051
	cKM_ECMQV_DERIVE                   = 0x00001052
	cKM_JUNIPER_KEY_GEN                = 0x00001060
	cKM_JUNIPER_ECB128                 = 0x00001061
	cKM_JUNIPER_CBC128                 = 0x00001062
	cKM_JUNIPER_COUNTER                = 0x00001063
	cKM_JUNIPER_SHUFFLE                = 0x00001064
	cKM_JUNIPER_WRAP                   = 0x00001065
	cKM_FASTHASH                       = 0x00001070
	cKM_AES_KEY_GEN                    = 0x00001080
	cKM_AES_ECB                        = 0x00001081
	cKM_AES_CBC                        = 0x00001082
	cKM_AES_MAC                        = 0x00001083
	cKM_AES_MAC_GENERAL                = 0x00001084
	cKM_AES_CBC_PAD                    = 0x00001085
	cKM_AES_CTR                        = 0x00001086
	cKM_BLOWFISH_KEY_GEN               = 0x00001090
	cKM_BLOWFISH_CBC                   = 0x00001091
	cKM_TWOFISH_KEY_GEN                = 0x00001092
	cKM_TWOFISH_CBC                    = 0x00001093
	cKM_DES_ECB_ENCRYPT_DATA           = 0x00001100
	cKM_DES_CBC_ENCRYPT_DATA           = 0x00001101
	cKM_DES3_ECB_ENCRYPT_DATA          = 0x00001102
	cKM_DES3_CBC_ENCRYPT_DATA          = 0x00001103
	cKM_AES_ECB_ENCRYPT_DATA           = 0x00001104
	cKM_AES_CBC_ENCRYPT_DATA           = 0x00001105
	cKM_DSA_PARAMETER_GEN              = 0x00002000
	cKM_DH_PKCS_PARAMETER_GEN          = 0x00002001
	cKM_X9_42_DH_PARAMETER_GEN         = 0x00002002
	cKM_VENDOR_DEFINED                 = 0x80000000
)
