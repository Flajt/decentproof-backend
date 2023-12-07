package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"os"
)

func main() {
	var outPutDir string
	var outPutFileName string
	flag.StringVar(&outPutDir, "out", "./", "Output directory")
	flag.StringVar(&outPutFileName, "name", "key", "Output file name e.g key, my_project_key etc. Adding .pem is not needed")
	flag.Parse()
	key, err := generateKeys()
	if err != nil {
		fmt.Printf("Something went wrong %d", err)
		return
	}
	err = saveKeys(key, outPutDir, outPutFileName)
	if err != nil {
		fmt.Printf("Something went wrong %d", err)
		return
	}
	fmt.Printf("Keys generated and saved to %s%s.pem and %spub_%s.pem", outPutDir, outPutFileName, outPutDir, outPutFileName)

}

func generateKeys() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}
func saveKeys(key *ecdsa.PrivateKey, outPutDir string, outPutFileName string) error {
	bytes, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return err
	}
	privBloc := pem.Block{Type: "EC PRIVATE KEY", Bytes: bytes}
	privKeyFile, err := os.Create(outPutDir + outPutFileName + ".pem")
	if err != nil {
		return err
	}
	defer privKeyFile.Close()
	err = pem.Encode(privKeyFile, &privBloc)
	if err != nil {
		return err
	}

	bytes, err = x509.MarshalPKIXPublicKey(&key.PublicKey)
	pubBloc := pem.Block{Type: "EC PUBLIC KEY", Bytes: bytes}
	pubKeyFile, err := os.Create(outPutDir + "pub_" + outPutFileName + ".pem")
	if err != nil {
		return err
	}
	defer pubKeyFile.Close()
	err = pem.Encode(pubKeyFile, &pubBloc)
	if err != nil {
		return err
	}

	return nil
}
