package ecdsa

import (
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"github.com/wumansgy/goEncrypt"
	"os"
)

const (
	privateKey = `MHcCAQEEIKozbXD9G6bGPJ26cCAfEdLrqAe697F8SiLRMdqxzNQ5oAoGCCqGSM49
AwEHoUQDQgAEk3/hltyR0r0J2Wkkhi4HaREJXS1vFooGpsKCbLvrdUW4peVIwKEW
+yC3/g2X7Q2A8ftJlYv2X4kDU180GhIQpA==`
	publicKey = `MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEk3/hltyR0r0J2Wkkhi4HaREJXS1v
FooGpsKCbLvrdUW4peVIwKEW+yC3/g2X7Q2A8ftJlYv2X4kDU180GhIQpA==`
)



const Kbegin = `-----BEGIN WUMAN ECC PRIVATE KEY-----
`
const Kend  = `
-----END WUMAN ECC PRIVATE KEY-----`


const privateKey2  =  `KyqxHcuh1Nam5ZevTdPcu8n4deg43SQhrFDP5iXCUM9qJWsq2Hjr`
const publicKey2  = `036f12ab6dd249fd1f381722948862f30c6a13c197e59a6fc4c88d638ba987f1d2`

const eccPrivateKeyPrefix  = "lee pri xy"
const eccPublicKeyPrefix  = "lee pub xy"
const eccPrivateFileName = "pri.pem"
const eccPublishFileName = "pub.pem"

func TransferEccKey(pri,pub string)([]byte,[]byte,error){

	prik,err := decode58(pri)
	if err!=nil{
		return nil,nil,err
	}

	block := pem.Block{
		Type:  eccPrivateKeyPrefix,
		Bytes: prik,
	}

	prikey := pem.EncodeToMemory(&block)


	pubk,err := hex.DecodeString(pub)
	if err!=nil{
		return nil,nil,err
	}

	publicBlock := pem.Block{
		Type:  eccPublicKeyPrefix,
		Bytes: pubk,
	}


	pubkey := pem.EncodeToMemory(&publicBlock)



	return prikey,pubkey,nil
}

func decode58(szPri string)([]byte,error){
	return Decode(szPri,BitcoinAlphabet)
}


func GetEccKey()error{

	prikey2 := privateKey2
	prik,err := decode58(prikey2)

	block := pem.Block{
		Type:  eccPrivateKeyPrefix,
		Bytes: prik,
	}
	file, err := os.Create(eccPrivateFileName)
	if err!=nil{
		return err
	}
	defer file.Close()
	if err=pem.Encode(file, &block);err!=nil{
		return err
	}

	pubk,err := hex.DecodeString(publicKey2)
	if err!=nil{
		return err
	}
	publicBlock := pem.Block{
		Type:  eccPublicKeyPrefix,
		Bytes: pubk,
	}
	publicFile, err := os.Create(eccPublishFileName)
	if err!=nil {
		return err
	}
	defer publicFile.Close()
	if err=pem.Encode(publicFile,&publicBlock);err!=nil{
		return err
	}
	return nil
}

func Myde() {
	GetEccKey()
	//prikey,pubkey,err := GetEccKey(privateKey2,publicKey2)
	_,_,err := TransferEccKey(privateKey2,publicKey2)



	plainText := []byte("窗前明月光，疑是地上霜,ECC加密解密")

	// 这里传入的私钥和公钥是要用GetECCKey里面得到的私钥和公钥，如果自己封装的话，
	// 获取密钥时传入的第一个参数是要用这条曲线elliptic.P256()，如果用别的会报无效公钥错误，
	// 例如用P521()这条曲线
	privateKey := []byte(Kbegin+privateKey+Kend)
	publicKey := []byte(Kbegin+publicKey+Kend)

	cryptText, _ := goEncrypt.EccEncrypt(plainText, publicKey)
	fmt.Println("ECC传入公钥加密的密文为：", hex.EncodeToString(cryptText))

	msg, err := goEncrypt.EccDecrypt(cryptText, privateKey)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ECC传入私钥解密后的明文为：", string(msg))
}

