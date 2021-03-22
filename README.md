# go-away-rtti

go-away-rtti is a RTTI obfuscator written in GO

(the name is a joke because the language its written in is GO, so its like go away rtti as in no more rtti but written in go)

# Usage

```shell
go get github.com/DucaRii/go-away-rtti
```

Example:

```go
package main

import (
	"fmt"

	"github.com/DucaRii/go-away-rtti/rtti"
)

func main() {
	// this will spit out a file dummy.dll.nor where the RTTI is obfuscated
	err := rtti.RunOnFile("dummy.dll")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// alternatively, you may feed it raw data and get a copy of that data with the RTTI obfuscated
	dummyData := []byte{0x90, 0x90, 0x90, 0x90}

	obfuscatedData, err := rtti.RunOnData(dummyData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

```
