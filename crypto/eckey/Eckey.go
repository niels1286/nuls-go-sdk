package eckey

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

//用于存储公私钥对的数据结构
type Eckey struct {
	PriKey           string
	PubKey           string
	CompressedPubKey string
}

func NewEckey() Eckey {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	var compressedPubKeyBytes []byte
	x, y := elliptic.Unmarshal(crypto.S256(), publicKeyBytes)

	val := y.Bit(y.BitLen() - 1)
	if val == 0 {
		compressedPubKeyBytes = append([]byte{2}, x.Bytes()...)
	} else {
		compressedPubKeyBytes = append([]byte{3}, x.Bytes()...)
	}

	return Eckey{
		PriKey:           hex.EncodeToString(privateKeyBytes),
		PubKey:           hex.EncodeToString(publicKeyBytes),
		CompressedPubKey: hex.EncodeToString(compressedPubKeyBytes),
	}
}

//根据私钥，得到包含公钥的eckey
func FromPriKey(prikey string) Eckey {
	key, err := crypto.HexToECDSA(prikey)
	if err != nil {
		log.Println("Import prikey failed :" + prikey)
		return Eckey{}
	}
	privateKeyBytes := crypto.FromECDSA(key)
	publicKey := key.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	var compressedPubKeyBytes []byte
	x, y := elliptic.Unmarshal(crypto.S256(), publicKeyBytes)

	fmt.Println("xxxx:", hex.EncodeToString(x.Bytes()))
	fmt.Println("yyyy:", hex.EncodeToString(y.Bytes()))

	val := y.Bit(y.BitLen() - 1)
	if val%2 == 0 {
		compressedPubKeyBytes = append([]byte{2}, x.Bytes()...)
	} else {
		compressedPubKeyBytes = append([]byte{3}, x.Bytes()...)
	}

	return Eckey{
		PriKey:           hex.EncodeToString(privateKeyBytes),
		PubKey:           hex.EncodeToString(publicKeyBytes),
		CompressedPubKey: hex.EncodeToString(compressedPubKeyBytes),
	}
}

func FromPubKey(pubkey string) Eckey {
	return Eckey{}
}

func (eckey Eckey) String() string {
	bytes, err := json.Marshal(eckey)
	if nil != err {
		log.Fatalf("wrong.")
	}
	return string(bytes)
}
