package main

import (
	"io/ioutil"
	"log"
)

func readFile(fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func writeFile(fileName string, ndung string) {
	err := ioutil.WriteFile(fileName, []byte(ndung), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// read file
	// fileName := "data.txt"
	// content := readFile(fileName)
	// fmt.Println(content)

	// write file
	fileName := "data2.txt"
	ndung := `[{"name":"nghia","age":20},{"name":"nguyen","age":21}]`
	writeFile(fileName, ndung)
}
