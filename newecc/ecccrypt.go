package newecc

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
)

/*
@Time : 2018/11/4 16:43
@Author : wuman
@File : EccCrypt
@Software: GoLand
*/



func WealedgerEccEncrypt(plainText, key []byte) (cryptText []byte, err error) {

	tempchainpubk, err := hex.DecodeString(string(key))
	pubkey, err := crypto.DecompressPubkey(tempchainpubk)
	if err != nil {
		return nil, err
	}

	// Convert to the public key in the ecies package in the ethereum package
	pubkey2 := ecies.ImportECDSAPublic(pubkey)
	crypttext, err := ecies.Encrypt(rand.Reader, pubkey2, plainText, nil, nil)
	//crypttext, err := ECCEncrypt([]byte(plainText), *pubkey2)
	if err != nil {
		return nil, err
	}

	return crypttext, err
}

func WealedgerEccDecrypt(cryptText, key []byte) (msg []byte, err error) {
	b58 := DecodeBase58(string(key))
	prikey, err := crypto.ToECDSA(b58[1 : len(b58)-5])
	prikey2 := ecies.ImportECDSA(prikey)

	plainText, err := prikey2.Decrypt(cryptText, nil, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}
