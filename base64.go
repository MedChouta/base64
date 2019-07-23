package main

import (
	b64 "encoding/base64"
	"os"
	"fmt"
	"strings"
)

func encode(input string) string{
	inputEnc := b64.StdEncoding.EncodeToString([]byte(input))
	
	return inputEnc
}

func decode(input string) string{
	inputDec, _ := b64.StdEncoding.DecodeString(input)

	return string(inputDec)
}

func showHelp() string{
	return "Usage: \n\tbase64 [option] string \nOptions: \n\t-e(--encode): Encodes the string\n\t-d(--decode): Decodes the string"
}

func main(){
	if len(os.Args) >= 2{

		input := strings.Join(os.Args[2:], " ")
		option := os.Args[1]


		if option == "-d" || option == "--decode" {
			decString := decode(input)
			fmt.Print(decString)
		} else if option == "-e" || option == "--encode"{
			encString := encode(input)
			fmt.Print(encString)
		}else{
			fmt.Print(showHelp())
		}
	}else{
		fmt.Print(showHelp())
	}
}
