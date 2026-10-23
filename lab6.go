
package main

import "fmt"

func xorEncrpt(text string, key int) string {
	data := []byte(text)
	for i :=0; i< len(data); i++{
		data[i] = data[i] ^ byte(key)

	}

	return string(data)
}

func main(){
	var message string
	var key int

	fmt.Print("Enter message to encrypt: ")
	fmt.Scan(&message)
	
	fmt.Print("Enter key (integer): ")
	fmt.Scan(&key)

	encryptedMessage := xorEncrpt(message, key)
	decryptedMessage := xorEncrpt(encryptedMessage, key)
	
	fmt.Println("Encrypted Message: ", encryptedMessage)
	fmt.Println("Decrypted Message: ", decryptedMessage)	
	
}