package btc

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"log"
	"math/big"
)

type EcKey struct {
	privateKey       *ecdsa.PrivateKey
	publicKey        []byte
	compressedPubKey []byte
}

func NewEcKey() (EcKey, error) {
	curve := secp256k1.S256()
	var err error
	pk, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	return getEcKeyByPriKey(pk)
}

var (
	secp256k1N, _  = new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	secp256k1halfN = new(big.Int).Div(secp256k1N, big.NewInt(2))
)

func FromPriKey(priHex string) (EcKey, error) {
	d, err := hex.DecodeString(priHex)
	if nil != err {
		return EcKey{}, err
	}
	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = secp256k1.S256()
	if 8*len(d) != priv.Params().BitSize {
		return EcKey{}, fmt.Errorf("invalid length, need %d bits", priv.Params().BitSize)
	}
	priv.D = new(big.Int).SetBytes(d)

	// The priv.D must < N
	if priv.D.Cmp(secp256k1N) >= 0 {
		return EcKey{}, fmt.Errorf("invalid private key, >=N")
	}
	// The priv.D must not be zero or negative.
	if priv.D.Sign() <= 0 {
		return EcKey{}, fmt.Errorf("invalid private key, zero or negative")
	}

	priv.PublicKey.X, priv.PublicKey.Y = priv.PublicKey.Curve.ScalarBaseMult(d)
	if priv.PublicKey.X == nil {
		return EcKey{}, fmt.Errorf("invalid private key")
	}
	return getEcKeyByPriKey(priv)
}

func getEcKeyByPriKey(pk *ecdsa.PrivateKey) (EcKey, error) {
	xbytes := pk.PublicKey.X.Bytes()
	ybytes := pk.PublicKey.Y.Bytes()
	pubKey := elliptic.Marshal(secp256k1.S256(), pk.PublicKey.X, pk.PublicKey.Y)
	var compressedPubKey []byte
	if ybytes[len(ybytes)-1]%2 == 0 {
		compressedPubKey = append([]byte{2}, xbytes...)
	} else {
		compressedPubKey = append([]byte{3}, xbytes...)
	}
	return EcKey{pk, pubKey, compressedPubKey}, nil
}

func (e EcKey) GetPriKeyHex() string {
	if e.privateKey == nil {
		return ""
	}
	bytes := PaddedBigBytes(e.privateKey.D, e.privateKey.Params().BitSize/8)

	return hex.EncodeToString(bytes)
}
func (e EcKey) GetPubKeyHex(compressed bool) string {
	if compressed {
		return hex.EncodeToString(e.compressedPubKey)
	}
	return hex.EncodeToString(e.publicKey)
}

const (
	// number of bits in a big.Word
	wordBits = 32 << (uint64(^big.Word(0)) >> 63)
	// number of bytes in a big.Word
	wordBytes = wordBits / 8
)

// PaddedBigBytes encodes a big integer as a big-endian byte slice. The length
// of the slice is at least n bytes.
func PaddedBigBytes(bigint *big.Int, n int) []byte {
	if bigint.BitLen()/8 >= n {
		return bigint.Bytes()
	}
	ret := make([]byte, n)
	ReadBits(bigint, ret)
	return ret
}

// ReadBits encodes the absolute value of bigint as big-endian bytes. Callers must ensure
// that buf has enough space. If buf is too short the result will be incomplete.
func ReadBits(bigint *big.Int, buf []byte) {
	i := len(buf)
	for _, d := range bigint.Bits() {
		for j := 0; j < wordBytes && i > 0; j++ {
			i--
			buf[i] = byte(d)
			d >>= 8
		}
	}
}
