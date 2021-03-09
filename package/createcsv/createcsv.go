package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	file, err := os.Create("戒菸.csv") //go mod or go path
	if err != nil {
		log.Println(`create file error:`, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	bomUtf8 := []byte{0xEF, 0xBB, 0xBF}
	writer.Write([]string{string(bomUtf8[:])})
	writer.Flush()
}