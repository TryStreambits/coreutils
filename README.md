# coreutils

This is a Go package containing useful functions used throughout our Go-based software.

## License

coreutils is licensed under Apache-2.0

## Usage

### Import

``` go
import "github.com/StroblIndustries/coreutils"
```

### Variables

```go
var GlobalFileMode os.FileMode
```
GlobalFileMode as a file mode we'll use for "global" operations such as when
doing IO as root

```go
var Separator string
```
Separator is the file system path separator

```go
var UniversalFileMode os.FileMode
```
UniversalFileMode as a file mode we'll wherever we can

### Functions

#### func  AbsDir

```go
func AbsDir(path string) string
```
AbsDir get the absolute directory path, cleaning out any file names, home
directory references, etc.

#### func  CopyDirectory

```go
func CopyDirectory(sourceDirectory, destinationDirectory string) error
```
CopyDirectory will copy a directory, sub-directories, and files

#### func  CopyFile

```go
func CopyFile(sourceFile, destinationFile string) error
```
CopyFile will copy a file and its relevant permissions

#### func  ExecCommand

```go
func ExecCommand(utility string, args []string, liveOutput bool) string
```
ExecCommand executes a utility with args and returning the stringified output

#### func  ExecutableExists

```go
func ExecutableExists(executableName string) bool
```
ExecutableExists checks if an executable exists

#### func  FindClosestFile

```go
func FindClosestFile(file string) (string, error)
```
FindClosestFile will return the closest related file to the one provided from a
specific path

#### func  GetFiles

```go
func GetFiles(path string) ([]string, error)
```
GetFiles will get all the files from a directory.

#### func  GetFilesContains

```go
func GetFilesContains(path, substring string) ([]string, error)
```
GetFilesContains will return any files from a directory containing a particular
string

#### func  InputMessage

```go
func InputMessage(message string) string
```
InputMessage fetches input after message

#### func  IsDir

```go
func IsDir(path string) bool
```
IsDir checks if the path provided is a directory or not

#### func  OutputStatus

```go
func OutputStatus(status bool, message string)
```
OutputStatus outputs a "check" or not check based on true / false status, along
with the message

#### func  Sha512Sum

```go
func Sha512Sum(content string, rounds int) string
```
Sha512Sum will create a sha512sum of the string

#### func  WriteOrUpdateFile

```go
func WriteOrUpdateFile(file string, fileContent []byte, sourceFileMode os.FileMode) error
```
WriteOrUpdateFile writes or updates the file contents of the passed file under
the leading filepath with the specified sourceFileMode
