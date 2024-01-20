package leaf

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func RemoveNote(directory string, reader io.Reader, writer io.Writer) error {
	fmt.Println("Which note would you like to remove?")

	var file_list []string
	file_list, err := walkTextFiles(directory, file_list, writer)
	if err != nil {
		fmt.Fprintln(writer, "Error walking files:", err)
		return err
	}

	printIndexedSlice(file_list, writer)

	choice, err := getSliceIndex(file_list, reader, writer)
	if err != nil {
		fmt.Fprintf(writer, "Error getting index: %v\n", err)
		return err
	}

	fmt.Fprintf(writer, "You chose index %d: %s \n", choice, file_list[choice])
	fmt.Fprintln(writer, "Are you sure you want to delete this note? (y/n)")

	if confirmDelete(reader, writer) {
		if err := os.Remove(file_list[choice]); err != nil {
			fmt.Fprintf(writer, "Error deleting note: %v\n", err)
			return err
		}
		fmt.Fprintln(writer, "Note deleted successfully")
	}
	return nil
}

func walkTextFiles(directory string, file_list []string, writer io.Writer) ([]string, error) {
	err := filepath.WalkDir(directory,
		func(path string, info fs.DirEntry, err error) error {
			if err != nil {
				fmt.Fprintln(writer, err)
				return err
			}
			if filepath.Ext(info.Name()) == ".txt" {
				file_list = append(file_list, path)
			}
			return nil
		})
	return file_list, err
}

func printIndexedSlice(slice []string, writer io.Writer) {
	for index, value := range slice {
		fmt.Fprintln(writer, index, " | name: ", value)
	}
}

func getSliceIndex(slice []string, reader io.Reader, writer io.Writer) (int, error) {
	var index int
	fmt.Fprintln(writer, "Which note would you like to remove? (enter the number)")
	_, err := fmt.Fscanln(reader, &index)
	if err != nil {
		return -1, err
	}
	if index < 0 || index >= len(slice) {
		return -1, fmt.Errorf("index out of range")
	}
	return index, err
}

// I could make this return error moving the error printing to the caller
func confirmDelete(reader io.Reader, writer io.Writer) bool {
	var answer string
	_, err := fmt.Fscanln(reader, &answer)
	if err != nil {
		fmt.Fprintln(writer, "Error reading input", err)
		return false
	}
	if answer == "y" {
		return true
	} else {
		return false
	}
}
