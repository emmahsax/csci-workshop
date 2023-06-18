# master-go-programming

This repo contains the lessons I completed when taking [Master Go Programming: The Complete Go Bootcamp 2023](https://www.udemy.com/course/master-go-programming-complete-golang-bootcamp/).

NOTE: If we want to run any lesson's code, we can pull the file out to the main directory, rename the package to `main`, and rename the lesson to `main`, and then run `go run FILE_NAME.go`.

### Running and building

To compile and run immediately:

```bash
go run main.go
```

To compile and build an executable, which you can then run separately:

```bash
go build # With optional -o output file
# Builds file master-go-programming
./master-go-programming
# Runs the executable
```

To compile for other architectures:

```bash
# Compile the application for Windows and 64bit CPUs:
GOOS=windows GOARCH=amd64 go build -o my_go_program.exe

# For Linux and Mac it's not necessary to use a file's extension.
# Compile the application for Linux and 64bit CPUs:
GOOS=linux GOARCH=amd64 go build -o  my_go_program

# Compile the application for Mac and 64bit CPUs:
GOOS=darwin GOARCH=amd64 go build -o  my_go_program
```

### Formatting

Running the following will format the go code according to standards, and write the output back to the file:

```bash
gofmt -w FILE_NAME.go
# OR
gofmt -w DIR_NAME
# OR
go fmt FILE_NAME.go
```

\*Adding `-l` will print out which files are modified, if passing in a directory


---

Tab character for reference

```go
import (
	"fmt"
	"os"
)
```
