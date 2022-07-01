# The Go programming language

## Install go on linux

1. Download the go tar.gz from golang.org

2. Extract it into /usr/local

```sh
~$ cd Downloads
$ tar -C /usr/local -xzf go*.linux-amd64.tar.gz
```

3. Set env vars: **$GOROOT** and **$GOPATH**

* GOROOT is where go is installed
* GOPATH is a workspace of your preference, where binaries and third party libraries are stored.

```sh
~$ nano .profile
o
~$ nano .bashrc
```

(*Add exports at the end of the file*)

```sh
export GOROOT=/usr/local/go
export GOPATH=$HOME/go            
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```

```sh
~$ source .profile
o
~$ source .bashrc
```

4. Check it out

```sh
$ echo $GOROOT
/usr/local/go

$ echo $GOPATH
/home/user/go

$ go env
```

5. Test go

```sh
$ nano main.go
```

(copy and paste)

```go
package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
}
```

```sh
$ go run main.go
hello, world
$ rm main.go
```

## Create a WORKSPACE

A workspace is where $GOPATH is pointing at and stores binaries, third party libs and the go code as well.

* bin -> binaries (.exe or .sh)
* pkg -> packages downloaded
* src -> source code

```sh
 $ mkdir go
 $ cd go
 $ mkdir bin pkg src
```

## The Go Tool

#### Version & help

```sh
$ go version
$ go help
```

#### Build an executable

```sh
$ go build x.go
$ ./x.sh
```

#### Cross compile

```text
Supported: aix/ppc64, darwin/386, darwin/amd64, dragonfly/amd64, freebsd/386, freebsd/amd64, freebsd/arm, freebsd/arm64, illumos/amd64, js/wasm, linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64, linux/ppc64le, linux/mips, linux/mipsle, linux/mips64, linux/mips64le, linux/riscv64, linux/s390x, netbsd/386, netbsd/amd64, netbsd/arm, netbsd/arm64, openbsd/386, openbsd/amd64, openbsd/arm, openbsd/arm64, plan9/386, plan9/amd64, plan9/arm, solaris/amd64, windows/386, windows/amd64, windows/arm
```

```sh
export GOOS=solaris
export GOARCH=amd64
go build main.go
```

#### Run

```sh
$ go run *.go
```

#### Clean object files

```sh
$ go clean
```

#### install (need set $GOBIN)

* This command installs the program in the computer (great in production)

```sh
$ go install main.go
```

#### list packages

```sh
$ go list ./...
```

#### Download (all) dependencie(s)

```sh
$ go get ./...
$ go get github.com/golang/example/hello
```

### Golint

#### Format (all)

```sh
$ go fmt *.go
```

#### Reports poor coding style

```sh
$ golint ./...
```

#### Reports suspicius constructs

```sh
$ golint ./...
```

## Go Doc

#### Docs for current package

```sh
$ go doc
```

#### Doc third-party packages

```sh
$ go doc pkg
$ go doc pkg.method

$ godoc pkg
$ godoc pkg method
$ godoc -src pkg method
```

#### Run godoc.org on localhost:8080

```sh
$ sudo apt install -y golang-golang-x-tools
$ godoc -http=:8080
```

## Testing

#### Command test

```sh
$ go test
$ go test ./..
```

#### test benchmark

```sh
$ go test -bench .
```

#### coverage

```sh
$ go test -cover .

$ go test -coverprofile cover.out
$ go tool cover -html=cover.out

$ go tool cover -h
```

## Go Modules

With Go modules, versions are referenced with semantic version tags:

* major.minor.patch

Ex: for v1.2.3

* major version is 1 (backwards incompatible)
* minor version is 2
* patch version is 3

#### Create new module

```sh
$ cd 09.modules/
$ go mod init github.com/manerajona/go-mod-example
$ ls
go.mod

$ go get all
(or)
$ go get go.uber.org/zap
$ ls
go.mod  go.sum

cat go.mod 
module github.com/manerajona/go-mod-example

go 1.14

require go.uber.org/zap v1.15.0 // indirect
```

#### Show all dependencies

```sh
$ go list -m all
```

#### Show dependency

```sh
$ go list -m rsc.io/q...
```

#### Get or upgrade dependencies

* v0 - Unstable
* v1 - Stable minor
* v2 - Stable major
* v0-beta - Experimental

```sh
$ go list -m -versions go.uber.org/zap
go.uber.org/zap v0.1.0-beta.1 v1.0.0-rc.1 v1.0.0-rc.2 v1.0.0-rc.3 v1.0.0 v1.1.0 v1.2.0 v1.3.0 v1.4.0 v1.4.1 v1.5.0 v1.6.0 v1.7.0 v1.7.1 v1.8.0 v1.9.0 v1.9.1 v1.10.0 v1.11.0 v1.12.0 v1.13.0 v1.14.0 v1.14.1 v1.15.0

$ go get go.uber.org/zap@v1.15.0

$ go doc zap
```

#### Specify tagged and untagged versions

**Tagged Latest**

```sh
$ go get golang.org/x/text@latest
$ go get golang.org/x/text
```

**Tagged Release**

```sh
$ go get golang.org/x/text@v1.3.1
```

**Untagged Master**

```sh
$ go get golang.org/x/text@master
```

**Untagged Commit**

```sh
$ go get golang.org/x/text@af044c0995fe
```

#### Remove unused dependencies

```sh
$ go list -m all
$ go mod tidy
$ go list -m all
```

#### Module path replacement

```sh
$ go mod edit -replace=github.com/user1/project=github.com/user2/project@master
```

#### Download all mod dependencies

```sh
$ go mod download
$ ls $GOPATH/pkg/mod
```

#### Clean local cache

```sh
$ go clean -cache -modcache -i -r
$ ls $GOPATH/pkg/mod/cache/download/
```

#### Vendor folder (Optional)

```sh
$ go mod vendor
```

#### Module mirror

A module mirror is a special kind of module proxy that caches metadata and source code in its own storage system,
allowing the mirror to continue to serve src that is not longer available. The Go team maintains a module mirror served
at *https://proxy.golang.org* (by default in +1.13v, for earlier versions set GOPROXY). For setting a custom
mirror ```export GOPROXY=proxy.example.com```.

#### Module private

The GOPRIVATE env var may be used by other tools as well to identify non-public modules. For example, if a company ran a
module proxy serving private modules, users would configure Go using:

* GOPRIVATE=*.corp1.example.com,*.corp2.example.com
* GOPROXY=proxy.example.com
* GONOPROXY=none

