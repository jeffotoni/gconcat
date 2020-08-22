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


#### Creators

Jefferson Otoni, [@jeffotoni](https://twitter.com/jeffotoni), [github.com/jeffotoni](https://github.com/jeffotoni), [linkedin.com/in/jeffotoni](https://www.linkedin.com/in/jeffotoni)   

João Vitor – [@ancogamer](https://twitter.com/ancogamer), [github.com/ancogamer](https://github.com/ancogamer), [linkedin.com/in/joão-vitor-astori-saletti](https://www.linkedin.com/in/joão-vitor-astori-saletti)


Distributed under the MIT license. See ``LICENSE`` for more information.

## Contributing

To contribuit is simples, just clone or fork the repository and send to us the pull request :D