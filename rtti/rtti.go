package rtti

import (
	"io/ioutil"
	"math/rand"
	"regexp"
	"time"
)

func RunOnData(data []uint8) ([]uint8, error) {
	rand.Seed(time.Now().UnixNano())

	// find strings starting with .?AV and ending with @@
	r, err := regexp.Compile("\\.\\?AV(.+?)@@\\0")
	if err != nil {
		return nil, err
	}

	obfuscatedData := r.ReplaceAllFunc(data,
		func(occurence []byte) []byte {
			// start replacing characters after the first 4 characters (.?AV)
			// during this process, skip '@' symbols

			chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
				"abcdefghijklmnopqrstuvwxyzåäö" +
				"0123456789")

			for i, b := range occurence[4 : len(occurence)-1] {
				if b == '@' {
					continue
				}

				occurence[4+i] = byte(chars[rand.Intn(len(chars))])
			}

			return occurence
		})

	return obfuscatedData, nil
}

func RunOnFile(filePath string) error {
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	obfuscatedData, err := RunOnData(fileData)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath+".nor", obfuscatedData, 0644)
}
