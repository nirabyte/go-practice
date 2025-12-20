package main

import (
	"os"
)

func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	//log the error
	// 	panic(err)
	// }
	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	//log the error
	// 	panic(err)
	// }
	// fmt.Println("file name: ", fileInfo.Name())
	// fmt.Println("file or folder: ", fileInfo.IsDir())
	// fmt.Println("file or folder: ", fileInfo.Size())
	// fmt.Println("file or folder: ", fileInfo.Mode())
	// fmt.Println("file or folder: ", fileInfo.ModTime())

	// read file

	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// buf := make([]byte, 12)
	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }

	//common method
	// for i := 0; i < len(buf); i++ {
	// 	println("data", d, string(buf[i]))
	// }

	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))

	//read folders

	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()
	// fileInfo, err := dir.ReadDir(-1)
	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	// create a file
	f, err := os.Create("example2.txt")
	if err != nil {
		panic(err)

	}
	defer f.Close()
	f.WriteString("this is it")

}
