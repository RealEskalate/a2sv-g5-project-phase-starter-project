package fs

import (
	"os"
)

/*
Checks if the file provided exists in the local fs and deletes it.
Function doesn't return an error if the file is not found
*/
func DeleteFile(filePath string) error {
	if _, err := os.Stat(filePath); err == nil {
		err := os.Remove(filePath)
		if err != nil {
			return err
		}
	}

	return nil
}
