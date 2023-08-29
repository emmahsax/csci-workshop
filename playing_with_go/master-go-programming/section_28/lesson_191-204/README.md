# `GOPATH`, Packages, Modules, and Code Organization

* Go requires you to organize your code in a specific way
* Go programmers typically keep all their Go code in a single workspace
* A _workspace_ is a directory in your file system whose path is stored in the environment variable called `GOPATH`
  * On Windows, it's in `%USERPROFILE%\go`
  * On Mac/Linux, it's in `~/go`

## Go Workspace

The workspace directory (`GOPATH`) contains the following 3 subdirectories at its root:

1. `src`: contains the source files for your own or other downloaded packages
2. `pkg`: contains go package objects - aka package archives. These are non-executable files or shared libraries ending in a `.a`
3. `bin`: contains compiled and executable Go programs. When you run `go install` on a program directory, Go will put the executable file under this folder.

## Go Packages

* A _package_ is a project directory containing `.go` files with the same package statement at the beginning. A package contains many source files each ending in `.go` extension and belonging to a single directory.
* There are two types of packages:
  * _Executable packages_ that generate files that can be run. The name of an executable package is predefined and is called `main`.
  * _Non-executable packages_ (libraries or dependencies) that are used by other packages and can have any name. They cannot be executed, only imported.

## Go Modules

> Modules are supported starting with Go v1.11, but the implementation was finalized in Go v1.13.

* A _module_ is a collection of related Go packages stored in a directory tree with a `go.mod` file at its root.
* A module contains one or more packages like a package contains one or more `.go` files.
* A file called `go.mod` defines the module's path, which is also the import path and its dependency requirements.
* Th `go` command has direct support to work with modules, including recording and resolving dependencies on other modules.
