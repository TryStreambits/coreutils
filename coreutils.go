package coreutils

import (
	"crypto/sha512"
	"encoding/hex"
	"os"
	"strconv"
)

// GlobalFileMode is a file mode we'll use for global IO operations.
var GlobalFileMode os.FileMode

// Separator is the file system path separator
var Separator string

// NonGlobalFileMode is the file mode we'll use for non-global IO operations.
var NonGlobalFileMode os.FileMode

func init() {
	GlobalFileMode = 0777 // Set to global read/write/executable
	Separator = strconv.QuoteRune(os.PathSeparator)
	NonGlobalFileMode = 0744 // Only read/write/executable by owner, readable by group and others
}

// Sha512Sum will create a sha512sum of the string
func Sha512Sum(content string, rounds int) string {
	var hashString string

	sha512Hasher := sha512.New()                           // Create a new Hash struct
	sha512Hasher.Write([]byte(content))                    // Write the byte array of the content
	hashString = hex.EncodeToString(sha512Hasher.Sum(nil)) // Return string encoded sum of sha512sum

	if (rounds != 0) && (rounds > 1) { // If we are cycling more than one rounds
		for currentRound := 0; currentRound < rounds; currentRound++ {
			hashString = Sha512Sum(hashString, 1) // Rehash the new hashString
		}
	}

	return hashString
}
