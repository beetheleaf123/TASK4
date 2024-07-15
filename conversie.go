package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	inputFile := "mat.in"

	start := time.Now()
	if strings.HasSuffix(inputFile, ".x") {
		convertToMatInWithoutCache(inputFile)
	} else {
		convertToMatInXWithoutCache(inputFile)
	}
	elapsed := time.Since(start)
	fmt.Printf("Non-cached version took %s\n", elapsed)
}

func convertToMatInWithoutCache(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	outputFileName := strings.TrimSuffix(inputFile, ".x")
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(outputFile)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		header := parts[0]
		hexData := parts[1]

		binaryData, err := hex.DecodeString(hexData)
		if err != nil {
			fmt.Println("Error decoding hex:", err)
			return
		}
		binaryString := fmt.Sprintf("%08b", binaryData)

		writer.WriteString(header + ":" + binaryString + "\n")
	}

	writer.Flush()
}

func convertToMatInXWithoutCache(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	outputFileName := inputFile + ".x"
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(outputFile)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		header := parts[0]
		binaryData := parts[1]

		var hexData string
		for i := 0; i < len(binaryData); i += 4 {
			end := i + 4
			if end > len(binaryData) {
				end = len(binaryData)
			}
			nibble := binaryData[i:end]
			nibbleByte, err := strconv.ParseUint(nibble, 2, 8)
			if err != nil {
				fmt.Println("Error converting binary to hex:", err)
				return
			}
			hexData += fmt.Sprintf("%X", nibbleByte)
		}

		writer.WriteString(header + ":" + hexData + "\n")
	}

	writer.Flush()
}
