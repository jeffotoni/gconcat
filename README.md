# gconcat.Build
>Just a simple lib, to gconcat string in #Go lang

#### Use

Example Gconcat Build
```go
package main

import (
	"github.com/jeffotoni/gconcat"
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

	println(gconcat.Build(ii))
	println(gconcat.Build(i))
	println(gconcat.Build("jeffotoni", "&", "joao", " ", 20, "/08/"))
	println(gconcat.Build([]string{"2017", " ", "2018", " ", "2020"}))
	println(gconcat.Build([]int{12, 0, 11, 0, 10, 11, 12, 23, 3}))
	println(gconcat.Build(10,9,10,20,30,40,"x", "&", "."))
	println(gconcat.Build("R$ ",23456.33, " R$ ",123.33, " R$ ",19.11))

}
```

Example on your project 
```go
package main

import (
	"log"
	"net/http"

	"github.com/jeffotoni/gconcat"
)

const PORT = ":8282"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping",
		func(w http.ResponseWriter, r *http.Request) {
			str := gconcat.Build(
                            []int{1, 2, 3, 4, 5}, " ", 
                            []string{"vamos testar nossa concat!!!"},
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

````
#### Install Using go mod in your project
```bash
$ go mod init <your-dir>
$ go run main.go
``````

#### Another possibility would be
```bash
$ go get -u github.com/jeffotoni/gconcat
```

#### Test Benchmarking

```bash
$ go test -bench . -benchmem
goos: linux
goarch: amd64
BenchmarkConcatBuildFast-4       4322348               301 ns/op             150 B/op          5 allocs/op
BenchmarkBuildFast-4             2482591               468 ns/op             160 B/op          6 allocs/op
BenchmarkBuild-4                  469720              2503 ns/op            3238 B/op         27 allocs/op
BenchmarkMarshal-4                293193              4128 ns/op             768 B/op          1 allocs/op
BenchmarkConcatMais-4                  8         133299313 ns/op        1086400464 B/op    10041 allocs/op
PASS
ok      _/home/jeffotoni/gitprojetos/gconcat     6.899s

```

#### Creators

Jefferson Otoni, [@jeffotoni](https://twitter.com/jeffotoni), [github.com/jeffotoni](https://github.com/jeffotoni), [linkedin.com/in/jeffotoni](https://www.linkedin.com/in/jeffotoni)   

João Vitor – [@ancogamer](https://twitter.com/ancogamer), [github.com/ancogamer](https://github.com/ancogamer), [linkedin.com/in/joão-vitor-astori-saletti](https://www.linkedin.com/in/joão-vitor-astori-saletti)


Distributed under the MIT license. See ``LICENSE`` for more information.

## Contributing

To contribuit is simples, just clone or fork the repository and send to us the pull request