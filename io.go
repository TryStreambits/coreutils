package codeutilsShared

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// IsDir
// Checks if the string provided is a directory or not (based on the current working directory)
func IsDir(path string) bool {
	var isDir bool
	currentWorkingDirectory, _ := os.Getwd()
	fileObject, fileOpenError := os.Open(currentWorkingDirectory + "/" + path) // Open currentDirectory + path

	if fileOpenError == nil { // If there was no error opening the file object
		fileStatistics, filePathError := fileObject.Stat() // Get any stats

		if filePathError == nil { // If we got the statistics properly
			isDir = fileStatistics.IsDir() // Set isDir to result from fileStatistics
		}
	}

	return isDir
}

// CopyDirectory
// This function will copy a directory, sub-directories, and files
func CopyDirectory(sourceDirectory, destinationDirectory string) error {
	var copyError error

	if IsDir(sourceDirectory) { // If sourceDirectory is a valid directory
		os.MkdirAll(destinationDirectory, UniversalFileMode)                     // Make all the needed directories to destinationDirectory
		sourceDirectoryFile, _ := os.Open(sourceDirectory)                       // Get the source directory "file" struct
		directoryContents, directoryReadError := sourceDirectoryFile.Readdir(-1) // Read the directory contents

		if directoryReadError == nil { /// If there was no read error on the directory
			if len(directoryContents) != 0 { // If there is content
				for _, contentItemFileInfo := range directoryContents { // For each FileInfo struct in directoryContents
					contentItemName := contentItemFileInfo.Name() // Get the name of the item
					sourceItemPath := sourceDirectory + "/" + contentItemName
					destinationItemPath := destinationDirectory + "/" + contentItemName

					if contentItemFileInfo.IsDir() { // If this is a directory
						copyError = CopyDirectory(sourceItemPath, destinationItemPath) // Copy this sub-directory and its contents
					} else { // If this is a file
						copyError = CopyFile(sourceItemPath, destinationItemPath) // Copy the directory
					}
				}
			}
		} else { // If there was a read error on the directory
			copyError = errors.New("Unable to read: " + sourceDirectory)
		}
	} else { // If sourceDirectory is not a valid directory
		copyError = errors.New(sourceDirectory + " is not a valid directory.")
	}

	return copyError
}

// CopyFile
// This function will copy a file and its relevant permissions
func CopyFile(sourceFile, destinationFile string) error {
	var copyError error

	sourceFileStruct, sourceFileError := os.Open(sourceFile) // Attempt to open the sourceFile

	if sourceFileError == nil { // If there was not an error opening the source file
		sourceFileStats, _ := sourceFileStruct.Stat() // Get the stats of the file

		if sourceFileStats.IsDir() { // If this is actually a directory
			copyError = errors.New(sourceFile + " is a directory. Please use CopyDirectory instead.")
		} else { // If it is indeed a file
			sourceFileMode := sourceFileStats.Mode() // Get the FileMode of this file
			sourceFileStruct.Close()                 // Close the file

			fileContent, _ := ioutil.ReadFile(sourceFile) // Read the source file
			copyError = WriteOrUpdateFile(destinationFile, fileContent, sourceFileMode)
		}
	} else { // If the file does not exist
		copyError = errors.New(sourceFile + " does not exist.")
	}

	return copyError
}

// WriteOrUpdateFile
// Function to write or update the file contents of the passed file under the leading filepath with the specified sourceFileMode
func WriteOrUpdateFile(file string, fileContent []byte, sourceFileMode os.FileMode) error {
	destinationFileDirectories := filepath.Dir(file) // Get the directories leading up to the destinationFileDirectories

	if sourceFileMode == 0777 { // If things are global rwe
		sourceFileMode = UniversalFileMode // No, I can't let you do that Dave. (Changes to 744)
	}

	if destinationFileDirectories != "." { // If this is not the same directory as working dir
		os.MkdirAll(destinationFileDirectories, sourceFileMode) // Make all the necessary directories
	}

	return ioutil.WriteFile(file, fileContent, sourceFileMode) // Write the file content
}
