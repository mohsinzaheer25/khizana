package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, password string) []byte {
	block, err := aes.NewCipher([]byte(createHash(password)))

	if err != nil {
		fmt.Println(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	//	Encrypt the data
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte, password string) []byte {
	block, err := aes.NewCipher([]byte(createHash(password)))
	if err != nil {
		fmt.Println(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err)
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println("Authentication Failed")
		os.Exit(1)
	}
	return plaintext
}

func EncryptFile(filename string, data []byte, password string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Failed To Create File")
		os.Exit(1)
	}
	defer f.Close()
	_, err = f.Write(encrypt(data, password))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func DecryptFile(filename string, password string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return decrypt(data, password)
}
