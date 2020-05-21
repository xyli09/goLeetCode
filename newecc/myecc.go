package newecc

import (
	"encoding/hex"
	"fmt"
)

func MyECC() {
	plainText := []byte("窗前明月光，疑是地上霜,ECC加密解密")
	fmt.Println("未加密原文：", string(plainText))

	//区块链钱包地址的私钥和公钥
	privateKey := []byte("KyqxHcuh1Nam5ZevTdPcu8n4deg43SQhrFDP5iXCUM9qJWsq2Hjr")
	publicKey := []byte("036f12ab6dd249fd1f381722948862f30c6a13c197e59a6fc4c88d638ba987f1d2")

	//公钥加密过程
	cryptText, _ := WealedgerEccEncrypt(plainText, publicKey)
	fmt.Println("区块链公钥：", string(publicKey))
	fmt.Println("ECC加密的密文：", hex.EncodeToString(cryptText))

	//私钥解密过程
	msg, err := WealedgerEccDecrypt(cryptText, privateKey)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("区块链私钥：", string(privateKey))
	fmt.Println("ECC解密后的明文：", string(msg))

}
