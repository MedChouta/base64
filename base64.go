
package main

import (
	b64 "encoding/base64"
	"os"
	"io/ioutil"
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

func checkError(e error) bool{
	if e != nil{
		panic(e)
		return false
	}

	return true
}

func EDfile(file string, option int) bool{

	inputData, err := ioutil.ReadFile(file)
	state := checkError(err)

	newFileName := strings.Join(strings.Split(file, "."), "_O.")
	outputFile, err := os.Create(newFileName)
	state = checkError(err)


	/*
	option 1: encode
	option 2: decode
	*/
							
	var outputData []byte

	if option == 1{
		outputData = []byte(encode(string(inputData)))
	}else if option == 2{
		outputData = []byte(decode(string(inputData)))
	}

	err = ioutil.WriteFile(newFileName, outputData, 0644)

	state = checkError(err)

	defer outputFile.Close()

	return state

}

func showHelp() string{
	return `
Usage: 
	base64 [option] string or file 
Options: 
	-e(--encode): Encodes the string
	-d(--decode): Decodes the string
	-Ef(--encodeFile): Encodes file
	-Df(--decodeFile): Decodes file
`
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
		}else if option == "-Ef" || option == "--encodeFile"{
			isSuccessful := EDfile(input, 1)
			if isSuccessful{
				fmt.Print("Encoding successful")
			}else{
				fmt.Print("An error has occured")
			}
		}else if option == "-Df" || option == "--decodeFile" {
			isSuccessful := EDfile(input, 2)
			if isSuccessful{
				fmt.Print("Decoding successful")
			}else{
				fmt.Print("An error has occured")
			}
		}else{
			fmt.Print(showHelp())
		}
	}else{
		fmt.Print(showHelp())
	}
}
