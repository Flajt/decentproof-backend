package main

import (
	"bufio"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"flag"
	"fmt"
	"os"
)

func main() {
	var signature string
	var hash string
	var pubKeyFilePath string
	flag.StringVar(&signature, "signature", "", "Your signature")
	flag.StringVar(&hash, "hash", "", "Your hash")
	flag.StringVar(&pubKeyFilePath, "pbKey", "./pub.pem", "Public Key file path (.pem)")
	flag.Parse()
	err := verifySignature(signature, hash, pubKeyFilePath)
	if err != nil {
		fmt.Printf("Something went wrong %d", err)
		return
	}
}

func verifySignature(signature string, hash string, pubKeyFilePath string) error {
	file, err := os.Open(pubKeyFilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	pemFileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	var size int64 = pemFileInfo.Size()
	pemBytes := make([]byte, size)
	buffer := bufio.NewReader(file)
	_, err = buffer.Read(pemBytes)
	if err != nil {
		return err
	}
	//parse pem file
	for block, rest := pem.Decode(pemBytes); block != nil; block, rest = pem.Decode(rest) {
		if block.Type == "EC PUBLIC KEY" || block.Type == "PUBLIC KEY" || block.Type == "EC Public KEY" {
			pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
			if err != nil {
				return err
			}
			ecdsaPubKey, ok := pubKey.(*ecdsa.PublicKey)
			if !ok {
				return fmt.Errorf("failed to parse ECDSA public key")
			}
			decodeSig, err := hex.DecodeString(signature)
			if err != nil {
				return err
			}
			bytes, err := hex.DecodeString(hash)
			isValid := ecdsa.VerifyASN1(ecdsaPubKey, bytes, decodeSig)
			fmt.Printf("\nIs signature valid: %v \n", isValid)
			fmt.Println()
		}
	}
	return nil
}
