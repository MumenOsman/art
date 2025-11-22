package main

import (
	"fmt"
	"os"

	h "art/helpers" // h is the alias for the package helpers
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Error")
		return
	}

	isMulti := false
	isEncode := false
	var uInput string

	for _, arg := range args { // the loop here goes through the range of args so with every round it checks the next argument, so in the case of multi encode it will still reads both
		if arg == "--help" || arg == "-h" {
			fmt.Println(`Usage: go run . [FLAGS] "<ENCODED_OR_RAW_ART>"

This tool converts compressed art data (like "[5 #]A") into text-based art, 
or converts raw text-art back into the compressed format.

Flags (Modes):
  --encode (-e)   Switches the tool to text-art encoding mode (compression).
  --multi (-m)    Enables multi-line processing. Input is split by the newline (\n) character 
                  and each line is processed independently.

Default Behavior:
If no flags are provided, the tool runs in single-line **Decode** mode.

Examples:
  # Decode a single line of art data:
  go run . "[5 #][2 -_]-"
  
  # Decode a multi-line resource:
  go run . --multi "$(cat art.encoded.txt)"
  
  # Encode a raw art string:
  go run . --encode "AAAAABBB"`)
			return
		} else if arg == "--encode" {
			isEncode = true
		} else if arg == "--multi" {
			isMulti = true
		} else {
			uInput = arg
		}
	}

	if uInput == "" {
		fmt.Println("Error")
		return
	}

	var output string
	var err error

	if isEncode {
		if isMulti {
			output, err = h.MultiLineEncode(uInput)
		} else {
			output, err = h.SingleEncode(uInput)
		}
	} else {
		if isMulti {
			output, err = h.MultiDecode(uInput)
		} else {
			output, err = h.SingleDecode(uInput)
		}
	}

	if err != nil {
		fmt.Println("Error")
	} else {
		fmt.Println(output)
	}
}
