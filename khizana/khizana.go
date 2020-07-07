package khizana

import (
	"fmt"
	"khizana/crypto"
	"os"
	"strings"
)

func Create(khizanaPath string, password string) {
	crypto.EncryptFile(khizanaPath, []byte("# Khizana\n"), password)
	fmt.Println("\nKhizana has initialized successfully")
}

func View(khizanaPath string, password string) string {
	return string(crypto.DecryptFile(khizanaPath, password))
}

func Add(khizanaPath string, key string, value string, password string) {
	//	khizana add -k username -v mohsinzaheer25@hotmail.com

	oldData := View(khizanaPath, password)
	if strings.Contains(oldData, key) {
		fmt.Printf("\n%s key already exists. Please try another key name.\n", key)
	} else {
		newData := fmt.Sprintf("%s%s: %s\n", oldData, key, value)
		crypto.EncryptFile(khizanaPath, []byte(newData), password)
		fmt.Println("\nKey Value Added")
	}
}

func Get(khizanaPath string, key string, password string) {
	// khizana get username

	data := View(khizanaPath, password)
	splitData := strings.Split(data, "\n")
	var value string
	for _, pair := range splitData {
		splitKeyvalue := strings.SplitN(pair, ":", 2)
		if splitKeyvalue[0] == key {
			value = splitKeyvalue[1]
		}
	}
	if value != "" {
		fmt.Printf("\n%s\n", value)
	} else {
		fmt.Printf("\n%s key not found\n", key)
	}
}

func Update(khizanaPath string, key string, value string, password string) {
	// khizana update -k username -v mohsinzaheer25@hotmail.com

	data := View(khizanaPath, password)
	splitData := strings.Split(data, "\n")
	var findKey string
	for _, pair := range splitData {
		splitKeyvalue := strings.SplitN(pair, ":", 2)
		if splitKeyvalue[0] == key {
			findKey = splitKeyvalue[0]
			// Giving space between key: & value
			value = " " + value
			newData := strings.ReplaceAll(data, splitKeyvalue[1], value)
			crypto.EncryptFile(khizanaPath, []byte(newData), password)
		}
	}
	if findKey != "" {
		fmt.Printf("\n%s value updated\n", key)
	} else {
		fmt.Printf("\n%s key not found\n", key)
	}
}

func Delete(khizanaPath string, key string, password string) {
	// khizana update -k username

	data := View(khizanaPath, password)
	splitData := strings.Split(data, "\n")
	var findKey string
	for _, pair := range splitData {
		splitKeyvalue := strings.SplitN(pair, ":", 2)
		if splitKeyvalue[0] == key {
			findKey = splitKeyvalue[0]
			// Giving space between key: & value
			keyValue := fmt.Sprintf("%s:%s\n", splitKeyvalue[0], splitKeyvalue[1])
			newData := strings.ReplaceAll(data, keyValue, "")
			crypto.EncryptFile(khizanaPath, []byte(newData), password)
		}
	}
	if findKey != "" {
		fmt.Printf("\n%s key deleted\n", key)
	} else {
		fmt.Printf("\n%s key not found\n", key)
	}
}

func Destroy(khizanaPath string, password string) {
	data := View(khizanaPath, password)
	if data != "" {
		err := os.Remove(khizanaPath)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Khizana destroyed")
	}
}
