package main

import (
	"fmt"
	"os"
)

func main() {
	// Name of the new folder
	folderName := "my_new_folder"

	// Create a new folder with permission 0755
	err := os.Mkdir(folderName, 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return
	}
	fmt.Println("Folder created successfully:", folderName)

	// Optional: Create a new file inside the folder
	filePath := folderName + "/example.txt"
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Println("File created successfully:", filePath)
}
