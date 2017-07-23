package coreutils

import (
	"errors"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// AbsPath get the absolute directory path, cleaning out any file names, home directory references, etc.
func AbsPath(path string) string {
	user, userGetErr := user.Current()

	if userGetErr == nil { // If we didn't fail getting the current user
		path = strings.Replace(path, "~", user.HomeDir+Separator, -1) // Replace any home directory reference
	}

	path, _ = filepath.Abs(path) // Get the absolute path of path

	var stripLastElement bool

	if file, openErr := os.Open(path); openErr == nil { // Attempt to open the path, to validate if it is a file or directory
		stat, statErr := file.Stat()
		stripLastElement = (statErr == nil) && !stat.IsDir() // Sets stripLastElement to true if stat.IsDir is not true
	} else { // If we failed to open the directory or file
		lastElement := filepath.Base(path)
		stripLastElement = filepath.Ext(lastElement) != "" // If lastElement is either a dotfile or has an extension, assume it is a file
	}

	if stripLastElement {
		path = filepath.Dir(path) + Separator // Strip out the last element and add the separator
	}

	return path
}

// CopyDirectory will copy a directory, sub-directories, and files
func CopyDirectory(sourceDirectory, destinationDirectory string) error {
	var copyError error

	if IsDir(sourceDirectory) { // If sourceDirectory is a valid directory
		os.MkdirAll(destinationDirectory, NonGlobalFileMode)                     // Make all the needed directories to destinationDirectory
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

// CopyFile will copy a file and its relevant permissions
func CopyFile(sourceFile, destinationFile string) error {
	var copyError error

	sourceFileStruct, sourceFileError := os.Open(sourceFile) // Attempt to open the sourceFile

	if sourceFileError == nil { // If there was not an error opening the source file
		sourceFileStats, _ := sourceFileStruct.Stat() // Get the stats of the file

		if sourceFileStats.IsDir() { // If this is actually a directory
			copyError = errors.New(sourceFile + " is a directory. Please use CopyDirectory instead.")
		} else { // If it is indeed a file
			var fileContent []byte
			sourceFileMode := sourceFileStats.Mode() // Get the FileMode of this file
			sourceFileStruct.Close()                 // Close the file

			fileContent, copyError = ioutil.ReadFile(sourceFile) // Read the source file
			copyError = WriteOrUpdateFile(destinationFile, fileContent, sourceFileMode)
		}
	} else { // If the file does not exist
		copyError = errors.New(sourceFile + " does not exist.")
	}

	return copyError
}

// GetFiles will get all the files from a directory.
func GetFiles(path string, recursive bool) ([]string, error) {
	var files []string      // Define files as a []string
	var getFilesError error // Define getFilesError as an error

	if directory, openErr := os.Open(path); openErr == nil {
		directoryContents, directoryReadError := directory.Readdir(-1)

		if directoryReadError == nil { // If there was no issue reading the directory contents
			for _, fileInfoStruct := range directoryContents { // For each FileInfo struct in directoryContents
				name := fileInfoStruct.Name()

				if recursive && fileInfoStruct.IsDir() { // If the FileInfo indicates the object is a directory and we're doing recursive file fetching
					additionalFiles, _ := GetFiles(path + Separator + name, true)
					files = append(files, additionalFiles...)
				} else if !fileInfoStruct.IsDir() { // FileInfo is not a directory
					files = append(files, path+ Separator + name) // Add to files the file's name
				}
			}
		} else { // If there was ano issue reading the directory content
			getFilesError = errors.New("Cannot read the contents of " + path)
		}
	} else { // If path is not a directory
		getFilesError = errors.New(path + " is not a directory.")
	}

	return files, getFilesError
}

// GetFilesContains will return any files from a directory containing a particular string
func GetFilesContains(path, substring string) ([]string, error) {
	var files []string                // Define files as the parsed files
	var getFilesError error           // Define getFilesError as an error
	var allDirectoryContents []string // Define allDirectoryContents as the contents returned (if any) from GetFiles

	allDirectoryContents, getFilesError = GetFiles(path, false) // Get all the files from the path

	if getFilesError == nil { // If there was no issue getting the directory contents
		for _, fileName := range allDirectoryContents { // For each file name in directory contents
			if strings.Contains(filepath.Base(fileName), substring) { // If the file name contains our substring
				files = append(files, fileName) // Append to files
			}
		}
	}

	return files, getFilesError
}

// IsDir checks if the path provided is a directory or not
func IsDir(path string) bool {
	var isDir bool
	fileObject, fileOpenError := os.Open(path) // Open currentDirectory + path

	if fileOpenError == nil { // If there was no error opening the file object
		stat, filePathError := fileObject.Stat() // Get any stats

		if filePathError == nil { // If we got the statistics properly
			isDir = stat.IsDir() // Set isDir to result from stat
		}
	}

	return isDir
}

// WriteOrUpdateFile writes or updates the file contents of the passed file under the leading filepath with the specified sourceFileMode
func WriteOrUpdateFile(file string, fileContent []byte, sourceFileMode os.FileMode) error {
	var writeDirectory string // Directory to write file

	currentDirectory, _ := os.Getwd()            // Get the working directory
	currentDirectory = AbsPath(currentDirectory) // Get the absolute path of the current working directory
	fileName := filepath.Base(file)

	if file == fileName { // If we did not specify a directory to write to
		writeDirectory = currentDirectory // Set to the current directory
	} else {
		writeDirectory = AbsPath(file)
	}

	if currentDirectory != writeDirectory { // If the currentDirectory is not the same directory as the writeDirectory
		if createDirsErr := os.MkdirAll(writeDirectory, sourceFileMode); createDirsErr != nil { // If we failed to make all the directories needed
			return errors.New("Failed to create the path leading up to " + fileName + ": " + writeDirectory)
		}
	}

	writeErr := ioutil.WriteFile(writeDirectory+fileName, fileContent, sourceFileMode)

	if writeErr != nil {
		writeErr = errors.New("Failed to write " + fileName + " in directory " + writeDirectory)
	}

	return writeErr
}
