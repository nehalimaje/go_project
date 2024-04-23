package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func createFile(filename string) error {
	// Create the file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func insertStringIntoFile(filename, content string) error {
	// Append content to the file
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

func updateStringInFile(filename, oldContent, newContent string) error {
	// Read the file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// Replace oldContent with newContent
	updatedContent := strings.ReplaceAll(string(content), oldContent, newContent)

	// Write the updated content back to the file
	err = ioutil.WriteFile(filename, []byte(updatedContent), 0644)
	if err != nil {
		return err
	}
	return nil
}

func deleteFile(filename string) error {
	// Delete the file
	if err := os.Remove(filename); err != nil {
		return err
	}
	return nil
}

func main() {
	filename := "example1940.txt"

	// Create a file
	if err := createFile(filename); err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	// Insert string into the file
	if err := insertStringIntoFile(filename, "Hello, World!\n"); err != nil {
		fmt.Println("Error inserting string into file:", err)
		return
	}

	// Update string in the file
	if err := updateStringInFile(filename, "Hello", "Hi"); err != nil {
		fmt.Println("Error updating string in file:", err)
		return
	}

	// Delete the file
	if err := deleteFile(filename); err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}

	fmt.Println("File operations completed successfully.")
}
