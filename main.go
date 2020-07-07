package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"golang.org/x/crypto/ssh/terminal"
	"khizana/khizana"
	"log"
	"os"
	"os/user"
	"strings"
)

var password string

var helpmsg = []byte(`Khizana is a personal vault for your sensitive information.

Options:

init				Initialize Your Khizana
view				To View Your Khizana
get KEY				To get Value of The Key
add KEY VALUE			To Add Key Value To Your Khizana
update KEY VALUE		To Updated Value Of The Key
delete KEY			To Remove Key From Khizana
destroy				To Destroy Khizana
help				To Get Help Of Khizana
`)

var khizanaPath string


func init(){
	usr, _ := user.Current()
	khizanaPath  =  usr.HomeDir + "/.khizana"
}

func main(){
	if len(os.Args) < 2 {
		fmt.Println(string(helpmsg))
	} else {
		if os.Args[1] == "init" {
			if _, err := os.Stat(khizanaPath); os.IsNotExist(err) {
				fmt.Println("Enter Password For Your Khizana")
				//_, err := fmt.Scan(&password)
				bytePassword, err := terminal.ReadPassword(0)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				password = string(bytePassword)
				khizana.Create(khizanaPath, password)
			} else {
				fmt.Println("Khizana already initialized")
			}
		} else if os.Args[1] == "view" {
			if _, err := os.Stat(khizanaPath); os.IsNotExist(err) {
				fmt.Println("You must initialized Khizana before you view any key value")
			} else {
				fmt.Println("Enter Password For Your Khizana")
				bytePassword, err := terminal.ReadPassword(0)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				password = string(bytePassword)
				fmt.Println("\n", khizana.View(khizanaPath, password))
			}
		} else if os.Args[1] == "add" {
			if _, err := os.Stat(khizanaPath); os.IsNotExist(err) {
				fmt.Println("You must initialized Khizana before you add any key value")
			} else {
				if len(os.Args) < 4 {
					fmt.Println(string(helpmsg))
				} else {
					if strings.Contains(os.Args[2], ":") {
						fmt.Printf("Key can't contain ':', Please try without adding : \n")
					} else {
						fmt.Println("Enter Password For Your Khizana")
						bytePassword, err := terminal.ReadPassword(0)
						if err != nil {
							fmt.Println(err)
							os.Exit(1)
						}
						password = string(bytePassword)
						// Fix this cobra flags
						key := os.Args[2]
						value := os.Args[3]
						khizana.Add(khizanaPath, key, value, password)
					}
				}
			}
		} else if os.Args[1] == "get" {
			if _, err := os.Stat(khizanaPath); os.IsNotExist(err) {
				fmt.Println("You must initialized Khizana before you get any key value")
			} else {
				if len(os.Args) < 3 {
					fmt.Println(string(helpmsg))
				} else {
					if strings.Contains(os.Args[2], ":") {
						fmt.Printf("Key can't contain ':', Please try without adding : \n")
					} else {
						fmt.Println("Enter Password For Your Khizana")
						bytePassword, err := terminal.ReadPassword(0)
						if err != nil {
							fmt.Println(err)
							os.Exit(1)
						}
						password = string(bytePassword)
						key := os.Args[2]
						khizana.Get(khizanaPath, key, password)
					}
				}
			}
		} else if os.Args[1] == "update" {
			if _, err := os.Stat(khizanaPath); os.IsNotExist(err) {
				fmt.Println("You must initialized Khizana before you update any key value")
			} else {
				if len(os.Args) < 4 {
					fmt.Println(string(helpmsg))
				} else {
					if strings.Contains(os.Args[2], ":") {
						fmt.Printf("Key can't contain ':', Please try without adding : \n")
					} else {
						fmt.Println("Enter Password For Your Khizana")
						bytePassword, err := terminal.ReadPassword(0)
						if err != nil {
							fmt.Println(err)
							os.Exit(1)
						}
						password = string(bytePassword)
						// Fix this cobra flags
						key := os.Args[2]
						value := os.Args[3]
						khizana.Update(khizanaPath, key, value, password)
					}
				}
			}
		}else if os.Args[1] == "delete" {
			if _, err := os.Stat(khizanaPath); os.IsNotExist(err) {
				fmt.Println("You must initialized Khizana before you delete any key value")
			} else {
				if len(os.Args) < 3 {
					fmt.Println(string(helpmsg))
				} else {
					if strings.Contains(os.Args[2], ":") {
						fmt.Printf("Key can't contain ':', Please try without adding : \n")
					} else {
						fmt.Println("Enter Password For Your Khizana")
						bytePassword, err := terminal.ReadPassword(0)
						if err != nil {
							fmt.Println(err)
							os.Exit(1)
						}
						password = string(bytePassword)
						// Fix this cobra flags
						prompt := promptui.Select{
							Label: "Are you sure you want to delete",
							Items: []string{"Yes", "No"},
						}
						_, result, err := prompt.Run()
						if err != nil {
							log.Fatalf("%v\n", err)
						}
						if result == "Yes"{
							key := os.Args[2]
							khizana.Delete(khizanaPath, key, password)
						} else {
							fmt.Println("Bailing out now, bye.")
						}
					}
				}
			}
		}else if os.Args[1] == "destroy" {
			if _, err := os.Stat(khizanaPath); os.IsNotExist(err) {
				fmt.Println("You must initialized Khizana before you destroy it")
			} else {
				fmt.Println("Enter Password For Your Khizana")
				bytePassword, err := terminal.ReadPassword(0)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				password = string(bytePassword)
				// Fix this cobra flags
				prompt := promptui.Select{
					Label: "Are you sure you want to destroy Khizana",
					Items: []string{"Yes", "No"},
				}
				_, result, err := prompt.Run()
				if err != nil {
					log.Fatalf("%v\n", err)
				}
				if result == "Yes"{
					khizana.Destroy(khizanaPath, password)
				} else {
					fmt.Println("Khizana Destroy Aborted!")
				}
			}
		}else if os.Args[1] == "help" || os.Args[1] == "--help" || os.Args[1] == "-h" {
			fmt.Println(string(helpmsg))
		}
	}
}
