# concat.Build
>Just a simple lib, to concat string in #Go lang

#### Use

Example on your project 
````go
package main

import (
	"log"
	"net/http"

	"github.com/jeffotoni/concat"
)

const PORT = ":8282"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping",
		func(w http.ResponseWriter, r *http.Request) {
			str := concat.Build(
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
#### Install Using go mod 
```bash
$ go mod init <your-dir>
$ go run main.go
``````

#### Another possibility would be
```bash
$ go get -u github.com/jeffotoni/concat

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
ok      _/home/jeffotoni/gitprojetos/concat     6.899s

```

#### Creators

Jefferson Otoni, [@jeffotoni](https://twitter.com/jeffotoni), [github.com/jeffotoni](https://github.com/jeffotoni), [linkedin.com/in/jeffotoni](https://www.linkedin.com/in/jeffotoni)   

João Vitor – [@ancogamer](https://twitter.com/ancogamer), [github.com/ancogamer](https://github.com/ancogamer), [linkedin.com/in/joão-vitor-astori-saletti](https://www.linkedin.com/in/joão-vitor-astori-saletti)


Distributed under the MIT license. See ``LICENSE`` for more information.

## Contributing

To contribuit is simples, just clone or fork the repository and send to us the pull request