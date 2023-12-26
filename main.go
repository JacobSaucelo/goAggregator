package main

import (
	"fmt"
	"path/filepath"
	"time"
)

var folderName = "data"

func main() {
	startNow := time.Now()
	a := []int{1, 43, 65, 87, 12, 65, 8, 1, -1, 54, 0, 93, -12}

	// c := make(chan int)
	d := make(chan []int)
	// go sum(a[:len(a)/2], c)
	// go sum(a[len(a)/2:], c)
	// x, y := <-c, <-c
	// fmt.Println("x", x, "y", y)

	fpath := filepath.Join(folderName, "data.json")
	go saveFile(fpath, a[:len(a)/2], d)
	go saveFile(fpath, a[len(a)/2:], d)
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

func saveFile(fpath string, data []int, d chan []int) {

	fmt.Println("fpath: ", fpath)
	fmt.Println("data:", data)

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
