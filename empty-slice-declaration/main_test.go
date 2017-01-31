package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func createByDeclaration() {
	var a []string
	fmt.Fprint(ioutil.Discard, a)
}

func createByLiteral() {
	a := []string{}
	fmt.Fprint(ioutil.Discard, a)
}

func createByMake() {
	a := make([]string, 0)
	fmt.Fprint(ioutil.Discard, a)
}

func Benchmark_createByDeclaration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createByDeclaration()
	}
}

func Benchmark_createByLiteral(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createByLiteral()
	}
}

func Benchmark_createByMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createByMake()
	}
}

/*
Result:

```
$ go version
go version go1.7.5 darwin/amd64

$ go1.8rc3 version
go version go1.8rc3 darwin/amd64

$ go test -bench . -benchmem
testing: warning: no tests to run
Benchmark_createByDeclaration-4   	10000000	       161 ns/op	      32 B/op	       1 allocs/op
Benchmark_createByLiteral-4       	10000000	       163 ns/op	      32 B/op	       1 allocs/op
Benchmark_createByMake-4          	10000000	       167 ns/op	      32 B/op	       1 allocs/op
PASS
ok  	github.com/keijiyoshida/go-code-snippets/empty-slice-declaration	5.471s

$ go1.8rc3 test -bench . -benchmem
Benchmark_createByDeclaration-4   	10000000	       161 ns/op	      32 B/op	       1 allocs/op
Benchmark_createByLiteral-4       	10000000	       169 ns/op	      32 B/op	       1 allocs/op
Benchmark_createByMake-4          	10000000	       217 ns/op	      32 B/op	       1 allocs/op
PASS
ok  	github.com/keijiyoshida/go-code-snippets/empty-slice-declaration	6.001s
```
*/
