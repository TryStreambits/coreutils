package codeutilsShared

import (
	"crypto/sha512"
	"encoding/hex"
	"os"
)

// GlobalFileMode as a file mode we'll use for "global" operations such as when doing IO as root
var GlobalFileMode os.FileMode

// UniversalFileMode as a file mode we'll wherever we can
var UniversalFileMode os.FileMode

func init() {
	GlobalFileMode = 0777    // Set to global read/write/executable
	UniversalFileMode = 0744 // Only read/write/executable by owner, readable by group and others
}

// Sha512Sum will create a sha512sum of the string
func Sha512Sum(content string) string {
	sha512Hasher := sha512.New()                     // Create a new Hash struct
	sha512Hasher.Write([]byte(content))              // Write the byte array of the content
	return hex.EncodeToString(sha512Hasher.Sum(nil)) // Return string encoded sum of sha512sum
}
