package leaf

import (
	"fmt"
	"os"
	"time"
)

func NewLeaf(text string) error {
	filename := fmt.Sprintf("journal/%s.txt",
		time.Now().Format("20060102-150405"))
	// file handle to the newly created file
	// methods on returned file can be used for I/O
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(text)
	return err
}
