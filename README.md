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
$ go test -bench=. -count=3

goos: linux
goarch: amd64
BenchmarkBuildFast-4   	 1938358	       661 ns/op
BenchmarkBuildFast-4   	 1976593	       610 ns/op
BenchmarkBuildFast-4   	 1970155	       639 ns/op
BenchmarkBuild-4       	  318112	      3816 ns/op
BenchmarkBuild-4       	  355080	      3438 ns/op
BenchmarkBuild-4       	  424045	      3465 ns/op
BenchmarkMarshal-4     	  234531	      4507 ns/op
BenchmarkMarshal-4     	  257199	      4652 ns/op
BenchmarkMarshal-4     	  279288	      4706 ns/op
PASS
ok  	_/home/jeffotoni/gitprojetos/concat	13.331s

```

#### Creators

Jefferson Otoni, [@jeffotoni](https://twitter.com/jeffotoni), [github.com/jeffotoni](https://github.com/jeffotoni), [linkedin.com/in/jeffotoni](https://www.linkedin.com/in/jeffotoni)   

João Vitor – [@ancogamer](https://twitter.com/ancogamer), [github.com/ancogamer](https://github.com/ancogamer), [linkedin.com/in/joão-vitor-astori-saletti](https://www.linkedin.com/in/joão-vitor-astori-saletti)


Distributed under the MIT license. See ``LICENSE`` for more information.

## Contributing

To contribuit is simples, just clone or fork the repository and send to us the pull request