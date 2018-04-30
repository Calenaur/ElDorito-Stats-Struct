# ElDorito-Stats-Struct
---
This library aims to provide an up to date version of the datastructure [ElDorito](https://github.com/ElDewrito/ElDorito) publishes at the end of a match.


## Installation

`go get github.com/Calenaur/ElDorito-Stats-Struct`

## Usage
```go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Calenaur/Eldorito-Stats-Struct/data"
)

func main(){
	raw, err := ioutil.ReadFile("PATH-TO-EXAMPLE-JSON")

	if (err != nil) {
		panic(err.Error())
	}

	srv := data.Server{}
	json.Unmarshal(raw, &srv)

	fmt.Println(srv.ServerName)
}
```
