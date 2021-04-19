# gconcat

[![GoDoc][1]][2]
[![GitHub release][3]][4]
[![Coverage Status][5]][6]
[![Build Status][7]][8]
[![Go Report Card][9]][10]
[![License][11]][11]

[1]: https://godoc.org/github.com/jeffotoni/gconcat?status.svg
[2]: https://godoc.org/github.com/jeffotoni/gconcat
[3]: https://img.shields.io/github/v/release/jeffotoni/gconcat?include_prereleases
[4]: https://github.com/jeffotoni/gconcat/releases
[5]: https://coveralls.io/repos/github/jeffotoni/gconcat/badge.svg?branch=master
[6]: https://coveralls.io/github/jeffotoni/gconcat?branch=master
[7]: https://travis-ci.com/jeffotoni/gconcat.svg?branch=master
[8]: https://travis-ci.com/jeffotoni/gconcat
[9]: https://goreportcard.com/badge/github.com/jeffotoni/gconcat
[10]: https://goreportcard.com/report/github.com/jeffotoni/gconcat
[11]: https://img.shields.io/github/license/jeffotoni/gconcat

>A simple lib for concatenation, it accepts several types as a parameter and returns a string. A battery of tests was done, there are some complexities that we cannot escape to have the best computational cost 

## Example 

#### Concatenating some basic types in Go
```go
package main

import (
	g "github.com/jeffotoni/gconcat"
)

func main() {
	str := g.Concat("/api/v1", "/", 39383838, "/", 129494, "/product/", 2012)
	println(str)
}
```

#### Concatenating types using interfaces
```go
package main

import (
	g "github.com/jeffotoni/gconcat"
)

func main() {
	var ii []interface{}
	ii = append(ii, "jeffotoni")
	ii = append(ii, " ")
	ii = append(ii, "joao")
	ii = append(ii, " ")
	ii = append(ii, 2021)

	var i interface{}
	i  = "jeffotoni"

	println(g.Concat(ii))
	println(g.Concat(i))
	println(g.Concat("jeffotoni", "&", "joao", " ", 20, "/08/"))
	println(g.Concat([]string{"2017", " ", "2018", " ", "2020"}))
	println(g.Concat([]int{12, 0, 11, 0, 10, 11, 12, 23, 3}))
	println(g.Concat(10,9,10,20,30,40,"x", "&", "."))
	println(g.Concat("R$ ",23456.33, " R$ ",123.33, " R$ ",19.11))
}
```

## Example on your project 
```go
package main

import (
	"log"
	"net/http"

	g "github.com/jeffotoni/gconcat"
)

const PORT = ":8282"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping",
		func(w http.ResponseWriter, r *http.Request) {
			str := g.Concat(
                            []int{1, 2, 3, 4, 5}, " ", 
                            []string{"Let's test our concat!!!"},
                    )
		w.Write([]byte(str))
		})
	    
	server := &http.Server{
		Addr:    PORT,
		Handler: mux,
	}
	println("Start Run: ", PORT)
	log.Fatal(server.ListenAndServe())
}

```

### Some types allowed
> - bool
> - int
> - int32
> - int64
> - interface
> - string
> - uint
> - float32
> - float64
> - []int
> - []int32
> - []int64
> - []interface
> - []string
> - []uint
> - []float32
> - []float64

## Install Using go mod in your project
```bash
$ go mod init <your-dir>
$ go mod tidy
$ go run main.go
``````

#### Another possibility would be
```bash
$ go get -u github.com/jeffotoni/gconcat
```

#### Test Benchmarking

```bash
$ go test -bench . -benchmem
goarch: amd64
pkg: github.com/jeffotoni/gconcat
cpu: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
BenchmarkStringPlus-12         	22920072	        57.53 ns/op	      80 B/op	       1 allocs/op
BenchmarkStringLongPlus-12     	12399344	        91.54 ns/op	     416 B/op	       1 allocs/op
BenchmarkConcatNoPlus-12       	 3659144	       547.9 ns/op	     464 B/op	       7 allocs/op
BenchmarkConcatIntString-12    	 5304790	       217.8 ns/op	     128 B/op	       6 allocs/op
BenchmarkLongJoin-12           	10140075	       111.0 ns/op	     448 B/op	       1 allocs/op
BenchmarkLongSprintf-12        	 3758727	       322.5 ns/op	     480 B/op	       5 allocs/op
BenchmarkBuilder-12            	 4021644	       265.1 ns/op	    2396 B/op	       0 allocs/op
BenchmarkConcatVector-12       	  238450	      6137 ns/op	    4616 B/op	      84 allocs/op
BenchmarkMarshal-12            	  411598	      2906 ns/op	     768 B/op	       1 allocs/op
PASS
ok  	github.com/jeffotoni/gconcat	13.893s
```

#### Creators

Jefferson Otoni, [@jeffotoni](https://twitter.com/jeffotoni), [github.com/jeffotoni](https://github.com/jeffotoni), [linkedin.com/in/jeffotoni](https://www.linkedin.com/in/jeffotoni)   

João Vitor – [@ancogamer](https://twitter.com/ancogamer), [github.com/ancogamer](https://github.com/ancogamer), [linkedin.com/in/joão-vitor-astori-saletti](https://www.linkedin.com/in/joão-vitor-astori-saletti)


Distributed under the MIT license. See ``LICENSE`` for more information.

## Contributing

To contribuit is simples, just clone or fork the repository and send to us the pull request
