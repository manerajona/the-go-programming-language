# The Go programming language

### Install go on linux
-----
1. Download the go tar.gz from golang.org

2. Extract it into /usr/local
```sh
~$ cd Downloads
$ tar -C /usr/local -xzf go*.linux-amd64.tar.gz
```
3. Set **$GOROOT** and **$GOPATH** 
```sh
~$ nano .bashrc
```
(*Add exports at the end of the file*)
```sh
export PATH=$PATH:/usr/local/go/bin
```
```sh
~$ source .bashrc
```
```sh
~$ nano .profile
```
(*Add exports at the end of the file*)
```sh
export GOROOT=/usr/local/go
export GOPATH=$HOME/go            
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```
```sh
~$ source .profile
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

### CREATE YOUR WORKSPACE
-----
* bin -> binaries (.exe or .sh)
* pkg -> packages downloaded
* src -> source code
```sh
 $ mkdir go
 $ cd go
 $ mkdir bin pkg src
```

### The Go Tool
-----
```sh
$ go version
$ go help
```
#### Build & run:
```sh
$ go build *.go
$ go run *.go
```
#### Clean object files:
```sh
$ go clean
```
#### install (need set $GOBIN)
```sh
$ go intall main.go
```
#### Format (all):
```sh
$ go fmt ./...
```
#### list packages:
```sh
$ go list
```
#### Download (all) dependencies:
```sh
$ go get ./...
$ go get github.com/golang/example/hello
```

### MODULES
--------

#### Create new module:
```sh
$ go mod init example.com/username/repo
$ cat go.mod
$ cat go.sum
```
#### Testing:
```sh
$ go test 
```
#### Show all dependencies:
```sh
$ go list -m all
```
#### Get or upgrade dependencies:
```sh
$ go list -m -versions rsc.io/sampler
$ go get rsc.io/sampler@1.3.1
```
#### Clean dependencies:
```sh
$ go mod tidy
```
