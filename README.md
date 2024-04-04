# httpclient 

usage: 

```golang
package main

import "github.com/ituoga/httpclient"

func main() {
    type Response struct {
		Message string `json:"message"`
	}

	r, err := httpclient.Get[Response]("https://domain.tld")
	if err != nil {
		log.Fatal(err)
	}
    
	println(r.Message)
}
```
