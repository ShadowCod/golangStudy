package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	str := "MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDitaM+5F19XpHnKYqaCVEcpsAkGmQe6sJR0nlyg6q7TnIwo0/bdTftoQkkyiPx4xE1st14+BL6X8GxobhbVYRY7s4DHFa9cgjHUt0gZ4YXPLiMBrrv0upbwTowZk+ChGY7SQ8j3gxuVgDgEC4csB5AnWRvmFPwHF+up4hnSNZxueo10k2GnVQ6zMKdG1PIC/j5mA9+LYx367pngbgsDDMnPlprXCduKgvf+fe2abwKBqdGIJ2kBeI6P14MHzn/6NPjS7GxJTWlfXUKiqOW/zD5uPSncdAnyZAY7qAhUqy8GnDsQBgtYtcm3h9MFb3N1pKc7E2j9yD2y0wlRfhZfJenAgMBAAECgf9w6QXkzNfDGNaRyp69GObM3j/Tu1EiFMmy9qCF6A2gO/GjR7L+GeNA9nMXX5WHYw+vLuYe2W6aRBaz4VdwJadEo6r24aZRs/mrTBnVibm+2P8QvklcKC+Qgyv37vyqhEZmQHzJ4E1QpoTJE0XkA9MNW1YqPjOP4pHE6quExUIIYR/MWn+oM8FLyJ0a8JF3uiEFCFI/SQQ4nG4dGpWy9AAfYzcXOzq/PqawpUcAWQKb3l48rhcxavkxZHT8bypymFAORaVt4uPaqNq4AAnfTq1FANNm4UQK3KdjeQqTzH/Qi8HgwTl/ttuArifd5zVwb7Hv5sM1rRvEb5odPIP+IW0CgYEA8aEL8ab/a9qiYyIrN1kueotYv0dozBRsc4eVb1ytZVoR5hIXlNza5fIWHDAcS0poYjyR3TzkNS5OhGVwxGQXl7qlItk6kNCuc15Dhae7y7nbGT8cbYlDSJKZ617fnw05PhfIe73NS8cVJ+PaKQpsKpIXkcTp6bG/ZlFO+Xwb3ZsCgYEA8DFua+pBuTHoS5/5zU8mlgeF6D+XC6uGACYhGAN/iGD5lISLNT8KFP5cZMRLaM3qnSuYIgm/9hxF4e0spgi4nj2QEPL1C8ChaVI3ip1ASPw+CshMA2cNaii2Oi4wpAsmUxu4qPo/fuQ6F3AlV9cz7NBo6dQhqNODoHTqsBlR1OUCgYEApENQhFp8H81FQdFzwa/OCh77GSYplzt3Mt8Edu1xL4qYymiYQWoXIeV+pHEMex69cbMtklKX7ZNqa1Uu4UQOSebn5pX424QSZHm7u3v5DhluWm5uHAUJiaeoHbukmFL8DDtM3tp8WzyfJDwhDm94c0RX45ATWPyuWpeB8dcudmsCgYEAx4p5ToOdOCpC7linoS2pQ6haUXhKlnJXb1Y91gJJ99WAYia+s2x8hrZNsZT1hMrUpt+pklWBOQeB8tAjcIf8P9GrFrmQY8QTFDkuVSSQXFHZhQGjTIxXM6NAyBLJa+6rVw3HmfHTwCoALKqJC3GH/KujDOajU+rsBeg7dDKErhUCgYBCTR6dIKFq3vZWvjtBRbsdU1OUL7UfMXwBNfBFShqMpw6U4sja0nJts0wdWLaU8SSEihnKXATz8x7SvtbAMUq6IJykDOX6jWpjMpm9S8JB7CaZAAzKFcLaXLMPbOXuvGnu9wExQQZ1LQYlze+c1uoVIyt/1nOETfT48fpSnuxwZg=="
	r := strings.NewReader(str)
	pemBytes, err := ioutil.ReadAll(r)
	fmt.Println("pemBytes", pemBytes)
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode(pemBytes)
	if block == nil {
		fmt.Println("block", block)
	}
	// data := []byte(str)
	// // fmt.Println("data", data)
	// block := &pem.Block{
	// 	Type:  "RSA PRIVATE KEY",
	// 	Bytes: data,
	// }
	// prvkey := pem.EncodeToMemory(block)
	// fmt.Println("prvKey", string(prvkey))
	// str2 := string(data[:])
	// fmt.Println("str2", str2)

	//rsa 密钥文件产生
	// fmt.Println("-------------------------------获取RSA公私钥-----------------------------------------")
	// prvKey, _ := GenRsaKey()
	// fmt.Println("prvKey", string(prvKey))
	// fmt.Println(string(pubKey[:]))

	// fmt.Println("-------------------------------进行签名与验证操作-----------------------------------------")
	// var data1 = "卧了个槽，这么神奇的吗？？！！！  ԅ(¯﹃¯ԅ) ！！！！！！）"
	// fmt.Println("对消息进行签名操作...")
	// signData := RsaSignWithSha256([]byte(data1), prvkey)
	// fmt.Println("消息的签名信息： ", hex.EncodeToString(signData))
	// fmt.Println("\n对签名信息进行验证...")
	// if RsaVerySignWithSha256([]byte(data), signData, pubKey) {
	// 	fmt.Println("签名信息验证成功，确定是正确私钥签名！！")
	// }

	// fmt.Println("-------------------------------进行加密解密操作-----------------------------------------")
	// ciphertext := RsaEncrypt([]byte(data), pubKey)
	// fmt.Println("公钥加密后的数据：", hex.EncodeToString(ciphertext))
	// sourceData := RsaDecrypt(ciphertext, prvKey)
	// fmt.Println("私钥解密后的数据：", string(sourceData))
}

