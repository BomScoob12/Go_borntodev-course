package main

import (
	"os"
)

/*
0644 is a permission code on linux permission
r = 4
w = 2
x = 1
- = 0 (no permission)
*/
var permissionSet int16 = 0644

func customWriteFile() {
	inputData := []byte("Hello, BomScope")
	err := os.WriteFile("data.txt", inputData, os.FileMode(permissionSet))
	if err != nil {
		panic(err)
	}
}

func storeData(filename string, data ...string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, os.FileMode(permissionSet))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for i := 0; i < len(data); i++ {
		file.Write([]byte(data[i] + "\n"))
	}
}
func main() {
	customWriteFile()
	storeData("StudentName.txt", "bob", "Ryan", "Game")
}
