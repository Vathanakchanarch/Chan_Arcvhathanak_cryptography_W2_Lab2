
package main 

import (
	"fmt"
	"encoding/base64"
)

func encodeing(text string){
	fmt.Print("Binary: ")
	for i := 0; i < len(text); i++ {
        fmt.Printf("%08b ", text[i])
    }
	fmt.Printf("\nHexadecimal: ")
	for i := 0; i < len(text); i++ {
		fmt.Printf("%02x ", text[i])
	}
	base64Encoded := base64.StdEncoding.EncodeToString([]byte(text))
	fmt.Println("\nBase64 Encoded: ", base64Encoded)

}

func main() {
	var text string
	fmt.Print("Enter text to encode: ")
	fmt.Scan(&text)
	encodeing(text)
}