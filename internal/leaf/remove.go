package leaf

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func RemoveLeaf(journalPath string) {
	fmt.Println("Which leaf would you like to remove?")
	var file_list []string
	// List the leaves in the journal directory
	filepath.WalkDir(journalPath,
		func(path string, info fs.DirEntry, err error) error {

			if err != nil {
				fmt.Println(err)
			}
			if strings.Contains(path, ".txt") {
				file_list = append(file_list, path)
			}
			return nil
		})
	printSlice(file_list)
	choice := getSliceIndex(file_list)
	// hello this is not cursive
	fmt.Printf("You chose index %d: %s \n", choice, file_list[choice])
	fmt.Println("Are you sure you want to delete this leaf? (y/n)")

	var answer string
	fmt.Scanln(&answer)
	if answer == "y" {
		err := os.Remove(file_list[choice])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Leaf removed")
	} else {
		fmt.Println("Leaf not removed")
	}
}

func printSlice(slice []string) {
	for index, value := range slice {
		fmt.Println(index, " | name: ", value)
	}
}

// What is the best way to add an exit condition here?
func getSliceIndex(slice []string) int {
	var index int
	fmt.Println("Which leaf would you like to remove? (enter the number)")
	fmt.Scanln(&index)
	if index < 0 || index >= len(slice) {
		fmt.Println("Invalid index")
		return (getSliceIndex(slice))
	} else {
		return (index)
	}
}
