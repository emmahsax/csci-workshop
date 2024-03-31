# Testing in Go

To run tests:

```bash
go test -v
```

To see how much coverage:

```bash
go test -cover .
```

To get which parts are covered:

```bash
go test -coverprofile=coverage.out
```

To open coverage nice and pretty:

```bash
go tool cover -html=coverage.out
```

#### Note

I found it helpful to add this to my `~/.zshrc` temporarily while working on these projects:

```
export GOPATH="$HOME/git/csci-workshop/playing_with_go/testing-in-go"
```