//RSA公钥私钥产生
func GenRsaKey() (prvkey, pubkey []byte) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	derStream := []byte("MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDitaM+5F19XpHnKYqaCVEcpsAkGmQe6sJR0nlyg6q7TnIwo0/bdTftoQkkyiPx4xE1st14+BL6X8GxobhbVYRY7s4DHFa9cgjHUt0gZ4YXPLiMBrrv0upbwTowZk+ChGY7SQ8j3gxuVgDgEC4csB5AnWRvmFPwHF+up4hnSNZxueo10k2GnVQ6zMKdG1PIC/j5mA9+LYx367pngbgsDDMnPlprXCduKgvf+fe2abwKBqdGIJ2kBeI6P14MHzn/6NPjS7GxJTWlfXUKiqOW/zD5uPSncdAnyZAY7qAhUqy8GnDsQBgtYtcm3h9MFb3N1pKc7E2j9yD2y0wlRfhZfJenAgMBAAECgf9w6QXkzNfDGNaRyp69GObM3j/Tu1EiFMmy9qCF6A2gO/GjR7L+GeNA9nMXX5WHYw+vLuYe2W6aRBaz4VdwJadEo6r24aZRs/mrTBnVibm+2P8QvklcKC+Qgyv37vyqhEZmQHzJ4E1QpoTJE0XkA9MNW1YqPjOP4pHE6quExUIIYR/MWn+oM8FLyJ0a8JF3uiEFCFI/SQQ4nG4dGpWy9AAfYzcXOzq/PqawpUcAWQKb3l48rhcxavkxZHT8bypymFAORaVt4uPaqNq4AAnfTq1FANNm4UQK3KdjeQqTzH/Qi8HgwTl/ttuArifd5zVwb7Hv5sM1rRvEb5odPIP+IW0CgYEA8aEL8ab/a9qiYyIrN1kueotYv0dozBRsc4eVb1ytZVoR5hIXlNza5fIWHDAcS0poYjyR3TzkNS5OhGVwxGQXl7qlItk6kNCuc15Dhae7y7nbGT8cbYlDSJKZ617fnw05PhfIe73NS8cVJ+PaKQpsKpIXkcTp6bG/ZlFO+Xwb3ZsCgYEA8DFua+pBuTHoS5/5zU8mlgeF6D+XC6uGACYhGAN/iGD5lISLNT8KFP5cZMRLaM3qnSuYIgm/9hxF4e0spgi4nj2QEPL1C8ChaVI3ip1ASPw+CshMA2cNaii2Oi4wpAsmUxu4qPo/fuQ6F3AlV9cz7NBo6dQhqNODoHTqsBlR1OUCgYEApENQhFp8H81FQdFzwa/OCh77GSYplzt3Mt8Edu1xL4qYymiYQWoXIeV+pHEMex69cbMtklKX7ZNqa1Uu4UQOSebn5pX424QSZHm7u3v5DhluWm5uHAUJiaeoHbukmFL8DDtM3tp8WzyfJDwhDm94c0RX45ATWPyuWpeB8dcudmsCgYEAx4p5ToOdOCpC7linoS2pQ6haUXhKlnJXb1Y91gJJ99WAYia+s2x8hrZNsZT1hMrUpt+pklWBOQeB8tAjcIf8P9GrFrmQY8QTFDkuVSSQXFHZhQGjTIxXM6NAyBLJa+6rVw3HmfHTwCoALKqJC3GH/KujDOajU+rsBeg7dDKErhUCgYBCTR6dIKFq3vZWvjtBRbsdU1OUL7UfMXwBNfBFShqMpw6U4sja0nJts0wdWLaU8SSEihnKXATz8x7SvtbAMUq6IJykDOX6jWpjMpm9S8JB7CaZAAzKFcLaXLMPbOXuvGnu9wExQQZ1LQYlze+c1uoVIyt/1nOETfT48fpSnuxwZg==")
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	prvkey = pem.EncodeToMemory(block)
	// fmt.Println(prvkey)
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pubkey = pem.EncodeToMemory(block)
	return
}

//签名
func RsaSignWithSha256(data []byte, keyBytes []byte) []byte {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("private key error"))
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("ParsePKCS8PrivateKey err", err)
		panic(err)
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		panic(err)
	}

	return signature
}

//验证
func RsaVerySignWithSha256(data, signData, keyBytes []byte) bool {
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("public key error"))
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	hashed := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], signData)
	if err != nil {
		panic(err)
	}
	return true
}

// 公钥加密
func RsaEncrypt(data, keyBytes []byte) []byte {
	//解密pem格式的公钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("public key error"))
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, data)
	if err != nil {
		panic(err)
	}
	return ciphertext
}

// 私钥解密
func RsaDecrypt(ciphertext, keyBytes []byte) []byte {
	//获取私钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("private key error!"))
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 解密
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	if err != nil {
		panic(err)
	}
	return data
}
