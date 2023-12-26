package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type dataType struct {
	Data []int `json:"data"`
}

var folderName = "data"

func main() {
	startNow := time.Now()
	fpath := filepath.Join(folderName, "data.json")

	data, err := readFile(fpath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// c := make(chan int)
	d := make(chan []int)
	// go sum(a[:len(a)/2], c)
	// go sum(a[len(a)/2:], c)
	// x, y := <-c, <-c
	// fmt.Println("x", x, "y", y)

	fpath = filepath.Join(folderName+"/output", "data.json")
	go saveFile(fpath, data.Data[:len(data.Data)/2], d)
	go saveFile(fpath, data.Data[len(data.Data)/2:], d)
	data1, data2 := <-d, <-d

	fmt.Println("data1", data1, "data2", data2)
	fmt.Println("took: ", time.Since(startNow))
	// saveFile(fpath, data)

}

func sum(a []int, c chan int) {
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}

	fmt.Println(a)

	c <- sum
}

func readFile(fpath string) (dataType, error) {
	fmt.Println("reading... ", fpath)
	saveFile, err := os.Open(fpath)
	if err != nil {
		fmt.Println("error opening file")
		return dataType{}, err
	}
	defer saveFile.Close()

	decoder := json.NewDecoder(saveFile)
	var data dataType
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("error decoding file")
		return dataType{}, err
	}

	return data, nil
}

func saveFile(fpath string, data []int, d chan []int) {

	d <- data
	// saveFile, err := os.Open(fpath)
	// if err != nil {
	// 	fmt.Println("error opening save file")
	// 	return err
	// }
	// defer saveFile.Close()

	// encoder := json.NewEncoder(saveFile)
	// err = encoder.Encode(data)
	// if err != nil {
	// 	fmt.Println("error encoding data to json")
	// 	return err
	// }

	// return err
}
